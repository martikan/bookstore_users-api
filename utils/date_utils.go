package utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05"
	apiDbLayout   = "2006-01-02 15:04:05"
)

var (
	DateUtils dateUtilsInterface = &dateUtils{}
)

type dateUtils struct {
}

type dateUtilsInterface interface {
	GetNow() time.Time
	GetNowString() string
	GetNowDBFormat() string
}

func (d *dateUtils) GetNow() time.Time {
	return time.Now().UTC()
}

func (d *dateUtils) GetNowString() string {
	return DateUtils.GetNow().Format(apiDateLayout)
}

func (d *dateUtils) GetNowDBFormat() string {
	return DateUtils.GetNow().Format(apiDbLayout)
}
