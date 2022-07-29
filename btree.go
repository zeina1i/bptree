package btree

const order = 8

type BTree struct {
	root *Node
}

type Node struct {
	keys     []string
	pointers []string
}

func (bt *BTree) Set(key string, value string) error {
	r := bt.root
	if r == nil {
		bt.init(key, value)
		return nil
	}

	//find leaf
	l := bt.findLeaf(key)
	if len(l.keys) < order-1 {
		//	insert
		bt.insertIntoLeaf(l, key, value)

		return nil
	}

	//	splitLeaf
	return bt.splitLeaf(l)
}

func (bt *BTree) init(key string, val string) {
	bt.root = &Node{}
	bt.root.keys = []string{}
	bt.root.keys[0] = key
	bt.root.pointers = []string{}
	bt.root.pointers[0] = val
}

func (bt *BTree) findLeaf(val string) *Node {
	if curr := bt.root; curr == nil {
		return nil
	}

	return nil
}

func (bt *BTree) splitLeaf(leaf *Node) error {
	return nil
}

func (bt *BTree) insertIntoLeaf(leaf *Node, key string, value string) {
	i := 0
	keysLen := len(leaf.keys)
	for i < keysLen {
		if leaf.keys[i] < key {
			break
		}
		i++
	}

	for j := keysLen; j > i; j-- {
		leaf.keys[j] = leaf.keys[j-1]
		leaf.pointers[j] = leaf.pointers[j-1]
	}

	leaf.keys[i] = key
	leaf.pointers[i] = value
}
