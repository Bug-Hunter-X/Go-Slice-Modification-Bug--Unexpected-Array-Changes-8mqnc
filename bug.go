func myFunc(a []int) {for i := range a {
    a[i]++
}}

func main() {
a := []int{1, 2, 3}
myFunc(a)
fmt.Println(a) // Output: [2 3 4]

b := []int{1,2,3}
myFunc(b[:2])
fmt.Println(b) //Output: [2 3 3] 
}