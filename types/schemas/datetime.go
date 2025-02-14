package schemas

import (
	"time"
)

type Datetime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Datetime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Datetime(now)
	return
}

func (t Datetime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + time.Time(t).Format(timeFormart) + `"`), nil
}
