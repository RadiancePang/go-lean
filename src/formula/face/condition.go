package face

/*公式条件*/
type Condition struct {
	//条件字段
	Key string
	//条件值
	Value string
	//操作符(等于in等)
	Operator string
}
