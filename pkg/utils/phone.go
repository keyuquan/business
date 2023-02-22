package utils

import (
	"business/pkg/log"
	"github.com/xluohome/phonedata"
	"regexp"
)

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
func GetPhoneData(phone string) string {
	if !VerifyMobileFormat(phone) {
		return ""
	}
	pr, err := phonedata.Find(phone)
	if err != nil {
		log.Infof("获取号码归属地错误：%v", err)
		return ""
	}
	return pr.Province + "-" + pr.City
}
