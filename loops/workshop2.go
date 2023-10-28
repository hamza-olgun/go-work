package loops

import "fmt"

// Aklımdaki sayıyı tahmin etme uygulaması + kaç tahminde bildiğin
func Demo4() {

	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {

		if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Sayını küçült")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1

		}

		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Sayını büyült")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1

		}

	}
	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"

	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Printf("BİLDİN! %v tahmin: %v", tahminSayisi, basariDurumu)
}
