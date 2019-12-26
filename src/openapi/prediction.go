package openapi

import (
	"github.com/gin-gonic/gin"
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/config"
	"net/http"
)

func Prediction(engine *gin.RouterGroup) {

	engine.POST("/prediction", predict)

}

/*分数请求信息*/
type PredictionRequest struct {
	SysId string `json:"sysId" binding:"required"`

	OrderId string `json:"orderId" binding:"required"`

	TransId string `json:"transId" binding:"required"`

	ModelId string `json:"modelId" binding:"required"`

	User interface{} `json:"user"`
}

func predict(context *gin.Context) {

	var scoreReq PredictionRequest

	err := context.ShouldBindJSON(&scoreReq)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {

		modelFeature := config.ModelFeatureMap[scoreReq.ModelId]

		var apiResultInfo ApiResultInfo

		if tools.IsNotEmpty(modelFeature.ServiceName) {

		} else {

			apiResultInfo = NewApiResult(define.ModelSericeNotExist, nil, err)

		}

		context.JSON(http.StatusOK, apiResultInfo)

	}

}
