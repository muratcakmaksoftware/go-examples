package main

import "fmt"

type User struct {
	Name  string // Büyük harfle başlamak dışarıya bu veriyi açmak demektir - encoding/json
	age   int    // Küçük harfle başlamak sadece paket içi veri erişimi demektir
	Pages []Page
}

type Page struct {
	ID    int
	Title string
	URL   string
}

func main() {
	//Değişkenler
	test1 := "murat"
	test1 = test1 + " cakmak"
	test2 := &test1     //test2 de bellek adres değerinde test1 bellek adresi tutulur
	fmt.Println(*test2) //bellek adresindeki degeri yazdirmak için * koyularak adresdeki veriyi yazdırırız

	var test3 string //test4 ile aynı anlama gelmektedir. Default değeri "" eğer int olsaydı 0 olacaktı
	test4 := ""
	fmt.Println(test3, test4)

	var test5 *string //başlangıç değeri nil dir. Eğer nil vermek istiyorsak tipini belirtmek zorundayız bellekte açılacak tipi bilmelidir.
	fmt.Println(test5)

	//Veri Modelleme / Formatting
	user := User{}
	user.Name = "Murat"
	user.age = 28
	user = User{Name: "murat"}
	user.age = 25

	fmt.Println(user)

	// IF
	if 10 > 5 {
		fmt.Println("10 sayısı 5 den büyüktür")
	}

	if err := "murat"; err == "" {
		fmt.Println("err boş")
	} else if err == "murat" {
		fmt.Println("err = murat")
	} else {
		fmt.Println("err boş değil")
	}

}
