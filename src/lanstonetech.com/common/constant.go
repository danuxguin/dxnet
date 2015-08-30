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

//服务器类型定义
const (
	SERVER_TYPE_BEGIN         = 0
	SERVER_TYPE_LOGIN         = SERVER_TYPE_BEGIN //登陆服务器
	SERVER_TYPE_MAIN          = 1                 //主服务器
	SERVER_TYPE_FRIEND        = 2                 //好友服务器
	SERVER_TYPE_PUSH          = 3                 //推送服务器
	SERVER_TYPE_PUSHCENTER    = 4                 //推送中心服务器
	SERVER_TYPE_FILE          = 5                 //文件服务器
	SERVER_TYPE_GATE          = 6                 //网关服务器
	SERVER_TYPE_TREE          = 8                 //情感树服务器
	SERVER_TYPE_CLIENT_LOG    = 11                //客户端的日志Http服务器
	SERVER_TYPE_RESOURCE      = 13                //资源服务器
	SERVER_TYPE_FIND_FRIEND   = 15                //朋友计算服务器
	SERVER_TYPE_DATABASEMGR   = 19                //DatabaseMgr服务器
	SERVER_TYPE_WEBMANAGER    = 20                //WebManager服务器
	SERVER_TYPE_DNS           = 21                //短域名服务器
	SERVER_TYPE_SMS           = 22                //短信服务器
	SERVER_TYPE_WEB           = 23                //WEB服务器
	SERVER_TYPE_GAMEPLAY      = 24                //玩法服务器
	SERVER_TYPE_POSTGRESQLMGR = 25                //Postgresql管理器
	SERVER_TYPE_END           = 26                //结尾
	SERVER_TYPE_INVALID       = "INVALID_TYPE"    //常量定义
)

//消息类型定义
const (
	RES_TYPE_ALL     = 0
	RES_TYPE_IOS     = 1
	RES_TYPE_ANDROID = 2
	RES_TYPE_SERVER  = 88
)

//消息推送类型定义
const (
	MESSAGE_NOTIFY_MIN  = 0
	MESSAGE_NOTIFY_ALL  = MESSAGE_NOTIFY_MIN //全部
	MESSAGE_NOTIFY_CHAT = 1                  //聊天信息
	MESSAGE_NOTIFY_PUSH = 2                  //推送信息
	MESSAGE_NOTIFY_MAX  = 3
)

//聊天类型定义
const (
	CHAT_TYPE_NORMAL      = 0 //正常聊天
	CHAT_TYPE_SECRET_SEND = 1 //神秘人发送的聊天
	CHAT_TYPE_SECRET_RECV = 2 //神秘人接收的聊天
)

//短连接相关定义
const (
	SHORT_CONNECTION_READ_TIMEOUT      = 30
	SHORT_CONNECTION_WRITE_TIMEOUT     = 30
	SHORT_CONNECTION_READWRITE_TIMEOUT = 30
)

//官方小骚定义
const (
	COMMON_ACCOUNT            = "1b934c51830c547d4bd35d79ee1790942530381880" //官方小骚
	COMMON_ACCOUNT_CELL_PHONE = "+8613800000000"                             //官方小骚手机号
	COMMON_NICKNAME           = "官方小爱"                                       //官方小骚昵称
	COMMON_SECRETID           = 1                                            //官方小骚作为神秘人ID
	COMMON_QUESTIONID         = 0                                            //官方小骚的答题ID
)

//爱他状态定义
const (
	RELATION_ALREADY_LOVE = -2 //已经爱过
	RELATION_NOT_LOVE     = 0  //没有爱过
	RELATION_ACT_LOVE     = 1  //主动爱
	RELATION_BE_LOVED     = 2  //被爱
	RELATION_IN_LOVE      = 3  //互爱
)

