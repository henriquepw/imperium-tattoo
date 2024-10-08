package customid

import gonanoid "github.com/matoous/go-nanoid/v2"

func createID(size int) (string, error) {
	return gonanoid.Generate("123456789ABCDEFGHIJKLMNPQRSTUVWXYZ", size)
}

func New() (string, error) {
	return createID(18)
}

func NewTiny() (string, error) {
	return createID(5)
}
