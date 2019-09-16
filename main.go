package main

import (
	"github.com/astaxie/beego"
	_ "github.com/niceforbear/hdfs-mp4-player/routers"
)

func main() {
	beego.Run()
}

