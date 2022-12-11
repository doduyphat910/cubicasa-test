package entity

import (
	"strconv"
	"strings"
	"time"
)

type Team struct {
	ID          uint64
	HubID       uint64
	Hub         Hub
	GeoLocation GeoLocation
	Users       []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GeoLocation struct {
	Lat, Long float64
}

func (gl GeoLocation) GormDataType() string {
	return "geometry"
}

func (gl *GeoLocation) Scan(v interface{}) error {
	pointStr := v.(string)
	pointSplitted := strings.Split(pointStr, ",")
	long, err := strconv.ParseFloat(pointSplitted[0][1:], 64)
	if err != nil {
		return err
	}
	lat, err := strconv.ParseFloat(pointSplitted[1][:len(pointSplitted[1])-1], 64)
	if err != nil {
		return err
	}

	gl.Long = long
	gl.Lat = lat
	return nil
}
