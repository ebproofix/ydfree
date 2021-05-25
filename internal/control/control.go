package control

import (
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func CheckOnTimeOut(dir string) error {
	b := path.Base(dir)
	if !(utf8.RuneCountInString(b) > 9) {
		return errors.New(
			fmt.Sprintf("меньше 10 символов"),
			)
	}

	// dayStr := string(b[0]) + string(b[1])
	// mountStr := string(b[3]) + string(b[4])
	// yarStr := string(b[6]) + string(b[7]) + string(b[8]) + string(b[9])
	dayStr := string(b[:2])
	mountStr := string(b[3:4])
	yarStr := string(b[6:10])


	day, _ := strconv.ParseInt(dayStr, 10, 32)
	mount, _ := strconv.ParseInt(mountStr, 10, 32)
	yar, _ := strconv.ParseInt(yarStr, 10, 32)

	chiclo := time.Date(int(yar), time.Month(mount), int(day), 0, 0, 0, 0, time.UTC)
	diff := chiclo.Sub(time.Now())
	// 2190 - часа в 3х месяцах
	if !(diff.Hours() > float64(2190 * time.Hour)){
		return errors.New(
			fmt.Sprintf("Не прошло 3 месяца (2190 часа)"),
			)
	}
	return nil
}


func CheckOnNotDeleteWord(dir string) error {
	base := strings.ToLower(path.Base(dir))
	if strings.Contains(base, "не удалять"){
		return errors.New(
			fmt.Sprintf("Найдено \"не удалять\""),
			)
	}
	return nil
}


func CheckOnbackup(dir string, listBackups []string) error {
	base := path.Base(dir)
	for _, v := range listBackups{
		if base == v {
			return nil
		}
	}
	return errors.New(
		fmt.Sprintf("Нет бекапа"),
		)
}
