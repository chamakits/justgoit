package main

import (
	"fmt"
	"justgoit/files"
)

func main() {
	fmt.Println(files.Md5Sum("/home/chamakits/"))
	fmt.Println(files.Md5Sum("/home/chamakits/CantStopProcess.png"))

}