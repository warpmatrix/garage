func openLock(deadends []string, target string) int {
    source := "0000"
    if target == source { return 0 }
    deadTbl := map[string]struct{}{}
    for _, deadend := range deadends {
        deadTbl[deadend] = struct{}{}
    }
    if _, dead := deadTbl[source]; dead { return -1 }
    srcStat := map[string]struct{}{ source : struct{}{} }
    tarStat := map[string]struct{}{ target : struct{}{} }
    step, isVisited := -1, map[string]struct{}{}
    srchStat, waitStat := &srcStat, &tarStat
    for len(*srchStat) > 0 && len(*waitStat) > 0 {
        tmp := map[string]struct{}{}
        for stat := range *srchStat {
            if _, ok := isVisited[stat]; ok {
                return step
            }
            isVisited[stat] = struct{}{}
            neibors := getNeibors(stat)
            for _, neibor := range neibors {
                if _, dead := deadTbl[neibor]; dead { continue }
                if _, visited := isVisited[neibor]; visited { continue }
                tmp[neibor] = struct{}{}
            }
        }
        *srchStat = tmp
        srchStat, waitStat = waitStat, srchStat
        step++
    }
    return -1
}

func getNeibors(stat string) []string {
    ret := []string{}
    blob := []byte(stat)
    for i, ch := range blob {
        blob[i] = (ch - '0' + 1) % 10 + '0'
        ret = append(ret, string(blob))
        blob[i] = (ch - '0' + 9) % 10 + '0'
        ret = append(ret, string(blob))
        blob[i] = ch
    }
    return ret
}