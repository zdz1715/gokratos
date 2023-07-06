package zmap

import (
	"strconv"
	"testing"
)

func BenchmarkStringMap_Set(b *testing.B) {
	var (
		strMap StringMap[int]
	)

	for i := 0; i < b.N; i++ {
		strMap.Set("key_"+strconv.Itoa(i), i)
	}

	b.Logf("n: %d size: %d", b.N, strMap.Size())
}
