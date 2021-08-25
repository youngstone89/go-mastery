package form

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func DoMultipartFormTest()  {
	fileDir, _ := os.Getwd()
	fileName1 := "upload-file1.txt"
	fileName2 := "upload-file2.txt"
	fileName3 := "94711dd7-f91e-4d66-9e78-5e12fb69823a"
	filePath1 := path.Join(fileDir + "/pkg/form/", fileName1)
	filePath2 := path.Join(fileDir + "/pkg/form/", fileName2)
	filePath3 := path.Join(fileDir + "/pkg/form/", fileName3)

	file1, _ := os.Open(filePath1)
	file2, _ := os.Open(filePath2)
	file3, _ := os.Open(filePath3)
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("service","asraas")

	for i, file := range []string{
		filePath1,
		filePath2,
		filePath3,
	} {
		filebase := filepath.Base(file)
		fmt.Println("filebase:"+filebase)
		filePart, err := writer.CreateFormFile(fmt.Sprintf("file%d", i+1),
			filebase)
		if err != nil {
			panic(err)
		}
		//content := fmt.Sprintf("hi there %d",i)
		//contentInByteArr,err := json.Marshal(content)
		//if err != nil {
		//	panic(err)
		//}
		//filePart.Write(contentInByteArr)


		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(filePart, f)
		_ = f.Close()
		if err != nil {
			panic(err)
		}
	}

	err := writer.Close()
	if err != nil {
		panic(err)
	}


	//p1, err := writer.CreateFormField(fileName1)
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(p1,file1)
	//
	//p2, err := writer.CreateFormField(fileName2)
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(p2,file2)

	//
	//part2, _ := writer.CreateFormFile("audioFileName2", filepath.Base("2"+file.Name()))
	//io.Copy(part2, file)

	//writer.Close()

	ctx, cancel := context.WithTimeout(context.Background(),
		60*time.Second)
	defer cancel()


	//r, _ := http.NewRequest("POST", body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"http://127.0.0.1:8080/save", body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	fmt.Printf("%s",req.Body)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("expected status %d; actual status %d",
			http.StatusOK, resp.StatusCode)
	}

	fmt.Printf("\n%s", b)

}