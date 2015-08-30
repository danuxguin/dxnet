package ziparchive

import (
	"archive/zip"
	"bytes"
	//"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Writer struct {
	dir   string
	files map[string]string
}

func NewWriter() *Writer {
	var w Writer
	w.files = make(map[string]string, 0)
	return &w
}

func (this *Writer) AddFile(filepath, filepathinzip string) {
	this.files[filepath] = filepathinzip
}

func (this *Writer) AddDir(dir string) error {
	this.dir = dir
	if err := filepath.Walk(dir, this.foreachFile); err != nil {
		this.dir = ""
		return err
	}

	this.dir = ""
	return nil
}

func (this *Writer) Save(zipfile string) error {
	return this.zipFiles(this.files, zipfile)
}

func (this *Writer) foreachFile(path string, info os.FileInfo, err error) error {

	if info == nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if strings.Contains(path, ".DS_Store") {
		return nil
	}

	index := strings.IndexAny(path, "/\\")
	if index == -1 {
		return nil
	}

	fid := path[index+1:]
	filepath := path
	this.files[filepath] = fid
	return nil
}

//--------------------------------------------------------
// 		压缩几个文件
// 		filenames 		key:文件路径 val:在zip中显示的文件名
// 		tarfile 		目标文件路径
//--------------------------------------------------------
func (this *Writer) zipFiles(filenames map[string]string, tarfile string) error {

	// Add some files to the archive.
	type STFileCon struct {
		Name string
		Body []byte
	}
	var files = make([]STFileCon, 0)

	//读取文件列表的内容
	var err error
	for srcfile, tarname := range filenames {
		var filecon STFileCon
		filecon.Name = tarname
		filecon.Body, err = this.readFile(srcfile)

		if err != nil {
			return err
		}

		files = append(files, filecon)
	}

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = f.Write(file.Body)
		if err != nil {
			return err
		}
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return err
	}

	//创建文件
	filehandle, err := os.Create(tarfile)
	if err != nil {
		return err
	}
	defer filehandle.Close()

	//写文件
	_, err = filehandle.Write(buf.Bytes())

	return err
}

func (this *Writer) readFile(path string) ([]byte, error) {

	//打开文件
	file, err := os.Open(path) // For read access.
	if err != nil {
		return nil, err
	}

	//关闭文件
	defer file.Close()

	//读取内容
	rtndata := make([]byte, 0)

	for {
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil && err != io.EOF {
			return nil, err
		}

		rtndata = append(rtndata, data[:count]...)
		if err == io.EOF {
			return rtndata, nil
		}
	}

	return rtndata, nil
}
