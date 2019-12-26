package openapi

import (
	"github.com/gin-gonic/gin"
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/config"
	"go-learn/repository/sql"
	"net/http"
)

func (apiModel *ApiModel) Prediction(engine *gin.RouterGroup) {

	engine.POST("/prediction", apiModel.predict)

}

/*分数请求信息*/
type PredictionRequest struct {
	SysId string `json:"sysId" binding:"required"`

	OrderId string `json:"orderId" binding:"required"`

	TransId string `json:"transId" binding:"required"`

	ModelId string `json:"modelId" binding:"required"`

	User interface{} `json:"user"`
}

func (apiModel *ApiModel) predict(context *gin.Context) {

	var scoreReq PredictionRequest

	err := context.ShouldBindJSON(&scoreReq)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {

		modelFeature := config.ModelFeatureMap[scoreReq.ModelId]

		var apiResultInfo ApiResultInfo

		busiInfo := &sql.FintechBusinessInfo{
			OrderId: "test",
		}

		if tools.IsNotEmpty(modelFeature.ServiceName) {

			apiModel.Service.SqlTemplate.FetchEntity(busiInfo)

		} else {

			apiResultInfo = NewApiResult(define.ModelSericeNotExist, busiInfo, err)

		}

		context.JSON(http.StatusOK, apiResultInfo)

	}

}
