package date

import "time"

func FormatToISO(dt time.Time) string {
	return dt.UTC().Format(time.RFC3339)
}
