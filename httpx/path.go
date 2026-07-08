package httpx

import (
	"log"
	"net/http"
	"strconv"
)

func PathInt(key string, r *http.Request) int64 {
	id, err := strconv.ParseInt(r.PathValue(key), 10, 64)
	if err != nil {
		log.Panicf("Error getting path int (%s): %v", key, err)
	}
	return id
}
