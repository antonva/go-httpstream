package "go-httpstream/containers/matroska"

import (
        "bytes"
)
// WebM container interface.
// Author: Anton Vilhelm Ásgeirsson
// Date: 2015-09-11

// Known basic types for EBML are:
// Signed Integer   - Big endian, any size from 1 to 8 octets
// Unsigned Integer - Big endian, any size from 1 to 8 octets
// Float            - Big endian, defined for 4 and 8 octets (32 and 64 bits)
// String           - Printable ASCII (0x20 to 0x7E), zero-padded when needed
// UTF-8            - Unicode string, zero padded when needed (RFC 2279)[http://www.faqs.org/rfcs/rfc2279.html]

type EBML struct {

}

type EBMLHeader struct {
    EBMLVersion        uint   // Unsigned integer.
    EBMLReadVersion    uint   // Unsigned integer.
    EBMLMaxIDLength    uint   // Unsigned integer.
    EBMLMaxSizeLength  uint   // Unsigned integer.
    DocType            string // String. ( TODO: problematic conversion between ASCII and UTF?)
    DocTypeVersion     uint   // Unsigned integer.
    DocTypeReadVersion uint   // Unsigned integer.
}

type EBMLGlobal struct {
    Void    []byte
}
