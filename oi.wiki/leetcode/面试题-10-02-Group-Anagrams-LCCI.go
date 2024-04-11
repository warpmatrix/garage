func groupAnagrams(strs []string) [][]string {
    angsGrps := map[[26]int][]string{}
    for _, str := range strs {
       cnts := [26]int{}
       for _, ch := range str {
           cnts[ch - 'a']++
       }
       angsGrps[cnts] = append(angsGrps[cnts], str)
    }
    ret := [][]string{}
    for _, grp := range angsGrps {
        ret = append(ret, grp)
    }
    return ret
}
