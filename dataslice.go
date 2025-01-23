package main 
import "fmt" 

func main() {
  names := [...] string {"farhan", "maulana", "iqbal", "mcgilly", "cuddy"}
  slice := names[1:3] 

  fmt.Println(slice[0])
  fmt.Println(slice[1])

  slice1 := names[:2]
  fmt.Println(slice1)

  slice2 := names[2:]
  fmt.Println(slice2)

  slice3 := names[:]
  fmt.Println(slice3)

  //append
  days := [...] string {"senin","selasa","rabu","kamis","jumat","sabtu,","minggu"}
  dayslice1 := days[5:]
  dayslice1[0] = "sabtu baru"
  dayslice1[1] = "minggu baru"
  fmt.Println(days)

  dayslice2 = append(dayslice1, "libur baru")
  dayslice2[0] = "ups"
  fmt.Println(dayslice2)
  fmt.Println(days)

}
