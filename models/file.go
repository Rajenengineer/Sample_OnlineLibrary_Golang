package models

import "mime/multipart"

//FileInfo stores file data
type FileInfo struct {
	FileName string `json:"fileName"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Region   string `json:"region"`
	FileData []byte `json:"fileData"`
	File     multipart.File
}