//爱他状态定义
const (
	RELATION_ALREADY_APPLY_ADD_FRIEND = -2 //已经申请过成为好友
	RELATION_NOT_APPLY_ADD_FRIEND     = 0  //没有申请过成为好友
	RELATION_APPLY_ADD_FRIEND         = 1  //主动申请加对方为好友
	RELATION_BE_APPLIED_ADD_FRIEND    = 2  //被申请加为好友
	RELATION_BOTH_AGREE_ADD_FRIEND    = 3  //互相同意加为好友
)

const (
	FRIEND_SOURCE_CONTACTS         = 0 //好友来源:来自通讯录
	FRIEND_SOURCE_MEET             = 1 //好友来源:来自邂逅
	FRIEND_SOURCE_APPLY_ADD_FRIEND = 2 //好友来源:来自好友申请
	FRIEND_SOURCE_FRIEND_FRIEND    = 3 //好友来源:来自朋友的朋友
)

const (
	RELATION_DEL_FRIEND        = 1 //普通删除好友
	RELATION_PULL_BLACK        = 2 //拉黑
	RELATION_CANCEL_PULL_BLACK = 3 //取消拉黑
)

const (
	RELATION_NOT_FRIEND    = 0 //双方不是好友
	RELATION_A_IS_B_FRIEND = 1 //我是对方的好友
	RELATION_B_IS_A_FRIEND = 2 //对方是我的好友
	RELATION_BOTH_FRIEND   = 3 //互为好友
)

const (
	RELATION_ALREADY_APPLY_SUPER_LOVE = -2 //已经申请过成为恋爱模式
	RELATION_NOT_APPLY_SUPER_LOVE     = 0  //没有申请过成为恋爱模式
	RELATION_APPLY_SUPER_LOVE         = 1  //申请升级为恋爱模式
	RELATION_BE_APPLIED_SUPER_LOVE    = 2  //被申请升级为恋爱模式
	RELATION_BOTH_AGREE_SUPER_LOVE    = 3  //双方同意升级为恋爱模式
)

const (
	SEND_SMS_TYPE_ANSWER_QUESTION = 1 //签题类型
	SEND_SMS_TYPE_SEND_MUSIC      = 2 //点歌类型
	SEND_SMS_TYPE_OP_TREE         = 3 //情感树类型
	SEND_SMS_TYPE_CALL_UP1        = 4 //短信召唤一
	SEND_SMS_TYPE_CALL_UP2        = 5 //短信召唤二
)

const (
	MEET_NUM_ONE_DAY        = 50   //每天最多可以邂逅的数量
	MEET_NORMAL_NUM_ONE_DAY = 20   //每天正常可以邂逅的数量
	MEET_BUY_NUM_ONE_DAY    = 10   //每次购买获取的邂逅数量
	MEET_BUY_NEED_MONEY     = 1000 //每次购买邂逅次数花费的金钱
)

const (
	SEND_SMS_RECEIVE_SMS_MAX = 10 //收短信的限制数量
)

//客户端个人设置
const (
	SETTING_BIT_USER_EMOTION     = 7  //接收好友心情信息（回复，评论）				默认为0表示接收，如果为1的话，表示不接收
	SETTING_BIT_USER_ANONYMOUS   = 8  //接收不能说的秘密信息（回复，评论）			默认为0表示接收，如果为1的话，表示不接收
	SETTING_BIT_USER_RECEIVE_NEW = 9  //接收新消息通知 							默认为0表示接收，如果为1的话，表示不接收
	SETTING_BIT_USER_VOICE       = 10 //接收消息声音								默认为0表示接收，如果为1的话，表示不接收
	SETTING_BIT_USER_VIBRATION   = 11 //接收消息震动						    	默认为0表示接收，如果为1的话，表示不接收
)

//邂逅关系类型
const (
	MEET_TYPE_STRANGER      = 0 //陌生人
	MEET_TYPE_ATTENTIO      = 1 //关注你的有缘人
	MEET_TYPE_FRIEND_FRIEND = 2 //朋友的朋友
	MEET_TYPE_NEARBY        = 3 //附近的有缘人
)
