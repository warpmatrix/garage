func displayTable(orders [][]string) [][]string {
    tableOrders, tableArr := map[int]map[string]int{}, []int{}
    foodsMap, foodsArr := map[string]struct{}{}, []string{}
    for _, order := range orders {
        foodItem := order[2]
        table, _ := strconv.Atoi(order[1])
        if _, exist := foodsMap[foodItem]; !exist {
            foodsArr = append(foodsArr, foodItem)
            foodsMap[foodItem] = struct{}{}
        }
        if tableOrders[table] == nil {
            tableOrders[table] = map[string]int{}
            tableArr = append(tableArr, table)
        }
        tableOrders[table][foodItem]++
    }
    sort.Strings(foodsArr)
    sort.Ints(tableArr)
    
    ret := make([][]string, len(tableArr) + 1)
    ret[0] = make([]string, len(foodsArr) + 1)
    ret[0][0] = "Table"
    for i, foodItem := range foodsArr {
        ret[0][1 + i] = foodItem
    }
    for i, table := range tableArr {
        ret[1 + i] = make([]string, len(foodsArr) + 1)
        ret[1 + i][0] = strconv.Itoa(table)
        for j, foodItem := range foodsArr {
            ret[1 + i][1 + j] = strconv.Itoa(tableOrders[table][foodItem])
        }
    }
    return ret
}
