package lotr_sdk

import (
	"testing"
)

func TestGetMovies(t *testing.T) {
	testClient := NewLOTRClient()

	movies, _, err := testClient.GetMovies()
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if got, want := len(movies), 8; got != want {
		t.Errorf("Expected %d movies, but recived %d", want, got)
	}

}
