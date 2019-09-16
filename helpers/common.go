package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

var logger = logs.GetLogger("helper")

func IsClusterValidate(cluster string) bool {
	if cluster == "cluster1" || cluster == "cluster2"{
		return true
	}

	return false
}

func GetOffsetFromHeaderRange(headerRange string) int {
	if headerRange == "" {
		return 0
	}

	// e.g. bytes=0-
	words := strings.Split(headerRange, "=")
	offsets := strings.Split(words[1], "-")

	res, _ :=strconv.Atoi(offsets[0])

	return res
}

func GetDefaultPlayLength() int {
	res, err := beego.AppConfig.Int("DefaultPlayLength")
	if err != nil {
		logger.Printf("[ERROR] in get default play length: %v", err)
	}
	return res
}

func GetDefaultBufferSize() int {
	res, _ := beego.AppConfig.Int("DefaultBufferSize")
	return res
}
