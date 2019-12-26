package face

/*特征信息*/
type Feature struct {
	//特征id
	Id string `json:"id"`
	//模型id
	DataSetId string `json:"dataSetId"`
	//schema名称
	Name string `json:"name"`
	//中文名称
	Remark string `json:"remark"`
	//类型
	Type string `json:"type"`
	//是否输出字段
	Hidden bool `json:"hidden"`
	//计算公式
	Formula string `json:"formula"`
	//计算值
	Value interface{} `json:"value"`
	//格式 例如:日期
	Formate string `json:"formate"`
	//公式条件
	Condition string `json:"condition"`
	//过滤字段
	FilterField string `json:"filterField"`
	//值字段
	FetchField string `json:"fetchField"`
	//数学计算 加减乘除
	Compute string `json:"compute"`
	//默认值
	DefaultValue string `json:"defaultValue"`
	//默认值条件
	DefaultCondition string `json:"defaultCondition"`
}
