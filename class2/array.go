package class2

import ("fmt")

//prywatna
//array
var array = [7]int{}

//slice - powiększalny
var slice = []int{}
//incepcja
var inceptionSlice = [][]int{}

//___string to tablica bajtów___ | znaki polskie to 2 bajty

func ArrayFunc(){
	//print
	fmt.Println("-------------------------------------------")

	fmt.Println("array pusty (7 elementów): ",array)

	fmt.Println("slice pusty: ",slice)

	//dodaj
	slice = append(slice, 1,2,3)
	fmt.Println("slice dodany: ",slice)

	//wybierz
	fmt.Println("slice 2 elementy: ",slice[:2])

	str := "Owca Feministka Terrorystka" // szybkie zmienne tylko w funkcji
	fmt.Println("test: ",str)
}