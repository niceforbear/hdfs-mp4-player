# HDFS-video-player

* Play mp4 file in HDFS by Go(Based on Beego framework) and WebHDFS in WebBrowser.
* Support multi HDFS cluster.

### Implements

* Implement by Go which based on beego.
* Use HDFS's WebHDFS to fetch HDFS's mp4 file data.
* Support `play jump` which use HTTP status code 206 and header Content-Range.

### Some tips

Mp4 metadata needs to be in front of the video. Maybe you have to understand `ffmpeg` and use it to keep the previous condition.

ffmpeg command which could make metadata in front of mp4:

> ffmpeg -i input.mp4 -movflags faststart -acodec copy -vcodec copy output.mp4 

### Usage

* Exec `bee run` in this project dir.
* Visit http://localhost:8082 in your browser, select one cluster which you need to configure it before, then browse hdfs directories and choose one mp4 file to play.
* You've better use Google Chrome browser, and could drag play button backward by raw <video> player.

### Contributing

* Welcome to have a contributing!
* Please see CONTRIBUTING for details.

### License

The MIT License (MIT). Please see License File for more information.

### Reference

* Beego framework: https://beego.me/
* WebHDFS: https://hadoop.apache.org/docs/r1.0.4/webhdfs.html
