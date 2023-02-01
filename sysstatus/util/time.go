package util

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

var (
	AppStartTime time.Time
)

func init() {
	AppStartTime = time.Now().UTC()
}

// Tr translates content to locale language. fall back to default language.
func Tr(trKey string, trArgs ...interface{}) string {
	msg, err := Format(trKey, trArgs...)
	if err != nil {
		log.Printf("Error whilst formatting %q : %v\n", trKey, err)
	}
	return msg
}

func TimeSincePro(then, now time.Time) string {
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return Tr("将来")
	}
	if diff == 0 {
		return Tr("现在")
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiffFloor(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Format formats provided arguments for a given translated message
func Format(format string, args ...interface{}) (msg string, err error) {
	if len(args) == 0 {
		return format, nil
	}

	fmtArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		val := reflect.ValueOf(arg)
		if val.Kind() == reflect.Slice {
			// Previously, we would accept Tr(lang, key, a, [b, c], d, [e, f]) as Sprintf(msg, a, b, c, d, e, f)
			// but this is an unstable behavior.
			//
			// So we restrict the accepted arguments to either:
			//
			// 1. Tr(lang, key, [slice-items]) as Sprintf(msg, items...)
			// 2. Tr(lang, key, args...) as Sprintf(msg, args...)
			if len(args) == 1 {
				for i := 0; i < val.Len(); i++ {
					fmtArgs = append(fmtArgs, val.Index(i).Interface())
				}
			} else {
				err = errors.New("arguments error")
				break
			}
		} else {
			fmtArgs = append(fmtArgs, arg)
		}
	}
	return fmt.Sprintf(format, fmtArgs...), err
}

func computeTimeDiffFloor(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "现在"
	case diff < 2:
		diff = 0
		diffStr = "1 秒"
	case diff < 1*Minute:
		diffStr = Tr("%d 秒", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = Tr("1分钟")
	case diff < 1*Hour:
		diffStr = Tr("%d 分钟", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = Tr("1 小时")
	case diff < 1*Day:
		diffStr = Tr("%d 小时", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = Tr("1 天")
	case diff < 1*Week:
		diffStr = Tr("%d 天", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = Tr("1 周")
	case diff < 1*Month:
		diffStr = Tr("%d 周", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = Tr("1 个月")
	case diff < 1*Year:
		diffStr = Tr("%d 个月", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = Tr("1 年")
	default:
		diffStr = Tr("%d 年", diff/Year)
		diff -= (diff / Year) * Year
	}
	return diff, diffStr
}
