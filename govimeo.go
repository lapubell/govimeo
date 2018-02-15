package govimeo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var apiToken string
var video vimeoVideo

type vimeoVideo struct {
	URI      string `json:"uri"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Pictures struct {
		URI string `json:"uri"`
	} `json:"pictures"`
}

// SetToken allows you to pass your API token from a different package or the
// main function
func SetToken(s string) error {
	if s == "" {
		return errors.New("token can not be empty")
	}

	apiToken = s
	return nil
}

// GetVideoDuration will return the duration of the video by making a GET
// request to the Vimeo API
func GetVideoDuration(vID uint) (int, error) {
	if len(apiToken) == 0 {
		return 0, errors.New("Please set your token")
	}

	// check and see if we need to request new data
	if "/videos/"+strconv.Itoa(int(vID)) != video.URI {
		err := requestDataFromVimeo(vID)
		if err != nil {
			return 0, errors.New("Problem getting data from vimeo")
		}
		return video.Duration, nil
	}

	return video.Duration, nil
}

// GetVideoPictureID will return the ID of the picture by making a GET
// request to the Vimeo API, and parsing the pictures.uri value
func GetVideoPictureID(vID uint) (int, error) {
	if len(apiToken) == 0 {
		return 0, errors.New("Please set your token")
	}

	// check and see if we need to request new data
	if "/videos/"+strconv.Itoa(int(vID)) != video.URI {
		err := requestDataFromVimeo(vID)
		if err != nil {
			return 0, errors.New("Problem getting data from vimeo")
		}
		pID, err := getPictureID(video.Pictures.URI)
		if err != nil {
			return 0, err
		}
		return pID, nil
	}

	pID, err := getPictureID(video.Pictures.URI)
	if err != nil {
		return 0, err
	}
	return pID, nil
}

func getPictureID(uri string) (int, error) {
	parts := strings.Split(uri, "/")
	return strconv.Atoi(parts[len(parts)-1])
}

func requestDataFromVimeo(vID uint) error {
	v := vimeoVideo{}
	// fake data so i'm not doing a real http request during testing
	if apiToken == "testing" {
		v.Duration = 100
		v.Name = "fake video"
		v.Pictures.URI = "/videos/" + strconv.Itoa(int(vID)) + "/pictures/654321"
		v.URI = "/videos/" + strconv.Itoa(int(vID)) + ""

		video = v
		return nil
	}

	response, err := http.Get("https://api.vimeo.com/videos/" + strconv.Itoa(int(vID)))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &v)
	if err != nil {
		return errors.New("problem with the response from vimeo")
	}

	video = v
	return nil
}
