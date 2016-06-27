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
	"bytes"
	"encoding/xml"
	"fmt"
)

const (
	// Version of  tok package
	Version = "0.0.0"

	//Base token types, these are single letter tokens, no look ahead needed

	// Letter is an alphabetical letter (e.g. A-Z, a-z in English)
	Letter = "Letter"
	// Numeral is a single digit
	Numeral = "Numeral"
	// Punctuation is any non-number, non alphametical character, non-space (e.g. periods, colons, bang, hash mark)
	Punctuation = "Punctuation"
	// Space characters representing white space (e.g. space, tab, new line, carriage return)
	Space = "Space"

	// These are some common specialized token types provided for convientent.

	// Words a sequence of characters delimited by spaces
	Word = "Word"
	// OpenCurly bracket, e.g. "{"
	OpenCurlyBracket = "OpenCurlyBracket"
	// CloseCurly bracket, e.g. "}"
	CloseCurlyBracket = "CloseCurlyBracket"
	// CurlyBracket, e.g. "{}"
	CurlyBracket = "CurlyBracket"
	// OpenSquareBracket, e.g. "["
	OpenSquareBracket = "OpenSquareBracket"
	// CloseSquareBracket, e.g. "]"
	CloseSquareBracket = "CloseSquareBracket"
	// SquareBracket, e.g. "[]"
	SquareBracket = "SquareBracket"
	// OpenAngleBracket, e.g. "<"
	OpenAngleBracket = "OpenAngleBracket"
	// CloseAngleBracket, e.g. ">"
	CloseAngleBracket = "CloseAngleBracket"
	// AngleBracket, e.g. "<>"
	AngleBracket = "AngleBracket"
	// AtSign, e.g. "@"
	AtSign = "AtSign"
	// EqualSign, e.g. "="
	EqualSign = "EqualSign"
	// DoubleQuote, e.g. "\""
	DoubleQuote = "DoubleQuote"
	// SingleQuote, e.g., "'"
	SingleQuote = "SingleQuote"

	// EOF is an end of file token type. It is separate form Space only because of it being a common stop condition
	EOF = "EOF"
)

var (
	// Numerals is a map of numbers as strings
	Numerals = []byte("0123456789")

	// Spaces is a map space symbols as strings
	Spaces = []byte(" \t\r\n")

	// PunctuationMarks map as strings
	PunctuationMarks = []byte("~!@#$%^&*()_+`-=:{}|[]\\:;\"'<>?,./")

	// These map to the specialized tokens
	AtSignMark = []byte("@")
	// EqualMark, e.g. =
	EqualMark = []byte("=")
	// DoubleQuoteMark, e.g. "\""
	DoubleQuoteMark = []byte("\"")
	// SingleQuoteMark, e.g. "'"
	SingleQuoteMark = []byte("'")

	// OpenCurlyBrackets token
	OpenCurlyBrackets = []byte("{")
	// CloseCurlyBrackets token
	CloseCurlyBrackets = []byte("}")
	// CurlyBrackets tokens
	CurlyBrackets = []byte("{}")

	// OpenSquareBrackets token
	OpenSquareBrackets = []byte("[")
	// CloseSquareBrackets token
	CloseSquareBrackets = []byte("]")
	// SquareBrackets tokens
	SquareBrackets = []byte("[]")

	// OpenAngleBrackets token
	OpenAngleBrackets = []byte("<")
	// CloseAngleBrackets token
	CloseAngleBrackets = []byte(">")
	// AngleBrackets tokens
	AngleBrackets = []byte("<>")
)

