package app

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type localFile struct {
	Path string
}

type StorageService struct {
	storage string
}

func NewStorageService(storage string) *StorageService {
	return &StorageService{storage: strings.Trim(storage, string(os.PathSeparator))}
}

func (s *StorageService) SaveAs(name string, file multipart.File) (*localFile, error) {
	path := s.GetBasePath() + strings.TrimLeft(name, string(os.PathSeparator))

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &localFile{Path: path}, nil
}

func (s *StorageService) SaveAsTmp(name string, file multipart.File) (*localFile, error) {
	return s.SaveAs(s.GetTmpPath()+strings.TrimLeft(name, string(os.PathSeparator)), file)
}

func (s *StorageService) MoveToPermanent(tmpPath string) (*localFile, error) {
	n := s.GetBasePath() + path.Base(tmpPath)

	err := s.Move(tmpPath, n)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &localFile{Path: n}, nil
}

func (s *StorageService) Move(from, to string) error {
	err := os.Rename(from, to)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *StorageService) GetBasePath() string {
	fmt.Println(s.storage + string(os.PathSeparator))
	return s.storage + string(os.PathSeparator)
}

func (s *StorageService) GetTmpPath() string {
	return "tmp" + string(os.PathSeparator)
}
