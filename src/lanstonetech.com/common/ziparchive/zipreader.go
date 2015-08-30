package ziparchive

import (
	"archive/zip"
	"fmt"
	"io"
)

type FileHandler func(zr *Reader, name string, zf *zip.File)

type Reader struct {
	r *zip.ReadCloser
	l map[string]*zip.File
}

func (this *Reader) Open(archive string) error {
	this.Close()

	r, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	l := make(map[string]*zip.File, 0)
	for _, f := range r.File {
		l[f.Name] = f
	}

	this.r = r
	this.l = l
	return nil
}

func (this *Reader) Close() {
	if this.r != nil {
		this.r.Close()
		this.r = nil
		this.l = nil
	}
}

func (this *Reader) ExtractFile(filename string) ([]byte, error) {

	if this.r == nil {
		return nil, fmt.Errorf("Reader.ExtractFile failed! Reader invalid!! file=%s", filename)
	}

	//查找文件是否存在
	f, ok := this.l[filename]
	if !ok {
		return nil, fmt.Errorf("Reader.ExtractFile failed! file not exists!! file=%s", filename)
	}

	//解压文件
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	if f.UncompressedSize64 == 0 {
		return []byte{}, nil
	}

	out := make([]byte, f.UncompressedSize64)
	n, err := io.ReadFull(rc, out)
	if err != nil {
		return nil, err
	}

	//检查长度
	if len(out) != n {
		return nil, err
	}

	//返回结果
	return out, nil

}

func (this *Reader) ForEach(handler FileHandler) {

	for k, v := range this.l {
		handler(this, k, v)
	}
}
