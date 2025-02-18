package setup

import (
	"code-sandbox/internal/config"
	pb "code-sandbox/proto"
	"context"
	"log"
	"os/exec"

	storage_go "github.com/supabase-community/storage-go"
)

type Server struct {
	pb.UnimplementedFileDownloadServiceServer
}

func (s *Server) DownloadFiles(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	config.LoadEnv()
	folderName := req.GetFolderName()
	reference_id := config.GetEnv("SUPABASE_PROJECT_ID", "")
	api_key := config.GetEnv("SUPABASE_PROJECT_SERVICE_API_KEY", "")

	st := storage_go.NewClient("https://"+reference_id+".supabase.co/storage/v1", api_key, nil)

	exec.Command("mkdir " + folderName)

	baseFolder := "base_nextjs_page_router_js"

	q := []string{baseFolder}
	arr := []string{}

	for len(q) > 0 {
		filePath := q[0]
		arr = append(arr, filePath)
		q = q[1:]
		resp := GetFilesList(st, filePath)
		if len(resp) < 1 {
			exec.Command("touch " + folderName + "/" + filePath).Run()
		} else {
			exec.Command("mkdir " + folderName + "/" + filePath).Run()
			for i := 0; i < len(resp); i++ {
				q = append(q, filePath+"/"+resp[i].Name)
			}
		}
	}

	return &pb.FileResponse{
		FilesDownloaded: arr,
	}, nil
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
		log.Println("Error while Listing files/folder")
	}
	return resp
}
