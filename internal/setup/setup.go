package setup

import (
	"code-execution-sandbox/internal/config"
	pb "code-execution-sandbox/proto"
	"context"
	"log"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type Server struct {
	pb.UnimplementedFileDownloadServiceServer
}

type FileContent struct {
	FilePath string
	st       *storage_go.Client
}

func (s *Server) DownloadFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	config.LoadEnv()
	reference_id := config.GetEnv("SUPABASE_PROJECT_ID", "")
	api_key := config.GetEnv("SUPABASE_PROJECT_SERVICE_API_KEY", "")

	log.Println("Connection established .....")

	st := storage_go.NewClient("https://"+reference_id+".supabase.co/storage/v1", api_key, nil)

	baseFolder := "base_nextjs_page_router_js"

	q := []string{baseFolder}
	arr := []string{}

	// create channel
	jobs := make(chan FileContent)

	// create worker pool
	workerCount := 5
	for i := 0; i < workerCount; i++ {
		go worker(jobs)
	}

	for len(q) > 0 {
		filePath := q[0]
		arr = append(arr, filePath)
		q = q[1:]
		resp := GetFilesList(st, filePath)
		if len(resp) < 1 {
			pt := "/" + filePath
			log.Println("path to download ", pt)
			jobs <- FileContent{filePath, st}

		} else {
			p := "/" + filePath
			err := os.MkdirAll(p, 0755)
			if err != nil {
				log.Println("Error while creating folder")
			}
			for i := 0; i < len(resp); i++ {
				q = append(q, filePath+"/"+resp[i].Name)
			}
		}
	}

	close(jobs)

	return &pb.FileResponse{
		FilesDownloaded: arr,
	}, nil
}

func worker(jobs chan FileContent) {
	for job := range jobs {
		AddContentToFile(job.FilePath, job.st)
	}
}
func AddContentToFile(filePath string, st *storage_go.Client) {
	pt := "/" + filePath
	resp, err := st.DownloadFile("code_pilot", pt)
	if err != nil {
		log.Println("Error while downloading file from supabase ", err.Error())
		return
	}
	content := string(resp)
	p := "/" + filePath
	file, err := os.Create(p)
	if err != nil {
		log.Println("Error while creating file ", err.Error())
		return
	}
	defer file.Close()

	_, file_err := file.WriteString(content)
	if file_err != nil {
		log.Println("Error writing in file")
	}
	log.Println("Done writing content to ", filePath)
}

func GetFilesList(st *storage_go.Client, filePath string) []storage_go.FileObject {
	resp, err := st.ListFiles("code_pilot", filePath, storage_go.FileSearchOptions{
		Limit:  1000,
		Offset: 0,
		SortByOptions: storage_go.SortBy{
			Column: "",
			Order:  "",
		},
	})
	if err != nil {
		log.Println("Error while Listing files/folder " + err.Error())
	}
	return resp
}
