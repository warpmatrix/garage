var m = map[string]map[string]bool{}

func isScramble(s1 string, s2 string) bool {
    if _, ok := m[s1]; !ok {
        m[s1] = make(map[string]bool)
    }
    if tf, ok := m[s1][s2]; ok {
        return tf
    }
    if s1 == s2 {
        m[s1][s2] = true
        return true
    }
    cnt := [26]int{}
    for i := range s1 {
        cnt[s1[i] - 'a']++
        cnt[s2[i] - 'a']--
    }
    for _, v := range cnt {
        if v != 0 {
            m[s1][s2] = false
            return false
        }
    }
    for i := 1; i < len(s1); i++ {
        if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
            m[s1][s2] = true
            return true
        }
        if isScramble(s1[:i], s2[len(s2) - i:]) && isScramble(s1[i:], s2[:len(s2) - i]) {
            m[s1][s2] = true
            return true
        }
    }
    m[s1][s2] = false
    return false
}

