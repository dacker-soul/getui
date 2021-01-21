// 别名单推
package single

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// 别名单推参数
type PushSingleAliasParam struct {
	RequestId   string               `json:"request_id"`   // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Audience    *publics.Audience    `json:"audience"`     // 必须字段，别名数组，只能填一个别名
	Settings    *publics.Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *publics.PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *publics.PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// 别名单推返回
type PushSingleAliasResult struct {
	publics.PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 别名单推
func PushSingleByAlias(ctx context.Context, config publics.GeTuiConfig, token string, param *PushSingleAliasParam) (*PushSingleAliasResult, error) {

	url := publics.ApiUrl + config.AppId + "/push/single/alias"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushSingleAliasResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
