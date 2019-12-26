package openapi

import (
	"github.com/gin-gonic/gin"
	"go-learn/common/define"
	"go-learn/config"
	"net/http"
)

func (apiModel *ApiModel) Features(engine *gin.RouterGroup) {

	engine.POST("/features", apiModel.processFeatures)

}

/*特征请求信息*/
type FeaturesRequest struct {
	SysId string `json:"sysId" binding:"required"`

	OrderId string `json:"orderId" binding:"required"`

	TransId string `json:"transId" binding:"required"`

	ModelId string `json:"modelId" binding:"required"`
}

func (apiModel *ApiModel) processFeatures(context *gin.Context) {

	var featureRequest FeaturesRequest

	err := context.ShouldBindJSON(&featureRequest)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {

		var featureResultInfo ApiResultInfo

		modelInfo, ok := config.ModelFeatureMap[featureRequest.ModelId]

		if ok {

			featureResultInfo = NewApiResult(define.ServiceSucess, modelInfo, nil)

		} else {

			featureResultInfo = NewApiResult(define.ModelNotExist, nil, err)

		}

		context.JSON(http.StatusOK, featureResultInfo)

	}

}
