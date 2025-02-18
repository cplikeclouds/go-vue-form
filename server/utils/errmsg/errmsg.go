package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
)

var CodeMsg = map[int]string{
	SUCCESS: "成功",
	ERROR:   "出错",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
