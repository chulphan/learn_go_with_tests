package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

// Go에서의 map 은 runtime.hmap 이고 이는 포인터값이므로 &를 사용할 필요 없다(배열, 슬라이스와 달리)
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
