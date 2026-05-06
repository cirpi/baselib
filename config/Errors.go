package config

import (
	"errors"
	"fmt"
)

var invalidConfigPathEr = func(path string) error {
	msg := fmt.Sprintf("invalid format: %s, something like config.json", path)
	return errors.New(msg)
}

var noDecoderEr = func(ext string) error {
	msg := fmt.Sprintf("No Decoder found for %s. Add your custom decoder using AddConfigDecoder.", ext)
	return errors.New(msg)
}
