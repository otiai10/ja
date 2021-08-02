package ja

import (
	"regexp"
	"strings"
)

var (
	punctuation = regexp.MustCompile("[、。\n]")
	cleanset    = "\n  　"
)

func Cut(sentence string, clean ...bool) (chunks []string) {
	clean = append(clean, false)
	ba := []byte(sentence)
	locations := punctuation.FindAllStringIndex(sentence, -1)
	for i, loc := range locations {
		chunk := ""
		if i == 0 {
			chunk = string(ba[0:loc[1]])
		} else {
			chunk = string(ba[locations[i-1][1]:loc[1]])
		}
		if clean[0] {
			chunks = append(chunks, strings.Trim(chunk, cleanset))
		} else {
			chunks = append(chunks, chunk)
		}
	}
	return chunks
}
