package constant

const (
	WebUserTokenExpireIn       = 24 * 3600
	AccessTokenExpireIn        = 7200 //access_token有效时间2小时，有新token后，老token5分钟内有效
	AccessTokenTimestampExpire = 600  //申请access_token的时间戳默认有效期10分钟即600s
	AccessTokenOldExpireIn     = 600  //old access token expirein
)
