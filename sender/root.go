package sender

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gxanshu/go-transfer/utils"
)

func Send(fileName string) {
	ip, err := utils.GetLocalIP()
	if err != nil {
		fmt.Println("Unable to fetch system's IP", err)
		return
	}

	filePath, err := utils.GetFilePath(fileName)
	if err != nil {
		fmt.Println("Error get file path", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error reading file info:", err)
		return
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read file content")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

		w.Write(fileContent)
	})

	fmt.Printf("Server is started on \n %s \n you can copy this IP and start receving with \n go-transfer receive %s \n", ip, ip)

	http.ListenAndServe(":2595", nil)
}
