package disjointSet

type Element struct {
	rank int
	val  interface{}
	up   *Element
}

type DisjointSet struct {
	Data []*Element
	V    int
}

// insert as usual with &Element{0,val,nil}
func (d *DisjointSet) Initialize(V int) {
	d.V = V
	for i := 0; i < d.V; i++ {
		d.Data = append(d.Data, &Element{0, i, nil})
	}
}

func (e *Element) SameClass(e1 *Element) bool {
	// retrieve root
	var root, root1 *Element
	for root = e; root.up != nil; root = root.up {
	}
	for root1 = e1; root1.up != nil; root1 = root1.up {
	}
	// compression, all the elements point to them roots
	for tmp := e; tmp != root && tmp != nil; tmp = tmp.up {
		tmp.up = root
	}
	for tmp := e1; tmp != root1 && tmp != nil; tmp = tmp.up {
		tmp.up = root1
	}
	return root == root1
}

func (e *Element) Join(e1 *Element) {
	// retrieve root
	var root, root1 *Element
	for root = e; root.up != nil; root = root.up {
	}
	for root1 = e1; root1.up != nil; root1 = root1.up {
	}
	// union by rank
	var newRoot *Element
	switch {
	case root.rank > root1.rank:
		newRoot = root
		root1.up = root
	default:
		newRoot = root1
		root.up = root1
	}
	if root.rank == root1.rank {
		newRoot.rank++
	}
	// compressing the tree
	for tmp := e; tmp != newRoot && tmp != nil; tmp = tmp.up {
		tmp.up = newRoot
	}
	for tmp := e1; tmp != newRoot && tmp != nil; tmp = tmp.up {
		tmp.up = newRoot
	}
}
