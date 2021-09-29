func convertToTitle(num int) string {
    const base = 26
    arr := []byte{}
    for num > 0 {
        idx := num % base - 1
        if idx == -1 {
            idx = 25
            num = num / 26 - 1
        } else {
            num /= 26
        }
        arr = append(arr, byte(idx + 'A'))
    }
    return string(reverse(arr))
}

func reverse(arr []byte) []byte {
    for i := 0; 2 * i < len(arr); i++ {
        j := len(arr) - i - 1
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
}
