package utils

import (
	"business/pkg/log"
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

func IP2Long(ip string) uint32 {
	var long uint32
	err := binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	if err != nil {
		log.Infof("IP转long类型错误：%v", err.Error())
		return 0
	}
	return long
}
func LongToIP(ip interface{}) string {
	var ipInt int64
	switch ip.(type) {
	case int64:
		ipInt = ip.(int64)
	case string:
		ipInt, _ = strconv.ParseInt(ip.(string), 10, 64)
		if ipInt <= 0 {
			return ""
		}
	case int:
		ipInt = int64(ip.(int))
	default:
		return ""
	}
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}
