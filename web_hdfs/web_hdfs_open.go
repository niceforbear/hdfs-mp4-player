package web_hdfs

import (
	"github.com/niceforbear/hdfs-mp4-player/helpers"
	"strconv"
)

func (this *WebHdfsApi) Open() ([]byte, error) {
	address, _ := WrapperListDirUrl(this.RequestUrl, this.FilePath)
	logger.Printf("webHdfs Open address: %v", address)

	reqBody := map[string]string{
		"op":this.Operation,
		"offset":strconv.Itoa(this.Offset),
		"length":strconv.Itoa(this.Length),
		"buffersize":strconv.Itoa(helpers.GetDefaultBufferSize()),
	}

	respBody, err := helpers.HttpWrapperGet(address, reqBody, nil)
	if err != nil {
		logger.Printf("[ERROR] request error: %v", err)
	}
	return respBody, err
}
