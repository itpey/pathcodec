package pathcodec

import (
	"errors"
	"fmt"
)

var charToIndex = [256]byte{}
var indexToChar = "AACAAAAHAAALMAAAQASTAVAAAZaacaaaahaaalmaaaqastava.az0123456789-,"

func init() {
	for i := range charToIndex {
		charToIndex[i] = 255
	}
	for i, c := range indexToChar {
		charToIndex[c] = byte(i)
	}
}

func Compress(path string) ([]byte, error) {
	if len(path) < 2 || path[0] != 'M' || path[len(path)-1] != 'z' {
		return nil, errors.New("invalid input: must start with 'M' and end with 'z'")
	}

	length := len(path) - 2 // Exclude 'M' and 'z'
	encoded := make([]byte, length)

	for i := 1; i < len(path)-1; i++ {
		char := path[i]
		if char == ',' {
			encoded[i-1] = 128
		} else if char == '-' {
			encoded[i-1] = 64
		} else if char >= '0' && char <= '9' {
			encoded[i-1] = char - '0'
		} else if charToIndex[char] == 255 {
			return nil, fmt.Errorf("invalid character '%c' in input", char)
		} else {
			encoded[i-1] = charToIndex[char] + 192
		}
	}
	return encoded, nil
}

func Decompress(encoded []byte) (string, error) {
	if len(encoded) == 0 {
		return "", errors.New("invalid input: encoded data cannot be empty")
	}
	length := len(encoded) + 2
	path := make([]byte, length)
	path[0] = 'M'
	path[length-1] = 'z'

	for i, num := range encoded {
		num = num & 0xff
		if num >= 192 {
			if int(num-192) >= len(indexToChar) {
				return "", fmt.Errorf("invalid encoded byte: %d", num)
			}
			path[i+1] = indexToChar[num-192]
		} else if num >= 128 {
			path[i+1] = ','
		} else if num >= 64 {
			path[i+1] = '-'
		} else if num <= 9 {
			path[i+1] = '0' + num
		} else {
			return "", fmt.Errorf("invalid encoded byte: %d", num)
		}
	}
	return string(path), nil
}
