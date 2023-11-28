package main

import (
	"os"
	"time"

	"hello/mathTutorial/clockface"
)

func main() {
	ist,_:=time.LoadLocation("Asia/Kolkata")
	t := time.Now().In(ist)
	clockface.SVGWriter(os.Stdout, t)
}