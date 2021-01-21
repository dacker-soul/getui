// cid单推
package single

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// cid单推参数
type PushSingleParam struct {
	RequestId   string               `json:"request_id"`   // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Audience    *publics.Audience    `json:"audience"`     // 必须字段，cid数组，只能填一个cid
	Settings    *publics.Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *publics.PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *publics.PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// cid单推返回
type PushSingleResult struct {
	publics.PublicResult
	Data map[string]map[string]string `json:"data"`
}

// cid单推
func PushSingleByCid(ctx context.Context, config publics.GeTuiConfig, token string, param *PushSingleParam) (*PushSingleResult, error) {

	url := publics.ApiUrl + config.AppId + "/push/single/cid"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushSingleResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
