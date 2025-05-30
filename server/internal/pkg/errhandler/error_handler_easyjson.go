// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package errhandler

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

func easyjsonB52af4bdDecodeFilmerServerPkgUtils(in *jlexer.Lexer, out *errorResponse) {
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
func easyjsonB52af4bdEncodeFilmerServerPkgUtils(out *jwriter.Writer, in errorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v errorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB52af4bdEncodeFilmerServerPkgUtils(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v errorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB52af4bdEncodeFilmerServerPkgUtils(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *errorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB52af4bdDecodeFilmerServerPkgUtils(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *errorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB52af4bdDecodeFilmerServerPkgUtils(l, v)
}
