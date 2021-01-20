// 公共常量、结构体等
package publics

const (
	ApiUrl = "https://restapi.getui.com/v2/" // 个推开放平台接口前缀(BaseUrl)
)

// 个推配置文件
type GeTuiConfig struct {
	AppId        string `toml:"app_id"`
	AppKey       string `toml:"app_key"`
	AppSecret    string `toml:"app_secret"`
	MasterSecret string `toml:"master_secret"`
}

/*
 * 公共返回结构示例
 * {
 *	 "code": 0,
 *	 "msg": "",
 *	 "data": {
 *		 "$taskid": {
 *			 "$cid":"$status"
 *		 }
 * 	 }
 * }
 *
 * {
 *	 "code": 0,
 *	 "msg": "",
 *	 "data": {
 *		 "$taskid": ""
 * 	 }
 * }
 *
 * $taskid:任务编号
 * $cid:App的用户唯一标识
 * $status:
 * successed_offline: 离线下发(包含厂商通道下发)
 * successed_online: 在线下发
 * successed_ignore: 最近90天内不活跃用户不下发
 */
type PublicResult struct {
	Code int    `json:"code"` // code返回码，0为成功，其他请看http://docs.getui.com/getui/server/rest_v2/code/?id=doc-title-1
	Msg  string `json:"msg"`
}

// 推送目标用户
type Audience struct {
	Cid           []string `json:"cid,omitempty"`             // cid数组，单推只能填一个cid，批量推可以填写多个（数组长度小于200）
	Alias         []string `json:"alias,omitempty"`           // 别名数组，单推只能填一个别名，批量推可以填写多个（数组长度小于200）
	Tag           *[]Tag   `json:"tag,omitempty"`             // 推送条件
	FastCustomTag string   `json:"fast_custom_tag,omitempty"` // 使用用户标签筛选目标用户
}

type Tag struct {
	Key     string   `json:"key"`      // 必须字段，默认值：无，查询条件 phone_type 手机类型; region 省市; custom_tag 用户标签; portrait 个推用户画像使用编码
	Values  []string `json:"values"`   // 必须字段，默认值：无，查询条件值列表，其中 手机型号使用android和ios； 省市使用编号，
	OptType string   `json:"opt_type"` // 必须字段，默认值：无，or(或),and(与),not(非)，values间的交并补操作
	// 不同key之间是交集，同一个key之间是根据opt_type操作
	// eg. 需要发送给城市在A,B,C里面，没有设置tagtest标签，手机型号为android的用户，用条件交并补功能可以实现，city(A|B|C) && !tag(tagtest) && phonetype(android)
}

// 推送条件设置
type Settings struct {
	TTL          int64     `json:"ttl,omitempty"`           // 非必须，默认一小时，消息离线时间设置，单位毫秒，-1表示不设离线，-1 ～ 3 * 24 * 3600 * 1000(3天)之间
	Strategy     *Strategy `json:"strategy,omitempty"`      // 非必须，默认值：{"strategy":{"default":1}}，厂商通道策略
	Speed        int       `json:"speed,omitempty"`         // 非必须，定速推送，例如100，个推控制下发速度在100条/秒左右，0表示不限速
	ScheduleTime int64     `json:"schedule_time,omitempty"` // 非必须，定时推送时间，必须是7天内的时间，格式：毫秒时间戳
}

