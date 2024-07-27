package checkpoints

// AllCheckpoints all check points list
var AllCheckpoints = []*CheckpointDefinition{
	{
		Name:        "通用请求Header长度限制",
		Prefix:      "requestGeneralHeaderLength",
		Description: "通用Header比如Cache-Control、Accept之类的长度限制，防止缓冲区溢出攻击",
		HasParams:   false,
		Instance:    new(RequestGeneralHeaderLengthCheckpoint),
		Priority:    100,
	},
	{
		Name:        "客户端地址（IP）",
		Prefix:      "remoteAddr",
		Description: "试图通过分析X-Forwarded-For等Header获取的客户端地址，比如192.168.1.100",
		HasParams:   false,
		Instance:    new(RequestRemoteAddrCheckpoint),
		Priority:    100,
	},
	{
		Name:        "客户端源地址（IP）",
		Prefix:      "rawRemoteAddr",
		Description: "直接连接的客户端地址，比如192.168.1.100",
		HasParams:   false,
		Instance:    new(RequestRawRemoteAddrCheckpoint),
		Priority:    100,
	},
	{
		Name:        "客户端端口",
		Prefix:      "remotePort",
		Description: "直接连接的客户端地址端口",
		HasParams:   false,
		Instance:    new(RequestRemotePortCheckpoint),
		Priority:    100,
	},
	{
		Name:        "客户端用户名",
		Prefix:      "remoteUser",
		Description: "通过BasicAuth登录的客户端用户名",
		HasParams:   false,
		Instance:    new(RequestRemoteUserCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求URI",
		Prefix:      "requestURI",
		Description: "包含URL参数的请求URI，类似于 /hello/world?lang=go",
		HasParams:   false,
		Instance:    new(RequestURICheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求路径",
		Prefix:      "requestPath",
		Description: "不包含URL参数的请求路径，类似于 /hello/world",
		HasParams:   false,
		Instance:    new(RequestPathCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求URL",
		Prefix:      "requestURL",
		Description: "完整的请求URL，包含协议、域名、请求路径、参数等，类似于 https://example.com/hello?name=lily",
		HasParams:   false,
		Instance:    new(RequestURLCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求内容长度",
		Prefix:      "requestLength",
		Description: "请求Header中的Content-Length",
		HasParams:   false,
		Instance:    new(RequestLengthCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求体内容",
		Prefix:      "requestBody",
		Description: "通常在POST或者PUT等操作时会附带请求体，最大限制32M",
		HasParams:   false,
		Instance:    new(RequestBodyCheckpoint),
		Priority:    5,
	},
	{
		Name:        "请求URI和请求体组合",
		Prefix:      "requestAll",
		Description: "${requestURI}和${requestBody}组合",
		HasParams:   false,
		Instance:    new(RequestAllCheckpoint),
		Priority:    5,
	},
	{
		Name:        "请求表单参数",
		Prefix:      "requestForm",
		Description: "获取POST或者其他方法发送的表单参数，最大请求体限制32M",
		HasParams:   true,
		Instance:    new(RequestFormArgCheckpoint),
		Priority:    5,
	},
	{
		Name:        "上传文件",
		Prefix:      "requestUpload",
		Description: "获取POST上传的文件信息，最大请求体限制32M",
		HasParams:   true,
		Instance:    new(RequestUploadCheckpoint),
		Priority:    20,
	},
	{
		Name:        "请求JSON参数",
		Prefix:      "requestJSON",
		Description: "获取POST或者其他方法发送的JSON，最大请求体限制32M，使用点（.）符号表示多级数据",
		HasParams:   true,
		Instance:    new(RequestJSONArgCheckpoint),
		Priority:    5,
	},
	{
		Name:        "请求方法",
		Prefix:      "requestMethod",
		Description: "比如GET、POST",
		HasParams:   false,
		Instance:    new(RequestMethodCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求协议",
		Prefix:      "scheme",
		Description: "比如http或https",
		HasParams:   false,
		Instance:    new(RequestSchemeCheckpoint),
		Priority:    100,
	},
	{
		Name:        "HTTP协议版本",
		Prefix:      "proto",
		Description: "比如HTTP/1.1",
		HasParams:   false,
		Instance:    new(RequestProtoCheckpoint),
		Priority:    100,
	},
	{
		Name:        "主机名",
		Prefix:      "host",
		Description: "比如teaos.cn",
		HasParams:   false,
		Instance:    new(RequestHostCheckpoint),
		Priority:    100,
	},
	{
		Name:        "CNAME",
		Prefix:      "cname",
		Description: "当前网站服务CNAME，比如38b48e4f.goedge.cloud",
		HasParams:   false,
		Instance:    new(RequestCNAMECheckpoint),
		Priority:    100,
	},
	{
		Name:        "是否为CNAME",
		Prefix:      "isCNAME",
		Description: "是否为CNAME，值为1（是）或0（否）",
		HasParams:   false,
		Instance:    new(RequestIsCNAMECheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求来源",
		Prefix:      "refererOrigin",
		Description: "请求报头中的Referer或Origin值",
		HasParams:   false,
		Instance:    new(RequestRefererOriginCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求来源Referer",
		Prefix:      "referer",
		Description: "请求Header中的Referer值",
		HasParams:   false,
		Instance:    new(RequestRefererCheckpoint),
		Priority:    100,
	},
	{
		Name:        "客户端信息",
		Prefix:      "userAgent",
		Description: "比如Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103",
		HasParams:   false,
		Instance:    new(RequestUserAgentCheckpoint),
		Priority:    100,
	},
	{
		Name:        "内容类型",
		Prefix:      "contentType",
		Description: "请求Header的Content-Type",
		HasParams:   false,
		Instance:    new(RequestContentTypeCheckpoint),
		Priority:    100,
	},
	{
		Name:        "所有cookie组合字符串",
		Prefix:      "cookies",
		Description: "比如sid=IxZVPFhE&city=beijing&uid=18237",
		HasParams:   false,
		Instance:    new(RequestCookiesCheckpoint),
		Priority:    100,
	},
	{
		Name:        "单个cookie值",
		Prefix:      "cookie",
		Description: "单个cookie值",
		HasParams:   true,
		Instance:    new(RequestCookieCheckpoint),
		Priority:    100,
	},
	{
		Name:        "所有URL参数组合",
		Prefix:      "args",
		Description: "比如name=lu&age=20",
		HasParams:   false,
		Instance:    new(RequestArgsCheckpoint),
		Priority:    100,
	},
	{
		Name:        "单个URL参数值",
		Prefix:      "arg",
		Description: "单个URL参数值",
		HasParams:   true,
		Instance:    new(RequestArgCheckpoint),
		Priority:    100,
	},
	{
		Name:        "所有Header信息",
		Prefix:      "headers",
		Description: "使用\\n隔开的Header信息字符串",
		HasParams:   false,
		Instance:    new(RequestHeadersCheckpoint),
		Priority:    100,
	},
	{
		Name:        "所有请求报头名称",
		Prefix:      "headerNames",
		Description: "使用换行符（\\n）隔开的报头名称字符串，每行一个名称",
		HasParams:   false,
		Instance:    new(RequestHeaderNamesCheckpoint),
		Priority:    100,
	},
	{
		Name:        "单个报头值",
		Prefix:      "header",
		Description: "单个报头值",
		HasParams:   true,
		Instance:    new(RequestHeaderCheckpoint),
		Priority:    100,
	},
	{
		Name:        "请求报头最大长度",
		Prefix:      "headerMaxLength",
		Description: "最长的请求报头的长度。",
		HasParams:   false,
		Instance:    new(RequestHeaderMaxLengthCheckpoint),
		Priority:    100,
	},
	{
		Name:        "国家/地区名称",
		Prefix:      "geoCountryName",
		Description: "国家/地区名称",
		HasParams:   false,
		Instance:    new(RequestGeoCountryNameCheckpoint),
		Priority:    90,
	},
	{
		Name:        "省份名称",
		Prefix:      "geoProvinceName",
		Description: "中国省份名称",
		HasParams:   false,
		Instance:    new(RequestGeoProvinceNameCheckpoint),
		Priority:    90,
	},
	{
		Name:        "城市名称",
		Prefix:      "geoCityName",
		Description: "中国城市名称",
		HasParams:   false,
		Instance:    new(RequestGeoCityNameCheckpoint),
		Priority:    90,
	},
	{
		Name:        "ISP名称",
		Prefix:      "ispName",
		Description: "ISP名称",
		HasParams:   false,
		Instance:    new(RequestISPNameCheckpoint),
		Priority:    90,
	},
	{
		Name:        "CC统计（旧）",
		Prefix:      "cc",
		Description: "统计某段时间段内的请求信息",
		HasParams:   true,
		Instance:    new(CCCheckpoint),
		Priority:    10,
	},
	{
		Name:        "CC统计（新）",
		Prefix:      "cc2",
		Description: "统计某段时间段内的请求信息",
		HasParams:   true,
		Instance:    new(CC2Checkpoint),
		Priority:    10,
	},
	{
		Name:        "防盗链",
		Prefix:      "refererBlock",
		Description: "阻止一些域名访问引用本站资源",
		HasParams:   true,
		Instance:    new(RequestRefererBlockCheckpoint),
		Priority:    20,
	},
	{
		Name:        "通用响应Header长度限制",
		Prefix:      "responseGeneralHeaderLength",
		Description: "通用Header比如Cache-Control、Accept之类的长度限制，防止缓冲区溢出攻击",
		HasParams:   false,
		Instance:    new(ResponseGeneralHeaderLengthCheckpoint),
		Priority:    100,
	},
	{
		Name:        "响应状态码",
		Prefix:      "status",
		Description: "响应状态码，比如200、404、500",
		HasParams:   false,
		Instance:    new(ResponseStatusCheckpoint),
		Priority:    100,
	},
	{
		Name:        "响应Header",
		Prefix:      "responseHeader",
		Description: "响应Header值",
		HasParams:   true,
		Instance:    new(ResponseHeaderCheckpoint),
		Priority:    100,
	},
	{
		Name:        "响应内容",
		Prefix:      "responseBody",
		Description: "响应内容字符串",
		HasParams:   false,
		Instance:    new(ResponseBodyCheckpoint),
		Priority:    5,
	},
	{
		Name:        "响应内容长度",
		Prefix:      "bytesSent",
		Description: "响应内容长度，通过响应的Header Content-Length获取",
		HasParams:   false,
		Instance:    new(ResponseBytesSentCheckpoint),
		Priority:    100,
	},
}

// FindCheckpoint find a check point
func FindCheckpoint(prefix string) CheckpointInterface {
	for _, def := range AllCheckpoints {
		if def.Prefix == prefix {
			def.Instance.SetPriority(def.Priority)
			return def.Instance
		}
	}
	return nil
}

// FindCheckpointDefinition find a check point definition
func FindCheckpointDefinition(prefix string) *CheckpointDefinition {
	for _, def := range AllCheckpoints {
		if def.Prefix == prefix {
			return def
		}
	}
	return nil
}
