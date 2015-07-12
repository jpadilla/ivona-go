package ivona_test

import (
	"log"

	ivona "github.com/jpadilla/ivona-go"
)

func ExampleIvona_CreateSpeech() {
	client := ivona.New("IVONA_ACCESS_KEY", "IVONA_SECRET_KEY")
	options := ivona.NewSpeechOptions("Hello World")
	r, err := client.CreateSpeech(options)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", len(r.Audio))
	log.Printf("%v\n", r.ContentType)
	log.Printf("%v\n", r.RequestID)
}

func ExampleIvona_ListVoices() {
	client := ivona.New("IVONA_ACCESS_KEY", "IVONA_SECRET_KEY")

	r, err := client.ListVoices(ivona.Voice{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", len(r.Voices))
}
