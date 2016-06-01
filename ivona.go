// Package ivona provides the binding for IVONA Speech Cloud API
package ivona

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/aws4"
	"io/ioutil"
	"net/http"
)

// ivonaAPI is the public API IVONA Speech Cloud URL.
const ivonaAPI = "https://tts.eu-west-1.ivonacloud.com"

// createSpeechAPI is the public API IVONA Speech Cloud URL for the CreateSpeech action.
const createSpeechAPI = ivonaAPI + "/CreateSpeech"
const listVoicesAPI = ivonaAPI + "/ListVoices"

// Ivona is used to invoke API calls
type Ivona struct {
	AccessKey string
	SecretKey string
}

// New returns a new Ivona client.
func New(accessKey string, secretKey string) *Ivona {
	return &Ivona{AccessKey: accessKey, SecretKey: secretKey}
}

// CreateSpeech performs a synthesis of the requested text and returns the audio stream containing the speech.
func (client *Ivona) CreateSpeech(options SpeechOptions) (*SpeechResponse, error) {
	b, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", createSpeechAPI, bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	awsClient := aws4.Client{Keys: &aws4.Keys{
		AccessKey: client.AccessKey,
		SecretKey: client.SecretKey,
	}}

	resp, err := awsClient.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got non 200 status code: %s %q", resp.Status, data)
	}

	return &SpeechResponse{
		Audio:       data,
		RequestID:   resp.Header["X-Amzn-Ivonattsrequestid"][0],
		ContentType: resp.Header["Content-Type"][0],
	}, nil
}

// ListVoices retrieves list of voices from the api
func (client *Ivona) ListVoices(options Voice) (*ListResponse, error) {
	voiceOptions := struct {
		Gender 	 string `json:",omitempty"`
		Language string `json:",omitempty"`
		Name 	 string `json:",omitempty"`
	}{}
	voiceOptions.Gender = options.Gender
	voiceOptions.Language = options.Language
	voiceOptions.Name = options.Name

	b, err := json.Marshal(map[string]interface{}{ "Voice": voiceOptions })
	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", listVoicesAPI, bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	awsClient := aws4.Client{Keys: &aws4.Keys{
		AccessKey: client.AccessKey,
		SecretKey: client.SecretKey,
	}}

	resp, err := awsClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got non 200 status code: %s %q", resp.Status, data)
	}

	list := new(ListResponse)
	err = json.Unmarshal(data, list)
	if err != nil {
		return nil, err
	}

	list.RequestID = resp.Header["X-Amzn-Ivonattsrequestid"][0]
	list.ContentType = resp.Header["Content-Type"][0]

	return list, nil
}
