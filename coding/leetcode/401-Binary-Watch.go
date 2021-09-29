func readBinaryWatch(turnedOn int) []string {
    ret := []string{}
    for hour := 0; hour < 12; hour++ {
        for min := 0; min < 60; min++ {
            if bits.OnesCount(uint(hour)) + bits.OnesCount(uint(min)) == turnedOn {
                str := strconv.Itoa(hour) + ":" + fmt.Sprintf("%02d", min)
                ret = append(ret, str)
            }
        }
    }
    return ret
}
