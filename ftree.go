package ftreedepth

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type CallbackFunc func(path string, info os.FileInfo, err error)

// Similar to filepath.WalkFunc, but with a limited depth
func WalkTree(depth int, root string, callback CallbackFunc) {
	if depth <= 0 {
		return
	}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}
	for _, file := range files {
		fp := filepath.Join(root, file.Name())
		if file.IsDir() {
			WalkTree(depth - 1, fp, callback)
		} else {
			f, err := os.Open(fp)
			defer f.Close()
			if err != nil {
				callback(fp, nil, err)
			}
			s, err := f.Stat()
			callback(fp, s, err)
		}
	}
}