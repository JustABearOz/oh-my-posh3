package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/text/unicode/norm"
)

func lenWithoutANSI(text, shell string) int {
	rANSI := "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	stripped := replaceAllString(rANSI, text, "")
	switch shell {
	case zsh:
		stripped = strings.ReplaceAll(stripped, "%{", "")
		stripped = strings.ReplaceAll(stripped, "%}", "")
	case bash:
		stripped = strings.ReplaceAll(stripped, "\\[", "")
		stripped = strings.ReplaceAll(stripped, "\\]", "")
	}
	var i norm.Iter
	i.InitString(norm.NFD, stripped)
	var count int
	for !i.Done() {
		i.Next()
		count++
	}
	return count
}

// AnsiRenderer exposes functionality using ANSI
type AnsiRenderer struct {
	buffer  *bytes.Buffer
	formats *ansiFormats
	shell   string
}

func (r *AnsiRenderer) carriageForward() {
	r.buffer.WriteString(fmt.Sprintf(r.formats.left, 1000))
}

func (r *AnsiRenderer) setCursorForRightWrite(text string, offset int) {
	strippedLen := lenWithoutANSI(text, r.shell) + -offset
	r.buffer.WriteString(fmt.Sprintf(r.formats.right, strippedLen))
}

func (r *AnsiRenderer) changeLine(numberOfLines int) {
	position := "B"
	if numberOfLines < 0 {
		position = "F"
		numberOfLines = -numberOfLines
	}
	r.buffer.WriteString(fmt.Sprintf(r.formats.linechange, numberOfLines, position))
}

func (r *AnsiRenderer) creset() {
	r.buffer.WriteString(r.formats.creset)
}

func (r *AnsiRenderer) print(text string) {
	r.buffer.WriteString(text)
	r.clearEOL()
}

func (r *AnsiRenderer) clearEOL() {
	r.buffer.WriteString(r.formats.clearOEL)
}

func (r *AnsiRenderer) string() string {
	return r.buffer.String()
}

func (r *AnsiRenderer) saveCursorPosition() {
	r.buffer.WriteString(r.formats.saveCursorPosition)
}

func (r *AnsiRenderer) restoreCursorPosition() {
	r.buffer.WriteString(r.formats.restoreCursorPosition)
}
