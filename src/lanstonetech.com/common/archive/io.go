package archive

import (
	"io"
	"os"
)

/******************************************************************************
 @brief
 	判断文件是否存在
 @author
 	chenzhiguo
 @param
	filename			文件路径
 @return
 	bool				返回true表示存在，否则不存在
 @history
 	2015-05-16_09:47 	chenzhiguo		创建
*******************************************************************************/
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

/******************************************************************************
 @brief
 	读取文件到内存中
 @author
 	chenzhiguo
 @param
	path				文件路径
 @return
 	[]byte				文件内容
 	error				返回nil表示读取成功，否则表示失败原因
 @history
 	2015-05-16_09:47 	chenzhiguo		创建
*******************************************************************************/
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

/******************************************************************************
 @brief
 	写入数据到文件中
 @author
 	chenzhiguo
 @param
	tarpath				文件路径
	content				要保存的数据
 @return
 	error				返回nil表示写入成功，否则表示失败原因
 @history
 	2015-05-16_09:47 	chenzhiguo		创建
*******************************************************************************/
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
