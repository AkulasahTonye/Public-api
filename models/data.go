package models

import (
	"time"
)

type Data struct {
	// struct tag
	Email     string
	DateTime  time.Time
	GitHubUrl string
	Backlink  string
}

var data = []Data{}

func (d Data) Save() {

	// Later: add it to a database
	data = append(data, d)
}

func GetAllData() []Data {
	return data
}
