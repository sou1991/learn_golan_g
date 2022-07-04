package main 

import (
	"fmt"
	"time"
)

type MyErrors struct {
	time time.Time
	messege string
}

//fmtのerrorinterfaceの実装
// type error interface {
//     Error() string
// }

func (e *MyErrors) Error() string{
	return fmt.Sprintf("at %v, %s",
		e.time, e.messege)
}

func run() error {
	return &MyErrors{
		time.Now(),
		"it didn't work",
	}
}

func main(){
	if err := run(); err!=nil {
		fmt.Println(err)
	}
}