// Token structure for emitting simply tokens and value from Tok() and Tok2()
type Token struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:"type" json:"type"`
	Value   []byte   `xml:"value" json:"value"`
}

// TokenMap is a map of simple token names and associated array of possible bytes
type TokenMap map[string][]byte

// Tokenizer is a function that takes a current token, looks ahead in []byte and returns a revised token and remaining []byte
type Tokenizer func(*Token, []byte) (*Token, []byte)

// String returns a human readable Token struct
func (t *Token) String() string {
	return fmt.Sprintf("<%s> = %q", t.Type, t.Value)
}

// IsSpace checks to see if []byte is a space or not
func IsSpace(b []byte) bool {
	return bytes.Contains(Spaces, b)
}

// IsPunctuation checks to see if []byte is some punctuation or not
func IsPunctuation(b []byte) bool {
	return bytes.Contains(PunctuationMarks, b)
}

// IsNumeral checks to see if []byte is a number or not
func IsNumeral(b []byte) bool {
	return bytes.Contains(Numerals, b)
}

// IsTypeOf checks to see if a []byte is a specific type
func IsTypeOf(b []byte, typeMap []byte) bool {
	return bytes.Contains(typeMap, b)
}

// Tok is a naive tokenizer that looks only at the next character by shifting it off the []byte and returning a token found with remaining []byte
func Tok(buf []byte) (*Token, []byte) {
	var (
		s []byte
	)
	if len(buf) == 0 {
		return &Token{
			Type:  EOF,
			Value: []byte(""),
		}, nil
	}
	s, buf = buf[0:1], buf[1:]
	switch {
	case IsPunctuation(s) == true:
		return &Token{
			Type:  Punctuation,
			Value: s,
		}, buf
	case IsSpace(s) == true:
		return &Token{
			Type:  Space,
			Value: s,
		}, buf
	case IsNumeral(s) == true:
		return &Token{
			Type:  Numeral,
			Value: s,
		}, buf
	default:
		return &Token{
			Type:  Letter,
			Value: s,
		}, buf
	}
}

// TokenFromMap, revaluates token type against a map of type names and byte arrays
// returns modified Token
func TokenFromMap(t *Token, m map[string][]byte) *Token {
	for k, v := range m {
		if bytes.Contains(v, t.Value) {
			return &Token{
				Type:  k,
				Value: t.Value,
			}
		}
	}
	return &Token{
		Type:  t.Type,
		Value: t.Value,
	}
}

// Tok2 provides an easy to implement look ahead tokenizer by defining a look ahead function
func Tok2(buf []byte, fn Tokenizer) (*Token, []byte) {
	tok, rest := Tok(buf)
	return fn(tok, rest)
}

// Words is an example of implementing a Tokenizer function
func Words(tok *Token, buf []byte) (*Token, []byte) {
	if tok.Type == Letter || tok.Type == Word {
		// Get the next Token
		newTok, newBuf := Tok(buf)
		if newTok.Type == Letter {
			tok.Type = Word
			tok.Value = append(tok.Value, newTok.Value[0])
			tok, buf = Words(tok, newBuf)
		}
	}
	return tok, buf
}

// Bib is an example of a niave BibTeX Tokenizer function
func Bib(tok *Token, buf []byte) (*Token, []byte) {
	switch {
	case tok.Type == AtSign || tok.Type == "BibElement":
		// Get the next Token
		newTok, newBuf := Tok(buf)
		if newTok.Type == Letter {
			tok.Type = "BibElement"
			tok.Value = append(tok.Value, newTok.Value[0])
			tok, buf = Bib(tok, newBuf)
		}
	case tok.Type == Letter || tok.Type == Word:
		// Get the next Token
		newTok, newBuf := Tok(buf)
		if newTok.Type == Letter {
			tok.Type = Word
			tok.Value = append(tok.Value, newTok.Value[0])
			tok, buf = Bib(tok, newBuf)
		}
	default:
		// Revaluate token for more specific token types.
		tok = TokenFromMap(tok, map[string][]byte{
			OpenCurlyBracket:  OpenCurlyBrackets,
			CloseCurlyBracket: CloseCurlyBrackets,
			AtSign:            AtSignMark,
			EqualSign:         EqualMark,
			DoubleQuote:       DoubleQuoteMark,
			SingleQuote:       SingleQuoteMark,
			"Comma":           []byte(","),
		})
	}

	return tok, buf
}
