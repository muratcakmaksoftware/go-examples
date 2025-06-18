package main

import (
	"fmt"
)

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

const const1 = "Test"
const const2 string = "Test2"

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

	// if else de "else" in gereksiz olduğum durum. Fonksiyondan 2 tane data dönüyoruz ve error varsa direk hatalıya çekiyoruz.
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Sonuç:", result)

	//Constant
	const const3 = 100 / 2
	fmt.Println(const1)
	fmt.Println(const2)
	fmt.Println(const3)

	//FOR
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("infinity loop")
		break
	}

	fruits := map[string]int{
		"elma":     3,
		"muz":      5,
		"portakal": 2,
	}

	for key, value := range fruits {
		fmt.Printf("%s - %d\n", key, value) //Burada %s = string %d = decimal sayı anlamında basmak için C de kullanılır
	}

	for key := range fruits {
		if key == "elma" {
			delete(fruits, key)
		}
	}
	fmt.Println(fruits)

	for _, value := range fruits { // _ = underline ile key atlayabiliriz
		fmt.Println(value)
	}

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding // Burada karakterlerin unicode bilgisini görebiliyoruz "pos"
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	a := []int{1, 2, 3, 4, 5}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 { //Burada tek forda iki tane sayaç kullanılmıştır
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("Ters çevrilmiş:", a)

	//ARRAYS
	var isim [3]string
	isim[0] = "Ali"
	isim[1] = "Veli"
	isim[2] = "Ayşe"
	isim = [3]string{"Ali", "Veli", "Ayşe"}
	fmt.Println(isim)

	var sayilar = [...]int{1, 2, 3, 4, 5}
	fmt.Println(sayilar)

	matris := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(matris, matris[1][2])

	//Slice / Dinamik Array
	slice := []int{1, 2, 3} // [] boş slice demek
	fmt.Println(slice)
	fmt.Println(len(slice)) //Uzunluk
	fmt.Println(cap(slice)) //Kapasite
	slice = append(slice, 4)
	slice = append(slice, 5, 6)
	fmt.Println(slice)

	var sliceEmpty []string
	fmt.Println("uninit:", sliceEmpty, sliceEmpty == nil, len(sliceEmpty) == 0)

	s := make([]int, 0, 5) // type, 0 uzunlukta, 5 kapasiteli slice oluştur
	s = append(s, 10)
	s = append(s, 20, 30, 40, 50)
	fmt.Println(s)      // [10 20 30]
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 5

	//Slice kopyalama
	c := make([]int, len(s))
	copy(c, s)

	fmt.Println("cpy:", c)
	l := s[0:5] // 0 dan başla 5 tane al diyoruz. //[10 20 30 40 50]
	fmt.Println(l)

	l = s[:3] //en baştan başla 2 kadar al - başlanan yer dahildir. //[10 20 30]
	fmt.Println("sl2:", l)

	l = s[3:] //3 başla sona kadar al - başlanan yer dahil değildir. // [40 50]
	fmt.Println("sl3:", l)

	//Make
	//make(map[key-type]val-type)
	m := make(map[string]int) // map[string]int{"foo": 1, "bar": 2}
	m["test1"] = 1
	m["test2"] = 2
	fmt.Println("Map: ", m)

	delete(m, "test2")
	fmt.Println("Map: ", m)

	clear(m)
	fmt.Println("map:", m)

	_, exists := m["test2"] //identifier ile degere ihtiyacin olmadigini belirtirsin eger degere ihtiyacin varsa _ yerine val gibi degisken kullanilabilir
	if exists {
		fmt.Println("test exists")
	} else {
		fmt.Println("test doesnt exists ")
	}

	//Func
	res := plusPlus(1, 2, 3)
	fmt.Println(res)

	result1, result2 := murat() //Çoklu return dönüşü
	fmt.Println(result1, result2)

	//Variadic Functions
	sum(1, 2) // Çoklu parametre aktarımı
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//Closures
	nextInt := intSeq() // Fonksiyonu değişkende kullanım
	fmt.Println(nextInt())

	//Recursion // Tekrarlama
	fmt.Println(fact(7))

	//Anonymous functions
	var fib func(n int) int //Anonim Fonksiyon tanımlamak için önce declare edilmesi gerek.
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0'a bölme hatası")
	}
	return a / b, nil
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func murat() (int, int) { //Çoklu return dönüşü
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int { // İç fonksiyon yazımı
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
