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

	// Tilda as string
	Tilda = "~"
	// Bang as string
	Bang = "!"
	// AtSign as string
	AtSign = "@"
	// HashMark as string
	HashMark = "#"
	// DollarSign as string
	DollarSign = "$"
	// PercentSign as string
	PercentSign = "%"
	// Caret as string
	Caret = "^"
	// Ampersand as string
	Ampersand = "&"
	// Asterisk as string
	Asterisk = "*"
	// OpenParen as string
	OpenParen = "("
	// CloseParen as string
	CloseParen = ")"
	// MinusSign as string
	MinusSign = "-"
	// PlusSign as string
	PlusSign = "+"
	// BackTick as string
	BackTick = "`"
	// Underscore as string
	Underscore = "_"
	// EqualSign as string
	EqualSign = "="
	// OpenCurly (brackets) as string
	OpenCurly = "{"
	// CloseCurly (brackets) as string
	CloseCurly = "}"
	// Pipe symbol as string
	Pipe = "|"
	// OpenSquare (brackets) as string
	OpenSquare = "["
	// CloseSquare (brackets) as string
	CloseSquare = "]"
	// BackSlash as string
	BackSlash = "\\"
	// Colon as string
	Colon = ":"
	// SemiColon as string
	SemiColon = ";"
	// DoubleQuote as string
	DoubleQuote = "\""
	// SingleQuote as string
	SingleQuote = "'"
	// GreaterThan (sign) as string
	GreaterThan = ">"
	// LessThan (sign) as string
	LessThan = "<"
	// QuestionMark as string
	QuestionMark = "?"
	// Comma as string
	Comma = ","
	// Period as string
	Period = "."
	// Slash as string
	Slash = "/"

	// Spaces

	// Space as string
	Space = " "
	// Tab as string
	Tab = "\t"
	// CR (Carriage Return) as string
	CR = "\r"
	// LF (Line Feed) as string
	LF = "\n"

	// Numerals

	// Zero as string
	Zero = "0"
	// One as string
	One = "1"
	// Two as string
	Two = "2"
	// Three as string
	Three = "3"
	// Four as string
	Four = "4"
	// Five as string
	Five = "5"
	// Six as string
	Six = "6"
	// Seven as string
	Seven = "7"
	// Eight as string
	Eight = "8"
	// Nine as string
	Nine = "9"
)

var (
	// Numerals is a map of numbers as strings
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

	// Spaces is a map space symbols as strings
	Spaces = map[string]string{
		Space: " ",
		Tab:   "\t",
		CR:    "\r",
		LF:    "\n",
	}

	// PunctuationMarks map as strings
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

// Token structure for emitting simply tokens and value from Tok() and Tok2()
type Token struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:"type" json:"type"`
	Value   []byte   `xml:"value" json:"value"`
}

// Tokenizer is a function that takes a current token, looks ahead in []byte and returns a revised token and remaining []byte
type Tokenizer func(*Token, []byte) (*Token, []byte)

// String returns a human readable Token struct
func (t *Token) String() string {
	return fmt.Sprintf("<%s> = %q", t.Type, t.Value)
}

// IsSpace checks to see if []byte is a space or not
func IsSpace(b []byte) bool {
	_, ok := Spaces[string(b)]
	return ok
}

// IsPunctuation checks to see if []byte is some punctuation or not
func IsPunctuation(b []byte) bool {
	_, ok := PunctuationMarks[string(b)]
	return ok
}

// IsNumeral checks to see if []byte is a number or not
func IsNumeral(b []byte) bool {
	_, ok := Numerals[string(b)]
	return ok
}

// Tok is a naive tokenizer that looks only at the next character by shifting it off the []byte and returning a token found with remaining []byte
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

// Tok2 provides an easy to implement look ahead tokenizer by defining a look ahead function
func Tok2(buf []byte, fn Tokenizer) (*Token, []byte) {
	tok, rest := Tok(buf)
	return fn(tok, rest)
}
