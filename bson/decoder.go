package bson

import "github.com/pkg/bson"

type decoder struct {
	decoder *bson.Decoder
}

func (d *decoder) Decode(value interface{}) error {
	return d.decoder.Decode(value)
}
