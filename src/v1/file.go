package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"online-library/models"
	v1 "online-library/repo/v1"
)

//FileHandler handles the file upload & download & delete  request
func FileHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		data := models.FileInfo{
			FileName: req.FormValue("fileName"),
			Author:   req.FormValue("author"),
			Region:   req.FormValue("region"),
			Category: req.FormValue("category"),
		}
		data, err := v1.DownloadFile(data)
		if err != nil {
			log.Println("Failed to download the file", data.FileName, " err :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		byteData, err := json.Marshal(data)
		if err != nil {
			log.Println("Failed to marshal the data , err :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(byteData)
		if err != nil {
			log.Println("Failed to send file data for file name :", data.FileName)
		}
	case "POST":
		req.ParseMultipartForm(10 << 20) // 10 MB
		file, handler, err := req.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := models.FileInfo{
			FileName: handler.Filename,
			Author:   req.FormValue("author"),
			Region:   req.FormValue("region"),
			Category: req.FormValue("category"),
			File:     file,
		}
		if v1.UploadFile(data) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "DELETE":
		var data models.FileInfo
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Failed to read the delete request , err : ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal ther request data, err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = v1.DeleteFile(data)
		if err != nil {
			log.Println("Failed to delete the file ", data.FileName)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		log.Println("Method not supported")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
