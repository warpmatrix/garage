func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }
    return lhs
}

func trap(height []int) int {
    ret := 0
    lptr, rptr := 0, len(height) - 1
    lmax, rmax := 0, 0
    for lptr < rptr {
        lmax = max(lmax, height[lptr])
        rmax = max(rmax, height[rptr])
        if lmax < rmax {
            ret += lmax - height[lptr]
            lptr++
        } else {
            ret += rmax - height[rptr]
            rptr--
        }
    }
    return ret
}