package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := ListFilesOnDisk("C:\\Users\\Administrator\\Desktop\\go\\src\\github.com\\go-kratos\\kratos\\v2\\http")
	if err != nil {
		fmt.Println(err)
	}
}

func ListFilesOnDisk(root string) error {
	// filepath.Walk 强依赖于操作系统文件路径
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
}
