package csvfile

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

/******************************************************************************
 @brief
 	打开CSV文件
 @author
 	chenzhiguo
 @history
 	2015-05-16_10:04 	chenzhiguo		创建
*******************************************************************************/
type CSVFile struct {
	reader   *csv.Reader
	col      []string
	colcount int
	first    bool
	records  int
}

/******************************************************************************
 @brief
 	打开CSV文件
 @author
 	chenzhiguo
 @param
	file				文件名称
	colcount			预期要读取的文件列数，如果目标CSV文件列数少于他，文件读取将会失败
 @return
 	*CSVFile			返回CSVFile文件指针
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:04 	chenzhiguo		创建
*******************************************************************************/
func Open(file string, colcount int) (*CSVFile, error) {
	cntb, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(cntb)))
	reader.Comma = '\t'
	reader.Comment = '#'

	return &CSVFile{reader, []string{}, colcount, true, 0}, nil
}

/******************************************************************************
 @brief
 	从内存中打开CSV文件
 @author
 	chenzhiguo
 @param
	data				CSV文件内容
	colcount			预期要读取的文件列数，如果目标CSV文件列数少于他，文件读取将会失败
 @return
 	*CSVFile			返回CSVFile文件指针
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:04 	chenzhiguo		创建
*******************************************************************************/
func OpenBuffer(data []byte, colcount int) (*CSVFile, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.Comma = '\t'
	reader.Comment = '#'

	return &CSVFile{reader, []string{}, colcount, true, 0}, nil
}

/******************************************************************************
 @brief
 	读取第一行信息，一般CSV第一行信息都是注释性质的，读取程序会选择过滤掉，但是会通过这个验证CSV的列数是否满足读取要求
 @author
 	chenzhiguo
 @param
	-
 @return
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:04 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) fetchFirst() error {
	if !this.first {
		return nil
	}

	this.first = false

	col, err := this.reader.Read()
	if err == io.EOF {
		return io.EOF
	}

	if len(col) < this.colcount {
		return fmt.Errorf("colume not match!! records=%d colcount=%d hopecount=%d", this.records, len(col), this.colcount)
	}

	return nil
}

/******************************************************************************
 @brief
 	读取CSV文件数据，在这过程中会对每一行的列数进行数量校验，如果不能满足读取解析，将会导致失败，此接口一次只获取一行内容
 @author
 	chenzhiguo
 @param
	-
 @return
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:09 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) Fetch() error {

	//过滤第一行
	err := this.fetchFirst()
	if err != nil {
		return err
	}

	//开始正常读取
	col, err := this.reader.Read()
	if err == io.EOF {
		this.col = []string{}
		return io.EOF
	}

	if len(col) < this.colcount {
		this.col = []string{}
		return fmt.Errorf("colume not match!! records=%d record col = %#v hopecount = %d", this.records, col, this.colcount)
	}

	this.col = col
	this.records = this.records + 1
	return nil
}

/******************************************************************************
 @brief
 	读取指定列，并转换成string
 @author
 	chenzhiguo
 @param
	index				列索引
 @return
 	string				返回结果
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) GetString(index int) (string, error) {

	if index >= len(this.col) || index < 0 {
		return "", errors.New(fmt.Sprintln("OUT OF RANGE!!! index=", index, "From=0 To=", len(this.col), "Context=", this.col))
	}

	return this.col[index], nil
}

/******************************************************************************
 @brief
 	读取指定列，并转换成int
 @author
 	chenzhiguo
 @param
	index				列索引
 @return
 	int					返回结果
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) GetInt(index int) (int, error) {

	val, err := this.GetString(index)
	if err != nil {
		return -1, err
	}

	ret, err := strconv.Atoi(val)
	return ret, err
}

/******************************************************************************
 @brief
 	读取指定列，并转换成float
 @author
 	chenzhiguo
 @param
	index				列索引
 @return
 	float64				返回结果
 	error				返回nil表示成功，否则表示失败
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) GetFloat(index int) (float64, error) {

	val, err := this.GetString(index)
	if err != nil {
		return -1, err
	}

	ret, err := strconv.ParseFloat(val, 64)
	return ret, err
}

/******************************************************************************
 @brief
 	读取指定列，并转换成string，此接口不会报错，但可以指定出错以后，返回一个默认值替换
 @author
 	chenzhiguo
 @param
	index				列索引
	def					默认值，可以不设置
 @return
 	string				返回结果
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) MustString(index int, def ...string) string {
	v, err := this.GetString(index)
	if err != nil {
		if len(def) > 1 {
			return def[0]
		} else {
			return ""
		}
	}

	return v
}

/******************************************************************************
 @brief
 	读取指定列，并转换成int，此接口不会报错，但可以指定出错以后，返回一个默认值替换
 @author
 	chenzhiguo
 @param
	index				列索引
	def					默认值，可以不设置
 @return
 	int					返回结果
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) MustInt(index int, def ...int) int {
	v, err := this.GetInt(index)
	if err != nil {
		if len(def) > 1 {
			return def[0]
		} else {
			return 0
		}
	}

	return v
}

/******************************************************************************
 @brief
 	读取指定列，并转换成float，此接口不会报错，但可以指定出错以后，返回一个默认值替换
 @author
 	chenzhiguo
 @param
	index				列索引
	def					默认值，可以不设置
 @return
 	float				返回结果
 @history
 	2015-05-16_10:22 	chenzhiguo		创建
*******************************************************************************/
func (this *CSVFile) MustFloat(index int, def ...float64) float64 {
	v, err := this.GetFloat(index)
	if err != nil {
		if len(def) > 1 {
			return def[0]
		} else {
			return 0
		}
	}

	return v
}
