package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	copyf "main/copy"
	"os"
	"path/filepath"
	"strconv"
	"time"

	decode "github.com/lamouchedu94/ExifGO"
)

var path_error = errors.New("no such file or directory")

type Settings struct {
	Src     string
	Dest    string
	Verbose bool
}

func main() {
	s := args()

	//src := "/home/paul/Pictures/Photos/"
	//dest := "/home/paul/Pictures/TESTGO"
	/*
		_, err := final_dir(dest, t)

		if err != nil {
			fmt.Println(err)
		}
	*/

	err := run(s.Src, s.Dest, s.Verbose)

	if err != nil {
		fmt.Println(err)
	}

}

func run(src_path string, dest_path string, verb bool) error {
	file_count := 0
	err := filepath.Walk(src_path, func(img string, info fs.FileInfo, err error) error {
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
			if !check_dir(dest_path + "0000") {
				os.Mkdir("0000", 0750)
			}
			copyf.Copy_pictures(img, dest_path+"0000")
			return nil
		}
		date, err := decode.Image_date(image, ext)
		_ = date
		if err != nil {
			return err
		}

		final_path, err := final_dir(dest_path, date)
		if err != nil {
			return err
		}
		copyf.Copy_pictures(img, final_path)
		if verb {
			name := copyf.Get_image_name(img)
			fmt.Printf("%s -> %s/%s\n", img, dest_path, name)
		}

		file_count += 1
		return nil
	})
	fmt.Println(file_count, ": Photo in directory")
	return err
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
	if len(m) < 2 {
		m = "0" + m
	}
	dest_path += fmt.Sprintf("/%d.%s", y, m)
	create_dir(dest_path)

	d_temp := date.Day()
	d := strconv.Itoa(d_temp)
	if len(d) < 2 {
		d = "0" + d
	}

	dest_path += fmt.Sprintf("/%d.%s.%s", y, m, d)
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

func check_dir(destination_path string) bool {
	//false, no such file directory
	_, err := os.Stat(destination_path)
	return err == nil
}

func args() Settings {
	s := Settings{}

	flag.StringVar(&s.Src, "s", "", "source dir")
	flag.StringVar(&s.Dest, "d", "", "destination dir")
	flag.BoolVar(&s.Verbose, "v", false, "Display operations")
	flag.Parse()
	fmt.Println(s.Src, s.Dest, s.Verbose)
	return s
}
