package compute

import (
	"sort"
)

type SortDataArr []string

type SortDateCompator struct {
	SortDataArr

	Flag bool

	Formate string
}

func (p SortDataArr) Len() int { return len(p) }

func (compator SortDateCompator) Less(i, j int) bool {

	time1, err1 := tools.PraseTime(compator.SortDataArr[i], compator.Formate)

	time2, err2 := tools.PraseTime(compator.SortDataArr[j], compator.Formate)

	var result bool

	if err1 == nil && err2 == nil {

		if compator.Flag {

			return time1.Unix() > time2.Unix()

		} else {

			return time2.Unix() > time1.Unix()

		}

	}

	return result

}

func (p SortDataArr) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func SortDate(params []string, formate string, flag bool) string {

	sortCompator := SortDateCompator{

		Formate: formate,

		Flag: flag,

		SortDataArr: params,
	}

	sort.Sort(sortCompator)

	return sortCompator.SortDataArr[0]

}
