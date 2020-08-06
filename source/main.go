package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type sourceStr struct {
	name string
}

func main() {
	var (
		rootSource  string
		rootTarget  string
		filesSource []string
		filesTarget []string
		err         error
	)

	rootSource = "folder1/"
	rootTarget = "../target/folder1"

	// ioutil.ReadDIr get file path from source
	filesSource, err = IOReadir(rootSource)
	if err != nil {
		panic(err)
	}
	// ioutil.ReadDIr get file path from target
	filesTarget, err = IOReadir(rootTarget)
	if err != nil {
		panic(err)
	}

	diffStr := compareFolder(filesSource, filesTarget)
	for _, diffValue := range diffStr {
		fmt.Println(diffValue, "DELETED")
	}
}

/*The path/filepath package provides a handy way to scan all the files in a directory,
it will automatically scan each sub-directories in the directory.*/
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ioutil.ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename.
func IOReadir(root string) ([]string, error) {
	var files []string
	filesInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range filesInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

/*Readdir reads the contents of the directory associated with file and returns a slice of up to n FileInfo values,
as would be returned by Lstat, in directory order. Subsequent calls on the same file will yield further FileInfos.*/
func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func compareFolder(source []string, target []string) []string {
	var diff []string

	m := map[string]int{}
	for _, fileSrc := range source {
		m[fileSrc] = 1
	}
	for _, fileTrg := range target {
		m[fileTrg] = m[fileTrg] + 1
	}
	for mKey, mVal := range m {
		if mVal == 1 {
			diff = append(diff, mKey)
		}
	}
	return diff
}
