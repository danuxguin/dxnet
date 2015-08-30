package common

import (
	"io"
	"os"
)

//判断文件是否存在
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//--------------------------------------------------------
// 		读取文件内容
// 		path 		文件路径
//--------------------------------------------------------
func ReadFile(path string) ([]byte, error) {

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

//写文件到本地
func WriteFile(tarpath string, content []byte) error {

	//创建文件
	filehandle, err := os.Create(tarpath)
	if err != nil {
		return err
	}
	defer filehandle.Close()

	//写文件
	_, err = filehandle.Write(content)

	return err
}
