package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"sort"
	"strconv"
	"time"
)

func B64Encode(text string) string {
	encodeStr := base64.URLEncoding.EncodeToString([]byte(text))
	return encodeStr
}

func B64Decode(secret string) string {
	decodeStr, _ := base64.URLEncoding.DecodeString(secret)
	return string(decodeStr)
}

func Md5Encrypt(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}

func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	//index的取值：[0,len(str_array)]
	if index < len(strArray) && strArray[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

func NotNeedCookie(target string) bool {
	var strArray = []string{
		"/login", "/favicon.ico", "/register",
	}
	return in(target, strArray)
}

func CreatUid(username string, password string, timeStr string) string {
	uidStr := username + password + timeStr
	return Md5Encrypt(uidStr)
}

func CreateSecret(key []string) string {
	secret := ""
	for _, str := range key {
		secret += str
	}
	return Md5Encrypt(secret)
}

func GetNowTimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
