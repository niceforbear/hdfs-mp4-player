package controllers

import (
	"github.com/niceforbear/hdfs-mp4-player/helpers"
	"github.com/niceforbear/hdfs-mp4-player/services"
)

type FileListController struct {
	BaseController
}

func (this *FileListController) Get() {
	cluster := this.GetString("cluster")
	folder := this.GetString("folder", "/")

	logger.Printf("[REQUEST] FileList.Get, cluster: %v, folder: %v", cluster, folder)

	if helpers.IsClusterValidate(cluster) == false {
		this.Data["retCode"] = 1
		this.Data["msg"] = "Error cluster name"
		this.Data["files"] = []string{}
	} else {
		fileListService := services.FileList{
			Cluster:cluster,
			Folder:folder,
		}

		files, err := fileListService.Get()
		if err != nil {
			logger.Printf("[ERROR] in file list controller, %v", err)
			this.Data["retCode"] = 2
			this.Data["msg"] = err.Error()
			this.Data["files"] = []string{}
		}

		this.Data["retCode"] = 0
		this.Data["msg"] = "success"
		this.Data["files"] = files
	}

	this.Data["cluster"] = cluster
	if folder == "/" {
		this.Data["folder"] = ""
	} else {
		this.Data["folder"] = folder
	}

	logger.Printf("this.Data: %v", this.Data)
	this.TplName = "file_list.html"
}

func (this *FileListController) ApiGet() {
	result := map[string]interface{}{}
	cluster := this.GetString("cluster")
	folder := this.GetString("folder", "/")

	logger.Printf("[REQUEST] FileList.Get, cluster: %v, folder: %v", cluster, folder)

	if helpers.IsClusterValidate(cluster) == false {
		result["retCode"] = 1
		result["msg"] = "Error cluster name"
		result["files"] = []string{}
	} else {
		fileListService := services.FileList{
			Cluster:cluster,
			Folder:folder,
		}

		files, err := fileListService.Get()
		if err != nil {
			logger.Printf("[ERROR] in file list controller, %v", err)
			result["retCode"] = 2
			result["msg"] = err.Error()
			result["files"] = []string{}
		}

		result["retCode"] = 0
		result["msg"] = "success"
		result["files"] = files
	}

	result["cluster"] = cluster
	if folder == "/" {
		result["folder"] = ""
	} else {
		result["folder"] = folder
	}
	this.Data["json"] = result
	this.ServeJSON()
}