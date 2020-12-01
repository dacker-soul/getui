// 公共常量、结构体等
package getuiv2

const (
	ApiUrl = "https://restapi.getui.com/v2/" // 个推开放平台接口前缀(BaseUrl)

)

//var codeMsgList = map[int]string{
//	200: "成功",
//	400: "参数错误",
//	401: "权限相关",
//	403: "套餐相关",
//	404: "url错误",
//	405: "方法不支持",
//}

// 个推配置文件
type GeTuiConfig struct {
	AppId        string `toml:"app_id"`
	AppKey       string `toml:"app_key"`
	AppSecret    string `toml:"app_secret"`
	MasterSecret string `toml:"master_secret"`
}

// 公共返回结构
type PublicResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// clientId
type Cid struct {
	Cid []string `json:"cid"`
}

type Settings struct {
}

type PushMessage struct {
	Duration     string `json:"cid"` // 非必须，手机端通知展示时间段，格式为毫秒时间戳段，两个时间的时间差必须大于10分钟，例如："1590547347000-1590633747000"
	Notification        // 非必须，通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Transmission string `json:"transmission"` // 非必须，纯透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	revoke              // 非必须，撤回消息时使用，与notification、transmission三选一，都填写时报错
}

type PushChannel struct {
}
