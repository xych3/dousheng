package util

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func GetCoverFromVideo(userID int64, filename string, videoFilePath string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoFilePath).Filter("select", ffmpeg.Args{"gte(n,1)"}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return "", err
	}
	imgBytes := buf.Bytes()
	client, err := GetOSSClient()
	if err != nil {
		return "", err
	}
	imgPath := GetFileName(userID, filename) + ".jpeg"
	coverPath, err := client.UploadBytes(imgPath, imgBytes)
	if err != nil {
		return "", err
	}
	return coverPath, nil
}
