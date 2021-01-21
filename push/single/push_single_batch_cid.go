// cid批量推送消息
package single

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// cid批量单推参数
type PushSingleBatchCidParam struct {
	IsAsync bool               `json:"is_async"` // 非必须,默认值:false,是否异步推送，异步推送不会返回data,is_async为false时返回data
	MsgList []*PushSingleParam `json:"msg_list"` // 必须,默认值:无，消息内容，数组长度不大于 200
}

// cid批量单推返回
type PushSingleBatchCidResult struct {
	publics.PublicResult
	Data map[string]map[string]string `json:"data"`
}

// cid批量单推
func PushSingleByBatchCid(ctx context.Context, config publics.GeTuiConfig, token string, param *PushSingleBatchCidParam) (*PushSingleBatchCidResult, error) {

	url := publics.ApiUrl + config.AppId + "/push/single/batch/cid"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var push *PushSingleBatchCidResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
