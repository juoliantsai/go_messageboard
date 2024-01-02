package main
import (
"fmt"
"time"
"github.com/go-co-op/gocron"
)

func main() {
  sch:=gocron.NewScheduler(time.UTC)
  cronExp:="28 * * * *"
  _,err:=sch.Cron(cronExp).StartImmediately().Do(func(){
    fmt.Println("hello")
  })
  if err!=nil{
    panic(err)
  }
  sch.StartBlocking()
}
func newJob(schChan chan *gocron.Scheduler) {
  sch:=gocron.NewScheduler(time.UTC)
  cnt:=0
  _,err:=sch.Every(1).Seconds().Do(func(){
    cnt++
    fmt.Println(cnt)
  })
  if err!=nil {
    panic(err)
  }
  schChan<-sch
  sch.StartBlocking()
}
