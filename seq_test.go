package link

import (
	"sync"
	"testing"
)

func TestSeq(t *testing.T) {
	var startWith int64 = 10
	seq := NewSeq(SeqStartWith(startWith), SeqIncrementBy(2))

	t.Run("Next", func(t *testing.T) {
		n := seq.Next()
		if n != 10 {
			t.Fatalf("want next is %d but get %d", 10, n)
		}
	})

	t.Run("concurrently", func(t *testing.T) {
		count := 10
		var wg sync.WaitGroup
		wg.Add(count)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				seq.Next()
			}()
		}
		wg.Wait()

		n := seq.Next()
		if n != 32 {
			t.Fatalf("want next is %d but get %d", 32, n)
		}
	})
}
