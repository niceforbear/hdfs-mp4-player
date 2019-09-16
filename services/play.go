package services

import (
	"github.com/niceforbear/hdfs-mp4-player/consts"
	"github.com/niceforbear/hdfs-mp4-player/helpers"
	"github.com/niceforbear/hdfs-mp4-player/web_hdfs"
	"log"
)

func GetContent(cluster string, filePath string, headerRange string) ([]byte, error) {
	requestUrl, err := web_hdfs.GetRequestUrl(cluster)
	if err != nil {
		log.Printf("[ERROR] in file list get request url: %v", err)
		return nil, err
	}

	webHdfsApi := web_hdfs.WebHdfsApi{
		RequestUrl:requestUrl,
		FilePath:filePath,
		Offset: helpers.GetOffsetFromHeaderRange(headerRange),
		Length:helpers.GetDefaultPlayLength(),
		Operation:consts.OpStatusOpen,
	}

	return webHdfsApi.Open()
}
