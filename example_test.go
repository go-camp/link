package link

import "fmt"

func ExampleParseRandom() {
	id := Must(ParseRandom("a9e272b77827c32a785e27342acfd1680a1f8001"))
	p := fmt.Println
	p("ID:", id)
	p("Chain:", id.Chain())
	p("IsZero:", id.IsZero())
	p("Parent:", id.Parent())
	p("Index:", id.Index())
	p("Child(10):", id.Child(10))
	p("IsRoot:", id.IsRoot())
	p("Root:", id.Root())
	p("Deep:", id.Deep())
	p("Equal(NewRandom()):", id.Equal(Must(NewRandom())))
	// Output:
	// ID: a9e272b77827c32a785e27342acfd1680a1f8001
	// Chain: [a9e272b77827c32a785e27342acfd168 a9e272b77827c32a785e27342acfd1680a a9e272b77827c32a785e27342acfd1680a1f a9e272b77827c32a785e27342acfd1680a1f8001]
	// IsZero: false
	// Parent: a9e272b77827c32a785e27342acfd1680a1f
	// Index: 128
	// Child(10): a9e272b77827c32a785e27342acfd1680a1f80010a
	// IsRoot: false
	// Root: a9e272b77827c32a785e27342acfd168
	// Deep: 3
	// Equal(NewRandom()): false
}

func ExampleSeq() {
	seq := NewSeq()
	p := fmt.Println
	p(seq.Next())

	root := Must(NewRandom())
	id := NextChild(root, seq)
	p(id.Index())
	// Output:
	// 0
	// 1
}
