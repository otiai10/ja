package ja

import (
	"unicode/utf8"
)

// Strategy ...
type Strategy interface {
	Accept(chunk string) (ok bool)
	Newline(chunk string)
	Flush() []string
}

type WithinStrategy struct {
	current string
	lines   []string
	Max     int
}

func (ws *WithinStrategy) Accept(chunk string) (ok bool) {
	if utf8.RuneCountInString(ws.current+chunk) > ws.Max {
		return false
	}
	ws.current += chunk
	return true
}

func (ws *WithinStrategy) Newline(chunk string) {
	ws.lines = append(ws.lines, ws.current)
	ws.current = chunk
}

func (ws *WithinStrategy) Flush() []string {
	if len(ws.current) > 0 {
		ws.lines = append(ws.lines, ws.current)
	}
	lines := ws.lines
	ws.current = ""
	ws.lines = []string{}
	return lines
}

func Within(max int) Strategy {
	return &WithinStrategy{
		current: "",
		lines:   []string{},
		Max:     max,
	}
}

type MoreThanStrategy struct {
	current string
	lines   []string
	Min     int
}

func (ms *MoreThanStrategy) Accept(chunk string) (ok bool) {
	if utf8.RuneCountInString(ms.current) > ms.Min {
		return false
	}
	ms.current += chunk
	return true
}

func (ms *MoreThanStrategy) Newline(chunk string) {
	ms.lines = append(ms.lines, ms.current)
	ms.current = chunk
}

func (ms *MoreThanStrategy) Flush() []string {
	if len(ms.current) > 0 {
		ms.lines = append(ms.lines, ms.current)
	}
	lines := ms.lines
	ms.current = ""
	ms.lines = []string{}
	return lines
}

func MoreThan(min int) Strategy {
	return &MoreThanStrategy{
		current: "",
		lines:   []string{},
		Min:     min,
	}
}
