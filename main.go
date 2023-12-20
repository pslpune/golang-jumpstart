package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main(){
	fmt.Println("this is from inside my program")
	log.Debug("this is from inside my program")
}