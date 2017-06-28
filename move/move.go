package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

const dirPath = "../../userpages/content/articles"

func main() {
	count := 0
	// walk all files in directory
	filepath.Walk(dirPath, func(fpath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.Contains(info.Name(), "pali-chanting") {
				newpath := path.Join("../content/articles/", path.Base(fpath))
				newpath = strings.Replace(newpath, "-pali-chanting", "", -1)
				println(newpath)
				err := os.Rename(fpath, newpath)
				if err != nil {
					println(newpath)
					println(err)
				}
				count++
			}
		}
		return nil
	})
	println(count)
}
