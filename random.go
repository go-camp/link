package link

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/go-camp/link/internal"
)

// rid is an implementation of ID interface that based on random.
//
//   169 226 114 183 120 39 195 42 120 94 39 52 42 207 209 104         10        31       128      1
//  |------------------------> root <-------------------------|      child     child      |--child--|
//
//                           16 bytes                              varint64   varint64     varint64
//                           random                                   index     index       index
//                                                                     10        31          128
//
//                         32 hex chars                                hex       hex         hex
//                 a9e272b77827c32a785e27342acfd168                    0a        1f          8001
type rid []byte

// ridRootSize is root(min) rid bytes length.
const ridRootSize = 16

var ridZero rid

func (id rid) IsZero() bool { return len(id) == 0 }

func (id rid) copyN(n int) ID {
	return append(rid(nil), id[:n]...)
}

func (id rid) Parent() ID {
	if id.IsZero() || id.IsRoot() {
		return ridZero
	}
	return id.copyN(len(id) - internal.VarintSizeInLastBytes(id[ridRootSize:]))
}

func (id rid) Index() (index uint64) {
	if !id.IsZero() {
		index, _ = internal.ReadVarintLast(id[ridRootSize:])
	}
	return
}

func (id rid) Child(index uint64) ID {
	child := make(rid, len(id)+internal.VarintSize(index))
	copy(child, id)
	internal.PutVarint(child[len(id):], index)
	return child
}

func (id rid) IsRoot() bool { return len(id) == ridRootSize }

func (id rid) Root() ID {
	return id.copyN(ridRootSize)
}

func (id rid) Chain() (chain []ID) {
	if id.IsZero() {
		return nil
	}
	chain = append(chain, id.Root())
	for i := ridRootSize; i < len(id); {
		n := i + internal.VarintSizeInBytes(id[i:])
		chain = append(chain, id.copyN(n))
		i = n
	}
	return
}

func (id rid) Deep() (deep int) {
	if id.IsZero() {
		return -1
	}
	for i := ridRootSize; i < len(id); {
		i += internal.VarintSizeInBytes(id[i:])
		deep++
	}
	return
}

func (id rid) Equal(xid ID) bool {
	if xid == nil {
		return false
	}
	cxid, ok := xid.(rid)
	if !ok {
		return false
	}
	return bytes.Equal(id, cxid)
}

func (id rid) String() string {
	return hex.EncodeToString(id)
}

// NewRandom creates an new random root id based on given reader.
// If err is not nil, root is nil.
func NewRandomFromReader(reader io.Reader) (root ID, err error) {
	id := make(rid, ridRootSize)
	_, err = io.ReadFull(reader, id)
	if err != nil {
		return
	}
	root = id
	return
}

// NewRandom creates an new random root id based on crypto/rand.Reader.
// If err is not nil, root is nil.
func NewRandom() (root ID, err error) {
	return NewRandomFromReader(rand.Reader)
}

// ParseRandom decodes hex encoded sid into an id.
func ParseRandom(sid string) (ID, error) {
	var did []byte
	did, err := hex.DecodeString(sid)
	if err != nil {
		return nil, err
	}

	if len(did) < ridRootSize {
		return nil, fmt.Errorf("decoded id is shorter than root(min) id size %d", ridRootSize)
	}

	for i := ridRootSize; i < len(did); {
		s := internal.VarintSizeInBytes(did[i:])
		if s <= 0 {
			return nil, errors.New("invalid id format")
		}
		i += s
	}

	return rid(did), nil
}
