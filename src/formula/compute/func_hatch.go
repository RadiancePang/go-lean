package compute

import (
	"go-learn/common/define"
	"go-learn/components/tools"
	"strconv"
)

type HatchField struct {
	FuncName string

	Args []string
}

func (hactchField HatchField) AgeLevel() int {

	var level int

	agestr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeInt)

	if tools.IsEmpty(hactchField.Args[0]) {

		agestr = "0"

	}

	age, err := strconv.Atoi(agestr)

	if err == nil {

		switch {

		case age == 0:
			level = 0
		case age <= 30:
			level = 1
		case age <= 38:
			level = 2
		case age <= 45:
			level = 3
		case age > 45:
			level = 4
		default:
			level = 0

		}

	}

	return level
}

func (hactchField HatchField) DownPayment() int {

	paymentStr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeDecimal)

	var level int

	var payment float64

	if tools.IsNotEmpty(paymentStr) {

		payment, _ = strconv.ParseFloat(paymentStr, 64)

	}

	switch {

	case payment < 0:
		level = -1
	case payment == 0:
		level = 0
	case payment < 10000:
		level = 1
	case payment < 15000:
		level = 2
	case payment < 20000:
		level = 3
	case payment < 30000:
		level = 4
	case payment < 40000:
		level = 5
	case payment < 50000:
		level = 6
	case payment < 100000:
		level = 7
	case payment < 500000:
		level = 8
	case payment >= 500000:
		level = 9
	default:
		level = 0

	}

	return level
}

func (hactchField HatchField) CustChannelCode() int {

	custRangeStr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeInt)

	var level int

	if tools.IsNotEmpty(custRangeStr) {

		level, _ = strconv.Atoi(custRangeStr)

	}

	level++

	return level

}

func (hactchField HatchField) CustRate() int {

	custRangeStr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeDecimal)

	var level int

	var custRange float64

	if tools.IsNotEmpty(custRangeStr) {

		custRange, _ = strconv.ParseFloat(custRangeStr, 64)

	}

	switch {

	case custRange < 0:
		level = -1
	case custRange == 0:
		level = 0
	case custRange < 10:
		level = 1
	case custRange < 12.5:
		level = 2
	case custRange < 15:
		level = 3
	case custRange < 20:
		level = 4
	case custRange < 50:
		level = 5
	case custRange >= 50:
		level = 6
	default:
		level = 0

	}

	return level
}

func (hactchField HatchField) ProvinceLevel() int {

	level := 5

	provinceStr := hactchField.Args[0]

	level1map := map[string]string{"24": "", "31": "", "2": "", "26": "", "20": "", "27": "", "13": "", "98": "", "99": ""}

	level2map := map[string]string{"30": "", "3": "", "8": "", "7": "", "12": "", "25": "", "10": "", "15": ""}

	level3map := map[string]string{"1": "", "23": "", "18": "", "19": "", "28": "", "40": ""}

	level4map := map[string]string{"29": "", "6": "", "5": "", "16": "", "9": "", "21": "", "11": "", "22": "", "14": "", "17": ""}

	if _, ok := level1map[provinceStr]; ok {

		level = 1

	} else if _, ok := level2map[provinceStr]; ok {

		level = 2

	} else if _, ok := level3map[provinceStr]; ok {

		level = 3

	} else if _, ok := level4map[provinceStr]; ok {

		level = 4

	}

	return level
}

func (hactchField HatchField) PaymentRate() int {

	paymentStr := tools.DefaultStr(hactchField.Args[0], "0")

	var level int

	var custRange float64

	if tools.IsNotEmpty(paymentStr) {

		custRange, _ = strconv.ParseFloat(paymentStr, 64)

	}

	switch {

	case custRange < 0:
		level = -1
	case custRange == 0:
		level = 0
	case custRange < 20:
		level = 1
	case custRange == 20:
		level = 2
	case custRange < 30:
		level = 3
	case custRange == 30:
		level = 4
	case custRange < 40:
		level = 5
	case custRange < 50:
		level = 6
	case custRange >= 50:
		level = 7
	default:
		level = 0

	}

	return level

}

func (hactchField HatchField) CompareProvinceCode() int {

	var level int

	if hactchField.Args[0] == hactchField.Args[1] {
		level = 1
	}
	return level
}

func (hactchField HatchField) StyleIdLevel() int {

	rateStr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeDecimal)

	var rate float64

	var level int

	if tools.IsNotEmpty(rateStr) {

		rate, _ = strconv.ParseFloat(rateStr, 64)

	}

	switch {

	case rate <= 100000:
		level = 0
	case rate <= 10000000:
		level = 1
	case rate <= 94000000:
		level = 2
	case rate <= 900000000000000:
		level = 3
	case rate > 900000000000000:
		level = 4
	default:
		level = 0

	}

	return level
}

func (hactchField HatchField) MonthPayLevel() int {

	monthPayStr := tools.DefaultStr(hactchField.Args[0], define.MetaTypeDecimal)

	var level int

	var monthPay float64

	if tools.IsNotEmpty(monthPayStr) {

		monthPay, _ = strconv.ParseFloat(monthPayStr, 64)

	}

	switch {

	case monthPay < 0:
		level = 0
	case monthPay == 0:
		level = 0
	case monthPay < 1000:
		level = 1
	case monthPay < 1500:
		level = 2
	case monthPay < 2000:
		level = 3
	case monthPay < 2500:
		level = 4
	case monthPay < 3000:
		level = 5
	case monthPay < 4000:
		level = 6
	case monthPay < 5000:
		level = 7
	case monthPay >= 5000:
		level = 8
	default:
		level = 0

	}

	return level
}
