func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var i int
    for _, v := range hours {
        if v >= target {
            i++
        }
    } 
    return i
}
