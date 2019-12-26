package openapi

import (
	"go-learn/common/define"
)

/*api 统一响应信息*/
type ApiResultInfo struct {
	Returncode int `json:"returncode"`

	Pvid string `json:"pvid"`

	Msg string `json:"msg"`

	Score int `json:"score"`

	Data interface{} `json:"data,omitempty"`
}

type EngineResponse struct {
	Code int `json:"code"`

	Msg string `json:"msg"`

	Data interface{} `json:"data,omitempty"`
}

/*创建api结果*/
func NewApiResult(code int, context interface{}, err error) ApiResultInfo {

	apiResultInfo := ApiResultInfo{}

	apiResultInfo.Returncode = code

	if err != nil {

		apiResultInfo.Returncode = define.ServiceError

		apiResultInfo.Msg = err.Error()

	} else {

		if code == define.ServiceSucess {

			apiResultInfo.Data = context

		}

		apiResultInfo.Msg = define.ErrorMsg(code)

	}

	return apiResultInfo

}
