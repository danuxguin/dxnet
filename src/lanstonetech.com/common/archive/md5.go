package archive

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/******************************************************************************
 @brief
 	生成MD5，支持多字符串，每个字符串之间以"-"相连
 		例如：
 			md5 := MakeMD5("abc","cde","fgh") 在内部会拼成 abc-cde-fgh 然后对这个串进行md5计算
 @author
 	chenzhiguo
 @param
	arg					参数
 @return
 	string				返回md5
 @history
 	2015-05-16_09:47 	chenzhiguo		创建
*******************************************************************************/
func MakeMD5(arg ...string) string {
	info := strings.Join(arg, "-")
	m := md5.New()
	m.Write([]byte(info))
	return hex.EncodeToString(m.Sum(nil))
}

/******************************************************************************
 @brief
 	对一个内存数据进行MD5数据校验计算
 @author
 	chenzhiguo
 @param
	arg					要计算md5码的数据块
 @return
 	string				返回md5
 @history
 	2015-05-16_09:47 	chenzhiguo		创建
*******************************************************************************/
func MakeByteMd5(arg []byte) string {
	m := md5.New()
	m.Write(arg)
	return hex.EncodeToString(m.Sum(nil))
}
