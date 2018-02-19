package ink

import (
	"bytes"
	"fmt"
	"image"
)

func NewLog(r image.Rectangle, sz int) *Log {
	f := OpenFont(DefaultFontMono, sz, true)
	return &Log{
		clip: r, font: f, h: sz,
	}
}

type Log struct {
	clip    image.Rectangle
	font    *Font
	lines   []string // lines buffer
	w       int      // font width (since we use monospaced font)
	h       int      // font height
	Spacing int      // line spacing
}

func (l *Log) Close() {
	l.font.Close()
	l.lines = nil
}
func (l *Log) appendLine(line string) {
	h := l.clip.Size().Y
	lh := l.h + l.Spacing
	n := h / lh
	if h%lh != 0 {
		n++
	}
	if dn := len(l.lines) + 1 - n; dn > 0 {
		// remove exceeding lines
		copy(l.lines, l.lines[dn:])
		l.lines = l.lines[:len(l.lines)-dn]
	}
	l.lines = append(l.lines, line)
}
func (l *Log) Write(p []byte) (int, error) {
	lines := bytes.Split(p, []byte{'\n'})
	if len(lines) != 0 && len(l.lines) != 0 {
		// append string to the last line
		li := len(l.lines) - 1
		last := l.lines[li]
		l.lines = l.lines[:li]
		l.appendLine(last + string(lines[0]))
		lines = lines[1:]
	}
	// add new lines
	for _, line := range lines {
		l.appendLine(string(line))
	}
	return len(p), nil
}
func (l *Log) Draw() {
	l.font.SetActive(Black)
	if l.w == 0 {
		l.w = CharWidth('a')
	}
	FillArea(l.clip, White)
	h := l.clip.Size().Y
	if h < l.h {
		DrawString(l.clip.Min, "window size is too small")
		return
	}
	for i := len(l.lines) - 1; i >= 0; i-- {
		s := l.lines[i]
		h -= l.h + l.Spacing
		if h < 0 {
			break
		}
		if s == "" {
			continue
		}
		p := image.Pt(0, h)
		DrawString(p.Add(l.clip.Min), s)
	}
}

func (l *Log) WriteString(s string) error {
	_, err := l.Write([]byte(s))
	return err
}

func (l *Log) Println(args ...interface{}) error {
	return l.WriteString(fmt.Sprint(args...))
}

func (l *Log) Printf(format string, args ...interface{}) error {
	return l.WriteString(fmt.Sprintf(format, args...))
}
