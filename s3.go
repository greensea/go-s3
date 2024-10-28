package s3

import (
	"bytes"
	"io"
	"log"

	"github.com/rhnvrm/simples3"
)

var S3_AK = ""
var S3_SK = ""
var S3_BUCKET = ""
var S3_ENDPOINT = ""

func Init(endpoint, bucket, ak, sk string) {
	S3_AK = ak
	S3_SK = sk
	S3_BUCKET = bucket
	S3_ENDPOINT = endpoint
}

func Put(k string, v []byte) error {
	_, err := S3().FilePut(simples3.UploadInput{
		Bucket:    S3_BUCKET,
		ObjectKey: k,
		Body:      bytes.NewReader(v),
	})

	if err != nil {
		log.Printf("上传 S3 失败: %v", err)
	}

	return err
}

func Get(k string) (io.ReadCloser, error) {
	file, err := S3().FileDownload(simples3.DownloadInput{
		Bucket:    S3_BUCKET,
		ObjectKey: k,
	})
	if err != nil {
		log.Printf("下载 S3 文件失败(%s): %v", k, err)
	}

	return file, err
}

func Delete(k string) error {
	err := S3().FileDelete(simples3.DeleteInput{
		Bucket:    S3_BUCKET,
		ObjectKey: k,
	})

	if err != nil {
		log.Printf("删除 S3 文件失败: %v", err)
	}

	return err
}

func Head(k string) (simples3.DetailsResponse, error) {
	d, err := S3().FileDetails(simples3.DetailsInput{
		Bucket:    S3_BUCKET,
		ObjectKey: k,
	})
	if err != nil {
		log.Printf("读取 S3 文件信息失败: %v", err)
	}

	return d, err
}
