package gotrax

import "github.com/otrv4/ed448"

var DsaKeyType = []byte{0x00, 0x00}
var Ed448KeyType = []byte{0x00, 0x10}
var Ed448KeyTypeInt = uint16(0x0010)
var SharedPrekeyKeyType = []byte{0x00, 0x11}
var SharedPrekeyKeyTypeInt = uint16(0x0011)

const SymKeyLength = 57
const PrivKeyLength = 57
const FingerprintLength = 56
const SkLength = 64

var IdentityPoint = ed448.NewPoint([16]uint32{0x00}, [16]uint32{0x01}, [16]uint32{0x01}, [16]uint32{0x00})

const (
	ClientProfileTagInstanceTag           = uint16(0x0001)
	ClientProfileTagPublicKey             = uint16(0x0002)
	ClientProfileTagVersions              = uint16(0x0003)
	ClientProfileTagExpiry                = uint16(0x0004)
	ClientProfileTagDSAKey                = uint16(0x0005)
	ClientProfileTagTransitionalSignature = uint16(0x0006)
)

var kdfPrekeyServerPrefix = []byte("OTR-Prekey-Server")
var kdfPrefix = []byte("OTRv4")

const (
	usageFingerprint = byte(0x00)
	usageBraceKey    = byte(0x02)
	usageAuth        = byte(0x11)
)

var basePointBytes = []byte{
	0x14, 0xfa, 0x30, 0xf2, 0x5b, 0x79, 0x08, 0x98, 0xad, 0xc8, 0xd7, 0x4e,
	0x2c, 0x13, 0xbd, 0xfd, 0xc4, 0x39, 0x7c, 0xe6, 0x1c, 0xff, 0xd3, 0x3a,
	0xd7, 0xc2, 0xa0, 0x05, 0x1e, 0x9c, 0x78, 0x87, 0x40, 0x98, 0xa3, 0x6c,
	0x73, 0x73, 0xea, 0x4b, 0x62, 0xc7, 0xc9, 0x56, 0x37, 0x20, 0x76, 0x88,
	0x24, 0xbc, 0xb6, 0x6e, 0x71, 0x46, 0x3f, 0x69, 0x00,
}

var primeOrderBytes = []byte{
	0x3f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x7c, 0xca, 0x23, 0xe9, 0xc4, 0x4e, 0xdb, 0x49,
	0xae, 0xd6, 0x36, 0x90, 0x21, 0x6c, 0xc2, 0x72, 0x8d, 0xc5, 0x8f, 0x55,
	0x23, 0x78, 0xc2, 0x92, 0xab, 0x58, 0x44, 0xf3,
}
