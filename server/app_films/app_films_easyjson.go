// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package app_films

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

func easyjsonDea1808fDecodeServerAppFilms(in *jlexer.Lexer, out *Sort) {
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
		case "sortField":
			out.SortField = string(in.String())
		case "sortOrder":
			out.SortOrder = string(in.String())
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
func easyjsonDea1808fEncodeServerAppFilms(out *jwriter.Writer, in Sort) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SortField != "" {
		const prefix string = ",\"sortField\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.SortField))
	}
	if in.SortOrder != "" {
		const prefix string = ",\"sortOrder\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SortOrder))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Sort) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Sort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Sort) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Sort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms(l, v)
}
func easyjsonDea1808fDecodeServerAppFilms1(in *jlexer.Lexer, out *SetFilmCategoryIn) {
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
		case "MovieID":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.MovieID).UnmarshalText(data))
			}
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
func easyjsonDea1808fEncodeServerAppFilms1(out *jwriter.Writer, in SetFilmCategoryIn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"MovieID\":"
		out.RawString(prefix[1:])
		out.RawText((in.MovieID).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetFilmCategoryIn) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetFilmCategoryIn) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetFilmCategoryIn) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetFilmCategoryIn) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms1(l, v)
}
func easyjsonDea1808fDecodeServerAppFilms2(in *jlexer.Lexer, out *Pagination) {
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
		case "page":
			out.Page = int(in.Int())
		case "pages":
			out.Pages = int(in.Int())
		case "total":
			out.Total = int64(in.Int64())
		case "limit":
			out.Limit = int(in.Int())
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
func easyjsonDea1808fEncodeServerAppFilms2(out *jwriter.Writer, in Pagination) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"page\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Page))
	}
	{
		const prefix string = ",\"pages\":"
		out.RawString(prefix)
		out.Int(int(in.Pages))
	}
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix)
		out.Int64(int64(in.Total))
	}
	{
		const prefix string = ",\"limit\":"
		out.RawString(prefix)
		out.Int(int(in.Limit))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Pagination) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pagination) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Pagination) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pagination) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms2(l, v)
}
func easyjsonDea1808fDecodeServerAppFilms3(in *jlexer.Lexer, out *Filter) {
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
		case "ratingFrom":
			if in.IsNull() {
				in.Skip()
				out.RatingFrom = nil
			} else {
				if out.RatingFrom == nil {
					out.RatingFrom = new(float64)
				}
				*out.RatingFrom = float64(in.Float64())
			}
		case "yearFrom":
			out.YearFrom = int(in.Int())
		case "yearTo":
			out.YearTo = int(in.Int())
		case "type":
			out.Type = string(in.String())
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]string, 0, 4)
					} else {
						out.Genres = []string{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Genres = append(out.Genres, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonDea1808fEncodeServerAppFilms3(out *jwriter.Writer, in Filter) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RatingFrom != nil {
		const prefix string = ",\"ratingFrom\":"
		first = false
		out.RawString(prefix[1:])
		out.Float64(float64(*in.RatingFrom))
	}
	if in.YearFrom != 0 {
		const prefix string = ",\"yearFrom\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.YearFrom))
	}
	if in.YearTo != 0 {
		const prefix string = ",\"yearTo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.YearTo))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if len(in.Genres) != 0 {
		const prefix string = ",\"genres\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Genres {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Filter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Filter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Filter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Filter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms3(l, v)
}
func easyjsonDea1808fDecodeServerAppFilms4(in *jlexer.Lexer, out *CategoryFilmsOut) {
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
		case "filter":
			(out.Filter).UnmarshalEasyJSON(in)
		case "sort":
			(out.Sort).UnmarshalEasyJSON(in)
		case "pagination":
			(out.Pagination).UnmarshalEasyJSON(in)
		case "films":
			(out.Films).UnmarshalEasyJSON(in)
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
func easyjsonDea1808fEncodeServerAppFilms4(out *jwriter.Writer, in CategoryFilmsOut) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"filter\":"
		out.RawString(prefix[1:])
		(in.Filter).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"sort\":"
		out.RawString(prefix)
		(in.Sort).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"pagination\":"
		out.RawString(prefix)
		(in.Pagination).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"films\":"
		out.RawString(prefix)
		(in.Films).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CategoryFilmsOut) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CategoryFilmsOut) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CategoryFilmsOut) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CategoryFilmsOut) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms4(l, v)
}
func easyjsonDea1808fDecodeServerAppFilms5(in *jlexer.Lexer, out *CategoryFilmsIn) {
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
		case "page":
			out.Page = int(in.Int())
		case "pages":
			out.Pages = int(in.Int())
		case "total":
			out.Total = int64(in.Int64())
		case "limit":
			out.Limit = int(in.Int())
		case "sortField":
			out.SortField = string(in.String())
		case "sortOrder":
			out.SortOrder = string(in.String())
		case "ratingFrom":
			if in.IsNull() {
				in.Skip()
				out.RatingFrom = nil
			} else {
				if out.RatingFrom == nil {
					out.RatingFrom = new(float64)
				}
				*out.RatingFrom = float64(in.Float64())
			}
		case "yearFrom":
			out.YearFrom = int(in.Int())
		case "yearTo":
			out.YearTo = int(in.Int())
		case "type":
			out.Type = string(in.String())
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]string, 0, 4)
					} else {
						out.Genres = []string{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Genres = append(out.Genres, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonDea1808fEncodeServerAppFilms5(out *jwriter.Writer, in CategoryFilmsIn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"page\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Page))
	}
	{
		const prefix string = ",\"pages\":"
		out.RawString(prefix)
		out.Int(int(in.Pages))
	}
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix)
		out.Int64(int64(in.Total))
	}
	{
		const prefix string = ",\"limit\":"
		out.RawString(prefix)
		out.Int(int(in.Limit))
	}
	if in.SortField != "" {
		const prefix string = ",\"sortField\":"
		out.RawString(prefix)
		out.String(string(in.SortField))
	}
	if in.SortOrder != "" {
		const prefix string = ",\"sortOrder\":"
		out.RawString(prefix)
		out.String(string(in.SortOrder))
	}
	if in.RatingFrom != nil {
		const prefix string = ",\"ratingFrom\":"
		out.RawString(prefix)
		out.Float64(float64(*in.RatingFrom))
	}
	if in.YearFrom != 0 {
		const prefix string = ",\"yearFrom\":"
		out.RawString(prefix)
		out.Int(int(in.YearFrom))
	}
	if in.YearTo != 0 {
		const prefix string = ",\"yearTo\":"
		out.RawString(prefix)
		out.Int(int(in.YearTo))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if len(in.Genres) != 0 {
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Genres {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CategoryFilmsIn) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDea1808fEncodeServerAppFilms5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CategoryFilmsIn) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDea1808fEncodeServerAppFilms5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CategoryFilmsIn) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDea1808fDecodeServerAppFilms5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CategoryFilmsIn) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDea1808fDecodeServerAppFilms5(l, v)
}
