package constant

const (
	UserVerifyCode         = "User:VerifyCode:"          //用户验证码 prefix:功能:userid
	UserSendMsgCode        = "User:SendMsgCode:"         //短信验证码 prefix:功能:userid
	UserOnlineContainer    = "User:Online:"              //用户登录成功后记录userid、token、登录时间
	UserAccessTokenCurrent = "User:AccessToken:Current:" //通过access_token方式访问的，记录appid：access_token
	UserAccessTokenOld     = "User:AccessToken:Old:"     //通过access_token方式访问的，记录access_token:user
	UserAccessTokenLock    = "User:AccessToken:Lock:"    //通过access_token方式访问的，redis分布式锁key

	UserList    = "User:List:" //保存userlist（包含roles）-> prefix+
	UserApiList = "User:Api:"  //保存user拥有的api列表 -> redis hash appid uri uri
)
