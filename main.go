package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	decode "github.com/lamouchedu94/ExifGO"
)

func main() {
	err := run("/home/paul/Pictures/Photos/2023/")
	if err != nil {
		fmt.Println(err)
	}
}

func run(path string) error {
	file_count := 0
	err := filepath.Walk(path, func(img string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		image, err := decode.Read_img(img)
		if err != nil {
			return err
		}
		ext := filepath.Ext(img)
		if ext != ".CR3" && ext != ".JPG" {
			return nil
		}
		date, err := decode.Image_date(image, ext)
		_ = date

		if err != nil {
			return err
		}
		fmt.Println(date)
		file_count += 1
		return nil
	})
	fmt.Println(file_count, ": Photo in directory")
	return err
}
