type ThroneInheritance struct {
    kingName string
    deadNames map[string]struct{}
    children map[string][]string
}


func Constructor(kingName string) ThroneInheritance {
    return ThroneInheritance{
        kingName,
        map[string]struct{}{},
        map[string][]string{},
    }
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
    this.children[parentName] = append(this.children[parentName], childName)
}


func (this *ThroneInheritance) Death(name string)  {
    this.deadNames[name] = struct{}{}
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
    ret := []string{}
    var preOrder func(name string)
    preOrder = func(name string) {
        if _, found := this.deadNames[name]; !found {
            ret = append(ret, name)
        }
        for _, childName := range this.children[name] {
            preOrder(childName)
        }
    }
    preOrder(this.kingName)
    return ret
}


/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
