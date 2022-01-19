package gen

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
	"reflect"
	"regexp"
	"time"
	"unsafe"
)

// UUID 生成UUID
func UUID() string {
	return uuid.NewV4().String()
}

var (
	snowflake         = newSnowflake()
	numbers           = "0123456789"
	lowerLetter       = "abcdefghijklmnopqrstuvwxyz0123456789"
	capitalLetter     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	capitalLetterBase = []byte(capitalLetter)
	letterNum         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func newSnowflake() *sonyflake.Sonyflake {
	startTime, _ := time.ParseInLocation("2006-01-02", "2021-12-01", time.Local)

	snowflake := sonyflake.NewSonyflake(sonyflake.Settings{StartTime: startTime})
	if snowflake == nil {
		panic("创建snowflake失败, snowflake实例为nil")
	}

	return snowflake
}

// SnowflakeID 生成雪花ID
func SnowflakeID() (uint64, error) {
	return snowflake.NextID()
}

// ByteToString String and []byte buffers may converted without memory allocations
//This is an unsafe way, the result string and []byte buffer share the same bytes.
//Please make sure not to modify the bytes in the []byte buffer if the string still survives!
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToByte String and []byte buffers may converted without memory allocations
//This is an unsafe way, the result string and []byte buffer share the same bytes.
//Please make sure not to modify the bytes in the []byte buffer if the string still survives!
func StringToByte(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// ValidPhoneNumber 正则验证手机号
func ValidPhoneNumber(phone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
