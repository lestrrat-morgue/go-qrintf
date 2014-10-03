package qrintf

type ctx struct {
	str []byte
	off int
}

func (c *ctx) Finalize() int {
	return c.off
}

func (c *ctx) Byte(b byte) *ctx {
	i := c.off
	c.str[i] = b
	c.off = i + 1
	return c
}

func (c *ctx) UInt(u uint) *ctx {
	/* Max uint is 18446744073709551615 */
	var tmp = [len("18446744073709551615")]byte{}
	i := 0
	for {
		tmp[i] = byte(0x30 + u%10)
		i++
		u = u / 10
		if u == 0 {
			break
		}
	}

	var x int
	for {
		i--
		x = c.off
		c.str[x] = tmp[i]
		c.off = x + 1
		if i <= 0 {
			break
		}
	}
	return c
}
