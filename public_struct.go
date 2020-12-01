package getuiv2

const (
	API_URL      = "https://restapi.getui.com/v2/" // 个推开放平台接口前缀(BaseUrl)
	CODE_SUCCESS = 200
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
