package models

type UploadMetaDataBody struct {
	MetaData []FileMetaData `json:"metaData" xml:"metaData"`
}

type FileMetaData struct {
	Id      string `json:"id" xml:"id"`
	NewName string `json:"newName" xml:"newName"`
	Folder  string `json:"folder" xml:"folder"`
}
