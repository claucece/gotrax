package gotrax

import (
	"crypto/rand"

	. "gopkg.in/check.v1"
)

func (s *GotraxSuite) Test_RandomInto_fillsTheBuffer(c *C) {
	b := make([]byte, 3)
	res := RandomInto(ReaderIntoWithRandom(FixedRand([]string{"ABCDEF"})), b)
	c.Assert(res, IsNil)
	c.Assert(b, DeepEquals, []byte{0xAB, 0xCD, 0xEF})
}

func (s *GotraxSuite) Test_RandomInto_returnsErrorOnShortRead(c *C) {
	b := make([]byte, 3)
	res := RandomInto(ReaderIntoWithRandom(FixedRand([]string{"ABCD"})), b)
	c.Assert(res, ErrorMatches, "short read from random source")
}

func (s *GotraxSuite) Test_DefaultRandom_returnsWithRandomWithRandReader(c *C) {
	d := DefaultRandom()
	r := d.(*readerWithRandom).r
	c.Assert(r, Equals, rand.Reader)
}

func (s *GotraxSuite) Test_RandomUint32_GeneratesARandomNumber(c *C) {
	n := RandomUint32(ReaderIntoWithRandom(FixedRand([]string{"ABCDEF112233"})))
	c.Assert(n, Equals, uint32(0xABCDEF11))
}

func (s *GotraxSuite) Test_FixedRandBytes_returnsAWithRandom(c *C) {
	v := FixedRandBytes([]byte{0x01, 0x02}, []byte{0x03, 0x04})
	rr := v.(*fixedRandBytesReader)
	c.Assert(rr.at, Equals, 0)
	c.Assert(rr.data, DeepEquals, [][]byte{[]byte{0x01, 0x02}, []byte{0x03, 0x04}})
}

func (s *GotraxSuite) Test_readerWithRandom_RandReader_returnsTheGivenReader(c *C) {
	v := &readerWithRandom{r: rand.Reader}
	c.Assert(v.RandReader(), Equals, rand.Reader)
}

func (s *GotraxSuite) Test_fixedRandBytesReader_RandReader_returnsItself(c *C) {
	v := FixedRandBytes([]byte{0x01, 0x02}, []byte{0x03, 0x04})
	rr := v.(*fixedRandBytesReader)
	val := v.RandReader()
	c.Assert(val, Equals, rr)
}

func (s *GotraxSuite) Test_FixtureRand_ReturnsAWithRandomSeededWithFixtureData(c *C) {
	v := FixtureRand()
	c.Assert(v.(*fixedRandBytesReader).data[0], DeepEquals, []byte{0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd, 0xab, 0xcd})
}
