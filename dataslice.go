package main 
import "fmt" 

func main() {
  names := [...] String {"farhan", "maulana", "iqbal", "mcgilly", "cuddy"}
  slice := names[1:3] 

  fmt.Println(slice[0])
  fmt.Println(slice[1])

  slice1 := names[:2]
  fmt.Println(slice1)

  slice2 := names[2:]
  fmt.Println(slice2)

  slice3 := names[:]
  fmt.Println(slice3)

}
