// 创建消息体，并返回taskid，为批量推的前置步骤
package getuiv2

import (
	"context"
	"encoding/json"
)

// 创建消息体参数
type PushListMessageParam struct {
	RequestId   string       `json:"request_id"`   // 非必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失 例如：strconv.FormatInt(time.Now().UnixNano(), 10)
	GroupName   string       `json:"group_name"`   // 非必须，任务组名
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// 创建消息体返回
type PushListMessageResult struct {
	PublicResult
	Data map[string]string `json:"data"` // taskid:任务编号，用于执行cid批量推和执行别名批量推，此taskid可以多次使用，有效期为用户设置的离线时间
}

// 创建消息体
func PushListMessage(ctx context.Context, config GeTuiConfig, token string, param *PushListMessageParam) (*PushListMessageResult, error) {

	url := ApiUrl + config.AppId + "/push/list/message"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushListMessageResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
