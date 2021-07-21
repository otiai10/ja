package ja

import (
	"regexp"
	"strings"
)

var (
	punctuation = regexp.MustCompile("[、。]")
)

func Paragraph(section string, maxlen int) (blocks [][]string) {
	for _, sentence := range strings.Split(section, "\n") {
		lines := Linebreak(strings.TrimSpace(sentence), maxlen)
		if len(lines) != 0 {
			blocks = append(blocks, lines)
		}
	}
	return blocks
}

func Linebreak(sentence string, maxlen int) (lines []string) {
	ba := []byte(sentence)
	locs := punctuation.FindAllStringIndex(sentence, -1)
	templine := ""
	for i, l := range locs {
		if i == 0 {
			templine += string(ba[0:l[1]])
		} else {
			templine += string(ba[locs[i-1][1]:l[1]])
		}

		// Break if it's too long
		if len(templine) > maxlen {
			lines = append(lines, templine)
			templine = ""
		}
	}
	if templine != "" {
		lines = append(lines, templine)
	}
	return lines
}
