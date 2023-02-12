package models

import (
	"encoding/json"
	"time"
)

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

type TranslationDelivered struct {
	Timestamp Timestamp `json:"timestamp"`
	Duration  float64   `json:"duration"`
}

// GetTranslationDeliveredFromBytes return the TranslationDelivered from json bytes
func GetTranslationDeliveredFromBytes(b []byte) (TranslationDelivered, error) {
	var td TranslationDelivered
	err := json.Unmarshal(b, &td)
	if err != nil {
		return td, err
	}
	return td, nil
}
