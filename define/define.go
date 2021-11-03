package define

import "strconv"

const (
	ReqHeaderUserID    = "X-User-Id"
	ReqHeaderUserToken = "X-User-Token"
)

const (
	RespHeaderHostname = "X-Machine-Hostname"
	RespHeaderIP       = "X-Machine-Ip"
)

var (
	BinaryVersion         string // git hash，通过 build 注入
	BinaryBuildTime       int64  // 构建时间，通过 build 注入
	binaryBuildTimeString string // 只能注入变量到 string 类型，所以用一个 string 来中转
)

func init() {
	BinaryBuildTime, _ = strconv.ParseInt(binaryBuildTimeString, 10, 64)
}
