package ivona

type SpeechResponse struct {
	Audio       []byte
	RequestId   string
	ContentType string
}

type SpeechOptions struct {
	Input        *Input
	OutputFormat *OutputFormat
	Parameters   *Parameters
	Voice        *Voice
}

func NewSpeechOptions(data string) SpeechOptions {
	return SpeechOptions{
		Input: &Input{
			Data: data,
			Type: "text/plain",
		},
		OutputFormat: &OutputFormat{
			Codec:      "MP3",
			SampleRate: 22050,
		},
		Parameters: &Parameters{
			Rate:           "medium",
			Volume:         "medium",
			SentenceBreak:  400,
			ParagraphBreak: 640,
		},
		Voice: &Voice{
			Name:     "Salli",
			Language: "en-US",
			Gender:   "Female",
		},
	}
}

type Input struct {
	Data string
	Type string
}

type OutputFormat struct {
	Codec      string
	SampleRate int
}

type Parameters struct {
	Rate           string
	Volume         string
	SentenceBreak  int
	ParagraphBreak int
}

type Voice struct {
	Name     string
	Language string
	Gender   string
}
