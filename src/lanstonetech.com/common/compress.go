package common

import (
	"bytes"
	"compress/zlib"
	"io"
)

func ZlibCompress(buf []byte) []byte {

	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
	if err != nil {
		return buf
	}

	if n, err := w.Write(buf); err != nil || n != len(buf) {
		w.Close()
		return buf
	}

	w.Close()

	return b.Bytes()
}

func ZlibDecompress(buf []byte) []byte {

	r, err := zlib.NewReader(bytes.NewBuffer(buf))

	if err != nil {
		return buf
	}
	defer r.Close()

	outbuf := make([]byte, 0)
	temp := make([]byte, 4*1024)
	for {

		n, err := r.Read(temp[:])
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return buf
		}

		outbuf = append(outbuf, temp[:n]...)
	}

	return outbuf
}
