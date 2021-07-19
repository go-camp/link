package link

import (
	"sync"
	"testing"
)

func TestSeq(t *testing.T) {
	var startWith uint64 = 10
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

func TestNextChild(t *testing.T) {
	id := Must(NewRandom())
	seq := NewSeq(SeqStartWith(10))
	child := NextChild(id, seq)
	if !child.Parent().Equal(id) {
		t.Fatalf("want child's parent %s equals to id %s", child.Parent(), id)
	}
	if child.Index() != 10 {
		t.Fatalf("want index of child is 10 but get %d", child.Index())
	}
}
