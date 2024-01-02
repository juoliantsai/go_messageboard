package main
import (
"fmt"
"sync"
)
// 參考資料https://vocus.cc/article/64ff0f16fd8978000191a081
func main() {
  // 讓主程式等待所有的Goroutine完成
  var wg sync.WaitGroup

  // 增加計數器的值，表示有多少個Goroutine需要等待
  wg.Add(1)
  go routine1(&wg)
  wg.Add(1)
  go routine2(&wg)
  // 當計數器的值為零時，Wait()會返回，並允許主程式繼續執行
  wg.Wait()
  fmt.Println("All Goruntines have finished")
}
func routine1(wg *sync.WaitGroup){
  for i:=0;i<10;i++{
    fmt.Println("0")
  }
  // 減少計數器的值，表示一個Goroutine已完成
  defer wg.Done()
}
func routine2(wg *sync.WaitGroup){
  for i:=0;i<10;i++{
    fmt.Println("2")
  }
  defer wg.Done()
}