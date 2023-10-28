package slices

import "fmt"

func Demo2() {

	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[:3]) //İlk kısım boş bırakılırsa sıfırdan başla demektir.
	fmt.Println(sehirler[1:]) // Son kısım boş bırakılırsa kaçtane varsa sonunda kadar yazdır demektir.

}
