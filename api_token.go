package getuiv2

import (
	"context"
	"encoding/json"
	"fmt"
)

type TokenParam struct {
	Sign      string `json:"sign"`      // 加密算法: SHA256，格式:sha256(appkey+timestamp+mastersecret)
	Timestamp string `json:"timestamp"` // 毫秒时间戳，请使用当前毫秒时间戳，误差太大可能出错
	AppKey    string `json:"appkey"`    // 创建应用时生成的appkey
}

// Token返回结构
type TokenResult struct {
	PublicResult
	Data TokenResultData
}

// Token返回的data结构
type TokenResultData struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}

/*
 * token是个推开放平台全局唯一接口调用凭据，访问所有接口都需要此凭据，开发者需要妥善保管。
 * 注意：鉴权接口每分钟最大调用量为100次，每天最大调用量为10万次。
 * token的有效期截止时间通过接口返回参数expire_time来标识，目前是接口调用时间+1天的毫秒时间戳。
 * token过期后无法使用，开发者需要定时刷新。
 * 为保证高可用，建议开发者在定时刷新的同时做被动刷新，即当调用业务接口返回错误码10001时调用获取token被动刷新。
 */
func GetToken(ctx context.Context, config GeTuiConfig) (TokenResult, error) {
	tokenResult := TokenResult{}
	// 获取加密字符串和时间戳
	signStr, timestamp := Signature(config.AppKey, config.MasterSecret)

	param := &TokenParam{
		Sign:      signStr,
		Timestamp: timestamp,
		AppKey:    config.AppKey,
	}

	url := API_URL + config.AppId + "/auth"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return tokenResult, err
	}

	result, err := SendPost(ctx, bodyByte, url, "POST", "")
	if err != nil {
		return tokenResult, err
	}

	if err := json.Unmarshal([]byte(result), &tokenResult); err != nil {
		return tokenResult, err
	}

	return tokenResult, nil
}

// 删除token，为防止token被滥用或泄露，开发者可以调用此接口主动使token失效
func DelToken(ctx context.Context, token string, config GeTuiConfig) (PublicResult, error) {
	publicResult := PublicResult{}
	url := API_URL + config.AppId + "/auth/" + token
	fmt.Println("url:", url)
	result, err := SendPost(ctx, []byte{}, url, "DELETE", "")
	if err != nil {
		return publicResult, err
	}
	if err := json.Unmarshal([]byte(result), &publicResult); err != nil {
		return publicResult, err
	}
	return publicResult, nil
}
