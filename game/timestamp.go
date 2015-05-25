package game

import (
	"fmt"
	"strings"
	"time"
)

type Timestamp time.Time

const (
	time_format = time.RFC3339Nano
)

func NewTimestamp() *Timestamp {
	ts := Timestamp(time.Now())
	return &ts
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t)
	stamp := fmt.Sprintf("\"%s\"", ts.Format(time_format))
	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	bstr := strings.Replace(string(b), `"`, "", -1)
	ts, err := time.Parse(time_format, bstr)
	if err != nil {
		return err
	}
	*t = Timestamp(ts)
	return nil
}

func (t *Timestamp) String() string {
	return time.Time(*t).Format(time_format)
}
