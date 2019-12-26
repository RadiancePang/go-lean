package compute

import (
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/formula/face"
	"sort"
	"strconv"
)

type SortFieldArr []face.FetchFieldInfo

type SortFieldCompator struct {
	SortFieldArr

	Flag bool
}

func (p SortFieldArr) Len() int { return len(p) }

func (compator SortFieldCompator) Less(i, j int) bool {

	calvalue1 := tools.DefaultStr(compator.SortFieldArr[i].CalValue, define.MetaTypeDecimal)

	calvalue2 := tools.DefaultStr(compator.SortFieldArr[j].CalValue, define.MetaTypeDecimal)

	amount1, err1 := strconv.ParseFloat(calvalue1, 64)

	amount2, err2 := strconv.ParseFloat(calvalue2, 64)

	var result bool

	if err1 == nil && err2 == nil {

		if compator.Flag {

			return amount1 > amount2

		} else {

			return amount2 > amount1

		}

	}

	return result

}

func (p SortFieldArr) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func CalSortNumFetch(defalutValue string, params []face.FetchFieldInfo, flag bool) face.FetchFieldInfo {

	sortCompator := SortFieldCompator{

		Flag: flag,

		SortFieldArr: params,
	}

	sort.Sort(sortCompator)

	return sortCompator.SortFieldArr[0]

}
