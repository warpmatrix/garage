type ParkingSystem struct {
    int slots[4]
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{{0, big, medium, small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if slots[carType] > 0 {
        slots[carType]--
        return true
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */