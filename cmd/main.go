package main

import ("github/Ko4s/goCourse/class2"; "fmt")

func main() {
	class2.ArrayFunc()
	var f = class2.OpenFile("./class2/a.txt")
	fmt.Println(f)
	var fReadLines = class2.ReadLines("./class2/a.txt")
	fmt.Println(fReadLines)
}
