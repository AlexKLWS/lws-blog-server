package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	mutex        sync.Mutex
	fileMetaData []models.FileMetaData
)

func AddNewFileMetaData(c echo.Context) error {
	data := models.UploadMetaDataBody{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		log.Printf("Failed processing article submit request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	mutex.Lock()
	for i := range data.MetaData {
		u := uuid.Must(uuid.NewV4())
		data.MetaData[i].ReferenceId = u.String()
	}
	fileMetaData = append(fileMetaData, data.MetaData...)
	mutex.Unlock()

	return c.JSON(http.StatusOK, data)
}

func AddNewFiles(c echo.Context) error {
	referenceId := c.FormValue("referenceId")
	log.Printf("Uploading file with REFERENCE-ID: %s\n", referenceId)
	metaData, done := getMetaData(referenceId)
	if !done {
		log.Print("No file metadata!")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed extracting a file from formdata: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	src, err := file.Open()
	if err != nil {
		log.Printf("Failed opening filestream: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer src.Close()

	createFolderIfDoesntExist(metaData.Folder)
	filePath := getFilePath(metaData, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed creating a file at path: %s\nError: %s\n", filePath, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Printf("Failed writing a file: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, getFileURL(metaData, file.Filename))
}

func getMetaData(referenceId string) (models.FileMetaData, bool) {
	var metaData models.FileMetaData
	dataIndex := -1
	mutex.Lock()
	defer mutex.Unlock()
	// Potential bottleneck, need to look into splitting the slice and searching different parts in parallel
	for i := range fileMetaData {
		if fileMetaData[i].ReferenceId == referenceId {
			dataIndex = i
			break
		}
	}
	if dataIndex == -1 {
		return models.FileMetaData{}, false
	}
	metaData = fileMetaData[dataIndex]
	fileMetaData[dataIndex] = fileMetaData[len(fileMetaData)-1]
	fileMetaData[len(fileMetaData)-1] = models.FileMetaData{}
	fileMetaData = fileMetaData[:len(fileMetaData)-1]
	return metaData, true
}

func getLocalFilePath(metaData models.FileMetaData, fileName string) string {
	filePath := ""
	if metaData.Folder != "" {
		filePath += metaData.Folder + "/"
	}
	if metaData.NewName != "" {
		fileExtension := strings.Split(fileName, ".")[1]
		filePath += metaData.NewName + "." + fileExtension
	} else {
		filePath += fileName
	}
	return filePath
}

func getFilePath(metaData models.FileMetaData, fileName string) string {
	log.Printf("Getting path for file: %s\n", fileName)
	filePath := filepath.Join(".", "assets")
	return filepath.Join(filePath, getLocalFilePath(metaData, fileName))
}

func getFileURL(metaData models.FileMetaData, fileName string) string {
	fileURL := "/assets/"
	return fileURL + getLocalFilePath(metaData, fileName)
}

func createFolderIfDoesntExist(folder string) {
	path := filepath.Join(".", "assets")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	path = filepath.Join(path, folder)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
