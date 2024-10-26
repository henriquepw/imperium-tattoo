package date

import "time"

func ParseInput(input string) (time.Time, error) {
	return time.Parse(time.DateOnly, input)
}

func FormatToISO(dt time.Time) string {
	return dt.UTC().Format(time.RFC3339)
}

func FormatToFormInput(dt time.Time) string {
	return dt.UTC().Format(time.DateOnly)
}
