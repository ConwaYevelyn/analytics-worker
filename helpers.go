package analytics_worker

import (
	"log"
	"net/http"
	"fmt"
	"math"
)

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func clamp(value float64, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func calculateDistance(loc1, loc2 *Location) float64 {
	lat1, lon1 := math.Pi()*loc1.Latitude/180, math.Pi()*loc1.Longitude/180
	lat2, lon2 := math.Pi()*loc2.Latitude/180, math.Pi()*loc2.Longitude/180
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return 6371 * c
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type Stats struct {
	Count  int64
	Min    float64
	Max    float64
	Avg    float64
	Sum    float64
}

func calculateStats(values []float64) Stats {
	var sum float64
	for _, value := range values {
		sum += value
	}

	min := math.MaxFloat64
	max := math.MinFloat64
	for _, value := range values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return Stats{
		Count:  len(values),
		Min:    min,
		Max:    max,
		Avg:    sum / float64(len(values)),
		Sum:    sum,
	}
}

func getStatusCode(status string) int {
	switch status {
	case "success":
		return http.StatusOK
	case "error":
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}