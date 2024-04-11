type TimeMap struct {
    kv map[string][]pair
}

type pair struct {
    timestamp int
    value string
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{kv:map[string][]pair{}}
}


func (tm *TimeMap) Set(key string, value string, timestamp int)  {
    tm.kv[key] = append(tm.kv[key], pair{timestamp, value})
}


func (tm *TimeMap) Get(key string, timestamp int) string {
    slice := tm.kv[key]
    idx := sort.Search(len(slice), func(i int) bool {
        return slice[i].timestamp > timestamp
    }) - 1
    if idx < 0 { return "" }
    return slice[idx].value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
