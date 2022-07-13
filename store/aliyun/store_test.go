package aliyun_test

import (
	"github.com/JavierZhou/cloud-station/store"
	"github.com/JavierZhou/cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

var (
	uploader store.Uploader
)

func TestAliOssStore_Upload(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		t.Log("upload ok")
	}
}

func init() {
	ali, err := aliyun.NewAliOssStore(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		panic(err)
	}
	uploader = ali
}
