package ja

import (
	"strings"
)

func Paragraph(section string, strategy Strategy) (blocks [][]string) {
	for _, sentence := range strings.Split(section, "\n") {
		if lines := Linebreak(strings.TrimSpace(sentence), strategy); len(lines) != 0 {
			blocks = append(blocks, lines)
		}
	}
	return blocks
}

func Linebreak(sentence string, strategy Strategy) (lines []string) {
	for _, chunk := range Cut(sentence) {
		if ok := strategy.Accept(chunk); !ok {
			strategy.Newline(chunk)
		}
	}
	return strategy.Flush()
}
