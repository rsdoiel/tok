//
// Package tok is a niave tokenizer
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of tok nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package tok

import (
	"encoding/xml"
	"fmt"
)

const (
	// Punctuation Marks
	Tilda        = "~"
	Bang         = "!"
	AtSign       = "@"
	HashMark     = "#"
	DollarSign   = "$"
	PercentSign  = "%"
	Caret        = "^"
	Ampersand    = "&"
	Asterisk     = "*"
	OpenParen    = "("
	CloseParen   = ")"
	MinusSign    = "-"
	PlusSign     = "+"
	BackTick     = "`"
	Underscore   = "_"
	EqualSign    = "="
	OpenCurly    = "{"
	CloseCurly   = "}"
	Pipe         = "|"
	OpenSquare   = "["
	CloseSquare  = "]"
	BackSlash    = "\\"
	Colon        = ":"
	SemiColon    = ";"
	DoubleQuote  = "\""
	SingleQuote  = "'"
	GreaterThan  = ">"
	LessThan     = "<"
	QuestionMark = "?"
	Comma        = ","
	Period       = "."
	Slash        = "/"

	// Spaces
	Space = " "
	Tab   = "\t"
	CR    = "\r"
	LF    = "\n"

	// Numerals
	Zero  = "0"
	One   = "1"
	Two   = "2"
	Three = "3"
	Four  = "4"
	Five  = "5"
	Six   = "6"
	Seven = "7"
	Eight = "8"
	Nine  = "9"
)

var (
	Numerals = map[string]string{
		Zero:  "0",
		One:   "1",
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
	}

	Spaces = map[string]string{
		Space: " ",
		Tab:   "\t",
		CR:    "\r",
		LF:    "\n",
	}

	PunctuationMarks = map[string]string{
		Tilda:        "~",
		Bang:         "!",
		AtSign:       "@",
		HashMark:     "#",
		DollarSign:   "$",
		PercentSign:  "%",
		Caret:        "^",
		Ampersand:    "&",
		Asterisk:     "*",
		OpenParen:    "(",
		CloseParen:   ")",
		MinusSign:    "-",
		PlusSign:     "+",
		BackTick:     "`",
		Underscore:   "_",
		EqualSign:    ":",
		OpenCurly:    "{",
		CloseCurly:   "}",
		Pipe:         "|",
		OpenSquare:   "[",
		CloseSquare:  "]",
		BackSlash:    "\\",
		Colon:        ":",
		SemiColon:    ";",
		DoubleQuote:  "\"",
		SingleQuote:  "'",
		GreaterThan:  ">",
		LessThan:     "<",
		QuestionMark: "?",
		Comma:        ",",
		Period:       ".",
		Slash:        "/",
	}
)

type Token struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:"type" json:"type"`
	Value   []byte   `xml:"value" json:"value"`
}

type Tokenizer func(*Token, []byte) (*Token, []byte)

func (t *Token) String() string {
	return fmt.Sprintf("<%s> = %q", t.Type, t.Value)
}

func IsSpace(b []byte) bool {
	_, ok := Spaces[string(b)]
	return ok
}

func IsPunctuation(b []byte) bool {
	_, ok := PunctuationMarks[string(b)]
	return ok
}

func IsNumeral(b []byte) bool {
	_, ok := Numerals[string(b)]
	return ok
}

func Tok(buf []byte) (*Token, []byte) {
	var (
		s []byte
	)
	if len(buf) == 0 {
		return &Token{
			Type:  "EOF",
			Value: []byte(""),
		}, nil
	}
	s, buf = buf[0:1], buf[1:]
	switch {
	case IsPunctuation(s) == true:
		return &Token{
			Type:  "Punctuation",
			Value: s,
		}, buf
	case IsSpace(s) == true:
		return &Token{
			Type:  "Space",
			Value: s,
		}, buf
	case IsNumeral(s) == true:
		return &Token{
			Type:  "Numeric",
			Value: s,
		}, buf
	default:
		return &Token{
			Type:  "Alpha",
			Value: s,
		}, buf
	}
}

func Tok2(buf []byte, fn Tokenizer) (*Token, []byte) {
	tok, rest := Tok(buf)
	return fn(tok, rest)
}
