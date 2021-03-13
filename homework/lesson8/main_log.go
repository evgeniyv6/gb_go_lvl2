package main

import "log"

func main() {
	log.SetFlags(log.LUTC | log.Lmicroseconds | log.Lshortfile | log.Ldate)
	log.Println("test")

	//log.Panic("this is panic")

	log.Fatal("exit")
}
