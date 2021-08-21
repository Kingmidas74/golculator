package main

import (
	"io/fs"
	"io/ioutil"
)

type FileReader struct {

}

func(this *FileReader) GetListOfFiles(dir string) ([]fs.FileInfo,error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files,nil
}

func(this *FileReader) ReadFileToString(path string) (string,error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content),nil
}
