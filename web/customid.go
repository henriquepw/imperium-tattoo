package web

import gonanoid "github.com/matoous/go-nanoid/v2"

func createID(size int) (string, error) {
	return gonanoid.Generate("123456789ABCDEFGHIJKLMNPQRSTUVWXYZ", size)
}

func NewID() (string, error) {
	return createID(18)
}

func NewTinyID() (string, error) {
	return createID(5)
}
