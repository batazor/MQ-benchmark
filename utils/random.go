package utils

import "github.com/bxcodec/faker"

type Packet struct {
	Oid       string  `faker:"month_name" json:"_oid"`
	UserName  string  `faker:"username" json:"username"`
	UnixTime  int64   `faker:"unix_time" json:"unixtime"`
	Email     string  `faker:"email" json:"email"`
	Year      string  `faker:"year" json:"year"`
	DayOfWeek string  `faker:"day_of_week" json:"dayOfWeek"`
	Timestamp string  `faker:"timestamp" json:"timestamp"`
	IPV4      string  `faker:"ipv4" json:"IPv4"`
	Lat       float32 `faker:"lat" json:"lat"`
	Lon       float32 `faker:"long" json:"lon"`
	Bool      bool
	Int       int
}

func GetRandomPacket() (interface{}, error) {
	a := Packet{}
	err := faker.FakeData(&a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
