package controllers

import (
"fmt"
"time"
)

func SayHello(param string){
  fmt.Println(param, time.Now())
}
