package producer

import (
	"crypto/md5"
	"strings"
)

const zero32Str = "00000000000000000000000000000000"

func toHex(x byte) byte {
	if x < 10 {
		return '0' + x
	}

	return 'a' - 10 + x
}

func getBitValue(val byte, bits int) byte {
	if bits >= 16 {
		return val
	}

	if bits >= 8 {
		return val & (0xf - 1)
	}

	if bits >= 4 {
		return val & (0xf - 3)
	}

	if bits >= 2 {
		return val & (0xf - 7)
	}

	return val & (0xf - 15)
}

func AdjustHash(shardhash string, shardCount int) (string, error) {
	var (
		targetHash     = strings.Builder{}
		index      int = 0
	)
	targetHash.Grow(32)

	h := md5.New()
	h.Write([]byte(shardhash))
	cipherStr := h.Sum(nil)

	for index = 0; shardCount > 0 && index < len(cipherStr); index++ {
		if (index & 0x01) == 0 {
			targetHash.WriteByte(toHex(getBitValue(cipherStr[index>>1]>>4, shardCount)))
		} else {
			targetHash.WriteByte(toHex(getBitValue(cipherStr[index>>1]&0xF, shardCount)))
		}

		shardCount >>= 4
	}

	targetHash.WriteString(zero32Str[0 : 32-index])

	return targetHash.String(), nil
}
