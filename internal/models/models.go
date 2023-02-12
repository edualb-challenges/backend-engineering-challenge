package models

import (
	"encoding/json"
	"time"
)

type TranslationDelivered struct {
	Timestamp Timestamp `json:"timestamp"`
	Duration  float64   `json:"duration"`
}

type Timestamp struct {
	time.Time
}

func (ts *Timestamp) UnmarshalJSON(bytes []byte) error {
	var v string
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}

	ts.Time, err = time.Parse("2006-01-02 15:04:05.000000", v)
	if err != nil {
		return err
	}
	return nil
}
