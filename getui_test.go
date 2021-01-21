package getui

import (
	"context"
	"encoding/json"
	"github.com/dacker-soul/getui/auth"
	"github.com/dacker-soul/getui/publics"
	"github.com/dacker-soul/getui/push/single"
	"strconv"
	"time"

	"github.com/alecthomas/assert"
	"testing"
)

var (
	Conf = publics.GeTuiConfig{
		AppId:        "",
		AppSecret:    "",
		AppKey:       "",
		MasterSecret: "",
	}
	Cid   = "" // clientId
	Ctx   = context.Background()
	Token = ""
	Alias = "test_user"
)

// 测试-获取Token
func TestGetToken(t *testing.T) {
	// 注意这个token过期时间为一天+1秒，每分钟调用量为100次，每天最大调用量为10万次
	result, err := auth.GetToken(Ctx, Conf)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, result.Code)
	Token = result.Data.Token
}

// 测试-单次推送ByCid 厂商通道+纯透模板
func TestPushSingleByCidA(t *testing.T) {
	iosChannel := publics.IosChannel{
		Type: "",
		Aps: &publics.Aps{
			Alert: &publics.Alert{
				Title: "卡是谁？",
				Body:  "为什么我们每天都要打 TA ？",
			},
			ContentAvailable: 0,
		},
		AutoBadge:      "+1",
		PayLoad:        "",
		Multimedia:     nil,
		ApnsCollapseId: "",
	}
	stringIos, _ := json.Marshal(iosChannel)

	singleParam := single.PushSingleParam{
		RequestId: strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
		Audience: &publics.Audience{ // 目标用户
			Cid:           []string{Cid}, // cid推送数组
			Alias:         nil,           // 别名送数组
			Tag:           nil,           // 推送条件
			FastCustomTag: "",            // 使用用户标签筛选目标用户
		},
		Settings: &publics.Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &publics.Strategy{ // 厂商通道策略，具体看public_struct.go
				Default: 1,
				Ios:     4,
				St:      1,
				Hw:      1,
				Xm:      1,
				Vv:      1,
				Mz:      1,
				Op:      1,
			},
			Speed:        100, // 推送速度，设置100表示：100条/秒左右，0表示不限速
			ScheduleTime: 0,   // 定时推送时间，必须是7天内的时间，格式：毫秒时间戳
		},
		PushMessage: &publics.PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: nil,
			Transmission: string(stringIos),
			Revoke:       nil,
		},
		PushChannel: &publics.PushChannel{
			Ios: &iosChannel,
			Android: &publics.AndroidChannel{Ups: &publics.Ups{
				Notification: nil,
				TransMission: string(stringIos), // 透传消息内容，与notification 二选一
			}},
		},
	}
	result, err := single.PushSingleByCid(Ctx, Conf, Token, &singleParam)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, result.Code)

}

// 测试-单次推送ByCid 厂商通道普通模板+个推通道普通模板
func TestPushSingleByCidB(t *testing.T) {
	iosChannel := publics.IosChannel{
		Type: "",
		Aps: &publics.Aps{
			Alert: &publics.Alert{
				Title: "卡是谁啊？",
				Body:  "为什么我们每天都要打 TA ？",
			},
			ContentAvailable: 0,
		},
		AutoBadge:      "+1",
		PayLoad:        "",
		Multimedia:     nil,
		ApnsCollapseId: "",
	}
	notification := publics.Notification{
		Title:       "卡是谁啊？",
		Body:        "为什么我们每天都要打 TA ？",
		ClickType:   "startapp", // 打开应用首页
		BadgeAddNum: 1,
	}

	singleParam := single.PushSingleParam{
		RequestId: strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
		Audience: &publics.Audience{ // 目标用户
			Cid:           []string{Cid}, // cid推送数组
			Alias:         nil,           // 别名送数组
			Tag:           nil,           // 推送条件
			FastCustomTag: "",            // 使用用户标签筛选目标用户
		},
		Settings: &publics.Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &publics.Strategy{ // 厂商通道策略，具体看public_struct.go
				Default: 1,
				Ios:     4,
				St:      4,
				Hw:      4,
				Xm:      4,
				Vv:      4,
				Mz:      4,
				Op:      4,
			},
			Speed:        100, // 推送速度，设置100表示：100条/秒左右，0表示不限速
			ScheduleTime: 0,   // 定时推送时间，必须是7天内的时间，格式：毫秒时间戳
		},
		PushMessage: &publics.PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: &notification,
			Transmission: "",
			Revoke:       nil,
		},
		PushChannel: &publics.PushChannel{
			Ios: &iosChannel,
			Android: &publics.AndroidChannel{Ups: &publics.Ups{
				Notification: &notification,
				TransMission: "", // 透传消息内容，与notification 二选一
			}},
		},
	}
	result, err := single.PushSingleByCid(Ctx, Conf, Token, &singleParam)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, result.Code)

}

// 测试-删除Token
func TestDelToken(t *testing.T) {
	result, err := auth.DelToken(Ctx, Token, Conf)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, result.Code)
}
