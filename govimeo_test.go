package govimeo

import (
	"testing"
)

func TestSetToken(t *testing.T) {
	if apiToken != "" {
		t.Errorf("token should be empty by default")
	}

	err := SetToken("")
	if err.Error() != "token can not be empty" {
		t.Errorf("shouldn't be able to pass an empty string to SetToken")
	}

	err = SetToken("asdf")
	if err != nil {
		t.Errorf("this should have worked")
	}
	if apiToken != "asdf" {
		t.Errorf("the token didn't get set correctly")
	}
}

func TestGetVideoDuration(t *testing.T) {
	// reset the token to a zero state
	apiToken = ""
	_, e := GetVideoDuration(123)
	if e.Error() != "Please set your token" {
		t.Errorf("Token is required to make a call")
	}

	apiToken = "testing"
	d, e := GetVideoDuration(123)
	if d != 100 {
		t.Errorf("expected 100 from the fake data")
	}
	// call it with the hardcoded fake data make sure that it doesn't request data from vimeo unless video ID changes
	d, e = GetVideoDuration(123)
	if d != 100 {
		t.Errorf("expected 100 from the fake data")
	}
}

func TestGetPictureID(t *testing.T) {
	uri := ""
	result := getPictureID(uri)
	if result != "" {
		t.Errorf("Expected empty string")
	}

	uri = "a/s/d/f"
	result = getPictureID(uri)
	if result != "f" {
		t.Errorf("Expected f")
	}

	uri = "a/s/d/f/"
	result = getPictureID(uri)
	if result != "" {
		t.Errorf("Expected empty string")
	}
}

func TestGetVideoPictureID(t *testing.T) {
	// reset the token to a zero state
	apiToken = ""
	_, e := GetVideoPictureID(123)
	if e.Error() != "Please set your token" {
		t.Errorf("Token is required to make a call")
	}

	apiToken = "testing"
	pID, e := GetVideoPictureID(321)
	if pID != "654321" {
		t.Errorf("expected 654321 from the fake data")
	}
	// call it with the hardcoded fake data make sure that it doesn't request data from vimeo unless video ID changes
	pID, e = GetVideoPictureID(321)
	if pID != "654321" {
		t.Errorf("expected 654321 from the fake data")
	}
}
