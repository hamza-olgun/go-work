package variables

import "fmt"

func Demo1() {
	//Değişkenler
	var metin string = "Merhaba Türkiye!"
	fmt.Println(metin)

	var kdv int = 10 //int: tam sayılar için kullanılır.
	fmt.Println(kdv)

	var kdv2 float32 = 0.1 //float:32 ve 64 bit seçenekleri vardır. Ondalıklı sayılar için kullanılır.
	fmt.Print(kdv2)

	/* Bu değişkenler dışında kdv:=30 kdv2:="Hamza" şeklindede değişken tanımlayabiliriz.
	Değişkenin türünü go kendi anlar ve belirler.
	*/
	var durum bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2 // iki tane eşittir(==) işareti 'eşit mi?' anlamına gelir.(!=) ise farklı mı? anlamına gelir.

	fmt.Println(durum)

}
