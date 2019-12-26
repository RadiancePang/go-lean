package tools

import (
	"math"
	"strings"
	"time"
)

var Loc, _ = time.LoadLocation("Asia/Shanghai")

func TimeFormate(source string) string {

	source = strings.Replace(source, "yyyy", "2006", 1)

	source = strings.Replace(source, "MM", "01", 1)

	source = strings.Replace(source, "dd", "02", 1)

	source = strings.Replace(source, "HH", "15", 1)

	source = strings.Replace(source, "mm", "04", 1)

	source = strings.Replace(source, "ss", "05", 1)

	return source

}

func TimeNow() time.Time {
	return time.Now().In(Loc)
}

func PraseTime(param string, format string) (time.Time, error) {

	var err error

	var paramTIme time.Time

	param = strings.ReplaceAll(param, "-", ".")

	targetFormate := TimeFormate(format)

	paramTIme, err = time.ParseInLocation(targetFormate, param, Loc)

	return paramTIme, err

}

func DiffDays(sourceCalendar time.Time, targetCalendar time.Time) int64 {

	diffHours := sourceCalendar.Sub(targetCalendar).Hours() / 24

	return int64(diffHours)

}

func DiffMonths(sourceCalendar time.Time, targetCalendar time.Time, roundFlag bool) int64 {

	result := int(targetCalendar.Month() - sourceCalendar.Month())

	month := (targetCalendar.Year() - sourceCalendar.Year()) * 12

	diff := int64(math.Abs(float64(result + month)))

	if roundFlag {

		sourceDays := sourceCalendar.Day()

		targetDays := targetCalendar.Day()

		if sourceDays < targetDays {

			diff = diff - 1

		}

	}

	return diff

}

func DiffYears(sourceCalendar time.Time, targetCalendar time.Time) int {

	return sourceCalendar.Year() - targetCalendar.Year()

}
