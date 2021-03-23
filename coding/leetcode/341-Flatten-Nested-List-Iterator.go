type NestedIterator struct {
    stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
    queue := it.stack[len(it.stack)-1]
    val := queue[0].GetInteger()
    it.stack[len(it.stack)-1] = queue[1:]
    return val
}

func (it *NestedIterator) HasNext() bool {
    for len(it.stack) > 0 {
        queue := it.stack[len(it.stack)-1]
        if len(queue) == 0 { 
            it.stack = it.stack[:len(it.stack)-1]
            continue
        }
        nest := queue[0]
        if nest.IsInteger() {
            return true
        }
        it.stack[len(it.stack)-1] = queue[1:]
        it.stack = append(it.stack, nest.GetList())
    }
    return false
}