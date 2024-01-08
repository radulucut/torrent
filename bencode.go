package main

type Token interface{}

func readDict(start int, b []byte) (map[string]Token, int) {
	tokens := make(map[string]Token)
	i := start + 1
	for i < len(b) {
		switch {
		case b[i] == 'e':
			return tokens, i + 1
		default:
			l, end := readCount(i, b)
			i = end + l
			key := string(b[end:i])
			if key == "" {
				break
			}
			switch {
			case b[i] == 'd':
				t, end := readDict(i, b)
				tokens[key] = t
				i = end
			case b[i] == 'l':
				t, end := readList(i, b)
				tokens[key] = t
				i = end
			case b[i] == 'i':
				chunk := readInt(i, b)
				tokens[key] = chunk
				i = chunk[1] + 1 // skip 'e'
			default:
				l, end := readCount(i, b)
				i = end + l
				tokens[key] = [2]int{end, i}
			}
		}
	}
	return tokens, i
}

func readInt(start int, b []byte) [2]int {
	i := start + 1
	for i < len(b) && b[i] != 'e' {
		i++
	}
	return [2]int{start + 1, i}
}

func readCount(start int, b []byte) (int, int) {
	i := start
	var n int
	for i < len(b) && isDigit(b[i]) {
		n = n*10 + int(b[i]-'0')
		i++
	}
	if i < len(b) && b[i] == ':' {
		i++
	}
	return n, i
}

func readList(start int, b []byte) ([]Token, int) {
	tokens := make([]Token, 0)
	i := start + 1
	for i < len(b) {
		switch {
		case b[i] == 'e':
			return tokens, i + 1
		case b[i] == 'd':
			t, end := readDict(i, b)
			tokens = append(tokens, t)
			i = end
		case b[i] == 'l':
			t, end := readList(i, b)
			tokens = append(tokens, t)
			i = end
		case b[i] == 'i':
			chunk := readInt(i, b)
			tokens = append(tokens, chunk)
			i = chunk[1] + 1 // skip 'e'
		default:
			l, end := readCount(i, b)
			i = end + l
			tokens = append(tokens, [2]int{end, i})
		}
	}
	return tokens, i
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
