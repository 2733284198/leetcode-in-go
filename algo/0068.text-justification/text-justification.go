package problem0068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	l := len(words)

	r := []string{}
	w := []string{}
	t := 0

	for i := 0; i < l; i++ {
		t = t + len(words[i])
		w = append(w, words[i])

		// has words left
		if i+1 < l {
			// words-size + spaces-size + next-word-size
			if t+len(w)+len(words[i+1]) > maxWidth {
				r = append(r, typeSetting(w, t, maxWidth, false))
				t = 0
				w = []string{}
			}
		} else {
			r = append(r, typeSetting(w, t, maxWidth, true))
		}
	}

	return r
}

func typeSetting(words []string, t int, maxWidth int, end bool) string {
	if end {
		endLine := strings.Join(words, " ")
		return endLine + spaces(maxWidth-len(endLine))
	}

	if len(words) == 1 {
		return words[0] + spaces(maxWidth-t)
	}

	spacesCount := len(words) - 1
	spacesLength := int((maxWidth - t) / spacesCount)
	tails := (maxWidth - t) % spacesCount

	for i := 0; i < tails; i++ {
		words[i] = words[i] + " "
	}

	s := spaces(spacesLength)

	return strings.Join(words, s)
}

func spaces(n int) string {
	r := ""

	for n > 0 {
		r += " "
		n--
	}

	return r
}
