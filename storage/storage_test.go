package storage_test

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	. "github.com/gessnerfl/awsroutes/storage"
	"github.com/stretchr/testify/require"
)

const (
	location = "../temp/storage"
)

func TestShouldCreateStorageFolderOnceAndThenUseItForUpdatingRoutes(t *testing.T) {
	executeStorageTestWithFolderCleanup(t, func(t *testing.T) {
		testData1 := []byte("test-data-1")
		testData2 := []byte("test-data-2")
		filename := "test.txt"

		sut, err := NewStorage(location)

		require.Nil(t, err)

		err = sut.Update(filename, testData1)

		require.NoError(t, err)
		require.FileExists(t, location+"/"+filename)
		data1, _ := ioutil.ReadFile(location + "/" + filename)
		require.Equal(t, testData1, data1)

		err = sut.Update(filename, testData2)

		require.Nil(t, err)
		require.FileExists(t, location+"/"+filename)
		data2, _ := ioutil.ReadFile(location + "/" + filename)
		require.Equal(t, testData2, data2)
	})
}

func TestShouldWriteAndReadFile(t *testing.T) {
	executeStorageTestWithFolderCleanup(t, func(t *testing.T) {
		input := []byte("test-data")
		filename := "test.txt"

		sut, err := NewStorage(location)

		require.Nil(t, err)

		err = sut.Update(filename, input)

		require.NoError(t, err)
		require.FileExists(t, location+"/"+filename)

		result, err := sut.Read(filename)

		require.NoError(t, err)
		require.Equal(t, input, result)
	})
}

func TestShouldThrowErrorWhenReadingFileAndFileDoesNotExist(t *testing.T) {
	executeStorageTestWithFolderCleanup(t, func(t *testing.T) {
		filename := "test.txt"

		sut, err := NewStorage(location)

		require.Nil(t, err)

		_, err = sut.Read(filename)

		require.Error(t, err)
	})
}

func executeStorageTestWithFolderCleanup(t *testing.T, testFunc func(*testing.T)) {
	pc, _, _, _ := runtime.Caller(1)
	fullname := runtime.FuncForPC(pc).Name()
	name := strings.Split(fullname, "storage_test.")[1]
	t.Cleanup(func() {
		os.RemoveAll(location)
	})
	t.Run(name, testFunc)
}
