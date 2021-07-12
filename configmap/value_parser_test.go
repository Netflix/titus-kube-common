package configmap

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseBool(t *testing.T) {
	require.Equal(t, ParseBool(map[string]string{}, "myKey", true), true)
	require.Equal(t, ParseBool(map[string]string{"myKey": "true"}, "myKey", false), true)
	require.Equal(t, ParseBool(map[string]string{"myKey": "false"}, "myKey", true), false)
}

func TestParseDurationInSec(t *testing.T) {
	require.Equal(t, ParseDurationInSec(map[string]string{}, "myKey", time.Second*5), time.Second*5)
	require.Equal(t, ParseDurationInSec(map[string]string{"myKey": "10"}, "myKey", time.Second*5), time.Second*10)
}

func TestParseFloat64(t *testing.T) {
	require.Equal(t, ParseFloat64(map[string]string{}, "myKey", 10.1), 10.1)
	require.Equal(t, ParseFloat64(map[string]string{"myKey": "123.4"}, "myKey", 10), 123.4)
}

func TestParseInt64(t *testing.T) {
	require.Equal(t, ParseInt64(map[string]string{}, "myKey", 10), int64(10))
	require.Equal(t, ParseInt64(map[string]string{"myKey": "123"}, "myKey", 10), int64(123))
}

func TestParseStringList(t *testing.T) {
	require.Equal(t, ParseStringList(map[string]string{}, "myKey", []string{"a"}), []string{"a"})
	require.Equal(t, ParseStringList(map[string]string{"myKey": "a,b,c"}, "myKey", []string{"a"}), []string{"a", "b", "c"})
}
