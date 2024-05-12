package receiver

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

func PrintDownloadPercent(done chan int64, path string, total int64) {
	stop := false
	bar := progressbar.Default(100)

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
			}

			fi, err := file.Stat()
			if err != nil {
				fmt.Println(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			percent := float64(size) / float64(total) * 100
			bar.Set(int(percent))
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

func Receive(ip, dest string) {
	addr := "http://" + ip + ":2595"
	headResp, err := http.Head(addr)
	if err != nil {
		fmt.Println("Unable to process request")
		return
	}

	defer headResp.Body.Close()

	fileName := strings.Replace(headResp.Header.Get("Content-Disposition"), "attachment; filename=", "", 1)

	fmt.Printf("Receiving file %s from %s\n", fileName, ip)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(fileName)

	start := time.Now()

	out, err := os.Create(path.String())
	if err != nil {
		fmt.Println(path.String())
		fmt.Println("unable to create file")
		return
	}

	defer out.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println("Receving file size not found")
		return
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, path.String(), int64(size))

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println("Unable to connect to sender. please check IP")
		return
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("unable to write in file")
		return
	}

	done <- n

	elapsed := time.Since(start)
	fmt.Printf("Download completed in %s", elapsed)
}
