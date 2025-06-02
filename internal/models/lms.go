package models

type LMSData struct {
	ID       string `json:"id" bson:"id"`
	Title    string `json:"title" bson:"title"`
	Duration int    `json:"duration" bson:"duration"`
}
