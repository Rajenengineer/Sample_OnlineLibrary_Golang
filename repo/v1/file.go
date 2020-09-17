package v1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"online-library/config"
	"online-library/models"
	"os"
	"strings"
)

//DeleteFile deletes the file from system
func DeleteFile(data models.FileInfo) error {
	if validate(data, false) == false {
		return errors.New("Bad request or param received in file")
	}
	rawFileName := strings.Split(data.FileName, ".")[0]
	relFileName := fmt.Sprintf("%s/%s_%s_%s_%s.txt", config.Conf.StoragePath, rawFileName, data.Author, data.Category, data.Region)
	err := os.Remove(relFileName)
	if err != nil {
		log.Println("Failed to delete the file ,err ", err)
		return err
	}
	return nil
}

//UploadFile stores the file in storagePath (see config)
func UploadFile(data models.FileInfo) error {
	if validate(data, true) == false {
		return errors.New("Bad request or param received in file")
	}
	rawFileName := strings.Split(data.FileName, ".")[0]
	relFileName := fmt.Sprintf("%s/%s_%s_%s_%s.txt", config.Conf.StoragePath, rawFileName, data.Author, data.Category, data.Region)
	fp, err := os.Create(relFileName)
	if err != nil {
		log.Println("Failed to create file err :", err)
		return err
	}
	defer fp.Close()
	fileBytes, err := ioutil.ReadAll(data.File)
	if err != nil {
		fmt.Println("failed to upload file , err", err)
	}
	_, err = fp.Write(fileBytes)
	if err != nil {
		log.Println("Failed to upload a file while writing, err ", err)
		return err
	}
	log.Println(data.FileName, " is uploaded successfully....")
	return nil
}

//DownloadFile gives the byte data in respone or user can ask in string form
func DownloadFile(data models.FileInfo) (models.FileInfo, error) {
	if validate(data, false) == false {
		return data, errors.New("Bad request or param received in file")
	}
	rawFileName := strings.Split(data.FileName, ".")[0]
	relFileName := fmt.Sprintf("%s/%s_%s_%s_%s.txt", config.Conf.StoragePath, rawFileName, data.Author, data.Category, data.Region)
	fp, err := ioutil.ReadFile(relFileName)
	if err != nil {
		log.Println("Failed to read a file")
		return data, err
	}
	data.FileData = fp
	return data, nil
}

func validate(data models.FileInfo, isUpload bool) bool {
	if data.FileName == "" || data.Author == "" || data.Region == "" || data.Category == "" {
		return false
	}

	if isUpload && data.File == nil {
		return false
	}
	return true
}
