package models

import (
	"time"
)

type Data struct {
	// struct tag
	Email            string
	Current_datetime time.Time
	Github_url       string
}

//var data = []Data{}

func GetAllData() []Data {

	currentTime := time.Now().UTC()

	data := []Data{{
		Email:            "akulasaht@gmail.com",
		Current_datetime: currentTime,
		Github_url:       "https://github.com/AkulasahTonye/API_Testing",
	},
	}

	return data
}
