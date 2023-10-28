package loops

import "fmt"

//Aklımdaki sayıyı tahmin etme uygulaması
func Demo3() {

	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)

	for aklimdakiSayi != tahminEdilenSayi {

		if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Sayını küçült")
			fmt.Scanln(&tahminEdilenSayi)

		}

		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Sayını büyült")
			fmt.Scanln(&tahminEdilenSayi)

		}

	}
	fmt.Println("BİLDİN!")
}
