/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    m := make(map[int]*Employee)
    for _, v := range employees {
        m[v.Id] = v
    }
    queue := []*Employee{m[id]}
    ret := 0
    for len(queue) > 0 {
        e := queue[0]
        queue = queue[1:len(queue)]
        ret += e.Importance
        for _, subOrdId := range e.Subordinates {
            queue = append(queue, m[subOrdId])
        }
    }
    return ret
}
