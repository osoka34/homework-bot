package utils

import "time"

func GetLastSundayRange() (time.Time, time.Time) {
	now := GetMoscowTime()
	offset := int(now.Weekday())
	if offset == 0 {
		offset = 7 // If today is Sunday, go back to the previous Sunday
	}
	lastSunday := now.AddDate(0, 0, -offset)
	start := time.Date(
		lastSunday.Year(),
		lastSunday.Month(),
		lastSunday.Day(),
		12,
		0,
		0,
		0,
		lastSunday.Location(),
	)
	end := start.Add(15 * time.Hour) // Add 15 hours to reach Monday 03:00
	return start, end
}

func GetMoscowTime() time.Time {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Now().UTC().Add(3 * time.Hour)
	}
	return time.Now().In(loc)
}
