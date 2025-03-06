// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entity

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

func easyjson163c17a9DecodeFilmerServerInternalEntity(in *jlexer.Lexer, out *UserWithToken) {
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
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(User)
				}
				(*out.User).UnmarshalEasyJSON(in)
			}
		case "accessToken":
			out.AccessToken = string(in.String())
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
func easyjson163c17a9EncodeFilmerServerInternalEntity(out *jwriter.Writer, in UserWithToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		if in.User == nil {
			out.RawString("null")
		} else {
			(*in.User).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"accessToken\":"
		out.RawString(prefix)
		out.String(string(in.AccessToken))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity1(in *jlexer.Lexer, out *UserMovies) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UserMovies, 0, 0)
			} else {
				*out = UserMovies{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 UserMovie
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson163c17a9EncodeFilmerServerInternalEntity1(out *jwriter.Writer, in UserMovies) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserMovies) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserMovies) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserMovies) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserMovies) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity1(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity2(in *jlexer.Lexer, out *UserMovie) {
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
		case "status":
			out.Status = int8(in.Int8())
		case "stared":
			out.Stared = bool(in.Bool())
		case "movie":
			if in.IsNull() {
				in.Skip()
				out.Movie = nil
			} else {
				if out.Movie == nil {
					out.Movie = new(Movie)
				}
				(*out.Movie).UnmarshalEasyJSON(in)
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
func easyjson163c17a9EncodeFilmerServerInternalEntity2(out *jwriter.Writer, in UserMovie) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.Status))
	}
	{
		const prefix string = ",\"stared\":"
		out.RawString(prefix)
		out.Bool(bool(in.Stared))
	}
	if in.Movie != nil {
		const prefix string = ",\"movie\":"
		out.RawString(prefix)
		(*in.Movie).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserMovie) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserMovie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserMovie) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserMovie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity2(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity3(in *jlexer.Lexer, out *User) {
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
		case "email":
			out.Email = string(in.String())
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
func easyjson163c17a9EncodeFilmerServerInternalEntity3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity3(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity4(in *jlexer.Lexer, out *RawFilmStaffSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RawFilmStaffSlice, 0, 0)
			} else {
				*out = RawFilmStaffSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 RawFilmStaff
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson163c17a9EncodeFilmerServerInternalEntity4(out *jwriter.Writer, in RawFilmStaffSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v RawFilmStaffSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RawFilmStaffSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RawFilmStaffSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RawFilmStaffSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity4(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity5(in *jlexer.Lexer, out *RawFilmStaff) {
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
		case "staffId":
			out.StaffID = int(in.Int())
		case "nameRu":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "professionKey":
			out.ProfessionKey = string(in.String())
		case "posterUrl":
			out.ImgUrl = string(in.String())
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
func easyjson163c17a9EncodeFilmerServerInternalEntity5(out *jwriter.Writer, in RawFilmStaff) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"staffId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.StaffID))
	}
	{
		const prefix string = ",\"nameRu\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"professionKey\":"
		out.RawString(prefix)
		out.String(string(in.ProfessionKey))
	}
	{
		const prefix string = ",\"posterUrl\":"
		out.RawString(prefix)
		out.String(string(in.ImgUrl))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RawFilmStaff) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RawFilmStaff) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RawFilmStaff) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RawFilmStaff) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity5(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity6(in *jlexer.Lexer, out *RawFilmInfo) {
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
		case "kinopoiskId":
			out.KinopoiskID = int(in.Int())
		case "nameRu":
			out.Title = string(in.String())
		case "posterUrlPreview":
			out.PosterURL = string(in.String())
		case "webUrl":
			out.WebURL = string(in.String())
		case "ratingKinopoisk":
			out.RatingKinopoisk = float64(in.Float64())
		case "year":
			out.Year = int(in.Int())
		case "filmLength":
			out.FilmLenMinutes = int(in.Int())
		case "description":
			out.Description = string(in.String())
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
						out.Genres = make([]Genre, 0, 0)
					} else {
						out.Genres = []Genre{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Genre
					(v7).UnmarshalEasyJSON(in)
					out.Genres = append(out.Genres, v7)
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
func easyjson163c17a9EncodeFilmerServerInternalEntity6(out *jwriter.Writer, in RawFilmInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"kinopoiskId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.KinopoiskID))
	}
	{
		const prefix string = ",\"nameRu\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"posterUrlPreview\":"
		out.RawString(prefix)
		out.String(string(in.PosterURL))
	}
	{
		const prefix string = ",\"webUrl\":"
		out.RawString(prefix)
		out.String(string(in.WebURL))
	}
	{
		const prefix string = ",\"ratingKinopoisk\":"
		out.RawString(prefix)
		out.Float64(float64(in.RatingKinopoisk))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Int(int(in.Year))
	}
	{
		const prefix string = ",\"filmLength\":"
		out.RawString(prefix)
		out.Int(int(in.FilmLenMinutes))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		if in.Genres == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Genres {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RawFilmInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RawFilmInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RawFilmInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RawFilmInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity6(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity7(in *jlexer.Lexer, out *Person) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "role":
			if in.IsNull() {
				in.Skip()
				out.Role = nil
			} else {
				if out.Role == nil {
					out.Role = new(string)
				}
				*out.Role = string(in.String())
			}
		case "imgUrl":
			out.ImgUrl = string(in.String())
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
func easyjson163c17a9EncodeFilmerServerInternalEntity7(out *jwriter.Writer, in Person) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		if in.Role == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Role))
		}
	}
	{
		const prefix string = ",\"imgUrl\":"
		out.RawString(prefix)
		out.String(string(in.ImgUrl))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Person) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Person) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Person) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Person) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity7(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity8(in *jlexer.Lexer, out *Movie) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "kinopoiskID":
			out.KinopoiskID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "imgURL":
			out.ImgURL = string(in.String())
		case "webURL":
			out.WebURL = string(in.String())
		case "rating":
			out.Rating = float64(in.Float64())
		case "year":
			out.Year = int(in.Int())
		case "filmLength":
			out.FilmLength = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "personal":
			if in.IsNull() {
				in.Skip()
				out.Personal = nil
			} else {
				if out.Personal == nil {
					out.Personal = new(FilmStaff)
				}
				(*out.Personal).UnmarshalEasyJSON(in)
			}
		case "updatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]Genre, 0, 0)
					} else {
						out.Genres = []Genre{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Genre
					(v10).UnmarshalEasyJSON(in)
					out.Genres = append(out.Genres, v10)
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
func easyjson163c17a9EncodeFilmerServerInternalEntity8(out *jwriter.Writer, in Movie) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"kinopoiskID\":"
		out.RawString(prefix)
		out.Int(int(in.KinopoiskID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"imgURL\":"
		out.RawString(prefix)
		out.String(string(in.ImgURL))
	}
	if in.WebURL != "" {
		const prefix string = ",\"webURL\":"
		out.RawString(prefix)
		out.String(string(in.WebURL))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Int(int(in.Year))
	}
	if in.FilmLength != "" {
		const prefix string = ",\"filmLength\":"
		out.RawString(prefix)
		out.String(string(in.FilmLength))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Personal != nil {
		const prefix string = ",\"personal\":"
		out.RawString(prefix)
		(*in.Personal).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		if in.Genres == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Genres {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Movie) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Movie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Movie) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Movie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity8(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity9(in *jlexer.Lexer, out *Genre) {
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
		case "genre":
			out.Genre = string(in.String())
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
func easyjson163c17a9EncodeFilmerServerInternalEntity9(out *jwriter.Writer, in Genre) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"genre\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Genre))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Genre) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Genre) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Genre) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Genre) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity9(l, v)
}
func easyjson163c17a9DecodeFilmerServerInternalEntity10(in *jlexer.Lexer, out *FilmStaff) {
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
		case "directors":
			if in.IsNull() {
				in.Skip()
				out.Directors = nil
			} else {
				in.Delim('[')
				if out.Directors == nil {
					if !in.IsDelim(']') {
						out.Directors = make([]Person, 0, 1)
					} else {
						out.Directors = []Person{}
					}
				} else {
					out.Directors = (out.Directors)[:0]
				}
				for !in.IsDelim(']') {
					var v13 Person
					(v13).UnmarshalEasyJSON(in)
					out.Directors = append(out.Directors, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "actors":
			if in.IsNull() {
				in.Skip()
				out.Actors = nil
			} else {
				in.Delim('[')
				if out.Actors == nil {
					if !in.IsDelim(']') {
						out.Actors = make([]Person, 0, 1)
					} else {
						out.Actors = []Person{}
					}
				} else {
					out.Actors = (out.Actors)[:0]
				}
				for !in.IsDelim(']') {
					var v14 Person
					(v14).UnmarshalEasyJSON(in)
					out.Actors = append(out.Actors, v14)
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
func easyjson163c17a9EncodeFilmerServerInternalEntity10(out *jwriter.Writer, in FilmStaff) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"directors\":"
		out.RawString(prefix[1:])
		if in.Directors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.Directors {
				if v15 > 0 {
					out.RawByte(',')
				}
				(v16).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"actors\":"
		out.RawString(prefix)
		if in.Actors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Actors {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilmStaff) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeFilmerServerInternalEntity10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilmStaff) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeFilmerServerInternalEntity10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilmStaff) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeFilmerServerInternalEntity10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilmStaff) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeFilmerServerInternalEntity10(l, v)
}
