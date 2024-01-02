package main
import (
"fmt"
"sync"
)

func main() {
  var wg sync.WaitGroup

  wg.Add(1)
  go routine1(&wg)
  wg.Add(1)
  go routine2(&wg)
  
  wg.Wait()
  fmt.Println("All Goruntines have finished")
}
func routine1(wg *sync.WaitGroup){
  for i:=0;i<10;i++{
    fmt.Println("0")
  }
  defer wg.Done()
}
func routine2(wg *sync.WaitGroup){
  for i:=0;i<10;i++{
    fmt.Println("2")
  }
  defer wg.Done()
}
