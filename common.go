package appsflyer_sdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return fmt.Errorf("CustomFloat64: UnmarshalJSON: %v", err)
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return fmt.Errorf("CustomFloat64: UnmarshalJSON: %v", err)
		}
	}
	return nil
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

func (ct *CustomTimestamp) UnmarshalJSON(data []byte) error {
	var ts string
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return fmt.Errorf("CustomTimestamp: UnmarshalJSON: %v", err)
	}
	ct.Timestamp, err = time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		return fmt.Errorf("CustomTimestamp: UnmarshalJSON ParseTime: %v", err)
	}
	return nil
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

func (ct *CustomDate) UnmarshalJSON(data []byte) error {
	var ts string
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return fmt.Errorf("CustomTimestamp: UnmarshalJSON: %v", err)
	}
	ct.Date, err = time.Parse("2006-01-02", ts)
	if err != nil {
		return fmt.Errorf("CustomTimestamp: UnmarshalJSON ParseTime: %v", err)
	}
	return nil
}
