package model

/*特征信息*/
type ApiFeatureInfo struct {

	//特征id
	Id string `json:"id"`

	//特征备注
	Remark string `json:"remark"`

	//特征值
	Value string `json:"value"`
}
