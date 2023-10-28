package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	isimler[0] = "Engin"
	isimler[1] = "Derin"
	isimler[2] = "Salih"
	isimler = append(isimler, "Büşra") //append yeni bir array eklemek için kullanılır.
	fmt.Println(isimler)

}
