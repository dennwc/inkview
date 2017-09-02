package ink

import (
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
	lines   []string
	h       int
	Spacing int
}

func (l *Log) Close() {
	l.font.Close()
	l.lines = nil
}

func (l *Log) Draw() {
	l.font.SetActive(Black)
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
			n := len(l.lines)
			copy(l.lines, l.lines[i+1:])
			l.lines = l.lines[:n-(i+1)]
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
	l.lines = append(l.lines, s)
	return nil
}

func (l *Log) Println(args ...interface{}) error {
	return l.WriteString(fmt.Sprint(args...))
}

func (l *Log) Printf(format string, args ...interface{}) error {
	return l.WriteString(fmt.Sprintf(format, args...))
}
