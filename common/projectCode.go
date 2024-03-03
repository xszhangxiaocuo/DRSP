package common

/*业务模块错误代码*/

var (
	FileIsNull    = NewError(500001, "文件为空")
	DataParseFail = NewError(500002, "数据解析失败")
)
