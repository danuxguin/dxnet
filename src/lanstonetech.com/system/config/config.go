package config

import (
	"github.com/goconfig"
	"lanstonetech.com/common/logger"
)

//服务器信息
var (
	SERVER_TYPE       uint8  //服务器类型
	SERVER_INDEX      int    //服务器索引
	SERVER_NAME       string //服务器名称
	SERVER_IP         string //服务器IP
	SERVER_PORT       int    //服务器端口
	SERVER_LOCAL_IP   string //服务器内部IP，可选项
	SERVER_LOCAL_PORT int    //服务器内部PORT，可选项
	SERVER_DOMAIN     string //服务器域名，可选项
	SERVER_GROUP      string //服务器组别
	SERVER_GUID       string //服务器唯一编号，可选项
	SERVER_SIGNATURE  string //服务器签名
)

//配置信息
var ServerConfig *goconfig.ConfigFile

func init() {
	ini := "../../conf/conf.ini"
	var err error
	ServerConfig, err = goconfig.LoadConfigFile(ini)
	if err != nil {
		logger.Errorf("Err = %v", err)
		return
	}

	LoadServerInfo()
}

func LoadServerInfo() {
	// server_ip, err := ServerConfig.GetValue("", "server_ip")
	// if err != nil {
	// 	logger.Errorf("Err = %v", err)
	// 	return
	// }
	var err error
	SERVER_IP, err = ServerConfig.GetValue("Server", "IP_0")
	if err != nil {
		logger.Errorf("Err = %v", err)
		return
	}

	SERVER_PORT, err = ServerConfig.Int("Server", "PORT_0")
	if err != nil {
		logger.Errorf("Err = %v", err)
		return
	}
}
