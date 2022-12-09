const bitLen = 30

type trie struct {
    children [2]*trie
    min int
}

func (t *trie) insert(num int) {
    if t.min > num { t.min = num }
    node := t
    for bitIdx := bitLen - 1; bitIdx >= 0; bitIdx-- {
        bit := (num >> bitIdx) & 1
        if node.children[bit] == nil {
            node.children[bit] = &trie{ min: num }
        }
        node = node.children[bit]
        if num < node.min {
            node.min = num
        }
    }
}

func (t *trie) getMaxXorWithXor(num, limit int) int {
    if t.min > limit { return -1 }
    node, ret := t, 0
    for bitIdx := bitLen - 1; bitIdx >= 0; bitIdx-- {
        bit := num >> bitIdx & 1
        childIdx := bit ^ 1
        if node.children[childIdx] != nil && node.children[childIdx].min <= limit {
            ret |= 1 << bitIdx
        } else {
            childIdx = bit
        }
        node = node.children[childIdx]
    }
    return ret
}

func maximizeXor(nums []int, queries [][]int) []int {
    t := &trie{ min: math.MaxInt32 }
    for _, num := range nums {
        t.insert(num)
    }
    ret := make([]int, len(queries))
    for i, q := range queries {
        ret[i] = t.getMaxXorWithXor(q[0], q[1])
    }
    return ret
}