// 厂商通道策略
type Strategy struct {
	Default int `json:"default,omitempty"`
	/*
	 * default字段，非必须，默认值为 1
	 * 默认所有通道的策略选择1-4
	 * 1: 表示该消息在用户在线时推送个推通道，用户离线时推送厂商通道;
	 * 2: 表示该消息只通过厂商通道策略下发，不考虑用户是否在线;
	 * 3: 表示该消息只通过个推通道下发，不考虑用户是否在线；
	 * 4: 表示该消息优先从厂商通道下发，若消息内容在厂商通道代发失败后会从个推通道下发。
	 * 其中名称可填写: ios、st、hw、xm、vv、mz、op，
	 */
	Ios int `json:"ios,omitempty"` // 非必须，ios通道策略1-4，表示含义同上，要推送ios通道，需要在个推开发者中心上传ios证书，建议填写2或4，否则可能会有消息不展示的问题
	St  int `json:"st,omitempty"`  // 非必须，通道策略1-4，表示含义同上，需要开通st厂商使用该通道推送消息
	Hw  int `json:"hw,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Xm  int `json:"xm,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Vv  int `json:"vv,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Mz  int `json:"mz,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Op  int `json:"op,omitempty"`  // 非必须，通道策略1-4，表示含义同上
}

// 个推推送消息参数
type PushMessage struct {
	Duration string `json:"duration,omitempty"`
	/*
	 * duration字段，非必须
	 * 手机端通知展示时间段，格式为毫秒时间戳段，两个时间的时间差必须大于10分钟
	 * 例如："1590547347000-1590633747000"
	 */
	Notification *Notification `json:"notification,omitempty"` // 非必须，通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Transmission string        `json:"transmission,omitempty"`
	/*
	 * transmission字段，非必须，【建议选这个模式】
	 * 穿透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	 */
	Revoke *Revoke `json:"revoke,omitempty"` // 非必须，撤回消息时使用，与notification、transmission三选一，都填写时报错
}

// 通知消息内容，仅支持安卓系统(不建议选这个，建议用穿透模板【transmission】)
type Notification struct {
	Title        string `json:"title"`        // 必须，通知消息标题，长度 ≤ 50
	Body         string `json:"body"`         // 必须，通知消息内容，长度 ≤ 256
	BigText      string `json:"big_text"`     // 非必须，长文本消息内容，通知消息+长文本样式，与big_image二选一，两个都填写时报错，长度 ≤ 512
	BigImage     string `json:"big_image"`    // 非必须，大图的URL地址，通知消息+大图样式， 与big_text二选一，两个都填写时报错，长度 ≤ 1024
	Logo         string `json:"logo"`         // 非必须，通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png”，长度 ≤ 64
	LogoUrl      string `json:"logo_url"`     // 非必须，通知图标URL地址，长度 ≤ 256
	ChannelId    string `json:"channel_id"`   // 非必须，默认值：Default，通知渠道id，长度 ≤ 64
	ChannelName  string `json:"channel_name"` // 非必须，默认值：Default，通知渠道名称，长度 ≤ 64
	ChannelLevel int    `json:"channel_level,omitempty"`
	/*
	 * channel_level,非必须，默认值：3
	 * 设置通知渠道重要性（可以控制响铃，震动，浮动，闪灯等等）
	 * android8.0以下
	 * 0，1，2:无声音，无振动，不浮动
	 * 3:有声音，无振动，不浮动
	 * 4:有声音，有振动，有浮动
	 * android8.0以上
	 * 0：无声音，无振动，不显示；
	 * 1：无声音，无振动，锁屏不显示，通知栏中被折叠显示，导航栏无logo;
	 * 2：无声音，无振动，锁屏和通知栏中都显示，通知不唤醒屏幕;
	 * 3：有声音，无振动，锁屏和通知栏中都显示，通知唤醒屏幕;
	 * 4：有声音，有振动，亮屏下通知悬浮展示，锁屏通知以默认形式展示且唤醒屏幕;
	 */
	ClickType string `json:"click_type,omitempty"`
	/*
	 * click_type,必须，默认值：无
	 * 点击通知后续动作，目前支持以下后续动作：
	 * intent：打开应用内特定页面
	 * url：打开网页地址
	 * payload：自定义消息内容启动应用
	 * payload_custom：自定义消息内容不启动应用
	 * startapp：打开应用首页
	 * none：纯通知，无后续动作
	 */
	Intent string `json:"intent,omitempty"`
	/*
	 * click_type为intent时必填
	 * 点击通知打开应用特定页面，长度 ≤ 2048
	 * 示例：intent:#Intent;component=你的包名/你要打开的 activity 全路径;S.parm1=value1;S.parm2=value2;end
	 * 如何生成：https://github.com/GetuiLaboratory/getui-pushapi-java-demo/blob/master/intent%E7%94%9F%E6%88%90%E5%8F%82%E8%80%83%E7%A4%BA%E4%BE%8B.md
	 */
	Url         string `json:"url,omitempty"`       // click_type为url时必填,点击通知打开链接，长度 ≤ 1024
	PayLoad     string `json:"payload,omitempty"`   // click_type为payload/payload_custom时必填,点击通知加自定义消息，长度 ≤ 3072
	NotifyId    int64  `json:"notify_id,omitempty"` // 非必须，覆盖任务时会使用到该字段，两条消息的notify_id相同，新的消息会覆盖老的消息，范围：0-2147483647
	RingName    string `json:"ring_name,omitempty"` // 非必须，自定义铃声，请填写文件名，不包含后缀名(需要在客户端开发时嵌入)，个推通道下发有效 客户端SDK最低要求 2.14.0.0
	BadgeAddNum int    `json:"badge_add_num,omitempty"`
	/*
	 * badge_add_num,非必须，默认值：无
	 * 角标, 必须大于0, 个推通道下发有效
	 * 此属性目前仅针对华为 EMUI 4.1 及以上设备有效
	 * 角标数字数据会和之前角标数字进行叠加，也就是相加；
	 * 举例：角标数字配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3。
	 * 客户端SDK最低要求 2.14.0.0
	 */

	// options为push_channel厂商通道中安卓专有
	Options *Options `json:"options,omitempty"` // 第三方厂商通知扩展内容

}

type Options struct {
	Constraint string `json:"constraint,omitempty"` // 非必须，扩展内容对应厂商通道设置如：HW,MZ,...
	Key        string `json:"key,omitempty"`
	/*
	 * Key,必须，默认值：无
	 * 厂商内容扩展字段,单个厂商特有字段
	 * key目前支持的所有字段：
	 * hw角标设置：badgeAddNum
	 * badgeClass要设置hw角标，这两个字段需要配合使用 ，hw的icon，op私信 channel，op的消息去重 app_meaasge_id
	 * vv的消息分类classification， 0 代表运营消息，1代表系统消息，不填默认为0
	 * xm的channel:目前只有op和xm支持
	 */
	Value string `json:"value,omitempty"`
	/*
	 * value,必须，默认值：无
	 * value的设置根据key值决定。例如
	 * hw角标设置：key设为badgeAddNum，value：1（原来的角标值+1）
	 * key设为badgeClass，value：请写入角标设置的应用类名）
	 * key设为icon，value：请写⼊对应图标地址
	 */
}

// 撤回消息时使用，与notification、transmission三选一，都填写时报错
type Revoke struct {
	OldTaskId string `json:"old_task_id"`     // 必须，需要撤回的taskId
	Force     bool   `json:"force,omitempty"` // 非必须，【小心使用】在没有找到对应的taskId，是否把对应appId下所有的通知都撤回
}

// 厂商推送消息参数，包含ios消息参数，android厂商消息参数
type PushChannel struct {
	Ios     *IosChannel     `json:"ios,omitempty"`     // 非必须，ios通道推送消息内容
	Android *AndroidChannel `json:"android,omitempty"` // 非必须，android通道推送消息内容
}

// ios厂商通道消息
type IosChannel struct {
	// 具体参数含义详见苹果APNs文档
	// https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html
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

// 通知消息
type Alert struct {
	Title           string   `json:"title,omitempty"`             // 非必须，通知消息标题
	Body            string   `json:"body,omitempty"`              // 非必须，通知消息内容
	ActionLocKey    string   `json:"action-loc-key,omitempty"`    // 非必须，（用于多语言支持）指定执行按钮所使用的Localizable.strings
	LocKey          string   `json:"loc-key,omitempty"`           // 非必须，（用于多语言支持）指定Localizable.strings文件中相应的key
	LocArgs         []string `json:"loc-args,omitempty"`          // 非必须，如果loc-key中使用了占位符，则在loc-args中指定各参数
	LaunchImage     string   `json:"launch-image,omitempty"`      // 非必须，指定启动界面图片名
	TitleLocKey     string   `json:"title-loc-key,omitempty"`     // 非必须，(用于多语言支持）对于标题指定执行按钮所使用的Localizable.strings,仅支持iOS8.2以上版本
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`    // 非必须，对于标题,如果loc-key中使用的占位符，则在loc-args中指定各参数,仅支持iOS8.2以上版本
	SubTitle        string   `json:"sub_title,omitempty"`         // 非必须，通知子标题,仅支持iOS8.2以上版本
	SubTitleLocKey  string   `json:"subtitle-loc-key,omitempty"`  // 非必须，当前本地化文件中的子标题字符串的关键字,仅支持iOS8.2以上版本
	SubTitleLocArgs []string `json:"subtitle-loc-args,omitempty"` // 非必须，当前本地化子标题内容中需要置换的变量参数 ,仅支持iOS8.2以上版本
}

// 多媒体设置,最多可设置3个子项
type Multimedia struct {
	Url      string `json:"url"`                 // 必须，多媒体资源地址
	Type     int    `json:"type"`                // 必须，资源类型（1.图片，2.音频，3.视频）
	OnlyWifi bool   `json:"only_wifi,omitempty"` // 非必须，是否只在wifi环境下加载，如果设置成true,但未使用wifi时，会展示成普通通知
}

// Android厂商通道消息
type AndroidChannel struct {
	Ups *Ups `json:"ups"` // android厂商通道推送消息内容
}

// Android厂商通道推送消息内容
type Ups struct {
	Notification *Notification `json:"notification"` // 非必须,通知消息内容，与transmission 二选一，两个都填写时报错
	TransMission string        `json:"transmission"` // 非必须,透传消息内容，与notification 二选一，两个都填写时报错，长度 ≤ 3072
}
