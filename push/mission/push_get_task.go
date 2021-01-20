// 查询定时任务 该接口支持在推送完定时任务之后，查看定时任务状态，定时任务是否发送成功。
package mission

import (
	"context"
	"encoding/json"
	. "getui/publics"
)

// 查询定时任务参数
type PushGetTaskParam struct {
	TaskId string `json:"taskId"`
}

// 查询定时任务返回
type PushGetTaskResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 返回data示例
// "data": {
//	 "$taskid": {
//		 "create_time":"",			// 定时任务创建时间，毫秒时间戳
//		 "status":"success",		// 定时任务状态：success/failed
//		 "transmission_content":"",	// 透传内容
//		 "push_time":""				// 定时任务推送时间，毫秒时间戳
//	 }
// }

// 停止任务
func PushGetTask(ctx context.Context, config GeTuiConfig, token string, param *PushGetTaskParam) (*PushGetTaskResult, error) {
	url := ApiUrl + config.AppId + "/task/schedule" + param.TaskId
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "GET", token)
	if err != nil {
		return nil, err
	}

	var push *PushGetTaskResult
	if err := json.Unmarshal([]byte(result), &push); err != nil {
		return nil, err
	}

	return push, err
}
