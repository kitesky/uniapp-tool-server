package json

import jsoniter "github.com/json-iterator/go"

var (
	json                = jsoniter.ConfigCompatibleWithStandardLibrary
	Get                 = json.Get
	Valid               = json.Valid
	RegisterExtension   = json.RegisterExtension
	DecoderOf           = json.DecoderOf
	EncoderOf           = json.EncoderOf
	Marshal             = json.Marshal
	Unmarshal           = json.Unmarshal
	MarshalIndent       = json.MarshalIndent
	NewDecoder          = json.NewDecoder
	NewEncoder          = json.NewEncoder
	MarshalToString     = json.MarshalToString
	UnmarshalFromString = json.UnmarshalFromString
)
