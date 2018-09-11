package basemux

import (
	mc "github.com/dms3-mft/go-multicodec"
	mux "github.com/dms3-mft/go-multicodec/mux"

	b64 "github.com/dms3-mft/go-multicodec/base/b64"
	bin "github.com/dms3-mft/go-multicodec/base/bin"
	hex "github.com/dms3-mft/go-multicodec/base/hex"
)

func AllBasesMux() *mux.Multicodec {
	m := mux.MuxMulticodec([]mc.Multicodec{
		hex.Multicodec(),
		b64.Multicodec(),
		bin.Multicodec(),
	}, mux.SelectFirst)
	m.Wrap = false
	return m
}
