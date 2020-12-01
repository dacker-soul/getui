// 向单个用户推送消息
package getuiv2

import (
	"context"
	"encoding/json"
)

type PushSingleParam struct {
	RequestId   string `json:"requestid"` // 请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Audience    *Cid   // cid数组，只能填一个cid
	Settings    *Settings
	PushMessage *PushMessage
	PushChannel *PushChannel
}

// 执行cid单推
func PushSingleByCid(ctx context.Context, config GeTuiConfig, token string, param *PushSingleParam) (*PushSingleResult, error) {

	url := ApiUrl + config.AppId + "/push/single/cid"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var pushSingleResult *PushSingleResult
	if err := json.Unmarshal([]byte(result), &pushSingleResult); err != nil {
		return nil, err
	}

	return pushSingleResult, err
}
