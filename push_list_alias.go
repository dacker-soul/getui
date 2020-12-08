// 别名批量推
package getuiv2

import (
	"context"
	"encoding/json"
)

// 别名批量推参数
type PushListAliasParam struct {
	Audience *Audience `json:"audience"` // 必须字段，用alias数组，，注意这里 ！！数组长度不大于200
	IsAsync  bool      `json:"is_async"` // 非必须,默认值:false,是否异步推送，异步推送不会返回data,is_async为false时返回data
	TaskId   string    `json:"taskid"`   // 必须字段,默认值:无,使用创建消息接口返回的taskId，可以多次使用
}

// 别名批量推返回
type PushListAliasResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 别名批量推
func PushListAlias(ctx context.Context, config GeTuiConfig, token string, param *PushListAliasParam) (*PushListAliasResult, error) {

	url := ApiUrl + config.AppId + "/push/list/alias"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushListAliasResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
