package ivona

import (
	"encoding/json"
	// "fmt"
	"github.com/bmizerany/aws4"
	// "io"
	// "log"
	"net/http"
	// "os"
	// "strings"
	"bytes"
	"io/ioutil"
)

const (
	IvonaAPI        = "https://tts.eu-west-1.ivonacloud.com"
	CreateSpeechAPI = IvonaAPI + "/CreateSpeech"
)

type Ivona struct {
	AccessKey string
	SecretKey string
}

func New(accessKey string, secretKey string) *Ivona {
	return &Ivona{AccessKey: accessKey, SecretKey: secretKey}
}

func (client *Ivona) CreateSpeech(options SpeechOptions) (*SpeechResponse, error) {
	b, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", CreateSpeechAPI, bytes.NewReader(b))
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

	return &SpeechResponse{
		Audio:       data,
		RequestId:   resp.Header["X-Amzn-Ivonattsrequestid"][0],
		ContentType: resp.Header["Content-Type"][0],
	}, nil
}

// func main() {
// 	r, _ := http.NewRequest("POST", "https://tts.eu-west-1.ivonacloud.com/CreateSpeech", data)
// 	r.Header.Set("Content-Type", "application/json")

// 	Client := aws4.Client{Keys: &aws4.Keys{
// 		AccessKey: "GDNAI77EX24OLG2PG7XA",
// 		SecretKey: "zM9IBfwgyfleyy2QFT1mHR+W3SIuWHVZUii9iX4E",
// 	}}

// 	resp, err := Client.Do(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	out, err := os.Create("voice.mp3")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer out.Close()

// 	fmt.Println(resp)
// 	fmt.Println(resp.Body)

// 	io.Copy(out, resp.Body)
// }
