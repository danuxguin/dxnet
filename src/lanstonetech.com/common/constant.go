package common

const (
	_1K   = 1024
	_4K   = 4 * 1024
	_16K  = 16 * 1024
	_32K  = 32 * 1024
	_64K  = 64 * 1024
	_128K = 128 * 1024
	_256K = 256 * 1024
	_512K = 512 * 1024
	_1M   = 1024 * 1024
)

const (
	PACKET_HEAD_LEN   = 16
	PACKET_COMMON_LEN = 32
)

//最大长度定义
const (
	MAX_ACCOUNT_LEN                    = 48   //帐号
	MAX_PASSWORD_LEN                   = 256  //密码
	MAX_MD5_LEN                        = 32   //MD5
	MAX_TOKEN_LEN                      = 32   //Token
	MAX_SIGNATURE_LEN                  = 32   //签名
	MAX_DEVICENAME_LEN                 = 64   //设备名称
	MAX_IMEI_LEN                       = 64   //IMEI 设备唯一编号
	MAX_DEVICE_TOKEN_LEN               = 128  //设备Token
	MAX_COVER_LEN                      = 128  //头像
	MAX_FID_LEN                        = 128  //FID长度
	MAX_NICKNAME_LEN                   = 64   //昵称
	MAX_REALNAME_LEN                   = 64   //真名
	MAX_STATE_LEN                      = 128  //个人状态
	MAX_SPACE_LEN                      = 2048 //动态
	MAX_SPACE_REPLY_LEN                = 512  //动态回复的最大长度
	MAX_MESSAGE_LEN                    = 1024 //消息ID长度
	MAX_MESSAGE_EXT_LEN                = 1024 //消息扩展内容最大长度
	MAX_GEOHASH_LEN                    = 16   //GEOHASH最大长度
	MAX_SERVER_NAME_LEN                = 128  //服务器最大长度
	MAX_SERVER_HOST_LEN                = 128  //服务器主机最大长度，可能为ip也可能是一个域名
	MAX_PHONENUMBER_LEN                = 16   //电话号码
	MAX_LBS_STREET_LEN                 = 256  //LBS街道长度
	MAX_PID_LEN                        = 32   //PID最大长度
	MAX_SPACE_COUNT                    = 50   //单次请求最大数50个
	MAX_LOVE_FRIENDS                   = 10   //最多可以爱几个人
	MAX_UPLOAD_FRIEND_NUM              = 500  //通讯录一次上传最大数量
	MAX_SETPERMISSION_NUM              = 32   //每次设置权限最大数量
	MAX_GET_SECRET_FRIEND_GENDER_COUNT = 20   //最大获取神秘人性别人数
	MAX_ANSWER_LEN                     = 16   //真心话答案长度
	MAX_CHAT_SESSIONID_LEN             = 256  //匿名聊天SessionLen
	MAX_COLOR_LEN                      = 7    //颜色（#FF0000）
	MAX_TITLE_LEN                      = 128  //文本标题长度
	MAX_IP_LEN                         = 64   //IP最大长度
	MAX_OSVER_LEN                      = 8    //系统版本最大长度
	MAX_REMARK_LEN                     = 32   //备注最大长度
	MAX_LOVE_ATTENTION_COUNT           = 20   //每次可能喜欢我的人的数量
)
