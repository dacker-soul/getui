// 向单个用户推送消息 By别名
package getuiv2

import (
	"context"
	"encoding/json"
)

// 单推ByAlias参数
type PushSingleBatchAliasParam struct {
	IsAsync bool                    `json:"is_async"` // 非必须,默认值:false,是否异步推送，异步推送不会返回data,is_async为false时返回data
	MsgList []*PushSingleAliasParam `json:"msg_list"` // 必须,默认值:无，消息内容，数组长度不大于 200

}

// 单推ByAlias返回
type PushSingleBatchAliasResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 执行单推别名
func PushSingleByBatchAlias(ctx context.Context, config GeTuiConfig, token string, param *PushSingleBatchAliasParam) (*PushSingleBatchAliasResult, error) {

	url := ApiUrl + config.AppId + "/push/single/batch/alias"
	bodyByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	result, err := RestFulRequest(ctx, bodyByte, url, "POST", token)
	if err != nil {
		return nil, err
	}

	var pushSingleResult *PushSingleBatchAliasResult
	if err := json.Unmarshal([]byte(result), &pushSingleResult); err != nil {
		return nil, err
	}

	return pushSingleResult, err
}
