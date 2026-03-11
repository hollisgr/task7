package main

import (
	"strconv"
	"strings"
)

func Join(ids []int) string {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	return strings.Join(strIDs, ",")
}
