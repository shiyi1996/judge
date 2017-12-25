/**
 * Created by shiyi on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

package store

import "github.com/minio/minio-go"

//"log"
//
//"self/commons/g"
//
//"github.com/minio/minio-go"

var cli *minio.Client

//func InitMinio() {
//	cfg := g.Conf()
//	var err error
//	Client, err = minio.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.Secure)
//	if err != nil {
//		log.Fatalln("fail to connect minio ", cfg.Minio.Endpoint, err)
//		return
//	}
//}
//
//func MakeBucket(bucketName string) {
//	cli.MakeBucket("sad", "us-east-1")
//}
