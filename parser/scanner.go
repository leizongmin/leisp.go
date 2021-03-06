// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

import (
	"bufio"
	"bytes"
	"io"
)

const eof = rune(0)

type Position struct {
	Line   int
	Column int
}

type Scanner struct {
	r          *bufio.Reader
	Position   Position
	lastChar   rune
	lastColumn int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
		Position: Position{
			Line: 1,
		},
	}
}

func (s *Scanner) read() rune {

	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	s.lastColumn = s.Position.Column
	s.lastChar = ch
	if ch == '\n' {
		s.Position.Line++
		s.Position.Column = 0
	} else {
		s.Position.Column++
	}

	return ch
}

func (s *Scanner) unread() {

	_ = s.r.UnreadByte()

	if s.lastChar == '\n' {
		s.Position.Line--
		s.Position.Column = s.lastColumn
		s.lastColumn--
	} else {
		s.Position.Column--
	}
}

func (s *Scanner) Scan() (tok token, lit string) {

	ch := s.read()

	if isWhitesapce(ch) {
		s.unread()
		return s.scanWhitespace()
	}
	if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	switch ch {

	case eof:
		return tokenEOF, ""

	case '"':
		s.unread()
		return s.scanString()

	case ';':
		s.unread()
		return s.scanComment()

	case ':':
		s.unread()
		return s.scanKeyword()

	case '\'':
		return tokenQuote, string(ch)
	}

	if isPunctuation(ch) {
		return tokenPunctuation, string(ch)
	}

	s.unread()
	return s.scanSymbol()
}

func (s *Scanner) scanWhitespace() (tok token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitesapce(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokenWhitespace, buf.String()
}

func (s *Scanner) scanString() (tok token, lit string) {

	var buf bytes.Buffer
	s.read()

	for {
		if ch := s.read(); ch == eof {
			return tokenIllegal, buf.String()
		} else if ch == '\\' {
			buf.WriteRune(ch)
			ch = s.read()
			buf.WriteRune(ch)
		} else if ch == '"' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokenString, buf.String()
}

func (s *Scanner) scanNumber() (tok token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	isSlash := false
	isDot := false
	isE := false

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isDigit(ch) {
			buf.WriteRune(ch)
		} else if ch == '/' {
			if isSlash {
				return tokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isSlash = true
			}
		} else if ch == '.' {
			if isDot {
				return tokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isDot = true
			}
		} else if ch == 'e' || ch == 'E' {
			if isE {
				return tokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isE = true
			}
		} else if isWhitesapce(ch) || isPunctuation(ch) {
			s.unread()
			break
		} else {
			return tokenIllegal, string(ch)
		}
	}

	return tokenNumber, buf.String()
}

func (s *Scanner) scanSymbol() (tok token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitesapce(ch) {
			s.unread()
			break
		} else if isPunctuation(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokenSymbol, buf.String()
}

func (s *Scanner) scanKeyword() (tok token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitesapce(ch) {
			s.unread()
			break
		} else if ch == ':' {
			return tokenIllegal, string(ch)
		} else if isPunctuation(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokenKeyword, buf.String()
}

func (s *Scanner) scanComment() (tok token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '\n' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokenComment, buf.String()
}
