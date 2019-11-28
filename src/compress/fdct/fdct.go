// Ancient code extracted from https://golang.org/src/image/jpeg/fdct.go
package fdct

import (
  consts "../../consts"
)

const (
	fix_0_298631336 = 2446

	fix_0_390180644 = 3196

	fix_0_541196100 = 4433

	fix_0_765366865 = 6270

	fix_0_899976223 = 7373

	fix_1_175875602 = 9633

	fix_1_501321110 = 12299

	fix_1_847759065 = 15137

	fix_1_961570560 = 16069

	fix_2_053119869 = 16819

	fix_2_562915447 = 20995

	fix_3_072711026 = 25172
)

const (
	constBits     = 13
	pass1Bits     = 2
	centerJSample = 128
)


// fdct performs a forward DCT on an 8x8 block of coefficients, including a

// level shift.

func Fdct(b *consts.Block) {

	// Pass 1: process rows.

	for y := 0; y < 8; y++ {

		y8 := y * 8

		s := b[y8 : y8+8 : y8+8] // Small cap improves performance, see https://golang.org/issue/27857

		x0 := s[0]

		x1 := s[1]

		x2 := s[2]

		x3 := s[3]

		x4 := s[4]

		x5 := s[5]

		x6 := s[6]

		x7 := s[7]


		tmp0 := x0 + x7

		tmp1 := x1 + x6

		tmp2 := x2 + x5

		tmp3 := x3 + x4


		tmp10 := tmp0 + tmp3

		tmp12 := tmp0 - tmp3

		tmp11 := tmp1 + tmp2

		tmp13 := tmp1 - tmp2


		tmp0 = x0 - x7

		tmp1 = x1 - x6

		tmp2 = x2 - x5

		tmp3 = x3 - x4


		s[0] = (tmp10 + tmp11 - 8*centerJSample) << pass1Bits

		s[4] = (tmp10 - tmp11) << pass1Bits

		z1 := (tmp12 + tmp13) * fix_0_541196100

		z1 += 1 << (constBits - pass1Bits - 1)

		s[2] = (z1 + tmp12*fix_0_765366865) >> (constBits - pass1Bits)

		s[6] = (z1 - tmp13*fix_1_847759065) >> (constBits - pass1Bits)


		tmp10 = tmp0 + tmp3

		tmp11 = tmp1 + tmp2

		tmp12 = tmp0 + tmp2

		tmp13 = tmp1 + tmp3

		z1 = (tmp12 + tmp13) * fix_1_175875602

		z1 += 1 << (constBits - pass1Bits - 1)

		tmp0 *= fix_1_501321110

		tmp1 *= fix_3_072711026

		tmp2 *= fix_2_053119869

		tmp3 *= fix_0_298631336

		tmp10 *= -fix_0_899976223

		tmp11 *= -fix_2_562915447

		tmp12 *= -fix_0_390180644

		tmp13 *= -fix_1_961570560


		tmp12 += z1

		tmp13 += z1

		s[1] = (tmp0 + tmp10 + tmp12) >> (constBits - pass1Bits)

		s[3] = (tmp1 + tmp11 + tmp13) >> (constBits - pass1Bits)

		s[5] = (tmp2 + tmp11 + tmp12) >> (constBits - pass1Bits)

		s[7] = (tmp3 + tmp10 + tmp13) >> (constBits - pass1Bits)

	}

	// Pass 2: process columns.

	// We remove pass1Bits scaling, but leave results scaled up by an overall factor of 8.

	for x := 0; x < 8; x++ {

		tmp0 := b[0*8+x] + b[7*8+x]

		tmp1 := b[1*8+x] + b[6*8+x]

		tmp2 := b[2*8+x] + b[5*8+x]

		tmp3 := b[3*8+x] + b[4*8+x]


		tmp10 := tmp0 + tmp3 + 1<<(pass1Bits-1)

		tmp12 := tmp0 - tmp3

		tmp11 := tmp1 + tmp2

		tmp13 := tmp1 - tmp2


		tmp0 = b[0*8+x] - b[7*8+x]

		tmp1 = b[1*8+x] - b[6*8+x]

		tmp2 = b[2*8+x] - b[5*8+x]

		tmp3 = b[3*8+x] - b[4*8+x]


		b[0*8+x] = (tmp10 + tmp11) >> pass1Bits

		b[4*8+x] = (tmp10 - tmp11) >> pass1Bits


		z1 := (tmp12 + tmp13) * fix_0_541196100

		z1 += 1 << (constBits + pass1Bits - 1)

		b[2*8+x] = (z1 + tmp12*fix_0_765366865) >> (constBits + pass1Bits)

		b[6*8+x] = (z1 - tmp13*fix_1_847759065) >> (constBits + pass1Bits)


		tmp10 = tmp0 + tmp3

		tmp11 = tmp1 + tmp2

		tmp12 = tmp0 + tmp2

		tmp13 = tmp1 + tmp3

		z1 = (tmp12 + tmp13) * fix_1_175875602

		z1 += 1 << (constBits + pass1Bits - 1)

		tmp0 *= fix_1_501321110

		tmp1 *= fix_3_072711026

		tmp2 *= fix_2_053119869

		tmp3 *= fix_0_298631336

		tmp10 *= -fix_0_899976223

		tmp11 *= -fix_2_562915447

		tmp12 *= -fix_0_390180644

		tmp13 *= -fix_1_961570560


		tmp12 += z1

		tmp13 += z1

		b[1*8+x] = (tmp0 + tmp10 + tmp12) >> (constBits + pass1Bits)

		b[3*8+x] = (tmp1 + tmp11 + tmp13) >> (constBits + pass1Bits)

		b[5*8+x] = (tmp2 + tmp11 + tmp12) >> (constBits + pass1Bits)

		b[7*8+x] = (tmp3 + tmp10 + tmp13) >> (constBits + pass1Bits)

	}

}
