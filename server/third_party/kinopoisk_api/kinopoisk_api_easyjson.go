// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package kinopoisk_api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9e2e77eaDecodeFilmerServerThirdPartyKinopoiskApi(in *jlexer.Lexer, out *kinopoiskAPIError) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e2e77eaEncodeFilmerServerThirdPartyKinopoiskApi(out *jwriter.Writer, in kinopoiskAPIError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v kinopoiskAPIError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e2e77eaEncodeFilmerServerThirdPartyKinopoiskApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v kinopoiskAPIError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e2e77eaEncodeFilmerServerThirdPartyKinopoiskApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *kinopoiskAPIError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e2e77eaDecodeFilmerServerThirdPartyKinopoiskApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *kinopoiskAPIError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e2e77eaDecodeFilmerServerThirdPartyKinopoiskApi(l, v)
}
