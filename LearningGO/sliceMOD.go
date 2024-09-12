package main
import ("fmt")

func main() {
  prices := []int{10,20,30}
  prices[2] = 50 //change the third element in the prices slic
  fmt.Println(prices[0])
  fmt.Println(prices[1])
  fmt.Println(prices[2])
}