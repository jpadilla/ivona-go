package ivona

// SpeechResponse is the resource representing response from CreateSpeech action.
type SpeechResponse struct {
	Audio       []byte
	RequestID   string
	ContentType string
}

// ListResponse is the resource representing response from ListVoices action.
type ListResponse struct {
	Voices      []Voice
	RequestID   string
	ContentType string
}

// SpeechOptions is the set of parameters that can be used on the CreateSpeech action.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_actions.html#CreateSpeech.
type SpeechOptions struct {
	Input        *Input
	OutputFormat *OutputFormat
	Parameters   *Parameters
	Voice        *Voice
}

// NewSpeechOptions is the set of default parameters that can be used the CreateSpeech action.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_actions.html#CreateSpeech_DefaultValues.
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

// Input contains attributes describing the user input.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_data_types.html#DataTypes_Input.
type Input struct {
	Data string
	Type string
}

// OutputFormat contains attributes describing the audio compression and format in which the returned stream should be encoded.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_data_types.html#DataTypes_OutputFormat.
type OutputFormat struct {
	Codec      string
	SampleRate int
}

// Parameters contains additional attributes affecting the generated speech.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_data_types.html#DataTypes_Parameters.
type Parameters struct {
	Rate           string
	Volume         string
	SentenceBreak  int
	ParagraphBreak int
}

// Voice contains a filter for the voice selection that should be used for the speech synthesis.
// For more details see http://developer.ivona.com/en/speechcloud/api_ref_data_types.html#DataTypes_Voice.
type Voice struct {
	Name     string
	Language string
	Gender   string
}
