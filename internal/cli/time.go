package cli

import "time"

const dateTimeLayout = "2006-01-02 15:04:05"

func FormatUnixTime(sec int64) string {
	if sec == 0 {
		return "-"
	}
	return time.Unix(sec, 0).Local().Format(dateTimeLayout)
}
