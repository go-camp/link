package link

import "sync/atomic"

// Seq is a sequence number generator.
//
// All public methods can be safely called concurrently.
type Seq struct {
	seq uint64

	incrementBy uint64
}

func (seq *Seq) forceIncrementBy() uint64 {
	if seq.incrementBy == 0 {
		return 1
	}
	return seq.incrementBy
}

func (seq *Seq) init() {
	seq.seq -= seq.forceIncrementBy()
}

// Next advances the sequence to its next value and returns that value.
func (seq *Seq) Next() uint64 {
	return atomic.AddUint64(&seq.seq, seq.forceIncrementBy())
}

type SeqOption func(seq *Seq)

// SeqStartWith allows the sequence to begin anywhere. The default starting value is 0.
func SeqStartWith(startWith uint64) SeqOption {
	return func(seq *Seq) {
		seq.seq = startWith
	}
}

// SeqIncrementBy specifies which value is added to the current sequence value to create an new value.
// A positive value will make an ascending sequence, a negative one a descending sequence.
// The default value is 1.
func SeqIncrementBy(incrementBy uint64) SeqOption {
	return func(seq *Seq) {
		seq.incrementBy = incrementBy
	}
}

// NewSeq creates a sequence with a set of options.
func NewSeq(opts ...SeqOption) *Seq {
	seq := new(Seq)
	for _, opt := range opts {
		opt(seq)
	}
	seq.init()
	return seq
}

// NextChild returns the next derived child id.
func NextChild(id ID, seq *Seq) ID {
	return id.Child(seq.Next())
}
