package muxcodec

import (
	mc "github.com/dms3-mft/go-multicodec"
	cbor "github.com/dms3-mft/go-multicodec/cbor"
	json "github.com/dms3-mft/go-multicodec/json"
)

func StandardMux() *Multicodec {
	return MuxMulticodec([]mc.Multicodec{
		cbor.Multicodec(),
		json.Multicodec(false),
		json.Multicodec(true),
	}, SelectFirst)
}
