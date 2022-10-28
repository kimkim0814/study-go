package main

import "fmt"

//for

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3{
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 0 {
	// 	fmt.Println(point)
	// 	point ++
	// }

	// for i := 0; i < 10;i++{
	// 	if i == 3{
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//PHPのforeachみたいな使い方ができる
	// arr := [3]int{1,2,3}
	// for i:=0; i< len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1,2,3}

	// for _,v := range arr{
	// 	fmt.Println(v)
	// }

	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{"appple":100,"banana":200}

	for k,v := range m{
		fmt.Println(k,v)
	}
}
