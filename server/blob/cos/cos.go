package cos

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Service struct {
	client    *cos.Client
	secretID  string
	secretKey string
}

func NewService(rawURL, secretID, secretKey string) (*Service, error) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("rawURL invalid: %+v", err)
	}

	return &Service{
		client: cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				// 通过环境变量获取密钥
				// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretID: secretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
				// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretKey: secretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			},
		}),
		secretID:  secretID,
		secretKey: secretKey,
	}, nil
	//name := "exampleobject"
	//ctx := context.Background()
	//// 1. 通过普通方式下载对象
	//resp, err := client.Object.Get(ctx, name, nil)
	//if err != nil {
	//	panic(err)
	//}
	//bs, _ := ioutil.ReadAll(resp.Body)
	//resp.Body.Close()
	//// 获取预签名 URL
	//presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, name, ak, sk, time.Hour, nil)
	//if err != nil {
	//	panic(err)
	//}
	//// 2. 通过预签名 URL下载对象
	//resp2, err := http.Get(presignedURL.String())
	//if err != nil {
	//	panic(err)
	//}
	//bs2, _ := ioutil.ReadAll(resp2.Body)
	//resp2.Body.Close()
	//if bytes.Compare(bs2, bs) != 0 {
	//	panic(errors.New("content is not consistent"))
	//}
}

func (s *Service) SignURL(ctx context.Context, method string, path string, timeout time.Duration) (url string, err error) {
	u, err := s.client.Object.GetPresignedURL(ctx, method, path, s.secretID, s.secretKey, timeout, nil)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func (s *Service) Get(ctx context.Context, path string) (io.ReadCloser, error) {
	res, err := s.client.Object.Get(ctx, path, nil)
	var r io.ReadCloser
	if res != nil {
		r = res.Body
	}

	if err != nil {
		return r, err
	}

	if res.StatusCode >= 400 {
		return r, fmt.Errorf("got error response %+v", res)
	}
	return r, nil
}
