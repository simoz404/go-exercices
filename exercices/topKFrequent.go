func topKFrequent(nums []int, k int) []int {
    s := make(map[int]int)
    for _, v := range nums {
        s[v]++
    }
    var n, m int
    var arr []int
    for l := 0; l < k; l++ {
        for i, v := range s {
            if v > n {
                n, m = v, i
            }
        }
        arr = append(arr, m)
        delete(s, m)
        m, n = 0, 0
    }
    return arr
}
