package copyf

import (
	"errors"
	"os"
	"strings"
)

var AlreadyExist = errors.New("file already exist")

func Copy_pictures(src_path string, dest_path string) error {

	data, err := os.ReadFile(src_path)
	if err != nil {
		return err
	}

	img_name := Get_image_name(src_path)
	fi, _ := os.Stat(dest_path + "/" + img_name)

	if fi != nil {
		return AlreadyExist
	}

	dest_path = dest_path + "/" + img_name
	dst, err := os.Create(dest_path)
	if err != nil {
		return err
	}
	_ = dst

	defer dst.Close()
	os.WriteFile(dest_path, data, 0750)

	return nil
}

func Get_image_name(src_path string) string {
	tab := strings.Split(src_path, "/")
	return tab[len(tab)-1]
}
