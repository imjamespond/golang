package controllers

import (
	"http-server/service/ffmpeg"
	gin_util "http-server/utils/gin"

	"github.com/gin-gonic/gin"
)

func DownloadM3u8(c *gin.Context) {

	body := ffmpeg.JobDl{}
	if err := c.BindJSON(&body); gin_util.HandleBadRequest(c, err) {
		return
	}
	gin_util.HandleBadRequest(c, ffmpeg.GetInstance().Add(&ffmpeg.Job{Type: 1, Data: &body}))
}

func ConvertVideo(c *gin.Context) {

	body := ffmpeg.JobConv{}
	if err := c.BindJSON(&body); gin_util.HandleBadRequest(c, err) {
		return
	}
	gin_util.HandleBadRequest(c, ffmpeg.GetInstance().Add(&ffmpeg.Job{Type: 2, Data: &body}))
}

func GetLog(c *gin.Context) {
	log, err := ffmpeg.GetInstance().Log()
	if gin_util.HandleBadRequest(c, err) {
		return
	}
	gin_util.HandleJSON(c, map[string]interface{}{"log": log})
}

func Kill(c *gin.Context) {
	ffmpeg.GetInstance().Kill()
}
