package upload

import (
	"io"
)

type Client interface {
	Add(r io.Reader) (string, error)
}
