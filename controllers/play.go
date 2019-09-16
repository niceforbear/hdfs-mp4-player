package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/niceforbear/hdfs-mp4-player/helpers"
	"github.com/niceforbear/hdfs-mp4-player/services"
	"strconv"
	"strings"
)

type PlayController struct {
	BaseController
}

func (this *PlayController) Get() {
	cluster := this.GetString("cluster")
	filePath := this.GetString("file")

	logger.Printf("Request params: cluster: %v, filepath: %v", cluster, filePath)

	this.Data["FilePath"] = filePath
	this.Data["cluster"] = cluster
	this.TplName = "play.html"
}

func (this *PlayController) StreamV2() {
	cluster := this.GetString("cluster")
	filePath := this.GetString("filepath")

	fileSize := services.GetFileSizeFromFilePathCluster(cluster, filePath)
	logger.Printf("filesize from api: %v", fileSize)

	headerRange := this.Ctx.Input.Header("Range")  // e.g. bytes=0-

	logger.Printf("headerRange: %v", headerRange)

	rangeOffset := helpers.GetOffsetFromHeaderRange(headerRange)
	logger.Printf("range offset: %v", rangeOffset)

	logger.Printf("Request params: cluster: %v, filepath: %v, range: %v, filesize: %v", cluster, filePath, headerRange, fileSize)

	data, err := services.GetContent(cluster, filePath, headerRange)
	if err != nil {
		logger.Printf("[ERROR] get content error: %v", err)
	}

	startByte := 0
	endByte := fileSize - 1

	if headerRange != "" && strings.Contains(headerRange, "bytes=") && strings.Contains(headerRange, "-") {
		headerRange = strings.Replace(headerRange, "bytes=", "", -1)

		headerRangeArray := strings.Split(headerRange, "-")
		startByte, err = strconv.Atoi(headerRangeArray[0])
		if err != nil {
			panic(err)
		}

		if headerRangeArray[1] != "" {
			endByte, err = strconv.Atoi(headerRangeArray[1])

			if err != nil {
				panic(err)
			}
		}
	}

	contentLength := strconv.Itoa(endByte - startByte + 1)
	defaultPlaySize, _ := beego.AppConfig.Int("DefaultPlayLength")
	contentRange := fmt.Sprintf("bytes %v-%v/%v", rangeOffset, rangeOffset + defaultPlaySize - 1, fileSize)

	if strings.Contains(headerRange, "=0-") {
		this.Ctx.Output.Status = 200
	} else {
		this.Ctx.Output.Status = 206
	}

	this.Ctx.Output.Header("Content-Type", "video/mp4")
	this.Ctx.Output.Header("Content-Length", contentLength)
	this.Ctx.Output.Header("Accept-Ranges", "bytes")
	this.Ctx.Output.Header("Content-Range", contentRange)
	filename := strings.Split(filePath, "/")
	this.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("inline; filename=%v", filename[len(filename) - 1]))

	_ = this.Ctx.Output.Body(data)
}