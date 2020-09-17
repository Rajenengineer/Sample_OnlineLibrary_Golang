package v1

import (
	"io/ioutil"
	"log"
	"online-library/config"
	"online-library/models"
	"strings"
)

//GetFileInfo fetch the files from system basis on their search param
func GetFileInfo(fileName string) ([]models.FileInfo, error) {
	var resp []models.FileInfo
	fileNameTmp := strings.Split(fileName, ".")[0]
	files, err := ioutil.ReadDir(config.Conf.StoragePath)
	if err != nil {
		log.Println("Failed to get files info, err", err)
		return resp, err
	}
	for _, f := range files {
		fn := strings.Split(f.Name(), ".")[0]
		record := strings.Split(fn, "_")
		if record[0] == fileNameTmp {
			data := models.FileInfo{
				FileName: fileName,
				Author:   record[1],
				Category: record[2],
				Region:   record[3],
			}
			resp = append(resp, data)
		}

	}
	return resp, nil
}
