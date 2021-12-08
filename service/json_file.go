package service

import (
	"encoding/json"
	"io/ioutil"
	"kopuro/view"
	"os"
)

type JsonFileService struct {
	baseFilePath string
}

func NewJsonFileService(baseFilePath string) JsonFileService {
	return JsonFileService{baseFilePath}
}

func (j JsonFileService) WriteJSONFile(request view.WriteJsonFileRequest) error {
	f, err := os.Create(j.makeFilePath(request.Filename))
	if err != nil {
		return err
	}

	defer f.Close()

	b, err := json.Marshal(request.Content)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}

func (j JsonFileService) ReadJSONFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(j.makeFilePath(filename))
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (j JsonFileService) CheckJSONFileExistence(filename string) ([]byte, error) {
	f, err := os.Open(j.makeFilePath(filename))
	defer f.Close()

	fileExists := true
	if os.IsNotExist(err) {
		fileExists = false
	} else if err != nil {
		return nil, err
	}

	b, err := json.Marshal(map[string]interface{}{ "file_exists": fileExists })
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (j JsonFileService) makeFilePath(filename string) string {
	return j.baseFilePath + "/" + filename + ".json"
}