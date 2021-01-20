# 个推V2版本

## 前言
由于个推没有Go的SDK，因此写个

## 文件结构
```
getui
    |--auth         鉴权API：获取token，删除token
    |--publics      公共结构体，公共方法
    |--push         推送API
        |--single   单推送API：cid单推，别名单推，cid批量单推，别名批量单推
        |--list     批量推API：创建消息，cid批量推，别名批量推
        |--all      群推API：群推，根据条件筛选用户推送，使用标签快速推送
        |--mission  任务API：停止任务，查询定时任务，删除定时任务
    |--测试文件及其他

```
## 几个概念
推送都包括：推送参数，推送方法，推送返回结果,主要功能都在推送参数上面。
下面以单推为例子详细介绍（其他推送每个参数和方法里面都有详细注释）：

```
    singleParam := single.PushSingleParam{
        RequestId:   strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
        Audience:    &Audience{},       // 推送的目标用户
        Settings:    &Settings{},       // 推送条件：例如速度，定时推，厂商下发策略等
        PushMessage: &PushMessage{},    // 个推通道消息内容
        PushChannel: &PushChannel{      // 厂商通道消息内容
            Ios:     &IosChannel{},
            Android: &AndroidChannel{},
        },
    }
```
推送消息内容分为两种：个推通道消息内容（PushMessage）和厂商通道消息内容（PushChannel）

###个推通道消息内容（PushMessage）

PushMessage中的notification，transmission，revoke三选一，都填写报错

notification：仅支持安卓系统，iOS系统不展示个推通道下发的通知消息

transmission：纯透传消息内容，安卓和iOS均支持。纯透传消息是啥？假如你现在收到一条推送，你打开app
并且客户端配置了个推的sdk，你就可以收到一条弹窗消息。

revoke：撤回消息时使用(仅支持撤回个推通道消息，消息撤回请勿填写策略参数)

**总结一下：个推消息通道，普通消息推送仅支持安卓系统，ios不行。纯透传消息不分系统，都可以。**

###厂商通道消息内容（PushChannel）

先说ios，ios的推送用官方的[APNS](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html)

```
// ios厂商通道消息
type IosChannel struct {
	Type           string        `json:"type,omitempty"`             // 非必须，默认值：notify，voip：voip语音推送，notify：apns通知消息
	Aps            *Aps          `json:"aps,omitempty"`              // 推送通知消息内容
	AutoBadge      string        `json:"auto_badge,omitempty"`       // 非必须，用于计算icon上显示的数字，还可以实现显示数字的自动增减，如“+1”、 “-1”、 “1” 等，计算结果将覆盖badge
	PayLoad        string        `json:"payload,omitempty"`          // 非必须，增加自定义的数据
	Multimedia     *[]Multimedia `json:"multimedia,omitempty"`       // 非必须，该字段为Array类型,设置多媒体
	ApnsCollapseId string        `json:"apns-collapse-id,omitempty"` // 非必须，使用相同的apns-collapse-id可以覆盖之前的消息
}
// 推送通知消息内容
type Aps struct {
	Alert            *Alert `json:"alert"` // 非必须，通知消息
	ContentAvailable int    `json:"content-available"`
	/*
	 * content-available非必须，默认值：0
	 * 0表示普通通知消息(默认为0)
	 * 1表示静默推送(无通知栏消息)，静默推送时不需要填写其他参数
	 * 苹果建议1小时最多推送3条静默消息
	 */
	Sound    string `json:"sound,omitempty"`     // 非必须，通知铃声文件名，如果铃声文件未找到，响铃为系统默认铃声。 无声设置为“com.gexin.ios.silence”或不填
	Category string `json:"category,omitempty"`  // 非必须，在客户端通知栏触发特定的action和button显示
	ThreadId string `json:"thread-id,omitempty"` // 非必须，ios的远程通知通过该属性对通知进行分组，仅支持iOS 12.0以上版本
}
```



