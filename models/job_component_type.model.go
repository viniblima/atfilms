package models

type jobComponentType string

const (
	SIMPLE_TEXT           jobComponentType = "SIMPLE_TEXT"
	CITATION              jobComponentType = "CITATION"
	VIDEOS                jobComponentType = "VIDEOS"
	PHOTO_SLIDER          jobComponentType = "PHOTO_SLIDER"
	FILL_PHOTO_HORIZONTAL jobComponentType = "FILL_PHOTO_HORIZONTAL"
)
