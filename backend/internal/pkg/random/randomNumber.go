package random

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomRangeNumber(low, hi int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(low + rand.Intn(hi-low))
}
