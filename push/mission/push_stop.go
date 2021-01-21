// 停止任务,对正处于推送状态，或者未接收的消息停止下发
package mission

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// 停止任务参数
type PushStopParam struct {
	TaskId string `json:"taskId"`
}

// 停止任务返回
type PushStopResult struct {
	publics.PublicResult
}

// 停止任务
func PushStop(ctx context.Context, config publics.GeTuiConfig, token string, param *PushStopParam) (*PushStopResult, error) {
	url := publics.ApiUrl + config.AppId + "/task/" + param.TaskId
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "DELETE", token)
	if err != nil {
		return nil, err
	}

	var push *PushStopResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
