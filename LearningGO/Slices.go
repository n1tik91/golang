
package main
import ("fmt")

func main() {
  myslice1 := []int{} //empty slice of 0 length and 0 capacity
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  //slice of strings of length 4 and capacity 4
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}