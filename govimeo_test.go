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
