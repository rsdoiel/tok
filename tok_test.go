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
	"fmt"
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

func TestParse(t *testing.T) {
	// FIXME: Need to create a tokenizing function which takes a buffer and token mapping and produces a token structure
	// with a type field and the value of the token.
	src := []byte(`This is a test of the emergency broadcast system. It is only
a test. Should this be an actual emergency then this message
would be followed by instructions. Get it?
`)
	token := new(Token)
	l := len(src)
	for {
		l = len(src)
		token, src = Tok(src)
		if l != len(src)+1 {
			t.Errorf("Failed to shrink src: %s", src)
			t.FailNow()
		}
		fmt.Printf("DEBUG Tok: %s <-- %s\n", token, src)
		if l == 1 {
			break
		}
	}

}
