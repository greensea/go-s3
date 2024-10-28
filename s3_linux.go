package s3

import (
	"github.com/rhnvrm/simples3"
)

func S3() *simples3.S3 {
	s := simples3.New("", S3_AK, S3_SK)

	s.SetEndpoint(S3_ENDPOINT)

	return s
}
