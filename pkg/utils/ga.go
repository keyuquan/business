package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type GoogleAuthInterface interface {
	QrcodeUrl(account, secret string) string
	Secret() string
	code(secret string) (string, error)
	VerifyCode(secret, code string) (bool, error)
}
type googleAuth struct {
}

func NewGoogleAuth() GoogleAuthInterface {
	return &googleAuth{}
}
func (googleAuth) Secret() string {
	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bts = make([]byte, 16)
	_, _ = rand.Read(bts)
	for k, v := range bts {
		bts[k] = dictionary[v%byte(len(dictionary))]
	}
	return strings.ToUpper(string(bts))

}

// 为了考虑时间误差，判断前当前时间及前后30秒时间
func (this googleAuth) VerifyCode(secret, code string) (bool, error) {
	// 当前google值
	_code, err := this.code(secret)
	if err != nil {
		return false, err
	}
	fmt.Println(_code)
	return _code == code, nil
}

// 获取Google Code
func (this googleAuth) code(secret string) (string, error) {
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	number := this.oneTimePassword(key, this.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

func (this googleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (this googleAuth) toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func (this googleAuth) oneTimePassword(key []byte, value []byte) uint32 {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := this.toUint32(hashParts)
	return number % 1000000
}

// 获取动态码二维码内容
func (this *googleAuth) getQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// 获取动态码二维码图片地址,这里是第三方二维码api
func (this *googleAuth) QrcodeUrl(user, secret string) string {
	////secret = this.decode(secret)
	//qrcode := this.getQrcode(user, secret)
	//width := "200"
	//height := "200"
	//data := url.Values{}
	//data.Set("data", qrcode)
	//return "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
	qrcode := this.getQrcode(user, secret)
	width := "200"
	height := "200"
	return "https://chart.googleapis.com/chart?chs=" + width + "x" + height + "&chld=M|0&cht=qr&chl=" + qrcode
}
