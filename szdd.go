package szdd

import (
	"bytes"
	"fmt"
)

const WINDOW_SIZE = 4096;

var SZDD_SIGNATURE = []byte{0x53, 0x5A, 0x44, 0x44, 0x88, 0xF0, 0x27, 0x33};

func Expand(bs []byte) (data []byte, err error) {
	window := make([]byte, WINDOW_SIZE)
	header := bs[:14]
	pos := WINDOW_SIZE - 16
	res := bytes.Compare(header[:8], SZDD_SIGNATURE)
	chunk := bs[14:]
	data = make([]byte, 0, 3*len(bs))

	if res != 0 {
		return nil, fmt.Errorf("wrong signature")
	}
	for i := range window {
		window[i] = 0
	}

	for i := 0; i < len(chunk); {
		control := chunk[i]
		i++

		for cbit := byte(0x01); (cbit & 0xFF) != 0 && i < len(chunk); cbit <<= 1 {
			if (control & cbit) != 0 {
				data = append(data, chunk[i])
				window[pos] = chunk[i]
				pos++
				i++
				pos &= WINDOW_SIZE - 1
			} else {
				matchpos := int(chunk[i])
				i++
				matchlen := chunk[i]
				i++
				matchpos |= int((int(matchlen) & 0xF0) << 4)
				matchlen = (matchlen & 0x0F) + 3
				for ; matchlen > 0; matchlen-- {
					data = append(data, window[matchpos])
					window[pos] = window[matchpos]
					pos++
					matchpos++
					pos &= WINDOW_SIZE - 1
					matchpos &= WINDOW_SIZE - 1
				}
			}
		}
	}

	return
}

