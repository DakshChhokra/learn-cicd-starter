package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	sampleHeader := http.Header{}
	sampleHeader.Add("Authorization", "")

	want := ErrNoAuthHeaderIncluded
	apikey, got := GetAPIKey(sampleHeader)

	if want != got && !reflect.DeepEqual("", apikey) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
