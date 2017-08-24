package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func send(filename string, url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		return fmt.Errorf("error writing to buffer: %s", err)
	}

	fh, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	if _, err = io.Copy(fileWriter, fh); err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return nil
}

func main() {
	if err := send("upload.html", "http://localhost:9090/upload"); err != nil {
		log.Fatalln(err)
	}
}
