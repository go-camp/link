package internal

import (
	"fmt"
	"testing"
)

func TestVarint(t *testing.T) {
	for i := -64; i <= 63; i++ {
		var v int64
		if i < 0 {
			v = -1 << -i
		} else {
			v = (1 << i) - 1
		}

		t.Run(fmt.Sprintf("%d_%016x", i, v), func(t *testing.T) {
			buf := make([]byte, MaxVarintSize)

			n := PutVarint(buf, v)
			if n <= 0 {
				t.Errorf("want n > 0")
			}

			vs := VarintSize(v)
			if vs != n {
				t.Errorf("variant size is %d but want %d", vs, n)
			}

			vsb := VarintSizeInBytes(buf)
			if vsb != n {
				t.Errorf("variant size in bytes %d but want %d", vsb, n)
			}

			vslb := VarintSizeInLastBytes(buf[:n])
			if vslb != n {
				t.Errorf("variant size in last bytes %d but want %d", vsb, n)
			}

			rv, rvn := ReadVarint(buf)
			if rvn != n {
				t.Errorf("read varint n is %d but want %d", rvn, n)
			}
			if rv != v {
				t.Errorf("read varint v is %d but want %d", rv, v)
			}

			rvl, rvln := ReadVarintLast(buf[:n])
			if rvln != n {
				t.Errorf("read varint last n is %d but want %d", rvln, n)
			}
			if rvl != v {
				t.Errorf("read varint last v is %d but want %d", rvl, v)
			}
		})
	}
}
