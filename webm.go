package "go-httpstream"

import (
    "go-httpstream/containers/matroska/EBML"
)
// WebM container interface.
// Author: Anton Vilhelm Ásgeirsson
// Date: 2015-09-11

// Video Codecs:
// VP8: vp8.X codec is vp8, bitstream is version X.
//  - VP8 currently only has bitstream version 0.
//  - This will also match to a FourCC of VP8X for applications that need a FourCC
// VP9: vp8.X codec is vp9, bitstream is version X.
//  - VP9 currently only hasa bitstream version 0.
//  - This will also match to a FourCC of VP9X for applications that need a FourCC

type WebM struct {

}
