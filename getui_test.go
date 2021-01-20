package getui

import (
	"context"
	"encoding/json"
	"getui/push/single"
	"strconv"
	"time"

	"getui/auth"
	. "getui/publics"
	"github.com/alecthomas/assert"
	"testing"
)

var (
	Conf = GeTuiConfig{
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
	iosChannel := IosChannel{
		Type: "",
		Aps: &Aps{
			Alert: &Alert{
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
		Audience: &Audience{ // 目标用户
			Cid:           []string{Cid}, // cid推送数组
			Alias:         nil,           // 别名送数组
			Tag:           nil,           // 推送条件
			FastCustomTag: "",            // 使用用户标签筛选目标用户
		},
		Settings: &Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &Strategy{ // 厂商通道策略，具体看public_struct.go
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
		PushMessage: &PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: nil,
			Transmission: string(stringIos),
			Revoke:       nil,
		},
		PushChannel: &PushChannel{
			Ios: &iosChannel,
			Android: &AndroidChannel{Ups: &Ups{
				Notification: nil,
				TransMission: string(stringIos), // 透传消息内容，与notification 二选一
			}},
		},
	}
	result, err := single.PushSingleByCid(Ctx, Conf, Token, &singleParam)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, result.Code)

}

// 测试-单次推送ByCid 厂商通道+个推通道
func TestPushSingleByCidB(t *testing.T) {
	iosChannel := IosChannel{
		Type: "",
		Aps: &Aps{
			Alert: &Alert{
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
	notification := Notification{
		Title:       "卡是谁啊？",
		Body:        "为什么我们每天都要打 TA ？",
		ClickType:   "startapp", // 打开应用首页
		BadgeAddNum: 1,
	}

	singleParam := single.PushSingleParam{
		RequestId: strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
		Audience: &Audience{ // 目标用户
			Cid:           []string{Cid}, // cid推送数组
			Alias:         nil,           // 别名送数组
			Tag:           nil,           // 推送条件
			FastCustomTag: "",            // 使用用户标签筛选目标用户
		},
		Settings: &Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &Strategy{ // 厂商通道策略，具体看public_struct.go
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
		PushMessage: &PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: &notification,
			Transmission: "",
			Revoke:       nil,
		},
		PushChannel: &PushChannel{
			Ios: &iosChannel,
			Android: &AndroidChannel{Ups: &Ups{
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
