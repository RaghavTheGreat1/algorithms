func CreateTargetArray(nums []int, index []int) []int {
    n := len(nums)
    target := make([]int, 1)

    for i,v := range nums{
        indexToInsert := index[i]
        target = append(target[:indexToInsert+1], target[indexToInsert:]...)
        target[indexToInsert] = v
    }

    return target[:n]
}
