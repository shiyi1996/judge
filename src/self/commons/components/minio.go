/**
 * Created by leeezm on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

package components

import (
	"io"
	"self/commons/g"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/minio/minio-go"
)

type MinioCli struct {
	cli *minio.Client
}

func NewMinioCli() MinioCli {
	cfg := g.Conf()
	minioClient, err := minio.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.Secure)
	if err != nil {
		panic(err)
	}
	return MinioCli{cli: minioClient}
}
func (this MinioCli) GetImgName(userId int64, picType string) string {
	str := strconv.FormatInt(userId, 10)
	return str + "." + picType
}
func (this MinioCli) GetCodeName() string {
	var mutex sync.Mutex
	mutex.Lock()
	timestamp := time.Now().Unix()
	mutex.Unlock()
	str := strconv.FormatInt(timestamp, 10)
	return str + ".txt"
}
func (this MinioCli) GetPath(bucketName, objectName string) string {
	return "http://xupt1.fightcoder.com:9001/" + bucketName + "/" + objectName
}
func (this MinioCli) SaveImg(reader io.Reader, userId int64, picType string) string {
	cfg := g.Conf()
	str := this.GetImgName(userId, picType)
	_, err := this.cli.PutObject(cfg.Minio.ImgBucket, str, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
	return this.GetPath(cfg.Minio.ImgBucket, str)
}
func (this MinioCli) SaveCode(code string) string {
	cfg := g.Conf()
	str := this.GetCodeName()
	_, err := this.cli.PutObject(cfg.Minio.CodeBucket, str, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
	return this.GetPath(cfg.Minio.CodeBucket, str)
}
func (this MinioCli) DownloadCode(objectName, filePath string) {
	cfg := g.Conf()
	err := this.cli.FGetObject(cfg.Minio.CodeBucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
}

func (this MinioCli) DownloadCase(objectName, filePath string) {
	cfg := g.Conf()
	err := this.cli.FGetObject(cfg.Minio.CaseBucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
}
