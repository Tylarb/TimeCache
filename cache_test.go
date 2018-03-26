/*
Package timeCache is a simple implementation of an interface and two structures
which define aa in-memory cache where entries have a timeout


Released under MIT license, copyright 2018 Tyler Ramer
*/

package timeCache

import (
	"testing"
	"time"
)

const TIMEOUT = 10

var keys []string
var times []time.Time
var entries map[string]time.Time
var timeout = time.Duration(TIMEOUT) * time.Second
var sliceCache SliceCache
var dictCache DictCache

func TestSlice(t *testing.T) {
}
