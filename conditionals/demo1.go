package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız Hazırlanıyor")
		hesap = hesap - cekilmekIstenen

	}
	fmt.Printf("Bitti.Hesaptaki para : %v", hesap)
}
