package sim

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// return value based on probability density
func generateValue(cdf []int, values []interface{}) interface{} {
	r := rand.Intn(100)
	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return values[bucket]
}

func GenerateStatusCode() string {
	// 200: OK, 401: Unauthorized, 503: Service Unavailable
	status := generateValue([]int{80, 95, 100}, []interface{}{"200", "401", "503"})
	return status.(string)
}

func GenerateResponseSize() float64 {
	base := generateValue([]int{20, 60, 80, 90, 100}, []interface{}{200, 400, 600, 800, 1000})
	return float64(base.(int) + rand.Intn(200))
}

func GenerateRequestTime() float64 {
	base := generateValue([]int{80, 95, 100}, []interface{}{20, 500, 800})
	return time.Duration(time.Millisecond * time.Duration(base.(int)+rand.Intn(200))).Seconds()
}
