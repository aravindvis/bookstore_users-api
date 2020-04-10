package date_utils

import "time"

const (
	apiDateLayOut = "2006-01-02T15:04:05Z"
)

//GetNowString returns date in the defined format
func GetNowString() string {
	return GetNow().Format(apiDateLayOut)
}

//GetNow gets the current time in the UTC format
func GetNow() time.Time {
	return time.Now().UTC()
}
