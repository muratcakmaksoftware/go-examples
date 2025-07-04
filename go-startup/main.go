package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
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

	//Channels
	messages := make(chan string) //Chan ile channel oluşturulur veri tipi ise string aktarılacağı bellidir. Burada struct da kullanabiliriz farklı verileri göndermek istiyorsak

	//İş parçacağında işlem yapılıyor ve sonuç bilgisi channel dan gönderiliyor.
	go func() {
		//...
		//...
		messages <- "ping"
	}()

	msg := <-messages //İş parçacağındaki işlem bitene kadar beklemiş olur kısacası burada channeldan mesaj gelen kadar beklemiş olacak.
	fmt.Println(msg)

	//Channel Buffering
	ch := make(chan string, 2) //buradaki 2 parametre kapasitesidir 2 mesaj alabilir. Ayni anda 3 mesaj gelirse 3 mesajda durur ve ilerlemez taki 1 mesaj eriyene kadar
	ch <- "mesaj 1"
	ch <- "mesaj 2"

	go func() { //is parcacagini bloklayacak duruma getirdik 3 mesaj ekleyip.
		ch <- "mesaj 3" //bloklanacak
	}()

	fmt.Println("1 mesaj alınıyor:", <-ch) //1 mesaj eridigi icin 3 mesaj siraya girmiş oldu ve iş parçacağı devam etti.
	fmt.Println("2 mesaj alınıyor:", <-ch)
	fmt.Println("3 mesaj alınıyor:", <-ch)

	//Channel Synchronization -> Channel ile aynı fark olarak sadece sinyal olarak true/false dönülmesi.
	done := make(chan bool, 1) //bool tipinde channel başlatılıyor maks kapasite 1
	go worker(done)            //iş parçacağı başlatılıp channel bilgisiyle gönderiliyor.
	<-done                     //iş parçacağından veri dönüşü bekler.

	//Channel Directions
	//Burada data aktarımı yaparken fonksiyonlarda producer ve consumer belirteçleri kesinlik içermektedir. Hatalı gönderim yapılamasın diyedir.
	pings := make(chan string, 1) //Producer channel
	pongs := make(chan string, 1) //Transfer channel
	ping(pings, "passed message") //Mesaj ping işlenir. Burada önemli olan fonksiyondaki parametreler.
	pong(pings, pongs)            //Pings den Pongs mesaj taşınır. Burada önemli olan fonksiyondaki parametreler.
	fmt.Println(<-pongs)

	//Select // Birden fazla iş parçacağı ile çalışırken hangisi önce gönderirse yakalayıp işlem yapmamızı sağlar.
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() { //ilk iş parçacağı mesaj gönderimi
		time.Sleep(1 * time.Second) //1 saniye sonra
		channel1 <- "one"
	}()
	go func() { //ikinci iş parçacağı mesaj gönderimi
		time.Sleep(2 * time.Second) // 2 saniye sonra
		channel2 <- "two"
	}()

	for i := 0; i < 2; i++ { //For döngüsünü sonsuz yapsaydık sürekli olarak bu channel kontrolünde mesaj geldimi gibi bakacaktı.
		select { //Birden fazla channel ile çalışırken hangisinden önce mesaj gelirse o case gelip oradaki işlemleri çalıştırmasını sağlar.
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2: //buraya hiç girmeyecek çünkü timeout takılacak ikinci iş parçacağı 2 saniye sürmektedir.
			fmt.Println("received", msg2)
		case <-time.After(1 * time.Second): // Timeouts // her select kontrolünde eğer işlem 1 saniyeden uzun sürerse bu case'e girecek.
			fmt.Println("timeout 1")
		default: // Non-Blocking Channel Operations // Burada channelda herhangi bir mesaj yok ise default a girecektir. Böylece mesaj gelmesi beklemeden bu case işlenecektir. Bloklama olmayacaktır.
			fmt.Println("no message received")
		}
	}

	//Closing Channels
	channel := make(chan int, 3)
	channel <- 10
	channel <- 20
	channel <- 30
	close(channel)         // channel kapama // Channel kapansa bile veriler alınabilirdir. Ancak channel kapandığında ekleme yapılamaz.
	fmt.Println(<-channel) // 10
	fmt.Println(<-channel) // 20
	fmt.Println(<-channel) // 30
	fmt.Println(<-channel) // 0 buradakı sıfır boş ve kapalı anlamındadır

	//Range over Channels //Burada range ile channeldaki gönderdiğimiz dataları okuyabiliriz.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	//Timer
	timer1 := time.NewTimer(2 * time.Second) //2 saniye bekleyecek timer oluşturulur ve timer'ın süresi tamamlandığında channeldan dönüş sağlar. //Tek seferliktir tekrarlanmaz bu timer.
	<-timer1.C                               //.C = Channel 2 saniye sonra sinyal alacak channeldır bu yüzden burada channel dan dönüş bekler geldiğinde o süre geçmiştir.
	fmt.Println("2 saniye geçti")            // işlem biter.

	stop := timer1.Stop() //timer beklemesi iptal edildi. yukarıdaki timer iş parçacağında olsaydı ve daha erkenden timer kapatsaydık süreyi beklemeyecekti gibi.
	if stop {
		fmt.Println("Timer 1 stopped")
	}

	//Ticker
	ticker := time.NewTicker(1 * time.Second) //Ticker belirtilen zaman kadar sürekli tekrar çalışır ve channel bilgi yollar.

	// Ticker sonsuza kadar çalışmasın diye sonlandırmak için 3 kez mesaj aldıktan sonra ticker sonlandıralım
	count := 0
	for t := range ticker.C { //Ticker.C ile channeldan gelen bilgileri dinliyoruz
		fmt.Println("Zaman:", t)
		count++
		if count == 3 {
			ticker.Stop()
			fmt.Println("Ticker durduruldu.")
			break
		}
	}

	//Worker Pools //Birden fazla consumer oluşturabiliriz bu dinleyiciler iş gönderildikçe işleyecek
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3 tane worker başlattık ve id verdik worker'a hangi workerda işin işlendiğini görebileceğiz.
	for w := 1; w <= 3; w++ { //
		go worker2(w, jobs, results) //id, dinlenecek iş channel'ı, sonuçların aktarılacağaı results channel'ı
	}

	// 5 iş gönder
	for j := 1; j <= 5; j++ { // ilk 5 işin gönderilmesi ve workerlardan hangisi ilk yakalarsa o çalıştırılacak.
		jobs <- j
	}
	close(jobs) //Channel sonlandırıldı artık içeriye veri alamaz.

	for a := 1; a <= 5; a++ { // dönülen 5 sonucu bekletmek için
		<-results
	}

	//WaitGroups birden fazla goroutine (concurrent) yani eşzamanlı çalıştırıp onları işlerini bitmesini kontrol eden yapıdır.
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)          // her goroutine için 1 iş eklendi sayac artımı
		go worker3(i, &wg) //eklenen iş. wg gönderilmesinin nedeni iş parçacağın işi bitince sayaç düşümü sağlamak amaçlı
	}

	wg.Wait() // Hepsi iş parçacıklarının bitimini beklenir.
	fmt.Println("Tüm işler tamamlandı.")

	//Tick - Rate Limiting Örneği
	// Her 500 milisaniyede bir zaman sinyali gelir
	limiter := time.Tick(500 * time.Millisecond) //her 500 mili saniyede channelden veri gönderir

	//Rate Limiting örnek
	requests := []string{"İstek 1", "İstek 2", "İstek 3", "İstek 4", "İstek 5"}

	for _, req := range requests {
		<-limiter // channelden 500 milisaniye deki veriyi bekler ve devam eder.
		fmt.Println("İşleniyor:", req, "Zaman:", time.Now().Format("15:04:05.000"))
	}

	//Atomic Counters -- çoklu goroutine ile çalışırken güvenli şekilde sayaç artımı yapabiliriz.
	var counter uint64     // benzer sayaç örneği
	var wg2 sync.WaitGroup // Bekleyici grup
	var ops atomic.Uint64  // Başka bir sayaç örneği

	for i := 0; i < 50; i++ {
		wg2.Add(1)  //waitgroup sayacı
		go func() { //iş parçacağı
			for j := 0; j < 1000; j++ { // 1000 kez atomik artır
				atomic.AddUint64(&counter, 1) //counter++ yerine güvenli bir şekilde artırır. Eğer direk counter++ kullansaydık aynı anda yazım işlemlerinde çakışma olur ve yanlış bir artım sonucu elde ederdik.
				ops.Add(1)                    //Atomic Güvenli Artım
			}
			wg2.Done() //waitgroup dan sayac düşümü
		}()
	}

	wg2.Wait()                            //tüm işlemin bitmesini bekler.
	fmt.Println("Toplam Sayım:", counter) // 50 iş parçacağında her birinde 1000 artımı doğru orantıda sayacı arttırır. 50000 olur.
	fmt.Println("ops:", ops.Load())       // artımdan gelen datayı güvenli şekilde yazdırılması

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

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// Sadece channel veri aktarmak için //Procuder
func ping(pings chan<- string, msg string) { // pings chan<- string (sadece send-only) kanal tipi. Yani sadece veri gönderilir açıkça belirtilir.
	pings <- msg //pings channel'a msg gönderilir.
}

// Channeldaki verileri tüketmek için //Consumer
func pong(pings <-chan string, pongs chan<- string) { //pings <-chan string (sadece receive-only) kanal tipi. Yani sadece veri alınır açıkça belirtilir.
	msg := <-pings //channeldaki data alanır
	pongs <- msg   //alınan mesaj pongs channela yazılır
}

func worker2(id int, jobs <-chan int, results chan<- int) { //Sadece jobs dan gelen verileri dinlenir ve işlenen veriler results channel'a aktarılır.
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func worker3(id int, wg *sync.WaitGroup) {
	defer wg.Done() // iş bitince sayaç dan veriyi düşmek içindir. Ne olursa olsun hata olsada defer yapıldığı için bu fonksiyonda işlem sonlandığında bu işlem çalışmış olacak.

	fmt.Printf("Worker %d başlıyor\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d bitti\n", id)
}
