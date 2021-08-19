package form

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func DoMultipartFormTest()  {
	fileDir, _ := os.Getwd()
	fileName1 := "upload-file1.txt"
	fileName2 := "upload-file2.txt"
	filePath1 := path.Join(fileDir + "/pkg/form/", fileName1)
	filePath2 := path.Join(fileDir + "/pkg/form/", fileName2)

	file1, _ := os.Open(filePath1)
	file2, _ := os.Open(filePath2)
	defer file1.Close()
	defer file2.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("service","asraas")

	p1, err := writer.CreateFormField(fileName1)
	if err != nil {
		panic(err)
	}
	io.Copy(p1,file1)

	p2, err := writer.CreateFormField(fileName2)
	if err != nil {
		panic(err)
	}
	io.Copy(p2,file2)

	//
	//part2, _ := writer.CreateFormFile("audioFileName2", filepath.Base("2"+file.Name()))
	//io.Copy(part2, file)

	writer.Close()

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/save", body)

	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Status)
	fmt.Println(resp.Body)

	
}