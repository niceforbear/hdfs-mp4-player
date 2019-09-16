package web_hdfs

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var logger = logs.GetLogger("WebHDFS")

type WebHdfsApi struct {
	RequestUrl	string
	FilePath	string
	Offset		int
	Length		int
	BufferSize	int
	Operation	string	`json:"op"`
}

func GetRequestUrl(cluster string) (string, error) {
	requestUrl := ""
	if cluster == "cluster1" {
		requestUrl = beego.AppConfig.String("cluster1WebHdfsApi")
	} else if cluster == "cluster2" {
		requestUrl = beego.AppConfig.String("cluster2WebHdfsApi")
	} else {
		return requestUrl, errors.New("[ERROR] error cluster")
	}

	return requestUrl, nil
}

func WrapperListDirUrl(serverUrl string, filePath string) (requestUrl string, err error) {
	joinUrls := []string{
		serverUrl, filePath,
	}

	return strings.Join(joinUrls, "/"), nil
}