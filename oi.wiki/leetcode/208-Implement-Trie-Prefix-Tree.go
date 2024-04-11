type Trie struct {
    children [26]*Trie
    end bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &Trie{}
        }
        node = node.children[idx]
    }
    node.end = true
}

func (this *Trie) search(prefix string) *Trie {
    node := this
    for i := 0; i < len(prefix); i++ {
        idx := prefix[i] - 'a'
        if node.children[idx] == nil { return nil }
        node = node.children[idx]
    }
    return node
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.search(word)
    return node != nil && node.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    return this.search(prefix) != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */


// type Trie struct {
//     node [26]*Trie
//     end bool
// }


// /** Initialize your data structure here. */
// func Constructor() Trie {
//     return Trie{}
// }


// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string)  {
//     if len(word) == 0 {
//         this.end = true
//         return
//     }
//     if this.node[word[0] - 'a'] == nil {
//         this.node[word[0] - 'a'] = &Trie{}
//     }
//     this.node[word[0] - 'a'].Insert(word[1:])
// }


// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
//     if len(word) == 0 {
//         return this.end
//     }
//     if this.node[word[0] - 'a'] == nil { return false }
//     return this.node[word[0] - 'a'].Search(word[1:])
// }


// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
//     if len(prefix) == 0 {
//         return true
//     }
//     if this.node[prefix[0] - 'a'] == nil {return false}
//     return this.node[prefix[0] - 'a'].StartsWith(prefix[1:])
// }


// /**
//  * Your Trie object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Insert(word);
//  * param_2 := obj.Search(word);
//  * param_3 := obj.StartsWith(prefix);
//  */