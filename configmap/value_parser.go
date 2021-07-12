package configmap

import (
	"strconv"
	"strings"
	"time"
)

func ParseBool(data map[string]string, key string, defaultValue bool) bool {
	if valueStr, ok := data[key]; ok {
		return strings.EqualFold(valueStr, "true")
	}
	return defaultValue
}

func ParseInt64(data map[string]string, key string, defaultValue int64) int64 {
	if valueStr, ok := data[key]; ok {
		if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return value
		}
	}
	return defaultValue
}

func ParseFloat64(data map[string]string, key string, defaultValue float64) float64 {
	if valueStr, ok := data[key]; ok {
		if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
			return value
		}
	}
	return defaultValue
}

func ParseDurationInSec(data map[string]string, key string, defaultValue time.Duration) time.Duration {
	if valueStr, ok := data[key]; ok {
		if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return time.Duration(value) * time.Second
		}
	}
	return defaultValue
}

func ParseStringList(data map[string]string, key string, defaultValue []string) []string {
	if valueStr, ok := data[key]; ok {
		return strings.Split(valueStr, ",")
	}
	return defaultValue
}

