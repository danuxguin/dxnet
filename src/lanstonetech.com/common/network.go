package common

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func UrlEncode(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func UrlDecode(s string) string {
	url, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return s
	} else {
		return string(url)
	}
}

//--------------------------------------------------------
// 通过HTTP获取返回内容
//--------------------------------------------------------
func GetHttpResp(addr string) ([]byte, error) {

	//获取请求
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//读取返回数据
	return ioutil.ReadAll(resp.Body)
}

func GetHttpsResp(addr string) ([]byte, error) {

	//获取请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//读取返回数据
	return ioutil.ReadAll(resp.Body)

}

//--------------------------------------------------------
// 		获取上传地址
// 		ip 		文件服务器IP
// 		port 	文件服务器端口
//--------------------------------------------------------
func GetUploadUrl(ip string, port uint16, dir string) (string, string, error) {

	//获取请求
	u := fmt.Sprintf("http://%s:%d/dir/assign", ip, port)
	if len(dir) > 0 {
		u = u + fmt.Sprintf("?dc=%s", dir)
	}

	resp, err := http.Get(u)
	if err != nil {
		return "", "", fmt.Errorf("[GetUploadUrl] step=1 err=%v", err)
	}
	defer resp.Body.Close()

	//读取返回数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("[GetUploadUrl] step=2 err=%v", err)
	}

	//weeds文件请求结果结构
	type weedResp struct {
		Fid       string `json:"fid"`
		Url       string `json:"url"`
		PublicUrl string `json:"publicUrl"`
		Count     int    `json:"count"`
	}

	//解码返回数据
	var res weedResp
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return "", "", fmt.Errorf("[GetUploadUrl] step=3 body=%v err=%v", string(body), err)
	}

	//组装请求URL
	return "http://" + res.Url + "/" + res.Fid, res.Fid, nil
}

func UploadFileToFS(ip string, port uint16, filename string, dir ...string) (string, string, error) {

	data, err := ReadFile(filename)
	if err != nil {
		return "", "", err
	}

	url, fid, err := UploadFileToWeeds(ip, port, data, dir...)
	return url, fid, err
}

/******************************************************************************
 @brief
	上传头像文件
 @author
 	chenzhiguo
 @param
	ip  				IP地址
	port				服务器端口
	content_s			头像缩略图
	content				头像原图
 @return
 	string				上传路径
 	string				上传后文件id
 	error				如果返回nil，表示上传成功，否则，表示上传失败
 @history
 	2015-05-28 16:19 	chenzhiguo		创建
*******************************************************************************/
func UploadCoverToFS(ip string, port uint16, content_s, content []byte) (string, string, error) {

	uploadurl, fid, err := GetUploadUrl(ip, port, "")
	if err != nil {
		return uploadurl, "", fmt.Errorf("[UploadFileToWeeds] Step=1 err=%v", err)
	}

	if err = UploadCoverToFSByFID(content_s, uploadurl); err != nil {
		return "", "", err
	}

	if err = UploadCoverToFSByFID(content, fmt.Sprintf("%s_1", uploadurl)); err != nil {
		return "", "", err
	}

	return uploadurl, fid, err
}

/******************************************************************************
 @brief
 	上传文件
 @author
 	chenzhiguo
 @param
	content  			需要上传的文件
	uploadurl			文件上传路径
 @return
 	error				如果返回nil，表示上传成功成功，否则，表示上传失败
 @history
 	2015-05-28 16:19 	chenzhiguo		创建
*******************************************************************************/
func UploadCoverToFSByFID(content []byte, uploadurl string) error {

	//Upload
	body := &bytes.Buffer{}
	body_writer := multipart.NewWriter(body)

	//add file segment
	file_writer, err := body_writer.CreateFormFile("uploadfile", "uploadfile")
	if err != nil {
		return fmt.Errorf("[UploadCoverToFSByFID] Step=2 err=%v", err)
	}

	//设置文件内容
	_, err = file_writer.Write(content)
	if err != nil {
		return fmt.Errorf("[UploadCoverToFSByFID] Step=3 err=%v", err)
	}
	body_writer.Close()

	// 直接使用 NewRequest
	req, err := http.NewRequest("POST", uploadurl, body)
	if err != nil {
		return fmt.Errorf("[UploadCoverToFSByFID] Step=4 err=%v", err)
	}

	//类型
	content_type := body_writer.FormDataContentType()
	req.Header.Set("Content-Type", content_type)

	//Do upload
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[UploadCoverToFSByFID] Step=5 err=%v", err)
	}

	rsp.Body.Close()

	//OK
	return nil

}

//--------------------------------------------------------
// 		上传文件
// 		ip 		文件服务器IP
// 		port 	文件服务器端口
// 		content 要上传的文件的内容
//--------------------------------------------------------
func UploadFileToWeeds(ip string, port uint16, content []byte, dir ...string) (string, string, error) {

	dc := ""
	if len(dir) > 0 {
		dc = dir[0]
	}

	//Get Upload Url
	uploadurl, fid, err := GetUploadUrl(ip, port, dc)
	if err != nil {
		return uploadurl, "", fmt.Errorf("[UploadFileToWeeds] Step=1 err=%v", err)
	}

	//Upload
	body := &bytes.Buffer{}
	body_writer := multipart.NewWriter(body)

	//add file segment
	file_writer, err := body_writer.CreateFormFile("uploadfile", "uploadfile")
	if err != nil {
		return "", "", fmt.Errorf("[UploadFileToWeeds] Step=2 err=%v", err)
	}

	//设置文件内容
	_, err = file_writer.Write(content)
	if err != nil {
		return "", "", fmt.Errorf("[UploadFileToWeeds] Step=3 err=%v", err)
	}
	body_writer.Close()

	// 直接使用 NewRequest
	req, err := http.NewRequest("POST", uploadurl, body)
	if err != nil {
		return "", "", fmt.Errorf("[UploadFileToWeeds] Step=4 err=%v", err)
	}

	//类型
	content_type := body_writer.FormDataContentType()
	req.Header.Set("Content-Type", content_type)

	//Do upload
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("[UploadFileToWeeds] Step=5 err=%v", err)
	}

	rsp.Body.Close()

	//OK
	return uploadurl, fid, nil
}
