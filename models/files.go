package models

type UploadMetaDataBody struct {
	MetaData []FileMetaData `json:"metaData" xml:"metaData"`
}

type FileMetaData struct {
	Id          string `json:"id" xml:"id"`
	ReferenceId string `json:"referenceId" xml:"referenceId"`
	NewName     string `json:"newName" xml:"newName"`
	Folder      string `json:"folder" xml:"folder"`
}
