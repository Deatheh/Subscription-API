package pkg

import (
	"fmt"
	"time"
)

func IsValidMonthYear(s string) bool {
	_, err := time.Parse("01-2006", s)
	return err == nil
}

func MonthYearToDate(s string) (string, error) {
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02"), nil
}

func DateToMonthYear(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
	}
	var t time.Time
	var err error
	for _, layout := range formats {
		t, err = time.Parse(layout, s)
		if err == nil {
			return t.Format("01-2006"), nil
		}
	}
	return "", fmt.Errorf("не удалось распарсить дату: %s", s)
}

func IsValidMonthYearLength(startDate, endDate string) bool {
	startTime, _ := time.Parse("01-2006", startDate)
	endTime, _ := time.Parse("01-2006", endDate)
	if startTime.Before(endTime) {
		return true
	}
	return false
}
