package file

import (
	"archive/zip"
	"os"
	"io"
	"runtime"
	"path/filepath"
	"strings"
	"gostudy/src/common"
)


//压缩文件
func CreateZip(filename string, files []string) error {

	file, err := os.Create(filename);
	if (err != nil) {
		return err
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	for _, name := range files {
		if err := writeFileToZip(zipWriter, name); err != nil {
			return err
		}
	}
	return nil
}

func writeFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if (err != nil) {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if (err != nil) {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if (err != nil) {
		return err
	}
	header.Name = sanitizedName(filename)
	writer, err := zipWriter.CreateHeader(header)
	if (err != nil) {
		return err
	}
	_, err = io.Copy(writer, file)
	common.Out("write " + filename + " ...")
	return err
}
func sanitizedName(filename string) string {

	if len(filename) > 1 && filename[1] == ':' && runtime.GOOS == "windows" {
		filename = filename[2:]
	}
	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")
	return strings.Replace(filename, "../", "", -1)
}
