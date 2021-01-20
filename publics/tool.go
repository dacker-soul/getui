// 公共方法
package publics

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 获取加密后的字符串
func Signature(appKey string, masterSecret string) (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) //签名开始生成毫秒时间
	original := appKey + timestamp + masterSecret
	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), timestamp
}

//post请求
func RestFulRequest(ctx context.Context, bodyByte []byte, url, action, token string) (string, error) {
	//创建客户端实例
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	body := bytes.NewBuffer(bodyByte)

	//创建请求实例
	req, err := http.NewRequest(action, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Add("token", token)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//读取响应
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
