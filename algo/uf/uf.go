// uf is an implementation of UnionFind ADT.
package uf

type UnionFind struct {
	root        []int          // nodes' root
	size        []int          // each root's size
	counter     int            // stick ascending, used for new root id and index
	symbolTable map[string]int // symbol table
}

// return initial id
func (u *UnionFind) Insert(s string) {
	if !u.exist(s) {
		u.root = append(u.root, u.counter)
		u.size = append(u.size, 1)
		u.symbolTable[s] = u.counter
		u.counter++
	}
}

func (u *UnionFind) exist(s string) bool {
	if _, ok := u.symbolTable[s]; !ok {
		return false
	}
	return true
}

func (u *UnionFind) Union(l, r string) {
	u.Insert(l)
	u.Insert(r)
	lr := u.Find(l)
	rr := u.Find(r)

	if u.size[lr] < u.size[rr] {
		u.root[lr] = rr
		u.size[rr] += u.size[lr]
	} else {
		u.root[rr] = lr
		u.size[lr] += u.size[rr]
	}
}

func (u *UnionFind) Find(s string) (r int) {
	if _, ok := u.symbolTable[s]; !ok {
		return -1
	}
	// initial each node is itself's root
	r = u.symbolTable[s]
	tmp := r
	for tmp != u.root[tmp] {
		tmp = u.root[tmp]
	}

	// path compression
	u.root[r] = tmp
	return tmp
}

func (u *UnionFind) Connected(l, r string) bool {
	if _, ok := u.symbolTable[l]; !ok {
		return false
	}
	if _, ok := u.symbolTable[r]; !ok {
		return false
	}
	return u.Find(l) == u.Find(r)
}

// helper function
func NewUnionFind() *UnionFind {
	return &UnionFind{
		symbolTable: make(map[string]int),
	}
}
