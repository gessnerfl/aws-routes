package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//Storage interface definition of a storage for persisting files
type Storage interface {
	//Update updates the file of the given filename in the storage with the given data
	Update(filename string, data []byte) error
	//Read reads the file of the given filename from the storage
	Read(filename string) ([]byte, error)
}

const (
	tempDir        = ".awsroutes"
	userOnlyAccess = os.FileMode(0700)
)

//NewStorageAtUserHome creates a new storage for files at the user home directory
func NewStorageAtUserHome() (Storage, error) {
	homeDir, err := GetUserHome()
	if err != nil {
		return nil, err
	}
	return &baseStorage{storageDir: homeDir + "/" + tempDir}, nil
}

//NewStorage creates a new storage for file at the given directory
func NewStorage(storageDir string) (Storage, error) {
	dir, err := filepath.Abs(storageDir)
	return &baseStorage{storageDir: dir}, err
}

type baseStorage struct {
	storageDir string
}

//Update interface implementation of Storage.Update
func (s *baseStorage) Update(filename string, data []byte) error {
	err := s.prepareForUpdate(filename)
	if err != nil {
		return err
	}

	file := s.getFilePath(filename)
	return ioutil.WriteFile(file, data, userOnlyAccess)
}

func (s *baseStorage) prepareForUpdate(filename string) error {
	err := s.ensureStorageDirExists()
	if err != nil {
		return err
	}
	return s.deletePreviousFile(filename)
}

func (s *baseStorage) ensureStorageDirExists() error {
	return os.MkdirAll(s.storageDir, userOnlyAccess)
}

func (s *baseStorage) deletePreviousFile(filename string) error {
	filePath := s.getFilePath(filename)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return os.Remove(filePath)
}

//Read interface implementation of Storage.Read
func (s *baseStorage) Read(filename string) ([]byte, error) {
	file := s.getFilePath(filename)
	return ioutil.ReadFile(file)
}

func (s *baseStorage) getFilePath(filename string) string {
	return s.storageDir + "/" + filename
}
