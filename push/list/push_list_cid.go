// cid批量推
package list

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// cid批量推参数
type PushListCidParam struct {
	Audience *publics.Audience `json:"audience"` // 必须字段，用cid数组，多个cid，注意这里 ！！数组长度不大于200
	IsAsync  bool              `json:"is_async"` // 非必须,默认值:false,是否异步推送，异步推送不会返回data,is_async为false时返回data
	TaskId   string            `json:"taskid"`   // 必须字段,默认值:无,使用创建消息接口返回的taskId，可以多次使用

}

// cid批量推返回
type PushListCidResult struct {
	publics.PublicResult
	Data map[string]map[string]string `json:"data"`
}

// cid批量推
func PushListCid(ctx context.Context, config publics.GeTuiConfig, token string, param *PushListCidParam) (*PushListCidResult, error) {

	url := publics.ApiUrl + config.AppId + "/push/list/cid"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushListCidResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
