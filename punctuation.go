package ja

import "regexp"

var (
	punctuation = regexp.MustCompile("[、。]")
)

func Cut(sentence string) (chunks []string) {
	ba := []byte(sentence)
	locations := punctuation.FindAllStringIndex(sentence, -1)
	for i, loc := range locations {
		if i == 0 {
			chunks = append(chunks, string(ba[0:loc[1]]))
		} else {
			chunks = append(chunks, string(ba[locations[i-1][1]:loc[1]]))
		}
	}
	return chunks
}
