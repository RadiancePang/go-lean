package define

/*错误代码定义*/
const (
	ServiceSucess        = 10000
	ServiceError         = 20000
	ServicePanic         = 20001
	ModelNotExist        = 30001
	ModelSericeNotExist  = 30002
	ModelFeatureError    = 30003
	ModelAlgorithmError  = 30004
	ModelDecodeError     = 30005
	ModelFileWriteError  = 30006
	ModelFeatureNotExist = 40001
)

var errorMsgMap = map[int]string{
	ServiceSucess:        "服务处理成功",
	ServiceError:         "服务内部错误",
	ServicePanic:         "服务panic错误",
	ModelNotExist:        "模型不存在",
	ModelSericeNotExist:  "模型服务不存在",
	ModelFeatureError:    "模型特征计算错误",
	ModelAlgorithmError:  "模型算法执行错误",
	ModelDecodeError:     "模型base64数据转码错误",
	ModelFileWriteError:  "模型文件写入错误",
	ModelFeatureNotExist: "模型特征不存在",
}

func ErrorMsg(code int) string {

	return errorMsgMap[code]

}
