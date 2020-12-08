// 群推
package getuiv2

import (
	"context"
	"encoding/json"
)

// 群推参数
type PushAllParam struct {
	RequestId   string       `json:"request_id"`   // 必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name"`   // 非必须，任务组名
	Audience    string       `json:"audience"`     // 必须字段，必须为all
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// 群推返回
type PushAllResult struct {
	PublicResult
	Data map[string]string `json:"data"`
}

// 群推
func PushAll(ctx context.Context, config GeTuiConfig, token string, param *PushAllParam) (*PushAllResult, error) {

	url := ApiUrl + config.AppId + "/push/all"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushAllResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
