package util

import (
	"testing"
	"fmt"
)

func TestGetMD5Hash(t *testing.T) {
	str := "google"
	exper := GetMD5Hash(str)
	fmt.Println(str, exper)
}
