package httputil

import (
	"net/url"
	"strconv"
)

func GetQueryInt(q url.Values, name string, defaultVal int64) int64 {
	val, err := strconv.ParseInt(q.Get(name), 10, 64)
	if err != nil {
		val = defaultVal
	}

	return val
}
