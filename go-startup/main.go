package main

import (
	"errors"
	"fmt"
	"math"
	"time"

	"golang.org/x/exp/constraints"
)

type User struct {
	Name  string // Büyük harfle başlamak dışarıya bu veriyi açmak demektir - encoding/json
	age   int    // Küçük harfle başlamak sadece paket içi veri erişimi demektir
	Pages []Page // Struct Embedding
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

	//Switch
	x := 2
	age := 20
	mesaj := "merhaba"

	switch {
	case x == 1:
		fmt.Println("x: Bir")
	case x == 2 || x == 4 || x == 6: // çoklu kontrol
		fmt.Println("x: Çift sayılardan biri")
		fallthrough // bir sonraki case çalışsın (zorla geç)

	case age >= 18 && age < 65: // if gibi kullanım (switch true gibi)
		fmt.Println("age: Yetişkin")

	case mesaj == "merhaba":
		fmt.Println("mesaj: Selam!")

	default:
		fmt.Println("Hiçbirine uymadı")
	}

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

	//Pointer
	i = 1
	zeroptr(&i)
	fmt.Println("zeroval:", i)
	fmt.Println("pointer:", &i)

	//Struct & Methods // Burada Struct tanımlayıp struct bağlı fonksiyonlar desteklemektedir.
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	//Interfaces
	r2 := rect2{width: 3, height: 4}
	c2 := circle{radius: 5}

	//Interface ile 2 farklı struct birleştirdik ortak metodları kullanmış olduk
	measure(r2)
	measure(c2)

	//Burada ise if ile struct ayırma tespitini gördük.
	detectCircle(r2)
	detectCircle(c2)

	//ENUMS / Normalde Enum desteklemiyor ama uyduruyoruz
	var myRole Role = Editor
	fmt.Println("Rol:", myRole)

	//Named result parameters // Return boş gönderildi ancak içindeki değişkenler isimli şekilde fonksiyon dönüşüne yazıldı.
	carp, topla := carpVeTopla(3, 4)
	fmt.Println("Çarpım:", carp)  // 12
	fmt.Println("Toplam:", topla) // 7

	//Defer // İlgili fonksiyon işini bitirdiğinde defer edilen fonksiyon çalışır.
	fmt.Println("İlk fonksiyon çalıştı")
	defer fmt.Println("İkinci fonksiyon defer edildi en son çalışcak")
	fmt.Println("Üçüncü fonksiyon çalıştı")

	//Generics Burada birden fazla farklı tipi karşılaştırmalarını yapabildiğimiz metodu geliştirmektir.
	//Gönderdiğimiz veri tipleri farklı ve buna göre karşılaştırmayı yapabilecek fonksiyon da yazmaktadır.
	fmt.Println(Max(3, 7))           // int → 7
	fmt.Println(Max(4.2, 2.1))       // float64 → 4.2
	fmt.Println(Max("apple", "cat")) // string → cat

	//Errors Handling
	sonuc, err := bol(10, 0) //Örnek bölme hata fırlatımı
	if err != nil {
		fmt.Println("Hata:", err)
	}
	fmt.Println("Sonuç:", sonuc)

	err2 := kullaniciVerisiniOku() //Örnek hata sarmalama yapımı ve yazdırma.
	if err2 != nil {
		fmt.Println("En üst hata:", err2)

		// Zinciri çözerek her hatayı sırayla alalım
		for i := 1; err2 != nil; i++ {
			fmt.Printf("Seviye %d hata: %v\n", i, err2)
			err2 = errors.Unwrap(err2)
		}
	}

	//Custom Error
	err = checkAge(16)
	if err != nil {
		if ageErr, ok := err.(*AgeError); ok { // err.(*AgeError) mu kontrolü yapılır. ageErr hata mesajını verir, ok ise Error tipinin doğru olup olmadığı true false olarak yanıtlar.
			fmt.Println("Özel hata yakalandı!")
			fmt.Printf("Yaş çok küçük: %d\n", ageErr.Age)
			fmt.Println(ageErr.Error()) //Uzun yol basım
			fmt.Println(err)            //Kısayol basım
		} else {
			fmt.Println("Genel bir hata:", err)
		}
	} else {
		fmt.Println("Yaş uygun, devam edebilir.")
	}

	//Goroutines //Bunlar hafif iş parçacıkları çalıştıraiblir. Ana işlem(main) biterse tümü işlemdeyse bile kapanır.
	go f("goroutine") // Yeni bir goroutine olarak başlar

	go func(msg string) { // Anonymous (anonim) goroutine örneği
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) //Hafif iş parçacıkların mesaj basımı için bekletme yapıyoruz.
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

func zeroptr(iptr *int) {
	*iptr = 0
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type geometry interface {
	area() float64
	perim() float64
}

type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) { //geometry interface
	if c, ok := g.(circle); ok { //Circle struct mı ?
		fmt.Println("circle with radius", c.radius)
	}
}

type Role int

const (
	Admin  Role = iota //iota 0,1,2 gibi const girilen kadar indeksler ve Role ise buradaki değer tip int değil Role tipindedir.
	Editor             //1
	Viewer             //2
)

var roleNames = map[Role]string{
	Admin:  "Yönetici",
	Editor: "Düzenleyici",
	Viewer: "Görüntüleyici",
}

func (r Role) String() string { //Role türünde string basımı olduğu zaman otomatik çağrılan method dur. role.String() manuel çağrı
	if val, ok := roleNames[r]; ok {
		return val
	}
	return "Bilinmeyen"
}

func carpVeTopla(a int, b int) (carpim int, toplam int) { //Named result parameters
	carpim = a * b
	toplam = a + b
	return //carpim & toplam değişkenlerini göndermiş olur çünkü fonksiyon dönüşünde belirtilmiştir.
}

// Burada constraints.Ordered Interfaces kullanıldı detayına gidip baktığımızda. ~int gibi bir ifade göreceğiz.
// ~ en temelde int tipi varsa anlamındadır. "Ör: type MyInt int" bir tip tanımladım "var x MyInt = 5" int değerinde temelde değişken oluşturdum ancak tip adım MyInt ancak temelde int.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func bol(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("bölme sıfıra karşı yapılamaz")
	}
	return a / b, nil
}

func sistemOku() error {
	// En alt seviye hata
	return errors.New("disk okuma hatası")
}

func dosyaOku() error {
	err := sistemOku()
	if err != nil {
		// 2. seviye: wrap'liyoruz
		return fmt.Errorf("dosya erişimi başarısız: %w", err)
	}
	return nil
}

func kullaniciVerisiniOku() error {
	err := dosyaOku()
	if err != nil {
		// 3. seviye: tekrar wrap
		return fmt.Errorf("kullanıcı verisi alınamadı: %w", err)
	}
	return nil
}

// Custom hata mesajı için örnek struct
type AgeError struct {
	Age int
}

func (e *AgeError) Error() string { //Hata mesajını verebilmemiz için struct ait Error Metodu
	return fmt.Sprintf("Geçersiz yaş: %d", e.Age)
}

func checkAge(age int) error {
	if age < 18 {
		return &AgeError{Age: age} //Hata bilgisinin hazırlanıp, gönderilmesi
	}
	return nil
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}
