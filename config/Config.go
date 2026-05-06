package config

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path"

	"go.yaml.in/yaml/v3"
)

const (
	YAML_1 = "yaml"
	YAML_2 = "yml"
	JSON   = "json"
	XML    = "xml"
)

type DecodeFunc func([]byte, any) error

var configDecoders = map[string]DecodeFunc{
	YAML_1: yaml.Unmarshal,
	YAML_2: yaml.Unmarshal,
	JSON:   json.Unmarshal,
	XML:    xml.Unmarshal,
}

/*
Initialize the config using the predefined file formats.

Supported config type:
* YAML
* JSON
* XML

You can add your custom config decoder using AddConfigDecoder.
*/
func InitConfig[T any](cfgPath string, cfg *T) error {
	ext := path.Ext(cfgPath)
	if len(ext) == 0 {
		return invalidConfigPathEr(cfgPath)
	}
	ext = ext[1:]
	decoder, isPresent := configDecoders[ext]
	if !isPresent {
		return noDecoderEr(ext)
	}
	// TODO: will need stream reading if the config file is huge.
	bys, er := os.ReadFile(cfgPath)
	if er != nil {
		return er
	}
	er = decoder(bys, cfg)
	if er != nil {
		return er
	}
	return nil
}

// should be called before calling InitConfig.
func AddConfigDecoder(extension string, decoder DecodeFunc) {
	configDecoders[extension] = decoder
}
