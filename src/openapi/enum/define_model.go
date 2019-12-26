package enum

const (
	ConsumerLoan = "busi_type"

	CarSecondHand = "car_type"
)

var oldCardLoanMap = map[string]string{
	ConsumerLoan:  "1", //消费贷
	CarSecondHand: "2", //二手车
}

/*获取初始化值*/
func FetcholdCardCode(keycode string) string {

	return oldCardLoanMap[keycode]

}
