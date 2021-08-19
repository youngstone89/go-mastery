package form

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func saveHandler(w http.ResponseWriter, r *http.Request) {
	//MultipartReader를 이용해서 받은 파일을 읽는다
	reader, err := r.MultipartReader()

	//에러가 발생하면 던진다
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//for로 복수 파일이 있는 경우에 모든 파일이 끝날때까지 읽는다
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		formName := part.FormName()
		fmt.Println(formName)

		fileDir, _ := os.Getwd()
		//uploadedfile 디렉토리에 받았던 파일 명으로 파일을 만든다
		uploadedFile, err := os.Create(fileDir + "/pkg/form/downloads/" + part.FormName())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadedFile.Close()
			return
		}

		//만든 파일에 읽었던 파일 내용을 모두 복사
		_, err = io.Copy(uploadedFile, part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadedFile.Close()
			return
		}
	}
}

func DoReceiveTest()  {
	http.HandleFunc("/save",saveHandler)
	//서버 시작
	http.ListenAndServe(":8080", nil)
}
