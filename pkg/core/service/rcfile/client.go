package rcfile

import (
	"alat/pkg/core/address"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type SendFileStatus struct {
	Percent  float32
	Error    error
	Received bool
}

func SendFile(channel chan<- SendFileStatus, path string, address address.Address, token string) {
	fmt.Println("Client sending file ", path)
	var file *os.File
	info, err := os.Stat(path)
	if err == nil {
		file, err = os.Open(path)
	}
	if err != nil {
		channel <- SendFileStatus{
			Error: err,
		}
		close(channel)
		return
	}
	req, err := http.NewRequest("POST", "http://"+address.String()+"/rcfile/send", file)
	if err != nil {
		fmt.Println("Error", err.Error())
		channel <- SendFileStatus{
			Error: err,
		}
		close(channel)
		return
	}
	req.Header.Set("File-Size", fmt.Sprintf("%d", info.Size()))
	req.Header.Set("Pair-Token", token)
	req.Header.Set("File-Name", filepath.Base(path))
	client := http.Client{}
	fmt.Println("Client making request")
	res, err := client.Do(req)
	fmt.Println("Client request fulfilled!", err)
	if err != nil {
		channel <- SendFileStatus{
			Error: err,
		}
		close(channel)
		return
	} else {
		defer res.Body.Close()
		buf := make([]byte, 500)
		res.Body.Read(buf)
		os.Stderr.Write(buf)
		fmt.Println("Received satus ", res.StatusCode)
		channel <- SendFileStatus{
			Percent:  100,
			Error:    nil,
			Received: res.StatusCode == http.StatusOK,
		}
	}
}
