package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/niceforbear/hdfs-mp4-player/consts"
	"github.com/niceforbear/hdfs-mp4-player/web_hdfs"
	"log"
	"path"
	"strings"
	"time"
)

var logger = logs.GetLogger("FileListService")

type FileList struct {
	Cluster	string
	Folder	string
}

type ServiceFileStatus struct {
	Group				string		`json:"group"`
	Length				int			`json:"length"`
	ModificationTime	string		`json:"modificationTime"`
	Owner				string		`json:"owner"`
	PathSuffix			string		`json:"pathSuffix"`
	Permission			string		`json:"permission"`
	Type				string		`json:"type"`
	Op					string		`json:"op"`
}

func (this *FileList) Get() ([]ServiceFileStatus, error) {
	requestUrl, err := web_hdfs.GetRequestUrl(this.Cluster)
	if err != nil {
		log.Printf("[ERROR] in file list get request url: %v", err)
		return nil, err
	}

	webHdfsApi := web_hdfs.WebHdfsApi{
		RequestUrl: requestUrl,
		FilePath:this.Folder,
		Operation:consts.OpStatusListDir,
	}

	files, err := webHdfsApi.ListDir()
	var retFiles []ServiceFileStatus
	for _, item := range files {
		modTime := time.Unix(item.ModificationTime / 1000, 0).Format("2006-01-02 03:04:05PM")
		tmpFileStatus := ServiceFileStatus{
			Group:item.Group,
			Length:item.Length,
			ModificationTime:modTime,
			Owner:item.Owner,
			PathSuffix:item.PathSuffix,
			Permission:item.Permission,
			Type:item.Type,
			Op:GetOpByType(item.Type, item.PathSuffix),

		}

		retFiles = append(retFiles, tmpFileStatus)

	}

	return retFiles, nil
}

func GetOpByType(fileType string, fileName string) string {
	if fileType == "DIRECTORY" {
		return "DIRECTORY"
	} else {
		if strings.HasSuffix(fileName, "mp4") {
			return "VIDEO"
		} else {
			return "PASS"
		}
	}
}

func GetFileSizeFromFilePathCluster(cluster string, filePath string) int {
	folder := path.Dir(filePath)
	logger.Printf("folder: %v", folder)

	fileListService := FileList{
		Cluster:cluster,
		Folder:folder,
	}

	files, err := fileListService.Get()
	if err != nil {
		panic(err)
	}

	for _, hdfsFile := range files {
		if folder + "/" + hdfsFile.PathSuffix == filePath {
			logger.Printf("[FileSize] %v", hdfsFile.Length)
			return hdfsFile.Length
		}
	}

	logger.Printf("[ERROR] Can not find filesize!")
	return 0
}