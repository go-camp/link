package link

// ID is a self-describing identification interface.
type ID interface {
	// IsZero returns true if id is zero value.
	IsZero() bool

	// Parent returns the direct parent of id.
	// If id is zero or root, then parent is zero.
	Parent() ID

	// Index returns the index of id in the parant.
	// The index of root id is 0.
	Index() (index int64)

	// Child returns the derived child id according to the given index.
	Child(index int64) (child ID)

	// IsRoot returns true if id is a root id.
	IsRoot() bool

	// Root returns the root of id.
	Root() (root ID)

	// Chain returns the id derivation chain.
	// The order of chain from root to id.
	//   chain := []ID{root, ..., parent, id}
	Chain() (chain []ID)

	// Deep returns the number of parents in the id derivation chain.
	// The deep of root id is 0.
	Deep() (deep int)

	// Equal reports whether id and xid represent the same id.
	Equal(xid ID) bool

	// String returns string expression of id.
	String() string
}
