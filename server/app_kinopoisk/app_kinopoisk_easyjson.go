// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package app_kinopoisk

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

func easyjsonF8d3a567DecodeServerAppKinopoisk(in *jlexer.Lexer, out *SearchFilmsIn) {
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
		case "Query":
			out.Query = string(in.String())
		case "Page":
			out.Page = int(in.Int())
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
func easyjsonF8d3a567EncodeServerAppKinopoisk(out *jwriter.Writer, in SearchFilmsIn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Query\":"
		out.RawString(prefix[1:])
		out.String(string(in.Query))
	}
	{
		const prefix string = ",\"Page\":"
		out.RawString(prefix)
		out.Int(int(in.Page))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchFilmsIn) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF8d3a567EncodeServerAppKinopoisk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchFilmsIn) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8d3a567EncodeServerAppKinopoisk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchFilmsIn) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF8d3a567DecodeServerAppKinopoisk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchFilmsIn) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8d3a567DecodeServerAppKinopoisk(l, v)
}
func easyjsonF8d3a567DecodeServerAppKinopoisk1(in *jlexer.Lexer, out *GetFilmInfoIn) {
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
		case "KinopoiskID":
			out.KinopoiskID = int(in.Int())
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
func easyjsonF8d3a567EncodeServerAppKinopoisk1(out *jwriter.Writer, in GetFilmInfoIn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"KinopoiskID\":"
		out.RawString(prefix[1:])
		out.Int(int(in.KinopoiskID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFilmInfoIn) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF8d3a567EncodeServerAppKinopoisk1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFilmInfoIn) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8d3a567EncodeServerAppKinopoisk1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFilmInfoIn) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF8d3a567DecodeServerAppKinopoisk1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFilmInfoIn) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8d3a567DecodeServerAppKinopoisk1(l, v)
}
