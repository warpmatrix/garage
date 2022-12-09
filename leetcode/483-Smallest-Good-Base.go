func smallestGoodBase(n string) string {
    num, _ := strconv.Atoi(n)
    mMax := bits.Len(uint(num)) - 1
    for m := mMax; m > 1; m-- {
        base := int(math.Pow(float64(num), 1 / float64(m)))
		// wrong anwser: sum := int((math.Pow(float64(base), float64(m + 1)) - 1) / float64(base - 1))
        sum, prod := 1, 1
        for i := 0; i < m; i++ {
            prod *= base
            sum += prod
        }
        if sum == num {
            return strconv.Itoa(base)
        }
    }
    return strconv.Itoa(num - 1)
}
