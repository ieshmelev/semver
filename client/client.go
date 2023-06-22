package client

import (
	"bytes"

	"github.com/rs/zerolog"
)

func New() {
	b := []byte{}
	w := bytes.NewBuffer(b)
	zerolog.New(w)
}
