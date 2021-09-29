/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() (_ Codec) {
    return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    arr := []int32{}
    st := []*TreeNode{root}
    for len(st) > 0 {
        node := st[len(st) - 1]
        st = st[:len(st) - 1]
        if node == nil {
            arr = append(arr, math.MinInt32)
            continue
        }
        arr = append(arr, int32(node.Val))
        st = append(st, node.Right)
        st = append(st, node.Left)
    }
    buf := bytes.NewBuffer([]byte{})
    binary.Write(buf, binary.BigEndian, arr)
    return string(buf.Bytes())
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    blob := []byte(data)
    buf := bytes.NewBuffer(blob)
    arr := make([]int32, len(blob) / 4, len(blob) / 4)
    binary.Read(buf, binary.BigEndian, arr)
    var build func() *TreeNode
    build = func() *TreeNode {
        val := arr[0]
        if val == math.MinInt32 {
            arr = arr[1:]
            return nil
        }
        arr = arr[1:]
        return &TreeNode{int(val),build(),build()}
    }
    return build()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
