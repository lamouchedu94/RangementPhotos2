package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"time"

	decode "github.com/lamouchedu94/ExifGO"
)

var path_error = errors.New("no such file or directory")

func main() {
	r, _ := decode.Read_img("/home/paul/Pictures/Photos/2023/2023.04/2023.04.30/3H2A6646.CR3")
	t, _ := decode.Image_date(r, ".CR3")
	_, err := final_dir("/home/paul/Pictures/TESTGO", t)
	//err := run("")

	if err != nil {
		fmt.Println(err)
	}

}

func final_dir(dest_path string, date time.Time) (string, error) {
	//Creé les dossiers si manquants et revoie le chemin avec date

	if string(dest_path[len(dest_path)-1]) != "/" {
		dest_path += "/"
	}
	if !check_dir(dest_path) {
		fmt.Print(dest_path + ": ")
		return "", path_error
	}

	y := date.Year()
	dest_path += fmt.Sprintf("%d", y)
	create_dir(dest_path)

	m_temp := int(date.Month())
	m := strconv.Itoa(m_temp)
	if len(m) < 10 {
		m = "0" + m
	}
	dest_path += fmt.Sprintf("/%d.%s", y, m)
	create_dir(dest_path)

	d := date.Day()
	dest_path += fmt.Sprintf("/%d.%v.%d", y, m, d)
	create_dir(dest_path)

	return dest_path, nil
}

func create_dir(path string) error {
	if !check_dir(path) {
		err := os.Mkdir(path, 0750)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
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

func check_dir(destination_path string) bool {
	//false, no such file directory
	_, err := os.Stat(destination_path)
	return err == nil
}
