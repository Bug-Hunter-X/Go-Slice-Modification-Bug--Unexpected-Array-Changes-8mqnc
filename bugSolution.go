func myFunc(a []int) {
    b := make([]int, len(a))
    copy(b, a)
    for i := range b {
        b[i]++
    }
    fmt.Println("inside function:",b)
}

func main() {
a := []int{1, 2, 3}
myFunc(a)
fmt.Println(a) // Output: [1 2 3]

b := []int{1,2,3}
myFunc(b[:2])
fmt.Println(b) //Output: [1 2 3]
} 