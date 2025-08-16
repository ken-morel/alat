package rcfile

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func getFileDest(name string) string {
	var stem string
	ext := filepath.Ext(name)
	end := len(name) - len(ext) - 1
	if end >= len(name) || end < 0 {
		stem = "alat-unknown-file"
	} else {
		stem = name[0:end]
	}
	n := 0
	path := filepath.Join(config.Destination, name)
	for {
		_, err := os.Stat(path)
		if err != nil {
			break
		} else {
			n += 1
			path = filepath.Join(config.Destination, fmt.Sprintf("%s-%d.%s", stem, n, ext))
		}
	}
	return path
}

func HandleSendRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server: receiving request")
	var fileSize int
	var fileName string
	fmt.Println(r.Header)
	_, err := fmt.Sscanf(r.Header.Get("File-Size"), "%d", &fileSize)
	if err != nil {
		http.Error(w, "Could not parse File-Size"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = fmt.Sscanf(r.Header.Get("File-Name"), "%s", &fileName)
	if err != nil {
		http.Error(w, "Could not parse File-Name"+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Parsed parameters: ", fileName, fileSize)
	output, err := os.Create(getFileDest(fileName))
	if err != nil {
		http.Error(w, "Could not create local file: "+err.Error(), 500)
		return
	}
	defer output.Close()
	buffer := make([]byte, 1024)
	received := 0
	fmt.Println("Copying chunks")
	for received < fileSize {
		size, _ := r.Body.Read(buffer)
		if size == 0 {
			break
		}
		received += size
		fmt.Printf("Received %d of %d", received, fileSize)
		_, err = output.Write(buffer[0:size])
		if err != nil {
			http.Error(w, "Could not write to local file: "+err.Error(), 500)
			return
		}
	}
	w.WriteHeader(200)
	fmt.Println("File received succesfully")
}
