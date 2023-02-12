package models

import (
	"encoding/json"
	"fmt"
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

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", ts.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Event struct {
	Timestamp Timestamp `json:"timestamp"`
	Duration  float64   `json:"duration"`
}

// GetEventFromBytes return the Event from json bytes
func GetEventFromBytes(b []byte) (Event, error) {
	var e Event
	err := json.Unmarshal(b, &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

type AverageDeliveryTime struct {
	Date    Timestamp `json:"date"`
	Average float64   `json:"average_delivery_time"`
}
