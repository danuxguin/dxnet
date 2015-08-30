package archive

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const FHE = 0x874 //文件信息头结束标记

/******************************************************************************
 @brief
 	文件信息类
 @author
 	chenzhiguo
 @history
 	2015-05-16_09:23 	chenzhiguo		创建
*******************************************************************************/
type FileInfo struct {
	FID           string //文件ID
	MD5           string //md5
	Size          int64  //大小
	IsDelete      uint8  //是否已经删除
	IsReplication uint8  //是否是复制过来的
	TimeStamp     int64  //时间戳
	offset        int64  //私有偏移
}

/******************************************************************************
 @brief
 	读取文件信息头
 @author
 	chenzhiguo
 @param
	fp					文件名柄
	offset				偏移位置，从这个位置开始向后读
 @return
 	error				返回nil，读取成功，否则，表示失败原因
 @history
 	2015-05-16_09:23 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) Read(fp *os.File, offset int64) error {

	//读gob长度
	buf := make([]byte, 2)
	pos := offset
	n, err := fp.ReadAt(buf, pos)
	if err != nil || n != 2 {
		return io.EOF
	}
	pos += 2

	fil := binary.BigEndian.Uint16(buf)

	//读取gob buf
	buf = make([]byte, fil)
	n, err = fp.ReadAt(buf, pos)
	if err != nil {
		return err
	}
	pos += int64(n)

	if n != int(fil) {
		return fmt.Errorf("File.Read err! n=%d fil=%d", n, fil)
	}

	//gob解析
	if err := this.DecodeFileInfo(buf, this); err != nil {
		return err
	}

	//校验尾标记
	buf = make([]byte, 2)
	n, err = fp.ReadAt(buf, pos)
	if err != nil || n < 2 {
		return io.EOF
	}

	fhe := binary.BigEndian.Uint16(buf)
	if fhe != FHE {
		return io.EOF
	}

	pos += 2

	//计算偏移
	this.offset = pos

	//返回成功
	return nil
}

/******************************************************************************
 @brief
 	写入文件信息头
 @author
 	chenzhiguo
 @param
	fp					文件名柄
	offset				偏移位置，从这个位置开始写
 @return
 	int					返回写入到文件中的长度
 	error				返回nil，写入成功，否则，表示失败原因
 @history
 	2015-05-16_09:23 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) Write(fp *os.File, offset int64) (int, error) {

	buf, err := this.EncodeFileInfo(this)
	if err != nil {
		return 0, nil
	}

	//写入gob长度
	l := uint16(len(buf))
	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, l)
	if _, err := fp.WriteAt(temp, offset); err != nil {
		return 0, err
	}

	//写入gob
	n, err := fp.WriteAt(buf, offset+2)
	if err != nil {
		return 0, err
	}

	if n != int(l) {
		return 0, err
	}

	//写入成功标记
	temp = make([]byte, 2)
	binary.BigEndian.PutUint16(temp, FHE)
	n, err = fp.WriteAt(temp, int64(offset)+2+int64(l))
	if err != nil {
		return 0, err
	}

	if n != 2 {
		return 0, err
	}

	return (2 + int(l) + 2), nil
}

/******************************************************************************
 @brief
 	对文件信息进行编码，以生成一个可以方便写入到文件中的buf，并起到加密压缩的作用
 @author
 	chenzhiguo
 @param
	source				要编码对象
 @return
 	[]byte				返回编码结果
 	error				返回nil表示成功，否则表示失败的原因
 @history
 	2015-05-16_09:30 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) EncodeFileInfo(source interface{}) ([]byte, error) {

	b, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	return this.ZlibCompress(b)
}

/******************************************************************************
 @brief
 	对文件信息进行编码，以生成一个可以方便写入到文件中的buf，并起到加密压缩的作用
 @author
 	chenzhiguo
 @param
	data				待解码的数据源
	dest				解码后存放的地方
 @return
 	error				返回nil表示成功，否则表示失败的原因
 @history
 	2015-05-16_09:30 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) DecodeFileInfo(data []byte, dest interface{}) error {

	b, err := this.ZlibDecompress(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, dest)
}

/******************************************************************************
 @brief
 	zlib压缩
 @author
 	chenzhiguo
 @param
	buf					待压缩的数据
 @return
 	[]byte				压缩后的数据
 	error				返回nil表示压缩成功，否则表示压缩失败
 @history
 	2015-05-16_09:30 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) ZlibCompress(buf []byte) ([]byte, error) {

	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
	if err != nil {
		return nil, err
	}

	if n, err := w.Write(buf); err != nil || n != len(buf) {
		w.Close()
		return buf, fmt.Errorf("ZlibCompress failed! err=%v n=%d buflen=%d", err, n, len(buf))
	}

	w.Close()

	return b.Bytes(), nil
}

/******************************************************************************
 @brief
 	zlib解压
 @author
 	chenzhiguo
 @param
	buf					待解压的数据
 @return
 	[]byte				解压后的数据
 	error				返回nil表示解压成功，否则表示解压失败
 @history
 	2015-05-16_09:30 	chenzhiguo		创建
*******************************************************************************/
func (this *FileInfo) ZlibDecompress(buf []byte) ([]byte, error) {

	r, err := zlib.NewReader(bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	outbuf := make([]byte, 0)
	temp := make([]byte, 4*1024)
	for {

		n, err := r.Read(temp[:])
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return nil, err
		}

		outbuf = append(outbuf, temp[:n]...)
	}

	return outbuf, nil
}
