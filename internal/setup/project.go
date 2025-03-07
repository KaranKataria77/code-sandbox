package setup

import (
	"code-execution-sandbox/internal/config"
	pb "code-execution-sandbox/proto"
	"context"
	"log"
	"sync"

	"code-execution-sandbox/internal/utils"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Server struct {
	pb.UnimplementedFileDownloadServiceServer
}

func ReadFiles(srcPath string, delimiter string) (*s3.ListObjectsV2Output, error) {
	config.LoadEnv()
	svc, err := utils.GetS3Client()

	if err != nil {
		log.Println("Error while creating new s3 client " + err.Error())
		return nil, err
	}
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(config.GetEnv("AWS_BUCKET", "")),
		Delimiter: aws.String(delimiter),
		Prefix:    aws.String(srcPath), // "base/base_nextjs_page_router_js/"
	})

	if err != nil {
		log.Println("Error while listing folder from bucket ", err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *Server) CreateFolder(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {

	var requestJson struct {
		ProjectId string `json:"projecId" binding:"required"`
		AppName   string `json:"appName" binding:"required"`
	}

	requestJson.ProjectId = req.GetProjectId()
	requestJson.AppName = req.GetAppName()

	srcPath := requestJson.ProjectId + "/" + requestJson.AppName

	log.Println("srcpath = "+srcPath, requestJson)

	jobs := make(chan string, 10)

	var wg sync.WaitGroup

	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go createFolderWorker(jobs, &wg)
	}

	q := []string{srcPath}

	for len(q) > 0 {
		log.Println("queue = ", q)
		path := q[0]
		q = q[1:]
		jobs <- path
		resp, err := ReadFiles(path, "/")
		if err != nil {
			log.Println("Error reading folders from s3 " + err.Error())
			break
		}

		for _, p := range resp.CommonPrefixes {
			q = append(q, (*p.Prefix)[0:len(*p.Prefix)])
		}
	}

	close(jobs)

	wg.Wait()

	DownloadFile(requestJson.ProjectId, requestJson.AppName)

	return &pb.FileResponse{FilesDownloaded: []string{"all"}}, nil
}

func DownloadFile(projectId string, appName string) {

	srcPath := projectId + "/" + appName

	read_file_resp, err := ReadFiles(srcPath, "")

	if err != nil {
		log.Println("Error while reading files " + err.Error())
	}

	jobs := make(chan string)

	var wg sync.WaitGroup

	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go writeFileWorker(jobs, &wg)
	}

	for _, path := range read_file_resp.Contents {
		jobs <- (*path.Key)
	}

	close(jobs)

	wg.Wait()

	log.Println("File writing done ")
}

func WriteFiles(srcPath string) {
	config.LoadEnv()
	svc, err := utils.GetS3Client()

	if err != nil {
		log.Println("Error while downloading file " + err.Error())
		return
	}
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.GetEnv("AWS_BUCKET", "")),
		Key:    aws.String(srcPath),
	})

	if err != nil {
		log.Println("error while reading file from s3 ", err.Error())
		return
	}

	defer resp.Body.Close()

	file, err := os.Create(srcPath)

	if err != nil {
		log.Println("Failed to create file on local " + err.Error())
		return
	}

	defer file.Close()

	_, f_err := file.ReadFrom(resp.Body)

	if f_err != nil {
		log.Println("Error while writing in file " + f_err.Error())
		return
	}
}

func writeFileWorker(jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		WriteFiles(job)
	}
}

func createFolderWorker(jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		err := os.MkdirAll(job, os.ModePerm)
		if err != nil {
			log.Println("Error while creating folder " + err.Error())
		}
	}
}
