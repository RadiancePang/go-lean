package compute

import ()

func ProcessFunc(funcName string, params ...interface{}) interface{} {

	var funcResult interface{}

	if tools.IsNotEmpty(funcName) && tools.IsFunc(funcName) {

		fetchFuncInfo := FetchFuncInfo(funcName)

		funcAargs := tools.SplitComma(fetchFuncInfo.Args)

		switch fetchFuncInfo.FuncName {

		case "diff_months":
			funcResult = DiffMonths(params[0], params[1], funcAargs)
		case "greater":
			funcResult = Greater(params[1], funcAargs)
		case "isOverDue":
			funcResult = IsOverDue(params[1], funcAargs)
		case "in":
			funcResult = ElementsIn(params[1], funcAargs)
		case "max":
			funcResult = Max(params[0])
		case "total":
			funcResult = Total(params[0])
		case "isNotEmpty":
			funcResult = tools.IsNotEmpty(tools.DefaultStr(params[1], define.MetaTypeString))

		default:

		}

	}

	return funcResult

}

func FetchFuncInfo(funcName string) face.FetchFuncInfo {

	funcInfos := tools.SplitFunc(funcName)

	fetchFuncInfo := face.FetchFuncInfo{}

	fetchFuncInfo.FuncName = funcInfos[0]

	fetchFuncInfo.Args = funcInfos[1]

	return fetchFuncInfo

}
