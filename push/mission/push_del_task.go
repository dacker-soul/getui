// 删除定时任务 用来删除还未下发的任务，删除后定时任务不再触发(距离下发还有一分钟的任务，将无法删除，后续可以调用停止任务接口)
package mission

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/publics"
)

// 删除定时任务参数
type PushDelTaskParam struct {
	TaskId string `json:"taskId"`
}

// 删除定时任务返回
type PushDelTaskResult struct {
	publics.PublicResult
}

// 删除定时任务
func PushDelTask(ctx context.Context, config publics.GeTuiConfig, token string, param *PushDelTaskParam) (*PushDelTaskResult, error) {
	url := publics.ApiUrl + config.AppId + "/task/schedule" + param.TaskId
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := publics.RestFulRequest(ctx, bodyByte, url, "DELETE", token)
	if err != nil {
		return nil, err
	}

	var push *PushDelTaskResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
