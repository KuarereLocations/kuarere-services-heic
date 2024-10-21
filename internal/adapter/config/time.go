package config

import "time"

func GetTimeNowLocation() time.Time {
	location, errLoadLocation := time.LoadLocation("Europe/Madrid")
	if errLoadLocation == nil {
		return time.Now().In(location)
	}
	return time.Now().In(time.Local)
}
