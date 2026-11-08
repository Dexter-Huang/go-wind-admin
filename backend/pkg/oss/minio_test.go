package oss

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tx7do/go-utils/trans"
	bLogger "github.com/tx7do/kratos-bootstrap/logger"

	storageV1 "go-wind-admin/api/gen/go/storage/service/v1"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

// MinIO 测试端点与凭据默认值（与既有硬编码行为保持一致）。
const (
	defaultMinioTestEndpoint  = "127.0.0.1:9000"
	defaultMinioTestAccessKey = "root"
	defaultMinioTestSecretKey = "*Abcd123456"
)

// minioTestEndpoint 返回测试用 MinIO 端点：优先读环境变量 MINIO_TEST_ENDPOINT，
// 未设置时回退默认值。本机容器把 9000 映射到 19000 时，只需
// MINIO_TEST_ENDPOINT=127.0.0.1:19000 即可让全部实测指向该容器。
func minioTestEndpoint() string {
	if v := os.Getenv("MINIO_TEST_ENDPOINT"); v != "" {
		return v
	}
	return defaultMinioTestEndpoint
}

// minioTestAccessKey 返回测试用访问密钥：环境变量 MINIO_TEST_ACCESS_KEY，缺省回退默认值。
func minioTestAccessKey() string {
	if v := os.Getenv("MINIO_TEST_ACCESS_KEY"); v != "" {
		return v
	}
	return defaultMinioTestAccessKey
}

// minioTestSecretKey 返回测试用私钥：环境变量 MINIO_TEST_SECRET_KEY，缺省回退默认值。
func minioTestSecretKey() string {
	if v := os.Getenv("MINIO_TEST_SECRET_KEY"); v != "" {
		return v
	}
	return defaultMinioTestSecretKey
}

// createTestClient 构造指向测试端点（环境变量可覆盖）的 MinIOClient。
func createTestClient() *MinIOClient {
	endpoint := minioTestEndpoint()
	return NewMinIoClient(&conf.Bootstrap{
		Oss: &conf.OSS{
			Minio: &conf.OSS_MinIO{
				Endpoint:     endpoint,
				UploadHost:   endpoint,
				DownloadHost: endpoint,
				AccessKey:    minioTestAccessKey(),
				SecretKey:    minioTestSecretKey(),
			},
		},
	}, bLogger.NopLogger())
}

// 显式端点必须实测，连接或鉴权失败不能变成 skip。 / Explicit endpoints must fail on errors.
func minioAvailable() bool {
	if os.Getenv("MINIO_TEST_ENDPOINT") != "" {
		return true
	}
	return defaultMinioAvailable()
}

// 默认端点需完成 S3 请求，不能把空转发端口当作服务。 / Probe S3, not just TCP.
var defaultMinioAvailable = sync.OnceValue(func() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := createTestClient().GetClient().ListBuckets(ctx)
	if err == nil {
		return true
	}
	// S3 明确返回鉴权错误时仍运行测试，保留错误配置的失败信号。
	// Preserve failures when the S3 service rejects credentials.
	switch minio.ToErrorResponse(err).Code {
	case "AccessDenied", "InvalidAccessKeyId", "SignatureDoesNotMatch":
		return true
	}
	fmt.Fprintf(os.Stderr, "default S3 test endpoint unavailable: %v (set MINIO_TEST_ENDPOINT to require live tests)\n", err)
	return false
})

func TestMinIoClient(t *testing.T) {
	if !minioAvailable() {
		t.Skip("local MinIO (127.0.0.1:9000) not available, skipping environment-dependent test")
	}
	cli := createTestClient()
	assert.NotNil(t, cli)

	resp, err := cli.GetUploadPresignedUrl(t.Context(), &storageV1.GetUploadPresignedUrlRequest{
		Method:        storageV1.GetUploadPresignedUrlRequest_Put,
		ContentType:   trans.String("image/jpeg"),
		BucketName:    trans.String("images"),
		FileDirectory: trans.String("20221010"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestListFile(t *testing.T) {
	cli := createTestClient()
	assert.NotNil(t, cli)

	req := &storageV1.ListOssFileRequest{
		BucketName: trans.Ptr("users"),
		Folder:     trans.Ptr("1"),
		Recursive:  trans.Ptr(true),
	}
	files, err := cli.ListFile(t.Context(), req)
	assert.Nil(t, err)
	fmt.Println(files)
}

func TestDownloadFile(t *testing.T) {
	if !minioAvailable() {
		t.Skip("local MinIO (127.0.0.1:9000) not available, skipping environment-dependent test")
	}
	cli := createTestClient()
	assert.NotNil(t, cli)

	resp, err := cli.DownloadFile(t.Context(), &storageV1.DownloadFileRequest{
		Selector: &storageV1.DownloadFileRequest_StorageObject{
			StorageObject: &storageV1.StorageObject{
				BucketName: trans.Ptr("images"),
				ObjectName: trans.Ptr("DateTimePicker.png"),
			},
		},
		PreferPresignedUrl: trans.Ptr(false),
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
