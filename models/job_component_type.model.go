package models

import "database/sql/driver"

type jobComponentType string

const (
	SIMPLE_TEXT           jobComponentType = "SIMPLE_TEXT"
	CITATION              jobComponentType = "CITATION"
	VIDEOS                jobComponentType = "VIDEOS"
	PHOTO_SLIDER          jobComponentType = "PHOTO_SLIDER"
	FILL_PHOTO_HORIZONTAL jobComponentType = "FILL_PHOTO_HORIZONTAL"
)

func (ct *jobComponentType) Scan(value interface{}) error {
	*ct = jobComponentType(value.([]byte))
	return nil
}

func (ct jobComponentType) Value() (driver.Value, error) {
	return string(ct), nil
}
