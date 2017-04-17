package ivona_test

import (
	"os"
	"testing"

	ivona "github.com/omie/ivona-go"
)

var (
	ivonaAccessKey = os.Getenv("IVONA_ACCESS_KEY")
	ivonaSecretKey = os.Getenv("IVONA_SECRET_KEY")
	testText       = "Hello World"
)

func init() {
	if len(ivonaAccessKey) == 0 || len(ivonaSecretKey) == 0 {
		panic("IVONA_ACCESS_KEY and IVONA_SECRET_KEY environment variables are needed to run tests!\n")
	}
}

func TestIvona_CreateSpeech(t *testing.T) {
	client := ivona.New(ivonaAccessKey, ivonaSecretKey)
	options := ivona.NewSpeechOptions(testText)
	r, err := client.CreateSpeech(options)

	if err != nil {
		t.Error(err)
	}

	audioLength := len(r.Audio)
	expectedAudioLength := 6314
	expectedContentType := "audio/mpeg"

	if r.ContentType != expectedContentType {
		t.Errorf("ContentType %v does not match", r.ContentType)
	}

	if audioLength != expectedAudioLength {
		t.Errorf("Audio length %v does not match", audioLength)
	}
}

func TestIvona_ListVoices(t *testing.T) {
	client := ivona.New(ivonaAccessKey, ivonaSecretKey)

	r, err := client.ListVoices(ivona.Voice{})
	if err != nil {
		t.Error(err)
	}

	voicesLength := len(r.Voices)
	expectedVoicesLength := 51
	expectedContentType := "application/json"

	if voicesLength != expectedVoicesLength {
		t.Errorf("Voices length %v does not match", len(r.Voices))
	}

	if r.ContentType != expectedContentType {
		t.Errorf("ContentType %v does not match", r.ContentType)
	}
}
