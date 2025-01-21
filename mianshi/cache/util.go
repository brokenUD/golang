package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func ParseSize(size string)(int64, string) {
	// 默认大小为100MB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	default:
		num = 0
		byteNum = 0
	}

	if num == 0 {
		log.Println("parseSize XXX")
		num = 100
		byteNum = 100 * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}
