package azuretexttospeech

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/h2non/gentleman.v2"
)

// voiceListAPI is the source for supported voice list to region mapping
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#regions-and-endpoints
const voiceListAPI = "https://%s.tts.speech.microsoft.com/cognitiveservices/voices/list"

//go:generate enumer -type=voiceType -linecomment -json
type voiceType int

const (
	voiceStandard voiceType = iota // Standard
	voiceNeural                    // Neural
)

type regionVoiceListResponse struct {
	Name            string    `json:"Name"`
	ShortName       string    `json:"ShortName"`
	Gender          Gender    `json:"Gender"`
	Locale          Locale    `json:"Locale"`
	SampleRateHertz string    `json:"SampleRateHertz"`
	VoiceType       voiceType `json:"VoiceType"`
}

// supportedVoices represents the key used within the `localeToGender` map.
type supportedVoices struct {
	Gender Gender
	Locale Locale
}

type RegionVoiceMap map[supportedVoices]string

func (az *AzureCSTextToSpeech) buildVoiceToRegionMap() (RegionVoiceMap, error) {

	v, err := az.fetchVoiceList()
	if err != nil {
		return nil, err
	}

	m := make(map[supportedVoices]string)
	for _, x := range v {
		if x.VoiceType == voiceStandard {
			m[supportedVoices{Gender: x.Gender, Locale: x.Locale}] = x.ShortName
		}
	}
	return m, err
}

func (az *AzureCSTextToSpeech) fetchVoiceList() ([]regionVoiceListResponse, error) {

	// Create a new client
	cli := gentleman.New()

	// Define base URL
	cli.URL(az.voiceServiceListURL)

	// Create a new request based on the current client
	req := cli.Request()

	// // Define the URL path at request level
	// req.Path("/headers")

	// Set a new header field
	req.SetHeader("Authorization", "Bearer "+az.accessToken)

	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return nil, err
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, err
	}

	// Reads the whole body and returns it as string
	fmt.Printf("Body: %s", res.String())

	// request, err := http.NewRequest(http.MethodGet, az.voiceServiceListURL, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// request.Header.Set("Authorization", "Bearer "+az.accessToken)
	// client := &http.Client{Timeout: 2 * time.Second}
	// response, err := client.Do(request)
	// if err != nil {
	// 	return nil, err
	// }
	// defer response.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		var r []regionVoiceListResponse
		if err := json.Unmarshal(res.Bytes(), &r); err != nil {
			return nil, fmt.Errorf("unable to decode voice list response body, %v", err)
		}

		// if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		// 	return nil, fmt.Errorf("unable to decode voice list response body, %v", err)
		// }
		return r, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%d - A required parameter is missing, empty, or null. Or, the value passed to either a required or optional parameter is invalid. A common issue is a header that is too long", res.StatusCode)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%d - The request is not authorized. Check to make sure your subscription key or token is valid and in the correct region", res.StatusCode)
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("%d - You have exceeded the quota or rate of requests allowed for your subscription", res.StatusCode)
	case http.StatusBadGateway:
		return nil, fmt.Errorf("%d - Network or server-side issue. May also indicate invalid headers", res.StatusCode)
	}
	return nil, fmt.Errorf("%d - unexpected response code from voice list API", res.StatusCode)
}

func (az *AzureCSTextToSpeech) OfetchVoiceList() ([]regionVoiceListResponse, error) {

	request, err := http.NewRequest(http.MethodGet, az.voiceServiceListURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+az.accessToken)
	client := &http.Client{Timeout: 2 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var r []regionVoiceListResponse
		if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
			return nil, fmt.Errorf("unable to decode voice list response body, %v", err)
		}
		return r, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%d - A required parameter is missing, empty, or null. Or, the value passed to either a required or optional parameter is invalid. A common issue is a header that is too long", response.StatusCode)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%d - The request is not authorized. Check to make sure your subscription key or token is valid and in the correct region", response.StatusCode)
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("%d - You have exceeded the quota or rate of requests allowed for your subscription", response.StatusCode)
	case http.StatusBadGateway:
		return nil, fmt.Errorf("%d - Network or server-side issue. May also indicate invalid headers", response.StatusCode)
	}
	return nil, fmt.Errorf("%d - unexpected response code from voice list API", response.StatusCode)
}
