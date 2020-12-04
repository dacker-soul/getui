// 向单个用户推送消息
package getuiv2

import (
	"context"
	"encoding/json"
)

// 单推参数
type PushSingleParam struct {
	RequestId   string       `json:"request_id"`   // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失 例如：strconv.FormatInt(time.Now().UnixNano(), 10)
	Audience    *Cid         `json:"audience"`     // 必须字段，cid数组，只能填一个cid
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

//
type PushSingleResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 返回示例
//{
//	"code": 0,
//	"msg": "",
//	"data": {
//	"$taskid": {
//		"$cid":"$status"
//		}
//	}
//}

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
