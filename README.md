[![Go Report Card](http://goreportcard.com/badge/rsdoiel/tok)](http://goreportcard.com/report/rsdoiel/tok)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)


# tok

A niave tokenizer library

## Public Interface

+ Token - a simple structure 
    + properties
        + Type is a string holding the label of the token type
        + Value is a byte array holding the value of the token
+ Tokenizer - is a type of function that can be applied by Tok2, may be recursive
    + parameters
        + byte array
        + a Tokenizer function
    + returns
        + Token
        + byte array of remaining buffer
+ Tok - is a simple, non-look ahead tokenizer
    + parameter
        + a byte array representing the buffer to evaluate
    + returns
        + a Token of Type *Letter*, *Numeral*, *Punctuation* and *Space*
        + the remaining buffer byte array
+ Tok2 - is a function the take
    + parameters
        + a byte array representing the buffer to evaluate
        + A Tokenizer function
    + returns
        + a Token of Type defined by the Tokenizer function
        + the remaining buffer byte array
+ ToksWords - Is an example Tokenizer function
    + returns tokens of type *Numeral*, *Punctuation*, *Space* and *Word*

