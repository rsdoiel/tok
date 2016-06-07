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
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestPunctuation(t *testing.T) {
	alpha := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	for _, a := range alpha {
		if IsPunctuation([]byte(a)) == true {
			t.Errorf("Failed IsPunctuation(%q)", a)
		}
		if IsSpace([]byte(a)) == true {
			t.Errorf("Failed IsSpace(%q)", a)
		}
	}
}

func TestTok(t *testing.T) {
	// Test Tok()
	fname1 := path.Join("testdata", "sample-0.txt")
	src1, err := ioutil.ReadFile(fname1)
	if err != nil {
		t.Errorf("%s, %s", fname1, err)
		t.FailNow()
	}
	// FIXME: Load expected-0.txt to compare token types.
	fname2 := path.Join("testdata", "expected-0.txt")
	src2, err := ioutil.ReadFile(fname2)
	if err != nil {
		t.Errorf("%s, %s", fname2, err)
		t.FailNow()
	}
	expected := strings.Split(strings.TrimSpace(string(src2)), "\n")

	// FIXME: Need to create a tokenizing function which takes a buffer and token mapping and produces a token structure
	// with a type field and the value of the token.
	var (
		token *Token
		i     int
	)
	for i, expectedType := range expected {
		token, src1 = Tok(src1)
		// fmt.Printf("DEBUG i: %d, token: %s, expectedType: %s, src1: %s\n", i, token, expectedType, src1)
		if strings.Compare(token.Type, strings.TrimSpace(expectedType)) != 0 {
			t.Errorf("%d: %s != %s", i, token, expectedType)
		}
	}
	if len(src1) != 0 {
		t.Errorf("Expected to have len(src1) == 1, %d", i)
	}

	// Test Tok2()
	src1, _ = ioutil.ReadFile(fname1)
	for i, expectedType := range expected {
		token, src1 = Tok2(src1, func(t *Token, b []byte) (*Token, []byte) {
			// This is just a pass through function, normally you'd add additional analysis
			return t, b
		})
		// fmt.Printf("DEBUG i: %d, token: %s, expectedType: %s, src1: %s\n", i, token, expectedType, src1)
		if strings.Compare(token.Type, strings.TrimSpace(expectedType)) != 0 {
			t.Errorf("%d: %s != %s", i, token, expectedType)
		}
	}
	if len(src1) != 0 {
		t.Errorf("Expected to have len(src1) == 1, %d", i)
	}
}
