// 群推-使用标签快速推送
package all

import (
	"context"
	"encoding/json"
	. "getui/publics"
)

// 使用标签快速推送参数
type PushFastCustomTagParam struct {
	RequestId   string       `json:"request_id"`   // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	Audience    *Audience    `json:"audience"`     // 必须字段，tag数组
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// 使用标签快速推送返回
type PushFastCustomTagResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 根据条件筛选用户推送
func PushFastCustomTag(ctx context.Context, config GeTuiConfig, token string, param *PushFastCustomTagParam) (*PushFastCustomTagResult, error) {
	url := ApiUrl + config.AppId + "/push/fast_custom_tag"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushFastCustomTagResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
