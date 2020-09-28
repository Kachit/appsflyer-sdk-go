package appsflyer_sdk

import (
	"fmt"
	"strings"
	"time"
)

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) Value() float64 {
	return cf.Float64
}

type CustomTimestamp struct {
	Timestamp time.Time
}

func (ct *CustomTimestamp) Value() time.Time {
	return ct.Timestamp
}

func (ct *CustomTimestamp) UnmarshalCSV(csv string) error {
	if csv != "" {
		var err error
		ct.Timestamp, err = time.Parse("2006-01-02 15:04:05", csv)
		if err != nil {
			return fmt.Errorf("CustomTimestamp: UnmarshalJSON ParseTime: %v", err)
		}
	}
	return nil
}

type CustomDate struct {
	Date time.Time
}

func (ct *CustomDate) Value() time.Time {
	return ct.Date
}

type CustomBoolean struct {
	Boolean bool
}

func (cb *CustomBoolean) Value() bool {
	return cb.Boolean
}

func (cb *CustomBoolean) UnmarshalCSV(csv string) error {
	switch strings.ToLower(csv) {
	case "false":
		cb.Boolean = false
	case "true":
		cb.Boolean = true
	}
	return nil
}
