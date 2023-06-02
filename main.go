package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"time"

	decode "github.com/lamouchedu94/ExifGO"
)

func main() {
	r, _ := decode.Read_img("/home/paul/Pictures/Photos/2023/2023.04/2023.04.30/3H2A6646.CR3")
	t, _ := decode.Image_date(r, ".CR3")
	dir := final_dir("/home/paul/", t)
	//err := run("")
	test := check_dir(dir, time.Time{})
	fmt.Println(test)
	/*
		if err != nil {
			fmt.Println(err)
		}
	*/
}

func final_dir(dest_path string, date time.Time) string {
	y := date.Year()
	dest_path += strconv.Itoa(y)
	m := int(date.Month())
	dest_path += "/" + strconv.Itoa(m)
	d := date.Day()
	dest_path += "/" + strconv.Itoa(d)
	return dest_path
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

func check_dir(destination_path string, date time.Time) bool {
	//false, no such file directory
	_, err := os.Stat(destination_path)
	return err == nil
}
