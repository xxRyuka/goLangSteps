[Complex Numbers 2f4ad6137d248071aefadd1332926c16.md](https://github.com/user-attachments/files/24987782/Complex.Numbers.2f4ad6137d248071aefadd1332926c16.md)
# Complex Numbers

ASP.NET Core ile geliştirdiğin standart REST API projelerinde muhtemelen `System.Numerics.Complex` yapısını hiç kullanmamışsındır. Go'da ise karmaşık sayılar (Complex Numbers), dilin sözdizimine (syntax) gömülmüş **birinci sınıf vatandaşlardır (first-class citizens)**.

Bu, Go'nun sadece bir "Web API dili" değil, aynı zamanda C ve C++'ın domine ettiği "Sistem ve Bilimsel Hesaplama" alanına da talip olduğunun mimari bir göstergesidir.

Backend mühendisliği kariyerinde belki çok nadir kullanacaksın ama "Language Design" (Dil Tasarımı) açısından bu yapıyı bilmek genel vizyonuna katkı sağlar.

### 1. Mimari Yapı ve Bellek Düzeni

Go'da karmaşık sayılar, bellekte **ardışık duran iki adet float** sayıdan ibarettir. Başka bir "sihir" veya obje referansı yoktur.[Comma-ok Idiom 2f9ad6137d248063b9a1ea2bc83cc10a.md](https://github.com/user-attachments/files/24987873/Comma-ok.Idiom.2f9ad6137d248063b9a1ea2bc83cc10a.md)
[Bool 2f4ad6137d24807483dbd3a4cc17d629.md](https://github.com/user-attachments/files/24987871/Bool.2f4ad6137d24807483dbd3a4cc17d629.md)
[basically 2f8ad6137d2480799381e33294746fcd.md](https://github.com/user-attachments/files/24987869/basically.2f8ad6137d2480799381e33294746fcd.md)
[Arrays 2f4ad6137d2480eebf4cdf621c2fc6eb.md](https://github.com/user-attachments/files/24987867/Arrays.2f4ad6137d2480eebf4cdf621c2fc6eb.md)
[1 2 Multiple Return Values 2faad6137d2480dab0b4e2c19962e929.md](https://github.com/user-attachments/files/24987865/1.2.Multiple.Return.Values.2faad6137d2480dab0b4e2c19962e929.md)
[1 1 Function Basics Sözdizimi, Parametreler ve C#  2f9ad6137d248017ba4ef00b114c73b2.md](https://github.com/user-attachments/files/24987863/1.1.Function.Basics.Sozdizimi.Parametreler.ve.C.2f9ad6137d248017ba4ef00b114c73b2.md)
[Zero Values 2f4ad6137d24806f9509f32753c45eb5.md](https://github.com/user-attachments/files/24987861/Zero.Values.2f4ad6137d24806f9509f32753c45eb5.md)
[Variables and Constans 2f4ad6137d2480929f08decef5270838.md](https://github.com/user-attachments/files/24987859/Variables.and.Constans.2f4ad6137d2480929f08decef5270838.md)
[var & = 2f4ad6137d2480d3b6acdca23365d6f7.md](https://github.com/user-attachments/files/24987857/var.2f4ad6137d2480d3b6acdca23365d6f7.md)
[Untitled 2f9ad6137d2480898d30ca20c1ad4674_all.csv](https://github.com/user-attachments/files/24987855/Untitled.2f9ad6137d2480898d30ca20c1ad4674_all.csv)
[Untitled 2f9ad6137d2480898d30ca20c1ad4674.csv](https://github.com/user-attachments/files/24987852/Untitled.2f9ad6137d2480898d30ca20c1ad4674.csv)
[Type Conversion 2f4ad6137d2480f09d79fe3198876712.md](https://github.com/user-attachments/files/24987850/Type.Conversion.2f4ad6137d2480f09d79fe3198876712.md)
[Struct Tags 2f9ad6137d24802e8ed0dc7ac6a6cc25.md](https://github.com/user-attachments/files/24987848/Struct.Tags.2f9ad6137d24802e8ed0dc7ac6a6cc25.md)
[Struct JSON 2f9ad6137d24805098cde8fea115d4e2.md](https://github.com/user-attachments/files/24987846/Struct.JSON.2f9ad6137d24805098cde8fea115d4e2.md)
[Struct 2f9ad6137d24809cb9bff70ab822c148.md](https://github.com/user-attachments/files/24987845/Struct.2f9ad6137d24809cb9bff70ab822c148.md)
[Strings Review 2f8ad6137d2480c1abacf098ee07e0d6.md](https://github.com/user-attachments/files/24987843/Strings.Review.2f8ad6137d2480c1abacf098ee07e0d6.md)
[stringer ınterface (early) 2f4ad6137d2480a7b1a3c6f005fc0c5e.md](https://github.com/user-attachments/files/24987841/stringer.interface.early.2f4ad6137d2480a7b1a3c6f005fc0c5e.md)
[String Type conversion 2f4ad6137d248066875edf171d801f35.md](https://github.com/user-attachments/files/24987840/String.Type.conversion.2f4ad6137d248066875edf171d801f35.md)
[String 2f4ad6137d2480c0bb17d07f63db522b.md](https://github.com/user-attachments/files/24987838/String.2f4ad6137d2480c0bb17d07f63db522b.md)
[Slice Array Conversions 2f8ad6137d2480e395c1fafaf8af180e.md](https://github.com/user-attachments/files/24987836/Slice.Array.Conversions.2f8ad6137d2480e395c1fafaf8af180e.md)
[Slice 2f4ad6137d2480c7ae14dac4fca96c88.md](https://github.com/user-attachments/files/24987834/Slice.2f4ad6137d2480c7ae14dac4fca96c88.md)
[Slice 2 2f4ad6137d24803fbcb2fda8f740ec7f.md](https://github.com/user-attachments/files/24987832/Slice.2.2f4ad6137d24803fbcb2fda8f740ec7f.md)
[Sıralı Sekilde Basmak İçin 2f9ad6137d24808dac6af82476631ffc.md](https://github.com/user-attachments/files/24987831/Sirali.Sekilde.Basmak.Icin.2f9ad6137d24808dac6af82476631ffc.md)
[Senaryolar 2f9ad6137d2480178bfac1c72e17dd69.md](https://github.com/user-attachments/files/24987829/Senaryolar.2f9ad6137d2480178bfac1c72e17dd69.md)
[Scope & Shadowing 2f4ad6137d248087b943d6dfa9ea1b4d.md](https://github.com/user-attachments/files/24987827/Scope.Shadowing.2f4ad6137d248087b943d6dfa9ea1b4d.md)
[Rune 2f4ad6137d2480778bacf22046d237a8.md](https://github.com/user-attachments/files/24987825/Rune.2f4ad6137d2480778bacf22046d237a8.md)
[Proje 2f8ad6137d248066ab00dbffb1038627.md](https://github.com/user-attachments/files/24987823/Proje.2f8ad6137d248066ab00dbffb1038627.md)
[Pointer 2f4ad6137d24803ebac3f1e1d0340569.md](https://github.com/user-attachments/files/24987821/Pointer.2f4ad6137d24803ebac3f1e1d0340569.md)
[Noktali Virgül 2f9ad6137d24805989a4fb1bc181d34f.md](https://github.com/user-attachments/files/24987819/Noktali.Virgul.2f9ad6137d24805989a4fb1bc181d34f.md)
[Maps 2f8ad6137d248064a5eac652be8c3321.md](https://github.com/user-attachments/files/24987817/Maps.2f8ad6137d248064a5eac652be8c3321.md)
[Loops 2f9ad6137d24801680e5e23e33dd1815.md](https://github.com/user-attachments/files/24987814/Loops.2f9ad6137d24801680e5e23e33dd1815.md)
[Iterating Slice , Map , String 2f9ad6137d24801c84bbc168553dd515.md](https://github.com/user-attachments/files/24987812/Iterating.Slice.Map.String.2f9ad6137d24801c84bbc168553dd515.md)
[Integers (Signed & Unsigned) 2f4ad6137d24801ca557e9e5a6edbff9.md](https://github.com/user-attachments/files/24987810/Integers.Signed.Unsigned.2f4ad6137d24801ca557e9e5a6edbff9.md)
[Here 2f4ad6137d2480b788cdc3a189ccff95.md](https://github.com/user-attachments/files/24987808/Here.2f4ad6137d2480b788cdc3a189ccff95.md)
[Go prefix suffix verbs 2f4ad6137d24807d9332cea24d782624.md](https://github.com/user-attachments/files/24987805/Go.prefix.suffix.verbs.2f4ad6137d24807d9332cea24d782624.md)
[go docs 2f4ad6137d2480b4b70ee672d2706a90.md](https://github.com/user-attachments/files/24987803/go.docs.2f4ad6137d2480b4b70ee672d2706a90.md)
[Go Commands 2f3ad6137d2480448893debe21c46cea.md](https://github.com/user-attachments/files/24987801/Go.Commands.2f3ad6137d2480448893debe21c46cea.md)
[Go Basisc 2f3ad6137d2480efaf86f5077e41d38c_all.csv](https://github.com/user-attachments/files/24987798/Go.Basisc.2f3ad6137d2480efaf86f5077e41d38c_all.csv)
[Go Basisc 2f3ad6137d2480efaf86f5077e41d38c.csv](https://github.com/user-attachments/files/24987796/Go.Basisc.2f3ad6137d2480efaf86f5077e41d38c.csv)
[Func 2f9ad6137d248090afc3cde7babd51bf.md](https://github.com/user-attachments/files/24987794/Func.2f9ad6137d248090afc3cde7babd51bf.md)
[for - for range string 2f8ad6137d2480189b6ec1a535a9ff57.md](https://github.com/user-attachments/files/24987792/for.-.for.range.string.2f8ad6137d2480189b6ec1a535a9ff57.md)
[Float Decimal 2f4ad6137d248063b054fc5ae9ad77ef.md](https://github.com/user-attachments/files/24987790/Float.Decimal.2f4ad6137d248063b054fc5ae9ad77ef.md)
[DataTypes 2f4ad6137d2480fc8674e55978586176.md](https://github.com/user-attachments/files/24987789/DataTypes.2f4ad6137d2480fc8674e55978586176.md)
[constant and iota 2f4ad6137d24805d849dfba79b849041.md](https://github.com/user-attachments/files/24987787/constant.and.iota.2f4ad6137d24805d849dfba79b849041.md)
[Conditionals 2f9ad6137d248056bce5d4b8a96ba34f.md](https://github.com/user-attachments/files/24987785/Conditionals.2f9ad6137d248056bce5d4b8a96ba34f.md)


- **`complex64`**: İki adet `float32`'den oluşur. (Toplam 8 Byte)
    - Gerçel Kısım (Real): `float32`
    - Sanal Kısım (Imaginary): `float32`
- **`complex128`** (Varsayılan): İki adet `float64`'ten oluşur. (Toplam 16 Byte)
    - Gerçel Kısım: `float64`
    - Sanal Kısım: `float64`

Bu bellek düzeni, Go'nun C/C++ kütüphaneleriyle (örneğin eski bir Fortran matematik kütüphanesiyle) veri alışverişi yaparken (cgo) sıfır maliyetle (zero-overhead) çalışmasını sağlar.

### 2. Tanımlama Yöntemleri

Go'da karmaşık sayıları tanımlamanın iki yolu vardır:

### A. Literal Syntax (Doğrudan Yazım)

Sanal kısmı belirtmek için `i` soneki kullanılır. Bu C#'ta yoktur, Go'ya özgüdür.

Go

```go
func main() {
    // Type Inference: complex128 olur
    c := 3 + 4i 
    
    // Açık Tanımlama
    var x complex64 = 1 + 2i
}
```

### B. `complex()` Yapıcı Fonksiyonu

Eğer elinde iki ayrı değişken varsa (örneğin veritabanından `lat` ve `long` gibi iki float okuduysan ve bunları vektörel hesaplama için birleştireceksen) bu fonksiyonu kullanırsın.

Go

```go
var r float64 = 3.5
var im float64 = 4.2

z := complex(r, im) // z artık complex128 tipindedir.
```

### 3. İşlemler ve `math/cmplx` Paketi

Standart `math` paketi sadece `float64` ile çalışır. Karmaşık sayılar için **`math/cmplx`** paketi ayrılmıştır. Bu ayrım, derleyicinin kullanılmayan matematik fonksiyonlarını binary dosyaya eklemesini engellemek (Binary Size Optimization) içindir.

Go

```go
import (
    "fmt"
    "math/cmplx"
)

func main() {
    z := 3 + 4i

    // Parçalara ayırma
    r := real(z) // 3
    i := imag(z) // 4

    // Mutlak değer (Modül) hesaplama |z|
    // sqrt(3^2 + 4^2) = 5
    abs := cmplx.Abs(z) 
    
    fmt.Println(abs)
}
```

### Backend Mühendisi İçin Ne Zaman Gerekli?

Standart bir e-ticaret veya CRUD uygulamasında `complex` tipleri görmezsin. Ancak şu senaryolarda karşına çıkabilir:

1. **Coğrafi Hesaplamalar:** Bazı harita servisleri rotasyon hesaplamalarını karmaşık düzlemde yapar.
2. **Sinyal İşleme:** Backend tarafında ses veya görüntü işleyen bir mikroservis yazarsan (FFT - Fast Fourier Transform algoritmaları) bu veri tipiyle çalışmak zorundasın.
3. **Kriptografi:** Bazı şifreleme algoritmalarının matematiksel tabanında kullanılır.

---

**Sıradaki Adım:**

Sayısal tiplerin tamamını (Integers, Floats, Complex) bitirdik. Şimdi Go'nun C# geliştiricilerini en çok şaşırtan, "String aslında String değildir" dedirten ve Türkçe karakter manipülasyonunda hayati önem taşıyan **Strings, Runes & Bytes** (Metin İşleme Mimarisi) konusuna geçelim mi?
# Conditionals

Created: January 31, 2026 4:26 AM

[Senaryolar](Senaryolar%202f9ad6137d2480178bfac1c72e17dd69.md)

Algoritmanın iskeleti **Karar Yapıları (Statements)** ve **Döngülerdir (Loops)**. Fonksiyonların içine girdiğimizde karmaşık mantıklar kurabilmen için bu yapıların Go'casını (özellikle C#'tan farklarını) çok iyi bilmen lazım.

Sıralamayı senin dediğin gibi güncelliyoruz:

1. **If / Else** (Parantez yok, özel syntax var)
2. **Switch** (Break yok, Fallthrough var)
3. **Loops** (While yok, sadece For var)

Hadi **If/Else** ve **Switch** ile başlayalım. C#'tan gelen alışkanlıklarını biraz törpüleyeceğiz.

---

### 1. IF / ELSE: "Parantezlere Elveda" 👋

[Noktali Virgül](Noktali%20Virg%C3%BCl%202f9ad6137d24805989a4fb1bc181d34f.md)

Go'da `if` yazarken parantez `()` kullanmayız. Ama süslü parantez `{}` **zorunludur**. (C#'taki tek satırlık `if (a) return;` yapısı Go'da yoktur).

### A. Temel Kullanım

Go

```go
age := 20

// YANLIŞ: if (age > 18) { ... }  <- Go bunu sevmez
// DOĞRU:
if age > 18 {
    fmt.Println("Reşitsin")
} else {
    fmt.Println("Değilsin")
}
```

### B. Go'nun Süper Gücü: "Statement Initialization" 🦸‍♂️

C#'ta önce değişkeni tanımlar, sonra `if` ile kontrol edersin.

Go'da **"Değişkeni if satırında tanımlayıp, aynı anda kontrol etme"** özelliği vardır. Bu değişken **sadece if bloğu içinde** yaşar. Bellek yönetimi (Scope) için harikadır.

**C# Versiyonu:**

C#

```go
var not = GetNot();
if (not > 50) { ... }
// 'not' değişkeni burada hala bellekte yaşıyor. Gereksiz.
```

**Go Versiyonu:**

```go
// Syntax: if <tanımlama>; <koşul> { ... }

if not := 55; not > 50 {
    fmt.Println("Geçtin, Notun:", not)
} else {
    fmt.Println("Kaldın, Notun:", not)
}
// 'not' değişkeni burada artık YOK. Bellekten silindi. Temiz iş.
```

Bunu en çok **Map kontrolü** (`comma-ok`) ve **Hata kontrolü** (`err != nil`) yaparken kullanacağız.

---

### 2. SWITCH: "Break Zincirlerini Kırmak" ⛓️

C#'ta her `case` sonuna `break` koymazsan kızar (Fallthrough hatası).

Go der ki: **"Switch yazıyorsan %99 ihtimalle tek bir case çalışsın istersin. Neden her seferinde break yazasın?"**

### A. Otomatik Break (Varsayılan)

Go

```go
day := "Cuma"

switch day {
case "Pazartesi":
    fmt.Println("Sendrom")
    // break; <-- GEREK YOK! Go otomatik çıkar.
case "Cuma":
    fmt.Println("Hayırlı Cumalar")
case "Cumartesi", "Pazar": // Virgülle birden fazla koşul
    fmt.Println("Tatil!")
default:
    fmt.Println("Çalışmaya devam")
}
```

### B. Fallthrough (C# Davranışını Zorlamak)

Eğer C#'taki gibi "Bu case çalışsın, **sonra alttaki case de çalışsın**" istiyorsan, `fallthrough` anahtar kelimesini kullanmalısın.

Go

```go
n := 5
switch { // DİKKAT: switch'in yanına değişken yazmadık! (Tagless Switch)
case n > 0:
    fmt.Println("Pozitif")
    fallthrough // "Durma, alttaki case'e de gir" (Koşula bakmaksızın!)
case n > 10:
    fmt.Println("Büyük sayı") // 5, 10'dan büyük değil ama YAZILIR!
}
// Çıktı:
// Pozitif
// Büyük sayı
```

*Not: `fallthrough` nadiren kullanılır ama bilmekte fayda var.*

---
# constant and iota

Created: January 26, 2026 4:21 AM

Süper — `const` ve `iota` Go’da sık karıştırılan ama mantığını kavradığın anda çok faydalı olan iki konu. Adım adım, açık ve C# karşılıklarıyla anlatıyorum — örneklerle beraber. 🇹🇷

# Temel: `const` nedir?

- `const` derleme zamanında (compile-time) sabit olan değerdir. Değiştirilemez, adresi alınamaz.
- Örnek:

```go
const Pi = 3.14
const Greeting string = "merhaba"

```

C# karşılığı: `const` (C#’ta da compile-time sabit).

Özellikler:

- Compile-time expression olmalı: sabitler başka sabitlerin ifadeleriyle hesaplanabilir (`const A = 2; const B = A+3`).
- Tipli veya tip belirtilmemiş olabilir (`const X = 5` vs `const X int = 5`). Tip belirtilmeyen sabitler *untyped constant* olur ve gerektiğinde başka tiplere uyum sağlar.

# Neden `const` kullanırsın?

- Magic number’ları kaldırmak (ör. `const MaxRetries = 5`)
- Performans/derleme zamanı optimizasyonu
- Belirgin API sözleşmesi

---

# `iota` nedir? — Kısa cevap

`iota`, bir `const` blokunda satır başına artan, sıfırdan başlayan bir sayaçtır. Her yeni `const (` bloğunda sıfırlanır.

```go
const (
    A = iota // 0
    B        // 1 (aynı satır ifadesi tekrar edilir)
    C        // 2
)

```

C# karşılığı: C# `enum` ile benzer amaca hizmet eder ama `iota` bir ifade aracıdır (enum değil). Go’da `iota` ile enum-benzeri sabitler kolayca oluşturulur.

---

# Davranış kuralları — Önemli noktalar

1. **`iota` her `const` bloğunda sıfırlanır**:

```go
const (
    A = iota // 0
)
const (
    B = iota // 0 again
)

```

1. **Eğer bir satırda ifade yoksa, önceki satırın ifadesi tekrar kullanılır (yani `iota` de o ifadeyi etkileyecek şekilde ilerler)**:

```go
const (
    A = iota       // 0
    B              // 1 -> B = iota (previous expression omitted so becomes `iota`)
    C = 100        // 100
    D              // 100 (repeats previous expression, iota is 3 here but expression is 100)
    E = iota       // 4 (iota continues increasing per line regardless)
)

```

1. **`iota` tam sayı (untyped integer) değeridir** ve genelde sabit ifadelerde kullanılır.

---

# Pratik örnekler (sık kullanılan kalıplar)

## 1) Simple enum-like

```go
const (
    StatusOK = iota   // 0
    StatusError       // 1
    StatusUnknown     // 2
)

```

C#:

```csharp
enum Status { OK = 0, Error = 1, Unknown = 2 }

```

## 2) Bit flag / powers of two (çok sık kullanılır)

```go
const (
    FlagRead = 1 << iota  // 1 << 0 == 1
    FlagWrite             // 1 << 1 == 2
    FlagExec              // 1 << 2 == 4
    FlagAdmin             // 1 << 3 == 8
)

```

C# (Flags enum):

```csharp
[Flags]
enum Flags { Read = 1 << 0, Write = 1 << 1, Exec = 1 << 2, Admin = 1 << 3 }

```

## 3) `_` ile atlayıp iota’yı kullanmak

```go
const (
    _ = iota             // 0 (atlandı)
    KB = 1 << (10 * iota) // 1 << 10
    MB                    // 1 << 20
    GB                    // 1 << 30
)

```

Açıklama: `_` satırı iota = 0 kullanır ama değeri atlanır; sonraki satırda iota = 1 olur.

## 4) Karmaşık ama güçlü: farklı ifadeler içinde iota

```go
const (
    A = iota * 2  // 0
    B             // 2
    C = iota + 5  // 7  (iota==2)
    D             // 8  (iota==3 -> expression iota+5 repeated)
)

```

---

# Typed vs Untyped constants (ince ama önemli fark)

```go
const X = 5       // untyped constant
var a int = X     // derleyici a'ya uyar (implicit conversion)
const Y int = 5   // typed constant
var b int = Y     // yine uyuyor, ama tip sabit

```

- Untyped sabitler gerektiğinde uygun tipe “dönüştürülür” (örneğin float, int, byte vb. bağlama göre).
- Bu esneklik `iota` ile yapılan sabit setlerinde faydalıdır.

---

# Sık yapılan hatalar / tuzaklar

- `iota` her satır için 1 artar; aynı ifadeyi tekrarlasan bile `iota` artmaya devam eder (yani satır sayısına bağlı).
- `iota` yalnızca `const` bloklarında çalışır; `var` içinde yok.
- `const` değerlerinin adresi alınamaz (`&Pi` hata verir).
- `const` runtime’da değiştirilemez (değişken değil).

---

# Hızlı referans tabloları

### iota davranışı özet

- Başlangıç: 0
- Her satır için +1
- Yeni `const (...)` bloğunda sıfırlanır
- İfade yoksa önceki ifade tekrar kullanılır

### Patternler

- Enum-like: `A = iota`
- Powers-of-two flags: `1 << iota`
- Byte units: `KB = 1 << (10*iota)`

---

# Tam bir kısa örnek — gerçek hayattan (dosya izinleri gibi)

```go
const (
    ModeRead = 1 << iota  // 1
    ModeWrite             // 2
    ModeExecute           // 4
    ModeAdmin             // 8
)

func hasWrite(m int) bool {
    return m & ModeWrite != 0
}

```

Bu örnekte bit mask kullanımıyla çok verimli permission kontrolü yaparsın.

---

# Sonuç — kafana yatması için

- `const` sabittir; compile-time.
- `iota` bir satır sayacıdır; `const` bloklarında enum-like sabitler üretir.
- `iota` ile enum, bitmask, byte-unit gibi yapıları kısa ve hatasız yazarsın.
- C# tarafında `enum` ve `[Flags] enum` ile benzer amaçları sağlarsın; Go’da `iota` + `const` daha esnek (ve fonksiyonel) bir yol sağlar.
# DataTypes

Created: January 26, 2026 5:11 AM

ASP.NET Core ve C# dünyasından gelen bir mühendis için Go'nun veri tipleri tanıdık görünse de, "Memory Layout" (Bellek Düzeni) ve "Strict Typing" (Katı Tipleme) konularında çok keskin mimari farklar barındırır.

Bu farkları, performans ve güvenliği nasıl etkiledikleri üzerinden inceleyelim.

### 1. Integers: Platform Bağımlılığına Dikkat

<aside>
✅

[Integers (Signed & Unsigned)](Integers%20(Signed%20&%20Unsigned)%202f4ad6137d24801ca557e9e5a6edbff9.md)

</aside>

C#'ta `int` her zaman 32-bit'tir (`Int32`). Go'da ise `int` tipi **mimari bağımlıdır**.

- **`int` ve `uint`**: Çalıştığın işletim sistemi 32-bit ise 32-bit, 64-bit ise 64-bit yer kaplar.
- **Explicit Types (`int8`, `int16`, `int32`, `int64`)**: Eğer veritabanındaki bir sütun kesinlikle `bigint` ise veya binary bir protokol yazıyorsan, asla düz `int` kullanma. `int64` kullan.

**Mimari Tavsiye:**

Genel döngüler ve array indeksleri için `int` kullan (işlemci için en optimize boyuttur). Ancak veri taşıma modellerinde (DTO) boyutunu bildiğin tipleri (`int32`, `int64`) tercih et.

### 2. Strings & Runes: En Büyük Ayrım

C# stringleri UTF-16 (karakter başına 2 byte) tabanlıdır. Go stringleri ise **UTF-8** tabanlıdır. Bu, bellek yönetimi ve karakter işleme mantığını tamamen değiştirir.

- **`string`**: Immutable (değiştirilemez) bir byte dizisidir (slice of bytes).
- **`byte`**: `uint8` için bir takma addır (alias). Ham veriyi (raw data) temsil eder.
- **`rune`**: `int32` için bir takma addır. Bir Unicode karakterini (Code Point) temsil eder. C#'taki `char` tipinin karşılığıdır ama 4 byte'tır.

**Neden Önemli?**

Bir string'in uzunluğunu (`len(s)`) aldığında, Go sana karakter sayısını değil, **byte sayısını** verir. Eğer string "Göll" ise;

- Byte sayısı: 5 (çünkü 'ö' UTF-8'de 2 byte yer kaplar).
- Rune sayısı: 4 (G, ö, l, l).
- **Best Practice:** Karakter saymak veya manipüle etmek için string'i `[]rune` tipine çevirmelisin.

### 3. Floats: Varsayılan Tercih

<aside>
✅

[Float Decimal ??](Float%20Decimal%202f4ad6137d248063b054fc5ae9ad77ef.md)

</aside>

C#'taki `double`, Go'da `float64`'e karşılık gelir.

- `float32`: Bellek çok darsa (örn: binlerce 3D vektör tutuyorsan) kullanılır.
- `float64`: Varsayılan ve önerilen tiptir. Matematiksel işlemlerde hassasiyet kaybını önler
- 

> 
> 

[Complex Numbers](Complex%20Numbers%202f4ad6137d248071aefadd1332926c16.md)

### 4. Zero Implicit Conversion (Gizli Dönüşüm Yok)

Go'nun mimari güvenliğinin temeli burasıdır. C#'ta `int` bir değeri `long` (int64) bir değişkene atayabilirsin (implicit casting). **Go buna asla izin vermez.**

Veri kaybı ihtimali olmasa bile dönüşümü açıkça (explicit) yazmalısın.

Go

```go
var a int32 = 10
var b int64 = int64(a) // 'int64()' yazmazsan derlenmez!
```

Bu katılık, "arka planda sessizce gerçekleşen" performans maliyetlerini ve taşma (overflow) hatalarını engeller.

---

### Özet Tablo: C# vs Go

| **C# Tipi** | **Go Karşılığı** | **Mimari Not** |
| --- | --- | --- |
| `int` (her zaman 32-bit) | `int32` | Go'da düz `int` platforma göre değişir. |
| `long` | `int64` | ID'ler ve zaman damgaları için standart. |
| `float` / `double` | `float32` / `float64` | `:=` ile tanımlarsan `float64` olur. |
| `char` | `rune` | `rune` aslında bir `int32`'dir. |
| `string` | `string` | Go stringleri UTF-8'dir ve byte tabanlıdır. |

<aside>
✅

[Bool](Bool%202f4ad6137d24807483dbd3a4cc17d629.md)

</aside>

<aside>
✅

[String](String%202f4ad6137d2480c0bb17d07f63db522b.md)

</aside>
# Float Decimal ??

ASP.NET Core tarafında `decimal` veri tipine alışkın bir backend geliştiricisi olarak, Go'nun **Floating Point** (Kayan Noktalı Sayı) yaklaşımı senin için kritik bir mimari karar noktasıdır.

C#'ta parasal işlemler için `decimal` (128-bit precision) kullanmaya alışıksın. Ancak Go'da yerleşik (built-in) bir **`decimal` tipi yoktur.** Bu eksiklik, finansal sistemler tasarlarken nasıl hareket etmen gerektiğini değiştirir.

İşte mimari açıdan bilmen gerekenler:

### 1. Tipler ve Varsayılanlar

Go'da ondalıklı sayılar IEEE-754 standardına göre çalışır.

- **`float32`**: C#'taki `float`. ~7 basamak hassasiyet. Genelde grafik kütüphaneleri (OpenGL/Vulkan) için kullanılır. Backend'de nadiren tercih edilir.
- **`float64`**: C#'taki `double`. ~15 basamak hassasiyet.
    - **Mimari Not:** Go'da `:=` operatörü ile ondalıklı sayı tanımlarsan varsayılan tip **`float64`** olur. Modern işlemcilerde `float64` işlemleri `float32` kadar hızlıdır, bu yüzden varsayılan olarak bunu kullan.

### 2. Büyük Mimari Tuzak: Finansal Hesaplamalar

Go'ya yeni başlayan C# geliştiricilerinin en sık yaptığı hata, para birimini `float64` ile tutmaktır.

Floating point sayıları, ikili sistemde (binary) bazı ondalıklı sayıları tam olarak temsil edemezler.

Go

`package main

import "fmt"

func main() {
    var a float64 = 0.1
    var b float64 = 0.2
    
    // Beklenti: 0.3
    // Gerçekleşen: 0.30000000000000004
    fmt.Println(a + b) 
    
    if a+b == 0.3 {
        fmt.Println("Eşit")
    } else {
        fmt.Println("Eşit DEĞİL!") // Burası çalışır
    }
}`

**Çözüm Stratejileri:**

Go'da `decimal` olmadığı için iki ana yol izlenir:

1. **Integer Math (Kuruş Hesabı):** Parayı en küçük birimine çevirip `int64` olarak tutmak. (Örn: 10.50 TL yerine 1050 Kuruş). Stripe ve birçok ödeme sistemi API'larında bu yöntemi kullanır.
2. **Üçüncü Parti Kütüphane:** `github.com/shopspring/decimal` paketi, Go ekosisteminin fiili `decimal` standardıdır. C#'taki `decimal` davranışını simüle eder.

### 3. Karşılaştırma Mantığı (Epsilon)

`float` değerleri asla `==` operatörü ile karşılaştırma. Yukarıdaki örnekte gördüğün gibi, çok küçük bir sapma (precision loss) eşitliği bozar.

Bunun yerine bir **"Epsilon"** (çok küçük bir tolerans değeri) belirleyip farkın bu değerden küçük olup olmadığına bakmalısın.

Go

`const epsilon = 1e-9

func floatEquals(a, b float64) bool {
    // Mutlak fark epsilon'dan küçükse eşit sayılır
    return (a-b) < epsilon && (b-a) < epsilon
}`

### 4. Tip Dönüşümü (Type Conversion)

Integers konusunda bahsettiğimiz katı kurallar burada da geçerlidir. `float32` ile `float64` arasında işlem yapamazsın, explicit (açık) dönüşüm gerekir.

Go

`var f32 float32 = 10.5
var f64 float64 = 20.5

// HATA: invalid operation: f32 + f64 (mismatched types)
// sum := f32 + f64 

// DOĞRU:
sum := float64(f32) + f64`

---

### Özet: Backend Checklist

1. **Varsayılan:** Genel matematiksel işlemler, oran hesaplamaları ve istatistikler için **`float64`** kullan.
2. **Yasak Bölge:** Para ve muhasebe hesaplarında asla `float` kullanma. Ya `int64` (kuruş) ya da `shopspring/decimal` kütüphanesini kullan.
3. **Karşılaştırma:** Unit testlerinde veya iş mantığında `float` karşılaştırırken mutlaka bir tolerans aralığı (epsilon) kullan.

**Sıradaki Adım:**

Temel sayısal tipleri (Integers ve Floats) bitirdik. Şimdi Go'nun C#'tan tamamen farklılaştığı ve yeni başlayanları en çok zorlayan konulardan biri olan **Strings, Runes & Bytes** üçlüsüne geçelim mi? Stringlerin neden aslında bir "Byte Dizisi" olduğunu anlamak, API'larda Türkçe karakter sorunlarını çözmek için şarttır.
# for - for range string

### Makaleden Alınacak Son Kritik Ders: "For Range vs For Loop"

Rob Pike makalede şunu vurguluyor:

- **Normal `for` döngüsü (`i++`):** String üzerinde **Byte Byte** gezer. (Bunu Egzersiz 3'te gördük, Türkçe karakter bozuluyordu).
- **`for range` döngüsü:** String üzerinde **Rune Rune (Karakter Karakter)** gezer. Go, range kullanırken otomatik olarak UTF-8 decoding yapar.

Gel bunu son bir "Mini Egzersiz" ile teyit edip String/Slice defterini sonsuza dek kapatalım.

---

### 🧪 Final Egzersizi: "Döngülerin Savaşı"

Aşağıdaki kodu çalıştır ve `for` ile `range` arasındaki o devasa farkı kendi gözlerinle gör.

Go

```go
package main

import "fmt"

func main() {
    metin := "Selam" 
    // 'S' (1 byte), 'e' (1 byte), 'l' (1 byte), 'a' (1 byte), 'm' (1 byte)
    // Toplam 5 byte. Sorun yok.

    metinTR := "Sealam" 
    // 'S' (1), 'e' (1), 'a' (1), 'l' (1), 'a' (1), 'm' (1)
    // Bekle... "Selam" yazdım ama bir harf "ğ" olsun.
    
    metin2 := "Dağ" 
    // D (1) + a (1) + ğ (2) = 4 Byte.

    fmt.Println("--- YÖNTEM 1: Normal For (Byte Byte) ---")
    // Bu yöntem C#'taki for(int i=0...) gibidir ama burada tehlikelidir.
    for i := 0; i < len(metin2); i++ {
        fmt.Printf("%d. Byte: %x (Karakter mi? %c)\n", i, metin2[i], metin2[i])
    }

    fmt.Println("\n--- YÖNTEM 2: Range For (Akıllı Gezinti) ---")
    // Go burada her turda "Bir sonraki karakter kaç byte?" diye bakar ve ona göre atlar.
    for index, harf := range metin2 {
        fmt.Printf("%d. İndekste Harf: %c (Unicode: %U)\n", index, harf, harf)
    }
}
```

**Çıktıda Göreceğin Şey:**

1. **Normal For:** "ğ" harfini ikiye bölecek ve iki tane saçma sapan karakter basacak.
2. **Range For:** "ğ" harfini tek parça ve düzgün basacak.

---

### 🎉 Tebrikler! Slice/Array/String/Pointer Bitti!

Artık;

1. Veritabanından gelen veriyi `slice` ile tutmayı (`append`),
2. Performans için `make` ile kapasite vermeyi,
3. Fonksiyonlara veri gönderirken kopyalanmasın diye `pointer` mantığını,
4. Metin parçalarken `rune` kullanman gerektiğini,
5. Döngü kurarken `range` kullanmanın güvenli olduğunu,
    
    BİLİYORSUN.
    

Bu, bir backend mühendisi için **çok sağlam bir temeldir.**
# Func

Created: February 1, 2026 2:15 AM

### The Master Plan: Fonksiyonların Anatomisi

**Bölüm 1: Temeller ve Go'ya Has Özellikler (Isınma)**

1. **Function Basics:** Sözdizimi, parametreler, C# ile farklar.
2. **Multiple Return Values:** (C#'taki `Tuple` veya `out` hamallığını bitiren özellik).
3. **Named Return Values (Naked Returns):** Değişkeni imzada tanımlama (Best Practice mi, tuzak mı?).
4. **Variadic Functions:** Sınırsız parametre alma (`...` operatörü - C# `params`).

**Bölüm 2: Fonksiyonel Programlama Yaklaşımı 🧠**
5.  **Anonymous Functions (Anonim Fonksiyonlar):** İsimsiz, tek atımlık fonksiyonlar (Lambda).
6.  **First-Class Functions:** Fonksiyonu değişkene atama, parametre olarak geçme (Callback mantığı).
7.  **Closures (Kapanışlar):** Hafıza yönetimi ve Scope'un sihirli dünyası. (State tutan fonksiyonlar).

**Bölüm 3: Nesne Yönelimli Yaklaşım (OOP'nin Go Hali) 🏗️**
8.  **Methods (Metotlar):** Fonksiyonu bir Struct'a yapıştırmak.
9.  **Pointer Receiver vs Value Receiver:** **EN ÖNEMLİ KONU.** (Veriyi kopyalayarak mı çalışacağız, adresiyle mi?).
10. **Defer:** `finally` bloğunun Go hali (Stack yapısı).

---

[Untitled](Untitled%202f9ad6137d2480898d30ca20c1ad4674.csv)
Name,Created,Tags
Go Commands,"January 26, 2026 2:49 AM",
Variables and Constans,"January 26, 2026 3:02 AM",
"var & := ","January 26, 2026 3:24 AM",
Zero Values,"January 26, 2026 3:33 AM",
Go prefix suffix verbs,"January 26, 2026 3:39 AM",
stringer ınterface (early),"January 26, 2026 4:15 AM",
constant and iota,"January 26, 2026 4:21 AM",
Scope & Shadowing,"January 26, 2026 4:43 AM",
"DataTypes ","January 26, 2026 5:11 AM",
Rune,"January 27, 2026 12:08 AM",
Type Conversion,"January 27, 2026 12:14 AM",
go docs,"January 27, 2026 1:11 AM",
Arrays,"January 27, 2026 1:27 AM",
Pointer,"January 27, 2026 1:36 AM",
Slice,"January 27, 2026 1:44 AM",
Slice 2,"January 27, 2026 2:03 AM",
Slice Array Conversions,"January 31, 2026 1:21 AM",
Strings Review,"January 31, 2026 12:26 AM",
Maps,"January 31, 2026 12:45 AM",
Struct,"January 31, 2026 3:20 AM",
Struct JSON,"January 31, 2026 3:58 AM",
Struct Tags,"January 31, 2026 4:01 AM",
Conditionals,"January 31, 2026 4:26 AM",
Comma-ok Idiom,"February 1, 2026 12:29 AM",
Loops,"February 1, 2026 12:46 AM",
"Iterating Slice , Map , String","February 1, 2026 1:55 AM",
Func,"February 1, 2026 2:15 AM",
Name,Created,Tags
Go Commands,"January 26, 2026 2:49 AM",
Variables and Constans,"January 26, 2026 3:02 AM",
"var & := ","January 26, 2026 3:24 AM",
Zero Values,"January 26, 2026 3:33 AM",
Go prefix suffix verbs,"January 26, 2026 3:39 AM",
stringer ınterface (early),"January 26, 2026 4:15 AM",
constant and iota,"January 26, 2026 4:21 AM",
Scope & Shadowing,"January 26, 2026 4:43 AM",
"DataTypes ","January 26, 2026 5:11 AM",
Rune,"January 27, 2026 12:08 AM",
Type Conversion,"January 27, 2026 12:14 AM",
go docs,"January 27, 2026 1:11 AM",
Arrays,"January 27, 2026 1:27 AM",
Pointer,"January 27, 2026 1:36 AM",
Slice,"January 27, 2026 1:44 AM",
Slice 2,"January 27, 2026 2:03 AM",
Strings Review,"January 31, 2026 12:26 AM",
Maps,"January 31, 2026 12:45 AM",
Slice Array Conversions,"January 31, 2026 1:21 AM",
Struct,"January 31, 2026 3:20 AM",
Struct JSON,"January 31, 2026 3:58 AM",
Struct Tags,"January 31, 2026 4:01 AM",
Conditionals,"January 31, 2026 4:26 AM",
Comma-ok Idiom,"February 1, 2026 12:29 AM",
Loops,"February 1, 2026 12:46 AM",
"Iterating Slice , Map , String","February 1, 2026 1:55 AM",
Func,"February 1, 2026 2:15 AM",
# Go Commands

Created: January 26, 2026 2:49 AM

Go dünyasında `go` komutu, sadece bir derleyici değil; paket yönetimi, test, formatlama ve build işlemlerini tek bir çatı altında toplayan devasa bir **toolchain**'dir. ASP.NET dünyasındaki `dotnet CLI` ile benzer bir görev üstlense de, Go'nun felsefesi gereği çok daha hafif ve "fikir sahibi" (opinionated) bir yapıdadır.

Mimari açıdan en sık kullanacağın temel komutları ve neyi, neden yaptıklarını inceleyelim:

Go komutlarını bir backend mühendisi perspektifiyle, mimari süreçlerine nasıl dokunduğunu detaylandırarak not alalım. Bu komutlar sadece kod çalıştırmak için değil, projenin yaşam döngüsünü (SDLC) yönetmek için tasarlanmıştır.

### 1. Proje Başlatma ve Bağımlılık Yönetimi (`go mod`)

Modern Go projelerinde "Dependency Management" `go mod` ile yapılır. ASP.NET'teki NuGet'in aksine, Go bağımlılıkları doğrudan kaynak kod seviyesinde (genelde GitHub URL'leri üzerinden) yönetir.

- **`go mod init <module_path>`**: Projenin kimliğini oluşturur. Genelde `github.com/kullaniciadi/projeadi` şeklinde isimlendirilir. Bu, projenin diğer projeler tarafından import edilebilir olmasını sağlar.
- **`go mod tidy`**: **Hayat kurtarıcıdır.** Kaynak kodunu tarar, `import` edilip `go.mod` dosyasında olmayanları indirir, `go.mod`'da olup kodda kullanılmayanları siler.
- **`go mod vendor`**: Tüm dış bağımlılıkların bir kopyasını projenin içine (`/vendor` klasörüne) indirir. Bu, internetin olmadığı bir ortamda veya CI/CD süreçlerinde "bağımlılığın silinmesi" riskine karşı projenin tamamen taşınabilir (self-contained) olmasını sağlar.

### 2. Derleme ve Çıktı Yönetimi (`go build`)

Go derleyicisi oldukça akıllıdır. Sadece değişen paketleri derler ve sonuçta **statik** bir binary üretir.

- **`go build -o <isim>`**: Çıktı dosyasının adını belirler.
- **`go build -ldflags="-s -w"`**: Binary boyutunu küçültmek için kullanılır. Debug sembollerini ve DWARF tablolarını siler. Microservice mimarilerinde Docker image boyutunu düşürmek için kritiktir.
- **Cross-Compilation**: Go'da başka işletim sistemine derleme yapmak için environment variable set edilir:
    - `GOOS=linux GOARCH=amd64 go build` (Linux 64-bit için derler)
    - `GOOS=windows GOARCH=amd64 go build` (.exe üretir)

### 3. Statik Analiz ve Kod Kalitesi

Kodun mimari açıdan "doğru" ve "temiz" olduğundan emin olmak için bu komutlar CI pipeline'larının vazgeçilmezidir.

- **`go fmt ./...`**: Projedeki tüm dosyaları standartlara göre formatlar. `./...` operatörü "bu dizin ve altındaki tüm paketler" demektir.
- **`go vet ./...`**: Derleyicinin yakalayamadığı ama mantıksal hata olabilecek durumları (örneğin yanış format specifier kullanımı, unreachable code) raporlar.
- **`go fix`**: Dilin yeni versiyonlarıyla gelen değişiklikleri eski koduna otomatik uygular.

### 4. Test ve Profiling (`go test`)

ASP.NET Core'daki xUnit/nUnit karmaşası yerine Go'da her şey tek komuttur.

- **`go test ./...`**: Tüm unit testleri çalıştırır.
- **`go test -v -race ./...`**: **En önemli komutlardan biridir.** `race` bayrağı "Race Detector"ı açar. Concurrency (goroutines) kullanırken aynı veriye iki farklı yerin aynı anda erişip erişmediğini (race condition) tespit eder.
- **`go test -bench=.`**: Benchmark testlerini çalıştırır. Kodun milisaniye bazında ne kadar hızlı olduğunu ve her işlemde kaç `alloc` (bellek tahsisi) yaptığını gösterir.

### 5. Dokümantasyon ve Yardım (`go doc`)

Go, kodun içindeki yorum satırlarını otomatik olarak dokümana çevirir.

- **`go doc fmt.Println`**: Terminalden çıkmadan bir fonksiyonun nasıl kullanıldığını ve imzasını görmeni sağlar.

---

### Özet Tablo: Geliştirici İş Akışı

| **Komut** | **Ne Zaman Kullanılır?** | **Mimari Katkısı** |
| --- | --- | --- |
| `go mod tidy` | Yeni bir paket eklediğinde veya sildiğinde. | Bağımlılık temizliği ve güvenliği. |
| `go fmt` | Commit atmadan hemen önce. | Ekip içi kod standart birliği. |
| `go test -race` | Concurrency içeren kod yazıldığında. | Runtime hatalarının (deadlock vb.) önlenmesi. |
| `go build -ldflags` | Production için image hazırlarken. | Lightweight (hafif) dağıtım birimleri. |

Bir sonraki adımda, Go'nun **"Variables & Constants"** yapısına geçelim mi? C#'tan farklı olarak "Zero Value" kavramı ve değişkenlerin bellekteki yerleşimi (stack vs heap) mimari kararlarını nasıl etkiliyor, oraya odaklanabiliriz. Ne dersin?

### 1. Çalıştırma ve Derleme (Build & Run)

Go, kodu doğrudan makine diline çevirir. Bu da uygulamanın çalışması için hedef makinede bir runtime veya VM (.NET Runtime gibi) kurulu olmasına gerek bırakmaz.

- **`go run`**: Kodu derler ve doğrudan çalıştırır. Binary çıktısını diskte tutmaz. Genelde geliştirme anında hızlı testler için kullanılır.
- **`go build`**: Kodu derler ve bir binary (çalıştırılabilir dosya) oluşturur. Go'nun en güçlü yanlarından biri **cross-compilation** yeteneğidir. Örneğin, macOS üzerinde Linux için bir binary üretebilirsin:Bash
    
    `GOOS=linux GOARCH=amd64 go build -o myapp`
    

### 2. Modül ve Bağımlılık Yönetimi (`go mod`)

Modern Go projelerinde (v1.11+) bağımlılıklar `go.mod` dosyası ile yönetilir. Bu, NuGet (`.csproj`) sistemine benzer ancak çok daha yalındır.

- **`go mod init <module-name>`**: Yeni bir modül başlatır. Projenin kök dizininde `go.mod` dosyası oluşturur.
- **`go mod tidy`**: Kodunu analiz eder; kullanılmayan paketleri temizler, eksik olanları indirir. Proje mimarisini temiz tutmak için en sık kullanacağın komuttur.

### 3. Kod Kalitesi ve Format (`go fmt` & `go vet`)

Go topluluğu, "kodun nasıl görünmesi gerektiği" konusundaki tartışmaları yıllar önce bitirmiştir.

- **`go fmt`**: Kodu standart Go formatına getirir. Girintiler, boşluklar ve stil her Go projesinde aynıdır. Bu, "mimarisel tutarlılık" sağlar.
- **`go vet`**: Kodundaki şüpheli yapıları (örneğin, printf format hataları veya unreachable code) tespit eden statik bir analiz aracıdır.

### 4. Test Yönetimi (`go test`)

Go'da test, dilin içine gömülüdür. Harici bir kütüphane kurmana gerek kalmaz.

- **`go test`**: `_test.go` ile biten dosyalardaki testleri çalıştırır.
- **`go test -bench=.`**: Yazdığın kodun performansını (benchmark) ölçer. Bir backend mühendisi olarak yazdığın algoritmanın ne kadar bellek tükettiğini görmek için kritiktir.

---

### Go Komutlarının Mimari Etkisi

Go toolchain'inin bu kadar birleşik olmasının sebebi, **"Zero Dependency"** ve **"Repeatable Builds"** felsefesidir. Kodunu bir başkasına verdiğinde, sadece `go build` diyerek seninle tıpatıp aynı çıktıyı alacağından emin olursun.

**Bir sonraki adımda**, Go'nun C# dünyasından oldukça farklı yaklaştığı **"Variables & Constants"** konusuna girip, Go'daki "Zero Value" (Sıfır Değer) kavramının mimari hataları (Null Reference gibi) nasıl önlediğini görmek ister misin?

Go komutlarını bir backend mühendisi perspektifiyle, mimari süreçlerine nasıl dokunduğunu detaylandırarak not alalım. Bu komutlar sadece kod çalıştırmak için değil, projenin yaşam döngüsünü (SDLC) yönetmek için tasarlanmıştır.

### 1. Proje Başlatma ve Bağımlılık Yönetimi (`go mod`)

Modern Go projelerinde "Dependency Management" `go mod` ile yapılır. ASP.NET'teki NuGet'in aksine, Go bağımlılıkları doğrudan kaynak kod seviyesinde (genelde GitHub URL'leri üzerinden) yönetir.

- **`go mod init <module_path>`**: Projenin kimliğini oluşturur. Genelde `github.com/kullaniciadi/projeadi` şeklinde isimlendirilir. Bu, projenin diğer projeler tarafından import edilebilir olmasını sağlar.
- **`go mod tidy`**: **Hayat kurtarıcıdır.** Kaynak kodunu tarar, `import` edilip `go.mod` dosyasında olmayanları indirir, `go.mod`'da olup kodda kullanılmayanları siler.
- **`go mod vendor`**: Tüm dış bağımlılıkların bir kopyasını projenin içine (`/vendor` klasörüne) indirir. Bu, internetin olmadığı bir ortamda veya CI/CD süreçlerinde "bağımlılığın silinmesi" riskine karşı projenin tamamen taşınabilir (self-contained) olmasını sağlar.

### 2. Derleme ve Çıktı Yönetimi (`go build`)

Go derleyicisi oldukça akıllıdır. Sadece değişen paketleri derler ve sonuçta **statik** bir binary üretir.

- **`go build -o <isim>`**: Çıktı dosyasının adını belirler.
- **`go build -ldflags="-s -w"`**: Binary boyutunu küçültmek için kullanılır. Debug sembollerini ve DWARF tablolarını siler. Microservice mimarilerinde Docker image boyutunu düşürmek için kritiktir.
- **Cross-Compilation**: Go'da başka işletim sistemine derleme yapmak için environment variable set edilir:
    - `GOOS=linux GOARCH=amd64 go build` (Linux 64-bit için derler)
    - `GOOS=windows GOARCH=amd64 go build` (.exe üretir)

### 3. Statik Analiz ve Kod Kalitesi

Kodun mimari açıdan "doğru" ve "temiz" olduğundan emin olmak için bu komutlar CI pipeline'larının vazgeçilmezidir.

- **`go fmt ./...`**: Projedeki tüm dosyaları standartlara göre formatlar. `./...` operatörü "bu dizin ve altındaki tüm paketler" demektir.
- **`go vet ./...`**: Derleyicinin yakalayamadığı ama mantıksal hata olabilecek durumları (örneğin yanış format specifier kullanımı, unreachable code) raporlar.
- **`go fix`**: Dilin yeni versiyonlarıyla gelen değişiklikleri eski koduna otomatik uygular.

### 4. Test ve Profiling (`go test`)

ASP.NET Core'daki xUnit/nUnit karmaşası yerine Go'da her şey tek komuttur.

- **`go test ./...`**: Tüm unit testleri çalıştırır.
- **`go test -v -race ./...`**: **En önemli komutlardan biridir.** `race` bayrağı "Race Detector"ı açar. Concurrency (goroutines) kullanırken aynı veriye iki farklı yerin aynı anda erişip erişmediğini (race condition) tespit eder.
- **`go test -bench=.`**: Benchmark testlerini çalıştırır. Kodun milisaniye bazında ne kadar hızlı olduğunu ve her işlemde kaç `alloc` (bellek tahsisi) yaptığını gösterir.

### 5. Dokümantasyon ve Yardım (`go doc`)

Go, kodun içindeki yorum satırlarını otomatik olarak dokümana çevirir.

- **`go doc fmt.Println`**: Terminalden çıkmadan bir fonksiyonun nasıl kullanıldığını ve imzasını görmeni sağlar.

---

### Özet Tablo: Geliştirici İş Akışı

| **Komut** | **Ne Zaman Kullanılır?** | **Mimari Katkısı** |
| --- | --- | --- |
| `go mod tidy` | Yeni bir paket eklediğinde veya sildiğinde. | Bağımlılık temizliği ve güvenliği. |
| `go fmt` | Commit atmadan hemen önce. | Ekip içi kod standart birliği. |
| `go test -race` | Concurrency içeren kod yazıldığında. | Runtime hatalarının (deadlock vb.) önlenmesi. |
| `go build -ldflags` | Production için image hazırlarken. | Lightweight (hafif) dağıtım birimleri. |

Bir sonraki adımda, Go'nun **"Variables & Constants"** yapısına geçelim mi? C#'tan farklı olarak "Zero Value" kavramı ve değişkenlerin bellekteki yerleşimi (stack vs heap) mimari kararlarını nasıl etkiliyor, oraya odaklanabiliriz. Ne dersin?
# go docs

Created: January 27, 2026 1:11 AM

Roadmap'teki bu başlık ("Commands & Docs"), ASP.NET Core geliştiricilerinin **"Oh be dünya varmış!"** dediği yerdir. Çünkü C#'taki o karmaşık XML yorum satırları (`/// <summary>`) ve Sandcastle/Swagger konfigürasyonlarıyla boğuşmak zorunda değilsin.

Go'da dokümantasyon, dilin ve mimarinin **doğal bir parçasıdır**, sonradan eklenen bir süs değildir.

Gel, bu işin mühendislik tarafını ve araçlarını senin dilinden (backend mimarisi) anlatalım.

---

### 1. Go Dokümantasyon Felsefesi: "Sadelik"

C#'ta bir fonksiyonu dokümante etmek için XML tagleri açarız (`<summary>`, `<param name="...">`, `<returns>`). Bu bazen koddan daha uzun sürer.

Go der ki: **"Bana XML verme, HTML verme, Markdown bile verme. Sadece düz yazı yaz."**

### Kural 1: Adıyla Başla

Bir fonksiyonun yorumu, mutlaka o fonksiyonun **adıyla** başlamalıdır. Bu, `go doc` aracının o satırı parse edebilmesi için şarttır.

Go

```go
// YANLIŞ:
// Bu fonksiyon kullanıcıyı veritabanından siler.
func DeleteUser(id int) { ... }

// DOĞRU (Best Practice):
// DeleteUser, verilen ID'ye sahip kullanıcıyı veritabanından kalıcı olarak siler.
// Eğer kullanıcı bulunamazsa hata döner.
func DeleteUser(id int) { ... }
```

### 2. Araçlar: Terminal vs Web (`go doc` vs `godoc`)

Go, dokümantasyonu kodun içinden çıkarıp sunmak için iki ana araç sunar.

### A. `go doc` (Terminaldeki Hız)

Kod yazarken bağlamdan kopmadan, IDE'den çıkmadan hızlıca bir şeyin ne işe yaradığını görmek içindir. Linux `man` sayfaları gibidir.

Terminaline şunu yazmayı dene:

Bash

`go doc fmt.Println`

Sana direkt `Println` fonksiyonunun imzasını ve açıklamasını basar. İnternete gitmene gerek kalmaz.

### B. `godoc` (Kendi Local Web Sunucun)

Bu araç, projenin kaynak kodlarını tarar ve **yerel bir web sunucusu** (localhost:8080) ayağa kaldırır. Sanki resmi Go sitesindeymişsin gibi kendi kodlarını web arayüzünde gezebilirsin.

- **Kurulum:** `go install golang.org/x/tools/cmd/godoc@latest`
- **Çalıştırma:** `godoc -http=:8080`

### 3. Testable Examples (Mimari Şaheseri) ✨

İşte C# dünyasında olmayan, Go'nun en bayıldığım özelliği: **Test Edilebilir Örnekler.**

Genelde dokümantasyona örnek kod koyarız ama zamanla kod değişir, dokümandaki örnek eski kalır ve çalışmaz (Outdated Documentation). Go bunu mimari olarak engeller.

Dokümantasyon içine yazdığın örnek kodlar, **aynı zamanda birer Unit Test'tir.** `go test` çalıştırdığında bu örnekler de derlenir ve çalıştırılır. Eğer örnek kodun çıktısı değişirse test patlar!

**Örnek (`example_test.go`):**

Go

```go
func ExampleReverse() {
    result := Reverse("Go")
    fmt.Println(result)
    
    // Output:
    // oG
}
```

**Mimari Büyü:**

1. Kodun sonunda `// Output:` yorumu var.
2. Go, bu fonksiyonu çalıştırır.
3. Ekrana basılan (Stdout) değer ile `Output:` altındaki değeri kıyaslar.
4. Eşleşmezse **TEST BAŞARISIZ** der.
5. Eşleşirse, bu kodu dokümantasyonda "Örnek Kullanım" olarak gösterir.

Böylece dokümantasyonun **asla** eskimez ve yalan söylemez.

### 4. `doc.go` ve Paket Dokümantasyonu

Büyük paketlerde (örneğin bir SDK yazdığında), `main.go` veya diğer dosyaların tepesine uzun uzun açıklama yazmak kodu kirletir.

Mimari olarak temiz kalmak için, projenin kök dizinine sadece yorum satırlarından oluşan bir `doc.go` dosyası konulur.

Go

```go
/*
Package myapi, kullanıcı yönetimi için gelişmiş fonksiyonlar sunar.

Bu paket thread-safe çalışır ve yüksek performanslıdır.
Kullanım detayları için README.md dosyasına bakınız.
*/
package myapi
```

Go, paketi derlerken bu dosyayı görür ve "Hah, burası bu paketin giriş dokümanı" der.

### 5. `pkg.go.dev` (Go'nun NuGet'i)

Sen kodunu GitHub'a `public` olarak pushladığın anda, Go'nun merkezi dokümantasyon sunucusu (`pkg.go.dev`) senin projenin lisansını (açıksa) görür ve dokümantasyon sayfanı **otomatik** oluşturur.

Ekstra bir işlem yapmana, ReadTheDocs konfigüre etmene gerek kalmaz. "Convention over Configuration" (Yapılandırma yerine Kural) mantığı burada da işler.

---

### Özet: Backend Mühendisi İçin Notlar

1. **Format Yok:** HTML yok, Markdown yok. Sadece paragraf başı yapmak için bir satır boşluk bırak. Kod bloğu göstermek için satır başından girinti (indent) yap.
2. **Public/Private:** Sadece Büyük Harfle Başlayan (Exported) fonksiyonlar ve tipler dokümantasyonda görünür. Küçük harfliler (private) görünmez.
3. **Deprecated:** Bir metodu kullanımdan kaldırmak istersen başına `// Deprecated: Use NewFunction instead.` yazman yeterlidir. IDE bunu tanır ve üzerini çizer.

**Sıradaki Adım:**

Dokümantasyon kültürünü de cebimize koyduk. Şimdi Go'nun en temel yapı taşı olan ve C# projelerindeki `Solution/Project` yapısının karşılığı olan **Packages & Modules** konusuna geçelim mi? `go.mod` dosyası nedir, paketler birbirini nasıl görür, onu halledelim.
# Go prefix suffix verbs

Created: January 26, 2026 3:39 AM

Bu konuya girmeden önce, mekanizmanın arka planını tam olarak kavramak için şu mimari başlıkları zihninde oturtman gerekir:

### Bu Konuyu Anlamak İçin Gereken Altyapı

- **Variadic Functions:** Bir fonksiyonun (Print gibi) nasıl sınırsız sayıda ve farklı tipte parametre alabildiğini anlamak (`...interface{}`).
- **Reflection (Yansıma):** Derleyicinin çalışma zamanında (runtime) bir değişkenin tipini nasıl anlayıp ona göre formatladığı.
- **Interfaces (Stringer Interface):** C#'taki `ToString()` override etme mantığının Go'daki karşılığı olan `String()` metodunun `fmt` paketiyle ilişkisi.
- **I/O Streams (io.Writer):** Çıktının sadece konsola değil, dosyaya veya bir HTTP response body'sine nasıl yönlendirildiği.

---

Go'da `fmt` (format) paketi, C#'taki `Console.WriteLine` ve `String.Format` yapılarının çok daha gelişmiş ve granüler halidir. Go'da fonksiyonlar isimlendirme standartlarına göre gruplanmıştır.

### 1. Temel Print Fonksiyonları (Nereye Yazıyor?)

Go'da print fonksiyonları "nereye" yazdıklarına göre 3 ana aileye ayrılır.

| **Önek (Prefix)** | **Fonksiyonlar** | **Nereye Yazar?** | **C# Karşılığı** |
| --- | --- | --- | --- |
| **(Yok)** | `Print`, `Println`, `Printf` | **Standard Output (Konsol)** | `Console.Write/WriteLine` |
| **S** (String) | `Sprint`, `Sprintln`, `Sprintf` | **String Deöndürür** (Yazmaz) | `String.Format` / `Interpolation` |
| **F** (File) | `Fprint`, `Fprintln`, `Fprintf` | **io.Writer** (Dosya, Ağ akışı vb.) | `StreamWriter.Write` |
- **Mimari Not:** Backend geliştirirken `Fprintf` çok kritiktir. Bir HTTP handler yazarken `http.ResponseWriter` bir `io.Writer` olduğu için, JSON veya text cevabını direkt `Fprintf` ile memory allocation maliyetini düşürerek client'a basabilirsin.

---

### 2. Sonek (Suffix) Mantığı (Nasıl Yazıyor?)

Fonksiyonların sonundaki ekler, veriyi nasıl işleyeceğini belirler.

- **`Print`:** Argümanları yan yana basar. Otomatik satır atlamaz (`\n` yoktur).
- **`Println`:** Argümanlar arasına boşluk koyar ve sonuna mutlaka yeni satır (`\n`) ekler.
- **`Printf` (Print Formatted):** En güçlüsüdür. Senin bahsettiğin `%` işaretli "verb"leri (fiilleri) kullanarak şablon doldurmanı sağlar.

---

### 3. Formatting Verbs (Biçimlendirme Karakterleri)

Go, C#'taki `{0}`, `{1}` gibi index bazlı formatlama yerine, **Tip Bazlı (Type-Based)** formatlama kullanır. İşte bir backend mühendisinin bilmesi gereken en kritik olanlar:

### A. Genel Amaçlı (En Sık Kullanılanlar)

| **Verb** | **Açıklama** | **Örnek** |
| --- | --- | --- |
| **`%v`** | **Value.** En önemli verb'dür. Tipi ne olursa olsun (int, string, struct) varsayılan formatta basar. | `Printf("%v", user)` |
| **`%+v`** | Struct'lar için **field isimlerini** de gösterir. Debug yaparken hayat kurtarır. | `{Name:Ali ID:1}` |
| **`%#v`** | Go syntax formatında basar. Verinin tam olarak nasıl tanımlandığını (tipi dahil) gösterir. | `main.User{Name:"Ali", ID:1}` |
| **`%T`** | Değişkenin **Tipini (Type)** basar. | `int`, `*string`, `main.User` |

### B. Spesifik Tipler İçin

| **Verb** | **Kullanım Alanı** | **Açıklama** |
| --- | --- | --- |
| **`%d`** | Integer (Digit) | Tam sayılar için (Base 10). |
| **`%s`** | String | Metin ifadeleri veya byte slice'ları string olarak basar. |
| **`%f`** | Float | Ondalıklı sayılar (Örn: `%.2f` ile virgülden sonra 2 hane). |
| **`%p`** | Pointer | Bellek adresini hex formatında basar (Örn: `0xc000010200`). |
| **`%t`** | Boolean | `true` veya `false` basar. |

### Örnek Senaryo: Neyi Neden Kullanırız?

Bir API yanıtı döndüğünüzü veya log bastığınızı düşünelim.

Go

```go
type User struct {
    ID   int
    Name string
}

user := User{ID: 101, Name: "Gopher"}

// 1. Basit Loglama (Console.WriteLine)
fmt.Println("Kullanıcı:", user) 
// Çıktı: Kullanıcı: {101 Gopher}

// 2. Detaylı Debugging (Field adlarını görmek istiyorsun)
fmt.Printf("Detay: %+v\n", user)
// Çıktı: Detay: {ID:101 Name:Gopher}

// 3. String Oluşturma (Veritabanına kaydetmek veya API response için)
msg := fmt.Sprintf("Hata: %s kullanıcısı (ID: %d) bulunamadı.", user.Name, user.ID)
// msg değişkeni artık formatlanmış stringi tutar.
```

### Özet ve Best Practice

- **Debug yaparken:** Verinin yapısını tam görmek için daima `Printf` içinde `%+v` veya `%#v` kullan.
- **Basit mesajlar için:** `Println` yeterlidir.
- **HTTP Response/Dosya Yazma:** String birleştirip `Write` etmek yerine direkt `Fprintf` kullanarak stream'e yaz (Performans için).

C# geçmişinle karşılaştırırsan; `%v` aslında `object` alıp `ToString()` çağırmak gibidir ama Go bunu reflection ile çok daha akıllıca yönetir.

Sıradaki adımda, bu çıktıların özelleştirilmesini sağlayan **Interfaces (Stringer)** konusuna mı geçelim, yoksa bellekte verinin nasıl tutulduğunu daha iyi anlamak için **Pointers** konusuna mı?
# Here

Go'nun değişken ve sabit yönetimini, bir mühendis gözüyle mimari temellere oturtarak inceleyelim. Bu yaklaşım, C#'taki `var` kullanımından veya genel nesne yönelimli dillerden bazı radikal farklar içerir.

---

## 1. Memory Allocation (Stack vs. Heap) ve Scope

Go, performansı optimize etmek için **Escape Analysis** (Kaçış Analizi) algoritmasını kullanır. C#'ta bir `struct` stack'te, bir `class` heap'te yer alır gibi kesin kurallar varken; Go'da bir değişkenin nerede yaşayacağına derleyici karar verir.

- **Stack:** Eğer değişkenin yaşam döngüsü sadece o fonksiyon içindeyse, stack'te kalır. Çok hızlıdır.
- **Heap:** Eğer bir değişken fonksiyon dışına sızıyorsa (örneğin bir pointer döndürülüyorsa), değişken heap'e "kaçış" yapar.
- **Mimari Yaklaşım:** Go'da değişken tanımlarken "Bu değişkenin pointer'ına (referansına) gerçekten ihtiyacım var mı?" diye sormak gerekir. Gereksiz pointer kullanımı, heap üzerinde yük oluşturarak GC (Garbage Collector) performansını etkiler.

---

## 2. Static Typing ve Kısa Değişken Tanımlama (`:=`)

Go, tip güvenliğini (Type Safety) elden bırakmadan yazım kolaylığı sağlar.

- **Tip Belirleme:** `var name string = "Go"` şeklinde açıkça yazılabilir. Ancak fonksiyon içinde genellikle `name := "Go"` (short declaration) tercih edilir.
- **C# Farkı:** C#'taki `var` çalışma mantığına benzese de Go'da bu operatör sadece fonksiyon blokları içinde geçerlidir. Paket seviyesinde (global) `var` veya `const` kullanmak zorunludur.
- **Strictness:** Go'da tanımlanan ama kullanılmayan her değişken bir derleme hatasıdır (Compile Error). Bu, kodun temiz kalmasını ve bellek israfının önlenmesini sağlayan mimari bir disiplindir.

---

## 3. Zero Value (Sıfır Değer) Mekanizması

Go'da bir değişkeni başlattığınızda (initialize), ona bir değer atamazsanız sistem otomatik olarak **Zero Value** atar. `null` veya `undefined` hatalarının büyük bir kısmını mimari seviyede bu engeller.

| Veri Tipi | Zero Value |
| --- | --- |
| `int`, `float` | `0` |
| `string` | `""` (Boş string) |
| `bool` | `false` |
| `pointer`, `slice`, `map` | `nil` |

E-Tablolar'a aktar

**Neyi Neden Yapıyoruz?**
Bu mekanizma sayesinde programın durumu her zaman öngörülebilir kalır. Bir döngüye veya işleme girmeden önce değişkenin "çöp değer" taşıma ihtimali ortadan kalkar.

---

## 4. Sabit (Constant) Yönetimi

Go'da `const` sadece derleme zamanında (compile-time) bilinen değerler içindir.

- **Typed vs Untyped:** Go'da sabitler "tipi belirlenmemiş" (untyped) olabilir. Örneğin, `const Pi = 3.14` dediğinizde Pi, hem `float32` hem de `float64` ile işleme girebilir. Bu, C#'taki katı cast işlemlerinden daha esnek bir yapı sunar.
- **iota:** Go'da sabit grupları tanımlarken `iota` anahtar kelimesi kullanılır. Bu, otomatik artan sayılar oluşturarak `enum` benzeri yapıları çok daha temiz kurmanızı sağlar.
# Integers (Signed & Unsigned)

C# ve ASP.NET Core geliştiricisi olarak `int`, `long`, `byte` gibi tiplere hakimsin. Ancak Go'da Integer (Tamsayı) yönetimi, mimari açıdan **"Platform Bağımlılığı"** ve **"Tip Katılığı"** nedeniyle biraz daha fazla dikkat gerektirir.

İşte Go'daki tamsayıların anatomisi ve backend geliştirirken dikkat etmen gereken mimari kararlar:

### 1. İşaretli (Signed) vs İşaretsiz (Unsigned)

Temel fark bit seviyesindeki kullanımdır.

- **Signed (`int`):** Bir bit işaret (+/-) için ayrılır. Hem negatif hem pozitif değer alır.
- **Unsigned (`uint`):** Tüm bitler sayısal değer (magnitude) için kullanılır. Sadece 0 ve pozitif değer alır.

**Mimari Karar:**

ASP.NET'te `uint` nadiren kullanılır. Ancak Go'da;

- **`byte` (uint8):** Binary veri (dosya okuma, ağ paketleri) işlerken standarttır.
- **Bellek Protokolleri:** TCP başlıkları, dosya formatları gibi yerlerde kesin boyut gerektiğinde `uint` türevleri zorunludur.

> Kritik Uyarı: İş mantığında (Business Logic) "Yaş negatif olamaz, o yüzden uint kullanayım" deme! uint ile çıkarma işlemi yaparsan (örn: 2 - 5) sayı negatife düşmez, underflow (başa sarma) yaparak devasa bir pozitif sayıya dönüşür. Bu da ciddi mantık hatalarına yol açar. Genel mantık için int kullanmak daha güvenlidir.
> 

### 2. "Architecture Dependent" Tipler (`int` ve `uint`)

C# dünyasında `int` her zaman 32-bit'tir (`Int32`). Go'da ise `int` tipi, kodun çalıştığı işlemci mimarisine göre şekil alır.

- **32-bit Sistemde:** `int` $\rightarrow$ 32 bit (4 byte)
- **64-bit Sistemde:** `int` $\rightarrow$ 64 bit (8 byte)

**Bu Neden Önemli?**

Veritabanı modellerinde (Struct) veya dış servislerle (API) konuşurken asla düz `int` kullanmamalısın.

- Eğer veritabanındaki sütun `BIGINT` (8 byte) ise ve sen kodda `int` kullanırsan, kodunu 32-bit bir makinede (veya Docker konteynerinde) derlediğinde veri kaybı (overflow) yaşarsın.
- **Kural:** DTO'larda ve DB modellerinde her zaman açık boyut belirt: `int32`, `int64`.
- **Kural:** Döngülerde (`for i := 0...`) ve array indekslerinde düz `int` kullan (İşlemci için en optimize olan budur).

### 3. Sınırlar ve Boyutlar

| **Tip** | **Bit Boyutu** | **C# Karşılığı** | **Kullanım Alanı** |
| --- | --- | --- | --- |
| **`int8`** | 8-bit | `sbyte` | Çok düşük bellekli sistemler. |
| **`int16`** | 16-bit | `short` | Ses işleme vb. |
| **`int32`** | 32-bit | `int` | Standart sayılar. |
| **`int64`** | 64-bit | `long` | **DB ID'leri, Timestamp'ler.** |
| **`uint8`** | 8-bit | `byte` | **Raw Data, Stream işlemleri.** (`byte` alias'ıdır) |

### 4. Overflow (Taşma) Davranışı

Go, çalışma zamanında (runtime) tamsayı taşmalarını (overflow) varsayılan olarak **kontrol etmez** (panic vermez). Sayı sessizce başa sarar (Wraparound).

Go

```go
var x uint8 = 255
x = x + 1 
// x şu an 0 oldu! Hata fırlatmadı.
```

Bu davranış, performans için bilerek tercih edilmiştir ancak finansal hesaplamalarda risklidir.

---

### Özet: Backend Geliştiricisi İçin Kontrol Listesi

1. Veritabanı ID'leri için her zaman **`int64`** kullan.
2. Dosya, Resim, Network işlemleri için **`[]byte`** (slice of byte) kullan.
3. Genel döngü sayaçları için **`int`** kullan.
4. Asla "negatif olamaz" kuralını sağlamak için `uint` kullanma; bunu validasyon katmanında yap.
# Iterating Slice , Map , String

Created: February 1, 2026 1:55 AM

# 📚 Go İterasyon Rehberi: Array, Slice ve Map

Go'da `foreach`, `while` veya `do-while` yoktur. Elimizde sadece **`for`** ve onun en iyi arkadaşı **`range`** vardır.

Bu notlar; verinin bellekte nasıl gezildiğini, performans tuzaklarını ve "Best Practice" yöntemlerini içerir.

---

## 1. Array ve Slice İterasyonu (`[]T`)

Array (Sabit) ve Slice (Dinamik) üzerinde dönmek mantık olarak aynıdır. Go, `range` kullanıldığında sana her turda iki şey verir:

1. **Index (`i`):** Elemanın sırası.
2. **Value (`v`):** Elemanın **KOPYASI**.

### A. Temel Kullanım (Index ve Value)

Hem sıra numarasına hem de veriye ihtiyacın varsa:

```go
sayilar := []int{10, 20, 30}

// i: index, v: value
for i, v := range sayilar {
    // v, o anki sayının kopyasıdır.
    fmt.Printf("Index: %d, Değer: %d\n", i, v)
}
```

### B. Sadece Değer Lazımsa (`_` Kullanımı)

C#'taki `foreach (var item in list)` karşılığıdır. Index'i `_` (Blank Identifier) ile çöpe atarız.

```go
isimler := []string{"Ali", "Veli"}

for _, isim := range isimler {
    fmt.Println("Kullanıcı:", isim)
}
```

### C. Sadece Index Lazımsa

Bazen verinin kopyasını almak pahalıdır (büyük struct'lar) veya sadece kaçıncı sırada olduğumuz önemlidir.

```go
// İkinci değişkeni hiç yazmazsan, Go sadece index verir.
for i := range isimler {
    fmt.Println("Sıra No:", i)
    // Değere ihtiyacım olursa: isimler[i] diyerek alırım.
}
```

### ⚠️ Kritik Detay: "Value Copy" Tuzağı

Go'da `range` ile gelen `v` (value) değişkeni, orijinal verinin **bir kopyasıdır.**

**Hatalı Yöntem:**

```go
for _, v := range sayilar {
    v = v + 5 // Sadece 'v' kopyası değişir. Orijinal dizi (sayilar) AYNI KALIR.
}
```

**Doğru Yöntem (Orijinali Değiştirmek İçin):**

```go
for i := range sayilar {
    sayilar[i] = sayilar[i] + 5 // Direkt adrese müdahale ediyoruz.
}
```

---

## 2. Map İterasyonu (`map[K]V`)

Map üzerinde dönmek, Slice'a benzer ama çok önemli bir fark vardır: **Sırasızlık (Randomness).**

### A. Temel Kullanım (Key ve Value)

C#'taki `foreach (KeyValuePair<K,V> kvp in dictionary)` karşılığıdır.

```go
stoklar := map[string]int{
    "Laptop": 5,
    "Mouse":  100,
}

for urun, adet := range stoklar {
    fmt.Printf("Ürün: %s, Stok: %d\n", urun, adet)
}
```

### B. Sadece Key Lazımsa (Optimize Yöntem) 🚀

Genelde Map'in anahtarlarını alıp başka bir işlem yapmak için kullanılır. Value okunmadığı için performanslıdır.

```go
// Sadece Key döner
for urun := range stoklar {
    fmt.Println("Stokta kayıtlı ürün:", urun)
}
```

### ⚠️ Kritik Detay: Kaos Teorisi

Go, Map iterasyonunun başlangıç noktasını her çalıştırmada **rastgele** seçer.

- **Sebep:** Geliştiriciler yanlışlıkla sıraya güvenip kod yazmasın diye.
- **Sonuç:** Çıktı bir `A, B` olur, bir dahakine `B, A` olur.
- **Çözüm:** Sıralı çıktı istiyorsan, önce Key'leri bir Slice'a atıp `sort` paketiyle sıralamalısın.

[Sıralı Sekilde Basmak İçin](S%C4%B1ral%C4%B1%20Sekilde%20Basmak%20%C4%B0%C3%A7in%202f9ad6137d24808dac6af82476631ffc.md)

---

## 3. String İterasyonu (Gizli Tehlike 🐍)

String üzerinde `range` ile dönmek, karakterleri (rune) tek tek gezmek demektir. Bu, Türkçe karakterler için hayati önem taşır.

### A. Range ile (Doğru Yöntem ✅)

Go, `range` kullanıldığında UTF-8 karakterleri otomatik çözer.

Go

```go
mesaj := "Göl" // 'ö' harfi 2 byte yer kaplar.

for i, harf := range mesaj {
    // %c : Karakter (Rune) olarak basar.
    // %d : Byte pozisyonunu basar.
    fmt.Printf("%d. index: %c\n", i, harf)
}
// Çıktı:
// 0. index: G
// 1. index: ö  (Dikkat: 2 byte olduğu için bir sonraki index 2 değil 3 olacak)
// 3. index: l
```

### B. Standart For ile (Riskli Yöntem ❌)

Eğer `len(mesaj)` ve `mesaj[i]` kullanırsan, karakterleri değil **byte'ları** gezersin. Türkçe karakterler bozulur.

```go
for i := 0; i < len(mesaj); i++ {
    fmt.Printf("%c", mesaj[i]) 
}
// Çıktı: GÃ¶l (Gibi bozuk karakterler çıkar)
```

---

## 🎓 Özet Tablo (Cheat Sheet)

| **Veri Yapısı** | **Kod Şablonu** | **Açıklama** |
| --- | --- | --- |
| **Slice / Array** | `for i, v := range list` | Index ve Value (Kopya) döner. |
| **Slice (Index)** | `for i := range list` | Sadece Index döner. Orijinal veriyi değiştirmek için ideal. |
| **Slice (Value)** | `for _, v := range list` | Sadece Değer döner. Index önemsizse kullanılır. |
| **Map** | `for k, v := range m` | Key ve Value döner. **Sıra Rastgeledir!** |
| **Map (Key)** | `for k := range m` | Sadece Key döner. |
| **String** | `for i, r := range str` | Byte indexi ve Rune (Karakter) döner. Türkçe dostudur. |

Bu notlar, Go'daki döngülerin %99'unu kapsar. Bunları defterine geçirdiğinde, "Bu neden böyle çalışıyor?" sorusunun cevabı hep elinin altında olacak. 🚀
# Loops

Created: February 1, 2026 12:46 AM

Harika. Go'nun en sevdiğim yanı olan "Sadelikteki Güç" felsefesinin zirve yaptığı yere geldik: **Döngüler (Loops).**

C# geliştiricisi olarak kafanda `for`, `foreach`, `while`, `do-while` gibi 4 farklı yapı var.

Go mühendisleri masaya vurmuş ve demiş ki: **"Hepsini çöpe atın. Sadece `for` kalacak."**

Evet, Go'da sadece **`for`** vardır. Ama bu `for`, şekil değiştirebilen bir İsviçre Çakısı gibidir.

Gel, bu çakının tüm bıçaklarını tek tek açalım ve **bellekte neler döndüğünü** (under the hood) inceleyelim.

---

### 1. Anatomi: "Üç Silahşörler" ⚔️

Klasik `for` döngüsü 3 parçadan oluşur. Bu yapı C# ile aynıdır, sadece parantez yoktur.

```go
// 1. INIT (Başlatma): Sadece döngü başlarken 1 kez çalışır. (i := 0)
// 2. CONDITION (Koşul): Her turdan önce kontrol edilir. (i < 5)
// 3. POST (İşlem): Her tur bittikten sonra çalışır. (i++)

for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**Mühendislik Detayı (Scope):**

Burada tanımlanan `i` değişkeni, sadece döngü süslü parantezleri `{}` içinde yaşar. Döngü bittiği an Garbage Collector onu yok eder.

---

### 2. While Modu (Sadece Koşul) ⏳

Go'da `while` anahtar kelimesi yoktur. `for`'un sağındaki ve solundaki (init ve post) kısımları atarsan, o artık bir `while` döngüsüdür.

```go
dbConnectionActive := true
retryCount := 0

// C#: while(dbConnectionActive)
// Go:
for dbConnectionActive {
    retryCount++
    if retryCount > 3 {
        dbConnectionActive = false // Döngüyü kıracak koşul
    }
    fmt.Println("Bağlantı deneniyor...")
}
```

---

### 3. Sonsuz Döngü Modu (Server Pattern) ♾️

Eğer koşul kısmını da silersen? Go bunu `true` kabul eder ve sonsuza kadar döner.

Backend servislerinde (RabbitMQ dinleyicisi, Web Sunucusu) bu kullanılır.

```go
// C#: while(true) { }
// Go:
for {
    // Mesaj kuyruğunu dinle...
    // İşle...
    // Asla durma.
    
    // Çıkış için mecburen 'break' veya 'return' kullanmalısın.
}
```

---

### 4. Range Modu (Go'nun `foreach`'i) 🌟

En çok kullanacağın, en güçlü ama **en çok tuzak barındıran** mod budur.

Slice, Map, Array veya String üzerinde gezmek için kullanılır.

### A. Slice/Array Üzerinde

`range` sana her turda iki değer verir:

1. **Index:** Sıra numarası.
2. **Value:** O elemanın **KOPYASI**.

```go
sayilar := []int{10, 20, 30}

for i, v := range sayilar {
    fmt.Printf("Index: %d, Değer: %d\n", i, v)
}
```

⚠️ **Performans Tuzağı (Deep Dive):**

Eğer `v` (Value) çok büyük bir struct ise (örneğin içinde resim datası olan bir struct), Go her turda o koca veriyi kopyalar (Memory Copy). Bu performansı öldürür.

**Çözüm:** Büyük structlarda sadece index ile dönmek veya pointer slice kullanmak daha iyidir.

### B. Değişkeni Yoksayma (`_`)

Bazen index lazım değildir, bazen de değer.

Go

```go
// Sadece değer lazım (Index'i çöpe at)
for _, v := range sayilar {
    fmt.Println(v)
}

// Sadece index lazım (Değeri kopyalama maliyetinden kurtarır!)
for i := range sayilar {
    fmt.Println("Sıra:", i)
}
```

### C. Map Üzerinde (Random Order) 🎲

Bunu Map konusunda konuşmuştuk. `range` map üzerinde gezerken sırayı garanti etmez. Her çalıştırdığında farklı sırada gelir.

```go
notlar := map[string]int{"Ali": 50, "Veli": 100}

for k, v := range notlar {
    fmt.Printf("Key: %s, Value: %d\n", k, v)
}
```

---

### 5. Akış Kontrolü: Break, Continue ve LABELS 🏷️

`break` (çık) ve `continue` (pas geç) C# ile aynıdır.

Ancak Go'da **Nested Loop (İç İçe Döngü)** kırmak için muazzam bir özellik vardır: **Labels (Etiketler).**

Normalde `break` sadece içinde bulunduğu en yakın döngüyü kırar. Etiket ile "En dıştaki döngüyü patlat" diyebilirsin.

```go
func main() {
    fmt.Println("İşlem Başlıyor...")

    // Dış döngüye bir isim verdik: 'DisDongu'
DisDongu:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            
            // Kritik bir hata oldu, komple çıkmamız lazım!
            if i == 2 && j == 2 {
                fmt.Println("⛔ Kritik Hata! Tüm işlemler durduruluyor.")
                break DisDongu // Sadece j'yi değil, i'yi de kırar!
            }
            
            fmt.Printf("i:%d j:%d işleniyor\n", i, j)
        }
    }
    
    fmt.Println("Döngüden tamamen çıkıldı.")
}
```

Bu özellik, matris işlemleri veya karmaşık veri analizlerinde "Flag" (bayrak) değişkeni tutma hamallığını bitirir.

---

### 🧪 Derinlemesine Egzersiz: "Retry Mechanism with Backoff"

Bu öğrendiklerimizi birleştiren, gerçek bir Backend senaryosu kodlayalım.

**Senaryo:** Bir mikroservise istek atıyoruz. Servis bazen cevap vermiyor.

**Görevin:**

1. `urls` adında bir string slice oluştur (`"api.google.com"`, `"api.github.com"`, `"api.bozuk.com"`).
2. Bu URL'ler üzerinde `range` ile dön.
3. Her URL için **içeride** başka bir `for` döngüsü kur (Retry döngüsü). Bu döngü en fazla 3 kere denesin.
4. **Logic:**
    - Eğer URL `"api.bozuk.com"` ise: 3 kere "Deneniyor..." yazsın, sonra "❌ Başarısız" deyip bir sonraki URL'e geçsin (`break` kullanma, iç döngü bitsin).
    - Diğer URL'ler için: 1. denemede "✅ Bağlantı sağlandı" desin ve **Retry döngüsünü kırsın (`break`)**.
5. **Bonus (Label Kullanımı):** Eğer URL `"api.github.com"` ise, "GitHub bulundu, tüm taramayı bitir" deyip **en dıştaki** döngüyü patlat (`Label` kullan).

Bu görev, `range`, `while-style for`, `break` ve `label` konularını tek potada eritecek. Zorlanırsan adım adım gidelim. Kodunu bekliyorum! ⌨️🔥

---

# 📚 Döngü Kontrolü: Break vs Continue

Döngüler normalde şart bitene kadar (örneğin 1'den 100'e kadar) döner.

Ancak bazen bu akışa **manuel müdahale** etmemiz gerekir. İşte bu iki komut, döngünün "Direksiyonu" gibidir.

## 1. Continue: "Bunu Atla, Sıradakine Geç" ⏭️

`continue` komutu, döngüyü sonlandırmaz. Sadece **o anki turu (iterasyonu)** iptal eder ve hemen bir sonraki tura geçer.

### 🧠 Analoji: Müzik Listesi

Bir playlist dinliyorsun. Sevmediğin bir şarkı geldi.

- **Ne yaparsın?** "Sonraki Şarkı" tuşuna basarsın.
- **Sonuç:** Müzik kapanmaz. Sadece o kötü şarkı atlanır ve liste çalmaya devam eder.

### 💻 Nasıl Çalışır? (Under the Hood)

Kod `continue` satırını gördüğü an, o satırın **altındaki hiçbir kodu çalıştırmaz.** Direkt olarak döngünün başına (artırma işlemine `i++`) zıplar.

### Örnek Senaryo: "Sadece Çift Sayıları İşle"

Tek sayıları (1, 3, 5...) sevmiyoruz, onları atlayacağız.

```go
package main

import "fmt"

func main() {
    fmt.Println("--- Çift Sayı Filtresi ---")

    for i := 1; i <= 5; i++ {
        // Eğer sayı TEK ise (kalan 1 ise)
        if i % 2 != 0 {
            fmt.Println("Atlanan Sayı:", i)
            continue // 🛑 DUR! Aşağı inme, direkt i'yi artır ve başa dön.
        }

        // 'continue' çalışırsa burası ASLA çalışmaz.
        fmt.Printf("✅ İŞLENEN SAYI: %d\n", i)
    }
}
```

**Adım Adım İzleme (Trace):**

1. **i = 1:** Tek sayı -> `continue` çalıştı. (Ekrana "İŞLENEN" yazmadı. `i` 2 oldu).
2. **i = 2:** Çift sayı -> `continue` çalışmadı. (Ekrana "İŞLENEN: 2" yazdı).
3. **i = 3:** Tek sayı -> `continue` çalıştı. (Atladı).

---

## 2. Break: "Acil Durum Çıkışı" ⛔

`break` komutu, döngüyü **tamamen öldürür.** O andan itibaren döngü şartı (`i < 100`) doğru olsa bile artık çalışmaz.

### 🧠 Analoji: Yangın Alarmı

Bir binadasın ve iş yapıyorsun. Yangın alarmı çaldı.

- **Ne yaparsın?** İşini bırakıp binayı terk edersin.
- **Sonuç:** Artık diğer odalara bakmazsın. Süreç biter.

### Örnek Senaryo: "Arananı Bulunca Dur"

100 kutu var, içinde elmas olanı arıyoruz. Bulunca diğer 99 kutuya bakmaya gerek yok.

```go
package main

import "fmt"

func main() {
    kutular := []string{"Boş", "Boş", "ELMAS", "Boş", "Boş"}

    for i, icerik := range kutular {
        fmt.Printf("%d. kutuya bakılıyor...\n", i)

        if icerik == "ELMAS" {
            fmt.Println("🎉 Elmas bulundu! Aramayı bırak.")
            break // 🛑 DÖNGÜYÜ PATLAT! Artık dönme.
        }
    }
    
    fmt.Println("Döngüden çıkıldı, hayat devam ediyor.")
}
```

---

## 🧐 Karşılaştırma Tablosu (Özet)

| **Özellik** | **continue ⏭️** | **break ⛔** |
| --- | --- | --- |
| **Görevi** | O anki turu iptal eder. | Döngüyü tamamen iptal eder. |
| **Sonraki Adım** | Bir sonraki elemana/sayıya geçer. | Döngüden sonraki kod satırına geçer. |
| **Altındaki Kodlar** | O tur için çalışmaz. | Asla çalışmaz. |
| **Gerçek Hayat** | Şarkı atlamak. | Radyoyu kapatmak. |

---

**GOTO:**

- **Görevi:** Kodun akışını belirtilen etikete direkt ışınlar.
- **Durum:** Kullanılması şiddetle **ÖNERİLMEZ.**
- **Sebep:** Kodun okunabilirliğini yok eder (Spagetti Kod).
- **Alternatif:** Döngü kırmak için `break Label`, akış kontrolü için `if/else` kullan.
# Maps

Created: January 31, 2026 12:45 AM

C#'taki **`Dictionary<Key, Value>`** neyse, Go'daki **Map** odur.

Tek görevi şudur: **"Bana bir ANAHTAR ver, sana bir DEĞER vereyim."**

Adım adım gidiyoruz.

---

### 1. Map Nasıl Oluşturulur? (Doğrusu ve Yanlışı)

Map referans tipidir (Slice gibi). Yani arkada bir Hash Table tutar. Onu kullanabilmek için önce bellekte yerini hazırlaman gerekir.

### A. Güvenli Yöntem (`make` ile) - Bunu Kullan ✅

C#'taki `new Dictionary<string, int>()` gibidir.

Go

```go
// [AnahtarTipi]DeğerTipi
// Anahtar string olacak (Öğrenci Adı)
// Değer int olacak (Notu)
notlar := make(map[string]int)
```

### B. Hızlı Başlatma (Literal)

İçini dolu başlatmak istersen:

Go

```go
notlar := map[string]int{
    "Ahmet": 50,
    "Ayşe":  90,
}
```

### C. Tehlikeli Yöntem (`var` ile) - Uzak Dur ❌

Bunu yaparsan Map oluşmaz, sadece boş (`nil`) bir pointer oluşur. Yazmaya çalıştığında program patlar.

Go

```go
var notlar map[string]int // nil map!
// notlar["Ali"] = 50 // PANIC! Uygulama çöker.
```

---

### 2. Veri Ekleme ve Güncelleme

C# ile birebir aynıdır. Köşeli parantez `[]` kullanılır.

Go

```go
// 1. Ekleme
notlar["Mehmet"] = 70 // Mehmet yoksa ekler.

// 2. Güncelleme
notlar["Ahmet"] = 85 // Ahmet varsa notunu 85 yapar (Eskisini ezer).
```

---

### 3. Veri Okuma (C#'tan Çok Farklı!) ⚠️

Burası en kritik nokta.

**Senaryo:** "Veli" diye biri map'te yok. Notunu istersen ne olur?

- **C#:** Hata verir (`KeyNotFoundException`).
- **Go:** Hata vermez. **Sıfır Değeri (Zero Value)** döner. `int` için 0 döner.

Go

```go
not := notlar["Veli"] 
// Veli yok ama hata vermedi. 'not' değişkeni 0 oldu.
```

---

### 4. "Var mı Yok mu?" Kontrolü (Comma-ok Idiom)

Madem olmayan elemana `0` dönüyor, peki Veli'nin notu gerçekten `0` mı, yoksa Veli sınıfta mı yok? Bunu nasıl anlayacağız?

Go'da map'ten değer okurken **iki değişken** dönebilirsin:

1. **Value:** Değerin kendisi.
2. **Ok (bool):** Anahtar haritada var mı? (`true`/`false`)

Go

```go
// Syntax: değer, varMı := map[anahtar]
puan, ok := notlar["Veli"]

if ok {
    fmt.Println("Veli'nin notu:", puan)
} else {
    fmt.Println("Veli listede yok!")
}
```

---

### 5. Veri Silme (`delete`)

C#'taki `.Remove()` fonksiyonu yerine Go'da gömülü `delete` fonksiyonu vardır.

Eğer silmek istediğin anahtar zaten yoksa, hiçbir şey yapmaz (Hata vermez).

Go

```go
delete(notlar, "Mehmet") // Mehmet'i listeden siler.
```

---

### TOPLU ÖRNEK (Kopyala - Yapıştır - Çalıştır)

Aşağıdaki kodu editörüne al ve çalıştır. Map'in bütün hayat döngüsü burada.

Go

```go
package main

import "fmt"

func main() {
    // 1. OLUŞTURMA (String -> Int)
    // Öğrenci Adı -> Notu
    sinif := make(map[string]int)

    // 2. EKLEME
    sinif["Ahmet"] = 40
    sinif["Ayşe"] = 90
    sinif["Fatma"] = 75

    fmt.Println("Sınıf Listesi:", sinif)

    // 3. GÜNCELLEME
    // Ahmet sözlüden yüksek aldı, notunu artır.
    sinif["Ahmet"] = 60
    fmt.Println("Ahmet'in Yeni Notu:", sinif["Ahmet"])

    // 4. OLMAYAN VERİYİ OKUMA (Zero Value)
    // 'Can' diye biri yok. Go bize int'in sıfır değeri olan 0'ı verir.
    canNotu := sinif["Can"]
    fmt.Println("Can'ın Notu (Aslında yok):", canNotu)

    // 5. VAR MI YOK MU KONTROLÜ (Comma-ok)
    // Can gerçekten 0 mı aldı yoksa listede mi yok?
    deger, varMi := sinif["Can"]

    if varMi {
        fmt.Println("Can listede var, notu:", deger)
    } else {
        fmt.Println("Can diye bir öğrenci bulunamadı!")
    }

    // 6. SİLME
    delete(sinif, "Ayşe")
    fmt.Println("Silindikten sonra sınıf:", sinif)
}
```

---

### Durum Özeti

Map bu kadar.

1. `make` ile oluştur.
2. `m[key] = val` ile ekle.
3. `val, ok := m[key]` ile oku ve kontrol et.
4. `delete(m, key)` ile sil.

Şimdi bu temel üzerine, makaledeki o karmaşık gibi görünen konuları çok rahat konuşabiliriz.

**Bu kod ve mantık oturdu mu?** Önce bunu teyit edelim, sonra senin şu ana kadar öğrendiklerini pekiştirecek bir Map Projesi yapalım.

[Proje](Proje%202f8ad6137d248066ab00dbffb1038627.md)
# Noktali Virgül

Harika bir soru! Kodun doğru çalışıyor ama o noktalı virgülün (`;`) orada ne işi olduğunu anlamaman çok normal. Çünkü bu yapı C#'ta yok.

O noktalı virgül, Go derleyicisine şunu söyler: **"Dur! Sol taraf bitti, şimdi sağ tarafa (koşula) geç."**

Gel bunu "Under the Hood" (Kaputun altı) mantığıyla parçalayalım.

### 🧠 Mantık: İki Farklı Bölge

Go'daki `if` satırı aslında noktalı virgül sayesinde **iki bölgeye** ayrılır:

```go
//       BÖLGE 1 (HAZIRLIK)       ;      BÖLGE 2 (KARAR)
if isTokenValid := true           ;    role == "admin" && isTokenValid {
```

1. **Hazırlık Bölgesi (Init Statement):**
    - Noktalı virgülden **önceki** kısımdır.
    - Burada bir değişken tanımlarsın (`isTokenValid := true`).
    - Bu kod **sadece bir kez** çalışır (değişken oluşturulur).
2. **Ayraç (`;`):**
    - Derleyiciye "Değişken tanımlama işim bitti, şimdi asıl meseleye gel" der.
3. **Karar Bölgesi (Condition):**
    - Noktalı virgülden **sonraki** kısımdır.
    - Burası `true` veya `false` dönmek zorundadır.
    - `if` bloğunun çalışıp çalışmayacağına burası karar verir.

---

### ❓ "Neden Böyle Bir Şey Var?" (Mühendislik Sebebi)

C# geliştiricisi olarak soruyorsun: *"Neden üst satırda tanımlayıp alt satırda kontrol etmiyoruz?"*

Cevap: **SCOPE (Kapsam) Yönetimi ve Bellek Temizliği.**

### C# Yöntemi (Bellek Sızdırır)

```go
// C#
public void Login() {
    var isValid = CheckToken(); // Değişken burada doğdu.
    
    if (isValid) {
        // ... işlem ...
    }

    // SORUN: 'isValid' değişkeni hala burada yaşıyor! 
    // Hafızada yer kaplıyor ve yanlışlıkla tekrar kullanılabilir.
    Console.WriteLine(isValid); // Çalışır.
}
```

### Go Yöntemi (Cerrahi Kesim) 🔪

```go
// Go
func Login() {
    // Değişken if satırında doğdu.
    if isValid := CheckToken(); isValid {
        // ... işlem ...
    }

    // MÜKEMMEL: 'isValid' burada ARTIK YOK.
    // Garbage Collector (GC) beklemez, Scope bittiği an erişim kesilir.
    // fmt.Println(isValid) // HATA! Kod derlenmez bile.
}
```

### 🤓 Bir Benzetme (For Döngüsü)

Aslında bu yapıyı zaten **`for` döngüsünden** tanıyorsun. Mantık birebir aynıdır!

- **For Döngüsü:**
    
    `for i := 0; i < 10; i++`
    
    *(Başlat ; Koşul ; Artır)* -> Noktalı virgül başlatma ile koşulu ayırır.
    
- **If Yapısı:**
    
    `if i := 0; i < 10 { ... }`
    
    *(Başlat ; Koşul)* -> Noktalı virgül başlatma ile koşulu ayırır.
    

### Özet

O noktalı virgülün tek amacı **Syntax Ayrımıdır.**

- Sol taraf: **Değişkeni yarat.**
- `;` : **Tamam, yaratma bitti.**
- Sağ taraf: **Şimdi kontrol et.**

Bu sayede yarattığın değişken, süslü parantezler `{}` kapandığı an bellekten silinir (Scope dışına çıkar). Backend tarafında milyonlarca istek alan bir API yazdığında, bu minik bellek tasarrufları hayat kurtarır.

Kafana yattı mı şimdi? 💡
# Pointer

Created: January 27, 2026 1:36 AM

# Burda temel düzeyde pointer anlatisi mevcut ielrde detayina gircem

İşte "Mimariden Ödün Vermeyenler" için Go Pointer Özeti:

### 1. Pointer Nedir?

Çok basit: **Pointer, verinin kendisini değil, RAM'deki adresini tutan değişkendir.**

- **Değişken (`x`):** Evin kendisi. İçinde insan (veri) yaşar.
- **Pointer (`p`):** Evin adresi yazılı kağıt. (Örn: `0xc000014080`)

### 2. İki Sihirli Operatör: `&` ve

Go'da pointer kullanmak için sadece bu iki sembolü bilmen yeterli.

- **`&` (Adres Operatörü):** "Nerede oturuyorsun?" diye sorar. Değişkenin adresini verir.
- **(Dereference Operatörü):** "O adrese git ve kapıyı çal" demektir. Adresini bildiğin verinin içine erişir.

Go

```go
package main

import "fmt"

func main() {
	// 1. Normal Değişken (Evin İçindeki Veri)
	sayi := 50

	// 2. Pointer Oluşturma (Adresi Öğrenme)
	// 'p', int tipinde bir verinin adresini tutar (*int)
	var p *int = &sayi 

	fmt.Println("Sayının Adresi:", p) // Çıktı: 0xc0000... (Hexadecimal Adres)

	// 3. Dereferencing (Adrese Gidip Değeri Okuma)
	fmt.Println("Adresteki Değer:", *p) // Çıktı: 50

	// 4. Pointer ile Değer Değiştirme (Uzaktan Kumanda)
	*p = 100 // "p'nin gösterdiği adrese git, içine 100 yaz"
	
	fmt.Println("Sayı'nın Yeni Hali:", sayi) // Çıktı: 100 (Orijinal değişti!)
}
```

### 3. C++ ile Go Arasındaki Dev Fark (Mimari Güvenlik)

C++'ta `p++` diyerek bellekteki bir sonraki adrese (komşu eve) dalabilirdin. Bu, yanlışlıkla işletim sisteminin belleğini okumana veya "Buffer Overflow" hatalarına yol açardı.

**Go'da "Pointer Arithmetic" (Pointer Aritmetiği) YOKTUR.**

- `p = p + 1` diyemezsin.
- Pointer sadece senin ona verdiğin adresi tutar, kafasına göre bellekte gezintiye çıkamaz.
- Bu sayede Go, C++ kadar hızlı ama C# kadar güvenli (Memory Safe) olur.

### 4. Fonksiyonlarda Pointer (Pass by Reference)

C#'ta `ref` keyword'ü kullanırsın ya, Go'da bunun karşılığı Pointer göndermektir.

Neden Kullanırız?

1. **Performans:** 100 MB'lık bir struct'ı fonksiyona gönderirken kopyalamamak için (Sadece 8 byte'lık adres gider).
2. **Side Effect:** Fonksiyonun içinde yapılan değişikliğin orijinal veriyi de değiştirmesi için.

Go

```go
// Parametre olarak int'in KENDİSİNİ değil, ADRESİNİ (*int) istiyoruz
func ZamYap(maas *int) {
    *maas = *maas + 1000 // Adrese git, değeri artır.
}

func main() {
    benimMaas := 5000
    ZamYap(&benimMaas) // Adresimi gönderiyorum
    fmt.Println(benimMaas) // 6000
}
```

### 5. `nil` Pointer (Null Reference Exception)

C#'taki `null`, Go'da `nil`'dir.

Eğer bir pointer hiçbir yeri göstermiyorsa değeri `nil` olur. `nil` olan bir pointer'ın içine erişmeye çalışırsan (`*p`) uygulama çöker (Panic).

Go

```go
var p *int // Şu an nil
// *p = 10 // HATA! (panic: runtime error: invalid memory address or nil pointer dereference)
```

---

### Özet: Neyi Cebine Koymalısın?

1. **`&` (Ampersand):** Adres al.
2.  **(Asterisk):** Adrese git (Değeri oku veya yaz).
3. **Hız:** Büyük verileri taşırken pointer kullanmak, veriyi kopyalamaktan çok daha ucuzdur.

**Bağlantı Noktası:**

Artık "Birisi veriyi, diğeri adresini tutar" mantığını anladın. **Slice** dediğimiz yapı, işte bu mantık üzerine kuruludur. Slice, aslında gizli bir Pointer kullanarak arka plandaki Array'i yönetir.
# Proje

```go
package main

import "fmt"

func main() {

	// map[keyType]valueType => Go arkada bir hash table olusturdu
	serviceLatencies := make(map[string]int)

	serviceLatencies["AuthService"] = 45
	serviceLatencies["PaymentService"] = 150
	serviceLatencies["OrderService"] = 100
	serviceLatencies["NotifyService"] = 800

	// 3. TABLO FORMATINDA YAZDIRMA (Prefix/Suffix Kullanımı)
	// %-20s : Sola yasla, 20 karakter yer aç (Kalanı boşluk doldurur)
	// %10s  : Sağa yasla, 10 karakter yer aç
	fmt.Println("\n------------------------------------")
	fmt.Printf("| %-20s | %10s |\n", "SERVICE NAME", "LATENCY (ms)")
	fmt.Println("------------------------------------")

	// Elle tek tek yazdırıyoruz (Döngüye henüz girmiyoruz)
	// AuthService
	fmt.Printf("| %-20s | %10d |\n", "AuthService", serviceLatencies["AuthService"])
	// PaymentService
	fmt.Printf("| %-20s | %10d |\n", "PaymentService", serviceLatencies["PaymentService"])
	// OrderService
	fmt.Printf("| %-20s | %10d |\n", "OrderService", serviceLatencies["OrderService"])

	fmt.Println("\n[INFO] PaymentService optimize edildi...")
	serviceLatencies["PaymentService"] = 55 // 150'den 55'e düştü

	// %q kullanımı: String'i tırnaklı gösterir.
	fmt.Printf("Güncel Durum -> Servis: %q, Yeni Latency: %v\n", "PaymentService", serviceLatencies["PaymentService"])

	// 5. OLMAYAN VERİ KONTROLÜ (Comma-ok Idiom)
	targetService := "EmailService"
	fmt.Printf("\n[CHECK] %s durumu kontrol ediliyor...\n", targetService)

	// Go burada Exception fırlatmaz, latency 0 gelir. Ama servis var mı yok mu?
	latency, isOnline := serviceLatencies[targetService]

	if isOnline {
		fmt.Printf(" %s Aktif! Gecikme: %dms\n", targetService, latency)
	} else {
		// %T ile tipini görelim, %v ile değerini (false) görelim.
		fmt.Printf(" %s Bulunamadı! (isOnline tipi: %T, değeri: %v)\n", targetService, isOnline, isOnline)
	}

	//  SİLME İŞLEMİ (Delete)
	fmt.Println("\n[ALERT] OrderService kapatılıyor...")
	delete(serviceLatencies, "OrderService")

	// Silineni kontrol edelim
	_, varMi := serviceLatencies["OrderService"]
	fmt.Printf("OrderService hala var mı? : %t\n", varMi) // %t bool için özeldir
	serviceName := "SearchService"                        // İsmi değişkene alalım ki aşağıda kullanalım
	serviceLatencies[serviceName] = 15

	search, isOk := serviceLatencies["SearchService"]

	if isOk && search < 25 {
		fmt.Printf("[%s] cok hizliyiz %d\n", serviceName, search)
	}
}

```

---

Harika iş çıkardın. Kodun çalışması bir yana, asıl önemli olan **C# mimarisiyle Go mimarisi arasındaki farkı** parmak uçlarında hissetmiş olman.

Bu projeden (ve Map konusundan) cebine koyman gereken **5 Kritik Mühendislik Dersi** şunlardır:

### 1. "Zero Value" Felsefesi (Exception Yok!)

C# geliştiricisi olarak en büyük şokun bu olacak.

- **C#:** `dictionary["olmayan_key"]` -> **PATLAR (Exception).** `try-catch` gerekir.
- **Go:** `map["olmayan_key"]` -> **Sessizce `0`, `false` veya `""` döner.**
- **Ders:** Go, akışı kesmeyi sevmez. Eğer 0 değeri senin için "geçerli bir veri" ise (örneğin hava sıcaklığı 0 olabilir), mutlaka `comma-ok` idiomunu kullanmalısın. Yoksa veri geldi sanıp işlem yaparsın.

### 2. Defansif Kodlama: `comma-ok` Idiom

Backend servislerinde bir konfigürasyonu veya cache'i okurken asla direkt okuma yapma.

- **Kötü:** `timeout := config["timeout"]`
- **İyi:** `timeout, ok := config["timeout"]`
    
    Bu desen, Go'nun `try-catch` mekanizmasıdır. Bunu kas hafızana kazı.
    

### 3. Bellek Güvenliği: `nil` vs `Empty`

- `var m map...` dersen elinde **nil** bir pointer olur. Okursun ama yazamazsın (Panic).
- `make(map...)` dersen elinde **başlatılmış (initialized)** bir Hash Table olur.
- **Ders:** Slice'larda `var` kullanmak bazen okay'dir (`append` halleder), ama **Map'lerde her zaman `make` kullan.**

### 4. Sırasızlık (Randomness) İlkesi

Çıktıyı fark ettiysen, ekleme sırasına göre yazılmadı.

- Go, map iterasyonunu **bilerek** randomize eder.
- **Ders:** Bir API yanıtı dönerken JSON içindeki alanların sırasına asla güvenme. Sıralı veri lazımsa Map değil, Slice kullanmalısın.

### 5. Formatlama Gücü (`fmt` Paketi)

Backend loglarında hayat kurtarır.

- `%q`: String'i tırnaklı basar. (Boşluk hatalarını yakalar).
- `%T`: Gelen verinin tipini basar (`interface{}` ile çalışırken çok lazım olacak).
- `%v`: Ne verirsen basar.
- **Ders:** `Print` veya `Println` yerine, loglama yaparken `Printf` kullanmak sana debug sırasında röntgen çekme imkanı verir.

---
# Rune

Created: January 27, 2026 12:08 AM

Tamamdır, şimdi frekansı tam senin istediğin mühendis seviyesine çekiyorum. Ne anaokulu seviyesi ne de akademik makale; direkt "Under the hood" (kaputun altı) mantığı.

ASP.NET Core geçmişin olduğu için şu teknik gerçekten başlayacağız:

### 1. Teknik Tanım: Alias (Takma Ad)

Go'nun kaynak kodunda `rune` tam olarak şöyle tanımlıdır:

Go

`type rune = int32`

Yani `rune` sihirli bir tip değildir, bildiğimiz **32-bitlik bir tamsayıdır.** C#'taki `int` (System.Int32) ile aynı belleği kaplar (4 Byte).

### 2. Neden `char` değil de `int32`?

C# tarafında `char` tipi 16-bit'tir (UTF-16). Bu, çoğu karakteri tutar ama **her şeyi** tutamaz (örneğin bazı emojiler veya eski diller 2 tane `char` gerektirir).

Go mimarisi şöyle der: "Ben karakter kodlamasıyla (encoding) uğraşmak istemiyorum. Dünyadaki her karakterin (Unicode) bir kimlik numarası vardır. Ben bu numarayı direkt sayı olarak tutacağım."

- **`byte` (uint8):** 0-255 arası sayılar. (Sadece standart İngilizce karakterler, ASCII).
- **`rune` (int32):** 0-2.000.000+ arası sayılar. (Dünyadaki tüm diller, emojiler, semboller).

### 3. Kod Üzerinde Kanıt

Döngüye girmeden, sadece değişkenin tipine ve değerine bakalım.

Go

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var r1 rune = 'A'
	var r2 rune = 'Ğ'

	// reflect.TypeOf ile tipine bakıyoruz:
	fmt.Println("r1 Tipi:", reflect.TypeOf(r1)) // Çıktı: int32 (Şaşırtıcı değil mi?)
	
	// Değerine bakıyoruz:
	fmt.Println("r1 Değeri:", r1) // Çıktı: 65 (A'nın ASCII/Unicode kodu)
	fmt.Println("r2 Değeri:", r2) // Çıktı: 286 (Ğ'nin Unicode kodu)
}
```

### 4. Mimari Sonuç

Backend geliştirirken veri tabanından veya JSON'dan bir metin geldiğinde;

- Eğer **harf harf** işlem yapacaksan (örneğin "Parolanın 3. karakteri '?' mi?" diye kontrol edeceksen), Go derleyicisi o karakteri RAM'de bir `int32` sayısı olarak işler.
- C#'ta bu `char` (16-bit) iken, Go'da `rune` (32-bit) olarak işlenir. Bellek farkı budur.
# Scope & Shadowing

Created: January 26, 2026 4:43 AM

**Scope (Kapsam)** ve **Shadowing (Gölgeleme)**, Go'da temiz kod yazmak ile "nedenini bulamadığın bug'larla boğuşmak" arasındaki ince çizgidir.

ASP.NET Core'daki "Scope" mantığına (Class level, Method level, Block level) oldukça benzerdir, ancak Go'nun `:=` operatörü işleri biraz tehlikeli hale getirebilir.

Bu konuyu üç ana başlıkta mimari olarak inceleyelim.

### 1. Scope Hiyerarşisi (Matruşka Bebekleri)

Go'da değişkenlerin görünürlüğü "Blok" (`{ ... }`) bazlıdır. En dıştan en içe doğru şöyle sıralanır:

1. **Universe Scope:** `true`, `false`, `nil`, `int`, `error` gibi dilin göbeğindeki tanımlar. Her yerden erişilir.
2. **Package Scope:** Fonksiyonların dışında tanımlanan değişkenler.
    - **Mimari Fark:** C#'taki `public/private` keywordleri yerine; Baş harf **Büyükse** (`User`) diğer paketlerden erişilebilir (Public), **küçükse** (`user`) sadece bu paketten erişilebilir (Private/Internal).
3. **Function Scope:** Fonksiyon içinde tanımlananlar.
4. **Block Scope:** `if`, `for`, `switch` süslü parantezleri `{}` içinde tanımlananlar.

### 2. Shadowing (Gölgeleme) Nedir?

Shadowing, içteki bir blokta (örneğin bir `if` içinde), dışarıdaki bir değişkenle **aynı isimde** yeni bir değişken tanımlamaktır. Bu durumda, dışarıdaki değişken o blok boyunca "görünmez" olur (gölgelenir).

C# derleyicisi seni aynı metod içinde aynı ismi kullanırsan genellikle uyarır veya hata verir. **Go ise buna izin verir** ve en büyük tuzak buradadır.

### Tehlikeli Senaryo: `err` Gölgelemesi

Backend geliştirirken en çok yaşayacağın sorun budur. Aşağıdaki koda dikkatle bak:

Go

```go
func main() {
	x := 10 // Dışarıdaki (Outer) değişken
    
	if x > 5 {
		x := 99 // DİKKAT: ':=' kullandık. Bu YENİ bir değişken (Inner).
		// Dışarıdaki x şu an gölgelendi. Burada yapılan işlem dışarıyı etkilemez.
		fmt.Println("İçerideki x:", x) // Yazar: 99
	}

	fmt.Println("Dışarıdaki x:", x) // Yazar: 10 (HATA! 99 olmasını bekleyebilirdin)
}
```

**Mimari Risk:** Veritabanı transaction'larında veya hata yönetiminde, iç blokta `:=` ile yanlışlıkla yeni bir değişken oluşturursan, dışarıdaki transaction nesnesi veya hata değişkeni güncellenmez. Kod "başarılı" gibi çalışır ama veri kaydetmez.

### 3. Gerçek Hayat Örneği (Hata Yönetimi)

Aşağıdaki örnekte yapılan mimari hatayı bulabilir misin?

Go

```go
func databaseIslemi() error {
    var err error // Dışarıdaki hata değişkeni

    if true {
        // HATA BURADA: ':=' kullanıldığı için yeni bir 'err' oluştu.
        // Dışarıdaki 'err' hala nil!
        result, err := sahteBirIslem() 
        fmt.Println(result)
        
        if err != nil {
            return err // Burası çalışırsa sorun yok
        }
    }

    return err // BURASI HER ZAMAN 'nil' DÖNER! Çünkü dışarıdaki 'err' hiç değişmedi.
}
```

**Çözüm:** İçerideki atamada `:=` yerine `=` kullanmak veya değişken isimlerini farklılaştırmak.

Go

```go
// Doğrusu:
result, tempErr := sahteBirIslem()
if tempErr != nil {
    err = tempErr // Dışarıdakine atama yaptık (Assignment)
}
```

### Nasıl Korunursun?

1. **Gözlem:** Bloklar `{}` içinde `:=` kullanırken "Acaba üstteki bir değişkeni eziyor muyum?" diye düşün.
2. **Tooling:** `go vet` komutu bazı shadowing durumlarını yakalar. Ayrıca VS Code veya GoLand editörlerinde "Shadowing" uyarısını açabilirsin.

---

**Sıradaki Adım:**

Scope ve Shadowing konusundaki bu "tuzakları" anladıysan, Go'nun en temel veri tiplerini inceleyeceğimiz **Data Types** (Primitive Types) bölümüne geçelim. Özellikle `int` ve `float` kullanırken bellek boyutları (int32 vs int64) backend performansını nasıl etkiler, buna bakabiliriz. Devam edelim mi?
# Senaryolar

---

### 1. Senaryo: Map "Comma-ok" Kontrolü 🗺️

Bir Map'ten değer okurken, anahtarın gerçekten orada olup olmadığını anlamak için bu yapıyı **sürekli** kullanacaksın.

**Amacımız:** Kullanıcı varsa "Hoşgeldin X" demek. Yoksa "Kayıt bulunamadı" demek. Ve `user` değişkeni işimiz bitince bellekten silinsin istiyoruz.

```go
package main

import "fmt"

func main() {
    // Veritabanı simülasyonu
    users := map[string]string{
        "enes":  "Admin",
        "ahmet": "Editor",
    }

    // --- KLASİK (AMATÖR) YÖNTEM ---
    // Değişkeni dışarıda tanımladık.
    role, ok := users["mehmet"] 
    if ok {
        fmt.Println("Rol:", role)
    }
    // KÖTÜ: 'role' ve 'ok' değişkenleri hala burada yaşıyor!
    // Alt satırlarda yanlışlıkla kullanılabilirler.

    // --- GO PRO (STATEMENT INIT) YÖNTEMİ ---
    // Syntax: if <tanımlama>; <koşul> { ... }
    
    // 1. users["enes"] değerini 'u'ya, var mı bilgisini 'found'a atadık.
    // 2. noktalı virgülden sonra 'found'u kontrol ettik.
    if u, found := users["enes"]; found {
        fmt.Printf("Hoşgeldin %s! Sisteme giriş yapıldı.\n", u)
    } else {
        fmt.Println("Kullanıcı bulunamadı.")
    }

    // BURAYA DİKKAT:
    // fmt.Println(u) // HATA! 'u' değişkeni tanımlı değil. 
    // Garbage Collector bu değişkeni if bittiği an temizledi. 🧹
}
```

---

### 2. Senaryo: Hata Kontrolü (`err != nil`) ⚠️

Go'da `try-catch` yoktur. Fonksiyonlar hata döner.

Bir işlemi yapıp, sonucuna göre hemen hata fırlatacaksan değişkeni dışarı sızdırmanın anlamı yoktur.

Örneğin: String'i Sayıya çevirmeye çalışalım (`"123a"` -> Hata vermeli).

Go

```go
package main

import (
    "fmt"
    "strconv" // String convert paketi
)

func main() {
    gelenVeri := "123a" // Bozuk veri

    // --- GO PRO YÖNTEMİ ---
    // strconv.Atoi fonksiyonu (sayı, hata) döner.
    
    if sayi, err := strconv.Atoi(gelenVeri); err != nil {
        // Hata varsa (err nil değilse) burası çalışır
        fmt.Printf("HATA OLUŞTU: %v\n", err) // Çıktı: parsing "123a": invalid syntax
    } else {
        // Hata yoksa burası çalışır. 'sayi' değişkeni burada güvenle kullanılır.
        fmt.Printf("Dönüşüm Başarılı: %d\n", sayi)
    }

    // fmt.Println(sayi) // HATA! 'sayi' artık yok.
    // fmt.Println(err)  // HATA! 'err' artık yok.
}
```

### 🧠 Neden Bu Kadar Önemli? (Mühendislik Bakışı)

Backend yazarken, bir fonksiyonun içinde 10 farklı veritabanı sorgusu veya API isteği yapabilirsin.

Eğer hepsini klasik yöntemle yaparsan:

```go
data1, err := GetData1() // err değişkeni tanımlandı
// ...
data2, err := GetData2() // Aynı err değişkenini tekrar kullandık (tehlikeli olabilir)
```

Scope Init kullanırsan:

```go
if d1, err := GetData1(); err != nil { ... } // err burada öldü.
if d2, err := GetData2(); err != nil { ... } // Yepyeni bir err doğdu ve öldü.
```

Bu yöntem **Namespace Pollution (İsim Uzayı Kirliliği)** sorununu çözer ve değişkenlerin birbirine karışmasını engeller.

Bu iki örnek kafana yattıysa, o "Giriş Güvenliği" challenge'ını bu yeni süper gücünü kullanarak (özellikle If kısmında) yapmanı bekliyorum! 🚀
# Sıralı Sekilde Basmak İçin

# 📚 Özel Konu: Map Verisini Sıralı Okumak

Go'da Map verisi **KAOTİKTİR**. Her çalıştırdığında farklı sırada gelir.

C#'taki `SortedDictionary` gibi veriyi otomatik sıralayan bir yapı Go'da (standart kütüphanede) yoktur.

Sıralı çıktı almak istiyorsan, **"Dolaylı Erişim" (Indirect Access)** tekniğini kullanmalısın.

### 🛠️ Algoritma (Adım Adım)

Bu iş 4 aşamadan oluşur:

1. **Ayıkla:** Map'in anahtarlarını (Keys) alıp bir Slice'a koy.
2. **Sırala:** Bu Slice'ı `sort` paketiyle (A-Z veya 1-9) sırala.
3. **Dön:** Sıralanmış Slice üzerinde `for` ile dön.
4. **Eriş:** Elindeki sıralı anahtarı kullanarak, Map'ten asıl değeri çek.

---

### 💻 Örnek Kod (Şablon)

Aşağıdaki kod, ürünleri A'dan Z'ye sıralayarak raporlar.

Go

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 1. Veri Kaynağı (Sırasız Map)
    stoklar := map[string]int{
        "Mouse": 100,
        "Laptop": 5,
        "Klavye": 50,
    }

    // 2. Anahtarları Toplama (Slice Oluşturma)
    // PRO TİP: Kapasiteyi baştan veriyoruz (len) ki Go sürekli belleği büyütmekle uğraşmasın.
    keys := make([]string, 0, len(stoklar)) 

// slice'a mapteki keyleri atiyoruz
    for k := range stoklar {
        keys = append(keys, k)
    }

    // 3. Sıralama (Slice'ı Alfabetik Dizme)
    sort.Strings(keys) 
    // Artık keys şuna döndü: ["Klavye", "Laptop", "Mouse"]

    // 4. İterasyon ve Erişim (En Kritik Kısım)
    fmt.Println("--- Sıralı Rapor ---")
    
    for _, k := range keys {
c        deger := stoklar[k] // MAP'E GİT VE DEĞERİ GETİR
        
        fmt.Printf("%s : %d adet\n", k, deger)
    }
}
```

---

### ⚠️ POTANSİYEL HATALAR VE TUZAKLAR (Senin İçin Özel)

Az önce yaşadığın kafa karışıklığı tam olarak buradaki 1. maddedir.

### ❌ Hata 1: Index ve Value Karıştırmak

`range keys` yaptığında sana gelen ilk değer **Index**'tir (0, 1, 2...). Ürünün adı (Key) ikinci değerdedir.

Go

```go
// YANLIŞ KULLANIM
for i := range keys { 
    // Burada 'i' sadece 0, 1, 2 sayısıdır.
    // stoklar[i] dersen PATLAR! Çünkü map'te "0" diye bir anahtar yok.
    fmt.Println(stoklar[i]) // ERROR
}

// YANLIŞ KULLANIM 2 (Senin Hatan)
for i, k := range keys {
    // Ekrana 'k' (ürün adı) yerine 'i' (sıra no) basmak.
    fmt.Printf("Ürün: %v", i) // Çıktı: Ürün: 0, Ürün: 1...
}
```

### ✅ Doğrusu

Go

```go
for _, k := range keys { // Index lazım değil, '_' ile attık. 'k' artık "Klavye"dir.
    fmt.Printf("Ürün: %v", k) // Çıktı: Ürün: Klavye...
}
```

### ❌ Hata 2: Map'i Değiştirmeye Çalışmak

`sort.Strings(keys)` yaptığında **MAP DEĞİŞMEZ.** Sadece elimizdeki anahtar listesi (Slice) sıralanır. Map bellekte hala karışıktır. Biz sadece okuma sıramızı değiştirdik.

---

### 🧠 Mühendislik Notu: Performans (Big O)

- **Slice Oluşturma:** O(N)
- **Sıralama (Sort):** O(N log N) - (En maliyetli kısım burasıdır)
- **Döngü:** O(N)

**Sonuç:** Eğer milyonlarca verilik bir map'i her saniye sıralayıp ekrana basmaya çalışırsan işlemciyi yorarsın. Bu işlemi sadece raporlama anında veya gerekli olduğunda yapmalısın.

Bu notları kaydettiysen, **Map ve Döngü** defterini tamamen kapatıp, Go'nun C#'a en büyük farkı attığı **Fonksiyonlar** konusuna geçebiliriz. 🚀
# Slice 2

Created: January 27, 2026 2:03 AM

Özellikle `make` fonksiyonu ve `Capacity` yönetimi, "Asla basit düzey bir mühendis olmak istemiyorum" diyen senin için, Go mülakatlarında ve yüksek trafikli backend projelerinde hayat kurtarır.

Eksik kalan ve mutlaka cebine koyman gereken noktaları C# karşılıklarıyla tamamlayalım.

### 1. `make()` vs `new()` (C# Geliştiricisinin Kafasını Karıştıran İkili)

C#'ta her nesneyi oluşturmak için `new List<int>()` dersin. Go'da ise iki tane ayırma (allocation) fonksiyonu vardır ve görevleri çok farklıdır.

- **`new(T)`:** Sadece bellekte yer ayırır ve **Zero Value** döner. (Pointer döner). Slice için kullanırsan `nil` döner, işine yaramaz.
- **`make(T, ...)`:** Sadece Slice, Map ve Channel için kullanılır. Arka plandaki yapıları (Array vb.) **başlatır (initialize)** ve kullanıma hazır hale getirir.

**Backend Kuralı:** Slice, Map veya Channel oluşturacaksan **DAİMA `make`** kullan. `new`'i şimdilik unutabilirsin.

Go

```go
// YANLIŞ (Slice için):
// p := new([]int) // Sana *[]int döner (Slice'ın pointer'ı). İçi boştur (nil).
// *p = append(*p, 1) // Çok karmaşık ve gereksiz.

// DOĞRU:
s := make([]int, 0) // Kullanıma hazır, boş bir slice.
s = append(s, 1)    // Mis gibi çalışır.
```

### 2. Capacity Planlaması (Performansın Sırrı) 🚀

C#'ta `new List<int>(1000)` diyerek kapasite (Capacity) vermeyi biliyorsundur. Bunu neden yaparız? Liste büyürken sürekli arka planda "Yeni dizi aç -> Kopyala -> Eskiyi sil" maliyetine girmesin diye.

Go'da bu optimizasyon çok daha kritiktir. `make` fonksiyonu 3 parametre alır:

`make([]T, length, capacity)`

- **Length (len):** Şu an erişebildiğin eleman sayısı.
- **Capacity (cap):** Arkada ayrılan koltuk sayısı.

**Senaryo:** Veritabanından 1.000.000 satır çekeceksin.

Go

```go
// KÖTÜ YÖNTEM (Sürekli Reallocation)
var users []User 
// Kapasite başta 0. Go bunu 1, 2, 4, 8... diye sürekli büyütecek.
// Milyon kere kopyalama işlemi olabilir. CPU yanar.

// İYİ YÖNTEM (Tek Seferde Allocation)
// "Bana 0 elemanlı ama 1.000.000 kişilik yer ayır"
users := make([]User, 0, 1000000) 

// append yaptığında kopyalama olmaz, sadece boş koltuğa oturur.
// Çok hızlıdır.
```

### 3. `copy()` Operasyonu (Gerçek Kopyalama)

Az önce Slice'ların referans gibi davrandığını, `s2 := s1` dersen aynı veriye baktıklarını konuştuk.

Peki C#'taki `.ToList()` veya `.Clone()` gibi **bağımsız** bir kopya istiyorsan?

`copy(hedef, kaynak)` fonksiyonunu kullanmalısın.

Go

```go
src := []int{1, 2, 3}
dst := make([]int, len(src)) // Hedefte yer ayırmalısın!

copy(dst, src) // src'yi dst'ye kopyalar.

dst[0] = 999
fmt.Println(src[0]) // 1 (Orijinal bozulmadı!)
```

### 4. Dönüşümler (Conversion) - Tuzaklı Bölge ⚠️

Roadmap'te geçen "Array to Slice" ve "Slice to Array" kısımları.

### A. Array -> Slice (Çok Yaygın)

Elinizde sabit bir dizi var ama API sizden Slice istiyor.

Bunu yapmak bedavadır (Cost-free). Çünkü sadece bir pencere açarsınız.

Go

```go
arr := [3]int{10, 20, 30}
slice := arr[:] // Tüm diziyi kapsayan bir slice oluşturur.
// Kopyalama yoktur. Slice, arr'yi gösterir.
```

### B. Slice -> Array (Go 1.17+) (Dikkat!)

Bazen Slice'ı sabit bir diziye çevirmek istersiniz (örn: Hash fonksiyonu `[32]byte` ister).

Burada **PANIC (Runtime Error)** riski vardır.

Go

```go
s := []int{1, 2} // Uzunluğu 2

// HATA: 
// Slice 2 elemanlı ama biz onu 4'lük diziye çevirmeye çalışıyoruz.
// Go "Veri yetmedi" diyip programı patlatır (Panic).
// a := (*[4]int)(s) 

// DOĞRU KONTROL:
if len(s) >= 4 {
    a := (*[4]int)(s)
    // ...
}
```

---

### Özet: Listene Eklenecekler

Roadmap.sh'taki maddeleri senin bilginle harmanlarsak:

1. **`make`**: Listeni oluştururken `make([]User, 0, tahmin)` şeklinde kullanmayı alışkanlık edin. Performans farkı muazzamdır.
2. **`new`**: Slice için kullanma.
3. **Dönüşüm**: Slice'ı Array'e çevirirken (`slice -> array`) boyut kontrolü yapmazsan uygulaman **Production'da çökebilir.**

Slice defterini tamamen kapattık. ✅

C#'taki `List<T>`'nin Go'daki karşılığı olan bu yapıyı, arkasındaki pointer ve kapasite mantığıyla artık biliyorsun.

---

### 1. Kapasite Dolduğunda Ne Olur? (Senin Kodun)

Senin kodundaki kritik an şu satır:

Go

```go
// 1. Başlangıç: Kapasite 2.
slice2 := make([]int, 0, 2) 

// 2. Doldurma:
slice2 = append(slice2, 1) // [1],    Len:1, Cap:2
slice2 = append(slice2, 2) // [1, 2], Len:2, Cap:2 (DOLDU!)

// 3. KRİTİK AN (Taşma):
// Go bakıyor: "Yeni bir eleman (5) gelecek ama kapasite 2 idi ve doldu."
// KARAR: "Mecburen yeni ve daha büyük bir eve (Array) taşınmalıyım."
slice3 := append(slice2, 5)
```

**Arka Planda Olanlar:**

1. Go, RAM'de **yeni** bir array oluşturur (Kapasite genelde 2 katına çıkar: 4).
2. Eski array'deki `[1, 2]` verisini yeni array'e **kopyalar**.
3. Yeni elemanı (`5`) yeni array'e ekler.
4. `slice3` artık bu **yeni array'e** bakar.
5. `slice2` hala **eski array'e** bakmaya devam eder.

**Sonuç:** Bağ koptu. `slice3`'te yapacağın değişiklik `slice2`'yi etkilemez.

---

### 2. Peki Ya Kapasite Dolmasaydı? (Tehlikeli Senaryo) ⚠️

Eğer başlangıçta kapasiteyi **büyük** verseydin, işler değişirdi.

Go

```go
// Kapasiteyi baştan 4 verdik (Yerimiz bol)
slice2 := make([]int, 0, 4) 
slice2 = append(slice2, 1) // [1]
slice2 = append(slice2, 2) // [1, 2] (Hala 2 kişilik boş yer var)

// KRİTİK AN:
// Go bakıyor: "Yer var mı? Var."
// KARAR: "Taşınmaya gerek yok, aynı eve 3. kişiyi al."
slice3 := append(slice2, 5)
```

**Sonuç:** `slice2` ve `slice3` **AYNI ARRAY'e** bakar (Sadece gördükleri uzunluk farklıdır).

- `slice3` değiştirdiğinde, `slice2`'nin baktığı veriler de değişebilir (Eğer üst üste biniyorsa).

---

### 3. İspat Kodu (Laboratuvar Testi) 🧪

Hemen senin kodunun altına şu "Side Effect Testi"ni ekle ve çalıştır. Bakalım ayrılmışlar mı?

Go

```go
// ... senin kodların devamı ...

fmt.Println("--- Ayrılık Testi ---")
// slice3'ün ilk elemanını değiştiriyorum (Eskiden 1 idi)
slice3[0] = 999

fmt.Println("Slice3 (999 olmalı):", slice3)
fmt.Println("Slice2 (1 mi kaldı?):", slice2)

// EĞER Slice2 hala [1, 2] ise -> Ayrılmışlardır (Farklı Array).
// EĞER Slice2 [999, 2] olduysa -> Ayrılmamışlardır (Aynı Array).
```

### Özet Kural (Altın Bilgi) 🌟

Backend kodlarken `append` kullandığında şu kuralı ezberle:

1. **Reallocation (Yeniden Tahsis) Var:** Eğer `append` sonucu kapasite aşılırsa, **yeni** bir array oluşur. Bağlantı kopar. (Güvenli).
2. **Reallocation Yok:** Eğer kapasite yeterliyse, **eski** array kullanılır. Bağlantı devam eder. (Dikkatli olunmalı).

Bu yüzden Go'da best practice şudur:

`append` sonucunu genelde **aynı değişkene** atarız (`s = append(s, val)`). Böylece "Acaba yeni mi eski mi?" diye düşünmeyiz, hep en güncel ve geçerli olanı takip ederiz.

Bu ayrım oturdu mu? `make` ile kapasite vermenin aslında "ne zaman kopacağını" yönetmek olduğunu görebiliyor musun?
# Slice

Created: January 27, 2026 1:44 AM

C# geliştiricisi olarak `List<T>` kullanmaya alışıksın. Slice, kullanım olarak `List`'e benzer ama mimari olarak **Array üzerine açılmış, boyutu değişebilen bir penceredir.**

İşte kaputun altı:

### 1. Mimari Yapı: Slice Header

Go'da bir Slice değişkeni tanımladığında (`var s []int`), bu değişken verinin kendisini tutmaz. Sadece veriye ulaşmak için gereken 3 küçük bilgiyi tutan bir "Struct"tır.

Bellekte (Stack'te) şöyle görünür:

1. **Pointer (`ptr`):** Arka plandaki dev Array'in (Backing Array) *neresinden* başladığını gösteren adres.
2. **Length (`len`):** Pencerenin şu an kaç elemanı gördüğü.
3. **Capacity (`cap`):** Pointer'ın başladığı yerden, Array'in sonuna kadar kaç elemanlık yer olduğu.

**Mimari Sonuç:**

Sen devasa bir Slice'ı fonksiyona parametre olarak gönderdiğinde (`func foo(s []int)`), verinin kendisi kopyalanmaz. Sadece bu 24 byte'lık (64-bit sistemde) küçük başlık kopyalanır. Bu yüzden Slice kullanımı çok **ucuzdur (lightweight)**.

### 2. Slicing: Array'e Pencere Açmak

Bir Array oluşturup, üzerinden Slice aldığımızda ne olduğunu Pointer mantığıyla görelim.

Go

```go
package main

import "fmt"

func main() {
    // 1. BACKING ARRAY (Asıl Veri Deposu)
    // Bellekte 5 tane int yerleşti.
    arr := [5]int{10, 20, 30, 40, 50}

    // 2. SLICE OLUŞTURMA (Pencere Açma)
    // arr[1:3] -> 1. indeksten başla, 3'e kadar (3 dahil değil) al.
    // Yani {20, 30}
    slice := arr[1:3]

    // Slice Header'ın Durumu:
    // Ptr: arr[1]'in adresi (20'yi gösteriyor)
    // Len: 2 (20 ve 30)
    // Cap: 4 (20, 30, 40, 50 -> Array'in sonuna kadar 4 yer var)

    fmt.Println("Slice:", slice) // [20 30]

    // 3. REFERANS ETKİSİ (Pointer Gücü) ⚠️
    // Slice üzerinden veri değiştiriyoruz.
    slice[0] = 999 

    // Orijinal Array de değişti! Çünkü aynı belleğe bakıyorlar.
    fmt.Println("Array:", arr) // [10 999 30 40 50]
}
```

**Kritik Ders:** Slice, veriyi kopyalamaz, **paylaşır**.

### 3. `append` ve Dinamik Büyüme

Array'ler sabittir, Slice'lar büyüyebilir. Peki arkadaki Array sabitse Slice nasıl büyüyor?

İşte C# `List<T>` mantığı burada devreye giriyor.

`append(slice, eleman)` fonksiyonu çalıştığında Go şuna bakar:

1. **Yer Var mı? (`len < cap`):**
    - Evet: Arka plandaki Array'in boş yerine yazar. `len` değerini 1 artırır. Çok hızlıdır.
2. **Yer Yok mu? (`len == cap`):**
    - **Reallocation:** Go, bellekte (Heap) daha büyük (genelde 2 katı) yeni bir Array oluşturur.
    - **Copy:** Eski verileri yeni Array'e kopyalar.
    - **Swap:** Slice'ın Pointer'ını bu yeni Array'e çevirir.

Bu yüzden `append` kullanırken her zaman atama yapmalısın:

`s = append(s, 10)`

Eğer `s =` demezsen ve arkada yeni bir Array oluşursa, senin `s` değişkenin hala eski (ve artık geçersiz) Array'e bakar.

### 4. `make` Fonksiyonu: Performans Optimizasyonu

Eğer 1 milyon eleman ekleyeceğini biliyorsan, `append` ile tek tek eklersen Go sürekli "Dizi doldu -> Yeni Yer Aç -> Kopyala -> Eskiyi Sil" döngüsüne girer. Bu performansı öldürür.

C#'taki `new List<int>(capacity)` mantığı gibi, Go'da `make` kullanılır.

Go

```go
// Tipi []int olan, Uzunluğu 0, Kapasitesi 1000 olan slice
// Bellekte TEK SEFERDE 1000'lik yer ayrılır.
s := make([]int, 0, 1000) 

s = append(s, 1) // Kopyalama/Reallocation olmaz, çok hızlıdır.
```

### 5. `nil` Slice vs Empty Slice

Backend'de API dönerken bu ikisi JSON'da farklı görünür.

- **Nil Slice:** Başlatılmamış. Pointer `nil`.
    - `var s []int`
    - JSON Çıktısı: `null`
- **Empty Slice:** Başlatılmış ama boş. Pointer bellekte bir yeri gösteriyor.
    - `s := []int{}` veya `make([]int, 0)`
    - JSON Çıktısı: `[]`

### Özet: Backend Mühendisi İçin Slice Kuralları

1. **Pencere Mantığı:** Slice'lar aslında Array pointerlarıdır. Bir Slice'ı değiştirirsen, o Array'i kullanan diğer Slice'lar da etkilenebilir.
2. **`append` Dönüş Değeri:** Daima `s = append(s, val)` şeklinde kullan. Çünkü Array'in adresi değişmiş olabilir.
3. **Kapasite Planlaması:** Boyutu tahmin edebiliyorsan mutlaka `make([]T, 0, cap)` kullan.
4. **Memory Leak Riski:** 1 GB'lık bir diziden `s := arr[0:1]` diye küçücük bir Slice alırsan, o 1 GB'lık dizi RAM'den silinemez (GC toplayamaz). Çünkü senin küçük Slice'ın hala o devasa bloğa tutunuyordur. (Çözüm: `copy` kullanmak).

---

# .Net Karsiligi :

Senin C#'taki **`List<T>`** neyse, Go'daki **`[]T` (Slice)** tam olarak odur.

Veritabanından veri çekerken boyut vermene, array hesaplamana **gerek yok.** Go bunu senin yerine arkada halleder.

Gel, bir **Backend Senaryosu** üzerinden (Veritabanından User listesi çekme) C# ve Go kodlarını yan yana koyalım. O zaman "Hah, aynısıymış!" diyeceksin.

### Senaryo: Veritabanından Kullanıcıları Çekmek

Amacımız: `SELECT * FROM Users` sorgusunun sonucunu bellekte bir listeye doldurmak.

### C# (Alıştığın Yöntem)

C#

```csharp
// 1. Listeyi oluştur (Boyut vermedin, dinamik)
List<User> users = new List<User>(); 

while (reader.Read()) {
    var u = new User { Name = reader["Name"] };
    
    // 2. Ekle (List arkada kendi yerini büyütür)
    users.Add(u); 
}
```

### Go (Slice Yöntemi)

Gördüğün gibi **hiçbir fark yok.** Boyut vermiyoruz, sadece ekliyoruz.

Go

```go
// 1. Slice'ı oluştur (Henüz boş, boyut vermedik)
// C#'taki 'new List<User>()' ile aynıdır.
var users []User 

for rows.Next() {
    var u User
    rows.Scan(&u.Name) // DB'den okuduk
    
    // 2. Ekle (Append)
    // C#'taki 'Add' ile aynıdır.
    // users doldukça Go arkada otomatik büyütür.
    users = append(users, u) 
}

// Döngü bittiğinde elinde dinamik bir liste var.
```

### Senin Sorduğun Kritik Sorular

**Soru 1: "Boyutunu vermem gerekecek mi?"**

**Cevap:** **Hayır.**

Yukarıdaki örnekte `var users []User` dedik. Başlangıç boyutu 0'dır. Sen `append` yaptıkça Go arkada (tıpkı C#'ın `List` yapısı gibi) yer açar, genişletir. Senin "Şu an 10 kişi oldu, 20 yapayım" demene gerek yok.

**Soru 2: "DB'den çektiğim listeyi nerede tutacağım?"**

**Cevap:** İşte bu oluşturduğun **Slice** değişkeninde (`users`) tutacaksın.

**Soru 3: "List ile alakası yok ki?"**

**Cevap:** Aslında kullanımı **birebir aynıdır**.

- C#: `list.Add(item)`
- Go: `slice = append(slice, item)`

Tek fark sözdizimi (syntax).

---

### Peki Neden "Array" ve "Pointer" Anlattım? (İleri Seviye Bilgi)

Sen normal bir `List` (Slice) kullanırken bunları bilmene gerek yok. `append` yap geç.

Benim anlatma sebebim şuydu:

C#'ta `List` bir sınıftır (Class). Go'da `Slice` bir veri yapısıdır.

Eğer bir gün *"Bu Slice'ı 10 milyon kere append ettim, program yavaşladı, neden?"* dersen, o zaman sana *"Çünkü kapasite (Capacity) dolunca Go arkada kopyalama yapıyor, keşke baştan `make` ile yer ayırsaydın"* diyeceğim.

Ama şu an için sadece şunu bilmen yeterli:

**Slice = Dinamik Liste.**

Veritabanından veri çekerken `var s []Tip` de, sonra `append` ile doldur. Gerisini Go halleder.

Şimdi taşlar yerine oturdu mu? Slice'ın senin `List<T>` ihtiyacını tam olarak karşıladığını görebiliyor musun?
# Slice Array Conversions

Created: January 31, 2026 1:21 AM

[basically](basically%202f8ad6137d2480799381e33294746fcd.md)

### 1. Array -> Slice Dönüşümü: Referans Tuzağı

Az önce `myArr[:]` diyerek array'den slice ürettik. Bu işlemin mimari karşılığı **Zero-Copy** (Kopyalamasız) işlemdir.

- **Detay:** Go, bu işlem sırasında bellekte yeni bir veri alanı ayırmaz. Sadece o array'i gösteren bir "Slice Header" (Pointer + Len + Cap) oluşturur.
- **Risk (Side Effect):** Slice üzerinde yaptığın bir değişiklik, ana array'i bozar. Bu, C#'taki `Span<T>` mantığına çok benzer.

Go

```go
func main() {
    // Orjinal Array (Veri Kaynağı)
    sourceArr := [5]int{10, 20, 30, 40, 50}

    // Slice oluşturduk (Sadece bir pencere açtık)
    mySlice := sourceArr[1:3] // {20, 30}

    // DİKKAT: Slice'ı değiştirmek, kaynağı değiştirir!
    mySlice[0] = 999

    fmt.Println(sourceArr)
    // Çıktı: [10, 999, 30, 40, 50] -> 20 değeri kayboldu!
}
```

**Mühendislik Çıkarımı:** Eğer verinin orjinalini korumak istiyorsan, dönüşüm yaparken `copy()` fonksiyonunu kullanmalısın. Aksi takdirde, oluşturduğun slice "mutable view" (değiştirilebilir görünüm) olarak davranır.

---

### 2. Slice -> Array Dönüşümü: Neden Kopyalar?

Slice'tan Array'e geçerken (`[3]int(slice)`), Go veriyi **kopyalamak zorundadır**.

- **Neden:** Array, bellekte sabit bir bloktur (Stack veya Heap). Slice ise dinamik bir yapıdır. Slice'ın gösterdiği veriyi Array'in içine "gömmen" gerekir.
- **Maliyet:** Bu işlem O(N) maliyetlidir çünkü veriler yeni bir adrese taşınır. Ancak bu sayede, yeni Array üzerinde yaptığın değişiklikler eski Slice'ı etkilemez.

---

### 3. Asıl Soru: Neden Hep Slice Kullanmıyoruz? (Array'in Varolma Sebebi)

Go projelerinin %95'inde Slice kullanılır. Ancak %5'lik kısım, senin gibi "performans ve mimari" odaklı mühendisler için kritiktir.

Neden Array kullanırız?

### A. Stack Allocation (Yığın Tahsisi) ve GC Baskısı

Bu en önemli mühendislik sebebidir.

- **Array:** Boyutu derleme zamanında (compile-time) bellidir. Eğer fonksiyon içinde küçük bir array tanımlarsan (örn: `[64]byte`), Go bunu **Stack** üzerinde tutar. Fonksiyon bitince bellek otomatik temizlenir. Garbage Collector (GC) devreye girmez.
- **Slice:** Dinamiktir. Genellikle **Heap** üzerinde yer ayrılır (backing array). Bu da GC'nin o belleği takip etmesi ve temizlemesi gerektiği anlamına gelir.

**Örnek Senaryo:** Saniyede 1 milyon istek alan bir API yazdığını düşün. Her istekte geçici bir buffer (tampon) lazımsa:

- `[]byte` (Slice) kullanırsan: GC sürekli çalışır, CPU harcar (Latency artar).
- `[64]byte` (Array) kullanırsan: Bellek bedavaya temizlenir, GC yükü sıfırdır.

### B. Compile-Time Safety (Derleme Zamanı Güvenliği)

Array kullanırsan, boyutun değişmeyeceğini garanti edersin.

- Bir UUID her zaman 16 byte'tır (`[16]byte`).
- Bir IPv4 adresi her zaman 4 byte'tır (`[4]byte`).
- Bir SHA256 hash'i her zaman 32 byte'tır (`[32]byte`).

Bunu Slice (`[]byte`) yaparsan, birinin yanlışlıkla 3 elemanlı veya 5 elemanlı veri gönderme riskini çalışma zamanına (runtime) bırakırsın. Array ile bu hatayı derlerken yakalarsın.

### C. Memory Layout (Bellek Düzeni)

Array'ler bellekte kesinlikle bitişik (contiguous) durur. Slice da bitişik duran bir diziye bakar ama Slice'ın kendisi bir "header" yapısıdır. Çok boyutlu dizilerde (`[3][3]int` vs `[][]int`) Array kullanmak, CPU Cache (Önbellek) optimizasyonu sağlar çünkü veriler arasında pointer zıplamaları (pointer chasing) olmaz, hepsi peş peşedir.

### Özet Tablo

| **Özellik** | **Array ([N]T)** | **Slice ([]T)** |
| --- | --- | --- |
| **Bellek Yeri** | Genellikle Stack (Hızlı) | Genellikle Heap (GC Yükü Var) |
| **Boyut** | Sabit (Compile-time) | Değişken (Runtime) |
| **Esneklik** | Düşük | Çok Yüksek |
| **Kullanım Alanı** | Fixed-size buffer, UUID, Hash, Crypto Key | Listeler, Veritabanı sonuçları, JSON parse |

**Senin İçin Bir Sonraki Adım:**

Şu an Array ve Slice farkını mimari olarak anladın. Sırada **"Slice'ların `append` işlemi sırasında Kapasite (Capacity) dolarsa bellekte ne olur?"** konusunu incelemek ister misin? Backend performansını doğrudan etkileyen nokta burasıdır.
# String

Harika, C# bilgilerini Go mimarisiyle güncellemeye başlıyoruz. String konusu, Go'nun performans odaklı tasarımının en net görüldüğü yerdir.

C#'ta `string` dediğinde aklına "Karakterler Dizisi" (Array of Chars - UTF-16) gelir.

Go'da ise `string`, **"Read-Only Byte Slice" (Salt Okunur Byte Dilimi)** demektir.

İşte "String aslında string değildir" sözünün altındaki mimari gerçekler:

### 1. Bellek Yapısı (Under the Hood)

Bir string değişkeni tanımladığında, Go bellekte (Stack üzerinde) sadece 2 küçük bilgi tutar. Biz buna **String Header** deriz:

1. **Data (Pointer):** Heap belleğindeki asıl byte dizisinin başlangıç adresini gösterir.
2. **Len (Int):** Byte dizisinin uzunluğu (Byte sayısı, karakter sayısı değil!).

**Mimari Avantaj:**

Sen 1 GB'lık bir string verisini bir fonksiyona parametre olarak gönderdiğinde, Go o 1 GB veriyi kopyalamaz. Sadece bu küçük "Header"ı (Pointer + Len) kopyalar. Bu yüzden Go'da string taşımak çok ucuzdur (lightweight).

### 2. UTF-8 ve `len()` Tuzağı

C# (UTF-16) her karakteri genelde 2 byte olarak tutar. Go (UTF-8) ise **değişken uzunluklu** (variable-width) çalışır.

- İngilizce harfler (`a`, `b`) -> **1 Byte**
- Türkçe harfler (`ğ`, `ş`) -> **2 Byte**
- Emojiler (`🚀`) -> **4 Byte**

Bu durum, `len()` fonksiyonunu kullandığında backend geliştiricilerini tuzağa düşürür.

Go

```go
package main

import "fmt"

func main() {
    // 1. İngilizce
    str1 := "Go"
    fmt.Println(len(str1)) // Çıktı: 2 (2 karakter, 2 byte)

    // 2. Türkçe (MİMARİ TUZAK BURADA)
    str2 := "Göll" 
    // G (1) + ö (2) + l (1) + l (1) = 5 Byte
    fmt.Println(len(str2)) // Çıktı: 5 (Karakter sayısı 4 ama sonuç 5!)
}
```

**Kural:** `len(string)` sana karakter sayısını vermez, **byte sayısını** verir. Karakter sayısını istiyorsan `utf8.RuneCountInString()` kullanmalısın.

### 3. Immutability (Değiştirilemezlik)

C#'ta olduğu gibi Go'da da stringler **immutable**'dır. Bir string'in içindeki bir byte'ı değiştiremezsin.

Go

```go
s := "Hello"
// s[0] = 'h' // DERLEME HATASI! (cannot assign to s[0])
```

Bunun mimari sebebi **Thread Safety** (İş parçacığı güvenliği) ve bellek yönetimidir. Aynı string verisini binlerce goroutine (thread) aynı anda okuyabilir, kimse değiştiremeyeceği için kilit (lock) kullanmaya gerek kalmaz.

### 4. Raw String Literals (Backticks)

ASP.NET Core'da SQL query veya JSON yazarken `@""` (verbatim string) kullanırsın. Go'da bunun karşılığı **Backtick** (     `` `) karakteridir.

- **Çift Tırnak (`"`):** Özel karakterleri işler (`\n` alt satıra geçer).
- **Backtick (```):** Ne görüyorsan odur. Çok satırlı metinler için kullanılır.

Go

```go
// SQL veya JSON için idealdir
query := `
    SELECT * FROM users 
    WHERE id = 1
`
```

---

### Özet: Go String Mimarisi

1. Stringler aslında **`byte` dizisidir**, karakter dizisi değildir.
2. Varsayılan kodlama **UTF-8**'dir.
3. `len()` fonksiyonu **byte sayısını** döndürür. Türkçe karakterlerde dikkatli olunmalıdır.
4. Stringler değiştirilemez (**Immutable**).

---

# String Uzunlugu Almak için

ASP.NET Core'da `.Length` property'si çoğu zaman işini görürdü ama Go'da "karakter sayısı" (özellikle Türkçe harfler ve emojiler işin içindeyse) için standart kütüphaneden özel bir paket kullanıyoruz.

İşte mimari olarak **doğru** ve **en performanslı** yöntem:

### 1. Standart Yöntem: `utf8.RuneCountInString`

Go'nun standart kütüphanesindeki `unicode/utf8` paketi, string'in içindeki byte'ları tek tek analiz eder ve "Burada kaç tane Rune (karakter) var?" sorusunu cevaplar.

Go

```go
package main

import (
	"fmt"
	"unicode/utf8" // Bu paketi import etmelisin
)

func main() {
	metin := "Göll" 

	// YANLIŞ YÖNTEM (Byte Sayısı)
	// 'ö' harfi 2 byte olduğu için sonuç 5 çıkar.
	fmt.Println("Byte Sayısı:", len(metin)) 

	// DOĞRU YÖNTEM (Karakter/Rune Sayısı)
	// UTF-8 paketini kullanır, karakterleri doğru sayar. Sonuç 4 çıkar.
	fmt.Println("Karakter Sayısı:", utf8.RuneCountInString(metin))
}
```

### 2. Alternatif Yöntem: `[]rune` Dönüşümü

String'i önce bir Rune Dizisine (Slice) çevirip, sonra o dizinin uzunluğunu alabilirsin.

Go

```go
metin := "Göll"
runeDizisi := []rune(metin) // String'i rune dizisine çevirdik
fmt.Println(len(runeDizisi)) // Artık 4 verir.
```

### Hangisini Kullanmalısın?

Bir backend mühendisi olarak performans farkını bilmen gerekir:

1. **Sadece Sayacaksan (`utf8.RuneCountInString`):**
    - **Çok Hızlıdır.** Bellekte yeni bir dizi oluşturmaz (allocation yapmaz). Sadece string üzerinde gezinip sayar.
    - **Tavsiye:** Validasyon yaparken ("Kullanıcı adı en az 5 karakter olsun" gibi) bunu kullan.
2. **Manipüle Edeceksen (`[]rune`):**
    - **Daha Maliyetlidir.** Bellekte string'in bir kopyasını oluşturur.
    - **Tavsiye:** Eğer string'i ters çevireceksen veya "ilk 3 karakteri kes" diyeceksen bunu kullan.

**Özet:**

Metin uzunluğu lazım olduğunda ezbere `len()` yazma alışkanlığını bırakıp, **"Bu metinde Türkçe karakter olabilir mi?"** diye düşünmelisin. Cevap evetse -> `utf8.RuneCountInString()`.

---

# Ek Notlar

Harika, roadmap.sh temel sözdizimini vermiş ama bir backend mühendisi olarak **"Kodun Derlenmesi"** ve **"Platform Bağımlılığı"** (Windows/Linux farkı) konularında seni yakabilecek çok kritik eksikler var.

C# geliştiricisi olarak `@""` (Verbatim String) yapısına alışıksın. Go'daki karşılığı **Backtick** (```) olsa da davranışları birebir aynı değildir.

İşte "mimariden ödün vermeyenler" için eksik kalan parçalar:

### 1. Raw String Literals (``...``) ve "Indent" Tuzağı

Roadmap'te "formatı korur" yazıyor ama bunun bedelini söylemiyor.

SQL query yazarken kodu okunaklı olsun diye girinti (indentation) verirsen, o boşluklar string'in **bir parçası** olur.

Go

```go
func main() {
    // Kodu güzel göstermek için TAB bastın
    query := `
        SELECT * FROM users
    `
    // ÇIKTI: 
    // "\t\tSELECT * \n\t\tFROM users"
}
```

**Mimari Risk:**

Eğer bu string'i hassas bir hash işlemine (HMAC, Signature verification) sokarsan veya boşluğa duyarlı bir protokole gönderirsen hata alırsın.

- **C#:** Genellikle `@""` kullanırken de bu risk vardır ama C# 11 ile gelen `"""..."""` yapısı girintiyi otomatik temizler.
- **Go:** Otomatik temizleme yoktur. Ne yazarsan o gider.

### 2. Platform Bağımlılığı (CRLF vs LF)

Bu madde, Windows'ta geliştirme yapıp Linux container (Docker) içinde çalıştırılan projelerde **gizli bug'ların** kaynağıdır.

Raw String (```), kaynak dosyadaki satır sonu karakteri neyse onu alır.

- **Windows'ta kodluyorsan:** Satır sonları `\r\n` (CRLF) olarak stringe girer.
- **Linux'ta kodluyorsan:** Satır sonları `\n` (LF) olarak girer.

**Neden Önemli?**

Bir API'ye çok satırlı bir veri gönderiyorsan ve API sadece `\n` bekliyorsa, senin Windows makinen yüzünden request reddedilebilir. C#'ta bu genelde normalize edilir ama Go'da raw string "ham"dır.

### 3. Backtick İçinde Backtick Kullanamazsın

C#'ta `@ "..."` içinde çift tırnak kullanmak için iki kere basman (`""`) yeterlidir.

**Go'da Raw String içinde Backtick (```) kullanmanın bir yolu yoktur.** Kaçış karakteri (`\`) çalışmaz.

Go

```go
// HATA! Derlenmez.
sql := `SELECT `id` FROM users` 

// ÇÖZÜM: Mecburen (+) ile birleştirme veya Interpreted string kullanma
sql := `SELECT ` + "`id`" + ` FROM users`
```

### Özet Karşılaştırma

| **Özellik** | **Interpreted ("...")** | **Raw (`...`)** | **Mimari Kullanım Alanı** |
| --- | --- | --- | --- |
| **Kaçış Dizileri** | Çalışır (`\n`, `\t`, `\"`) | Çalışmaz (Ne görürse o) | JSON, Regex, HTML şablonları |
| **Satır Sonu** | `\n` yazarsan `\n` olur. | Dosya formatına bağımlıdır (`\r\n` veya `\n`). | Çok satırlı SQL sorguları |
| **Kısıtlama** | Çift tırnak için `\"` gerekir. | **İçinde backtick (`) asla kullanılamaz.** | String içinde SQL/JSON yazımı |

---

### Konsolide Örnek: Backend'de SQL Yazımı

Go

```go
package main

import "fmt"

func main() {
    // 1. Interpreted: Okuması zor, sürekli \n ve \" ile uğraşırsın.
    jsonBad := "{\"user\": \"Ahmet\", \"age\": 30}"

    // 2. Raw: Tertemizdir. JSON ve SQL için standart budur.
    jsonGood := `{"user": "Ahmet", "age": 30}`

    // SQL Örneği (Backtick sorunu yoksa)
    query := `
SELECT * FROM orders
WHERE status = 'active'
`
    fmt.Println(jsonGood)
    fmt.Println(query)
}
```

**Eksik Parça Tamamlandı.** ✅

Metin (String) ve Karakter (Rune) mimarisini tamamen oturttuk. Artık Go'da metinlerin nasıl saklandığını ve nasıl yazıldığını biliyorsun.

Sıradaki konu: **Type Conversion (Tip Dönüşümü).**
# String Type conversion

Elbette, Tip Dönüşümü (Type Conversion) başlığında **String** ile ilgili olan ve backend geliştirirken en sık kullanacağın (veya hata yapabileceğin) 3 kritik dönüşümü detaylandıralım.

C# tarafında `string` manipülasyonu genelde kolaydır ama bellekte neler olduğu gizlidir. Go'da ise bu dönüşümlerin maliyetini bilmek zorundasın.

### 1. String $\leftrightarrow$ Byte Dizisi (`[]byte`)

Backend dünyasında (Network, Dosya Okuma/Yazma, JSON) veriler "Byte" olarak akar. Ancak biz kodda "String" olarak çalışırız. Bu ikisi arasındaki geçiş, Go'nun en temel ama **maliyetli** işlemidir.

### Neden Dönüştürürüz?

- **String:** Değiştirilemez (Immutable). Sadece okursun.
- **Byte Slice:** Değiştirilebilir (Mutable). Dosyadan veri okurken buffer olarak kullanılır.

Go

```go
s := "Merhaba"

// 1. String -> Byte Slice (Bellek Kopyalaması Yapar)
// String'in içeriği tamamen kopyalanıp yeni bir byte dizisi oluşturulur.
bytes := []byte(s) 

// Byte'ları değiştirebiliriz (String'i değiştiremezdik)
bytes[0] = 'm' // Küçük 'm' yaptık

// 2. Byte Slice -> String (Bellek Kopyalaması Yapar)
// Byte dizisinden tekrar immutable bir string üretilir.
s2 := string(bytes)

fmt.Println(s2) // "merhaba"
```

**Mimari Not:** Büyük verilerde (örneğin 10MB'lık bir HTTP body) bu dönüşümü sürekli yapmak performansı öldürür. O yüzden `io.Reader` gibi arayüzler genelde direkt `[]byte` ile çalışır.

---

### 2. String $\leftrightarrow$ Rune Dizisi (`[]rune`)

Metin manipülasyonu (tersten yazdırma, 3. karakteri alma, kesme) yapacaksan ve metinde **Türkçe karakter** varsa bu dönüşüm şarttır.

### Neden Dönüştürürüz?

String byte tabanlıdır, Türkçe karakterleri bölebilir. Rune dizisi karakter tabanlıdır.

Go

```go
text := "Göll"

// YANLIŞ: Byte olarak kesmek
// 'ö' harfi 2 byte olduğu için, ilk 2 byte'ı alırsan "G" ve "ö"nün yarısı gelir.
// fmt.Println(text[0:2]) // Bozuk karakter veya sadece "G"

// DOĞRU: Rune olarak kesmek
runes := []rune(text) // ["G", "ö", "l", "l"]
runes[0] = 'g'        // Karakteri değiştirdik

// Rune -> String
newText := string(runes) 
fmt.Println(newText) // "göll"
```

---

### 3. "Sayıyı String Yapma" Tuzağı (`string(int)`)

C# geliştiricilerinin en çok düştüğü tuzak budur. Bir tamsayıyı string'e çevirmek (casting) ile onu metin olarak formatlamak (formatting) farklı şeylerdir.

Go

```go
num := 65

// TUZAK: Casting Kullanımı
// Go der ki: "Sen bana bir sayı verdin ve bunu string yap dedin.
// 65 sayısı Unicode tablosunda 'A' harfidir. Al sana 'A'."
str := string(num) 
fmt.Println(str) // Çıktı: "A" (Beklenti "65" idi!)

// DOĞRU: strconv Paketi
// "Bana 65 sayısının metin halini ver" demektir.
import "strconv"
str2 := strconv.Itoa(num) // Integer to ASCII
fmt.Println(str2) // Çıktı: "65"
```

### Özet Tablo: Hangisi Ne Zaman?

| **Amaç** | **Dönüşüm Yöntemi** | **Mimari Uyarı** |
| --- | --- | --- |
| **Network/Dosya İşlemleri** | `[]byte("veri")` | Yeni bellek tahsisi (allocation) yapar. |
| **Metin Parçalama/Düzenleme** | `[]rune("veri")` | Türkçe karakterler için zorunludur. |
| **Sayı $\rightarrow$ Metin** | `strconv.Itoa(123)` | `string(123)` yaparsan karakter kodu olarak algılar! |
| **Metin $\rightarrow$ Sayı** | `strconv.Atoi("123")` | Dönüşüm hatası (`error`) dönebilir, kontrol etmelisin. |

Bu string dönüşümleri oturduysa, Go'nun akış kontrolü olan **Control Structures (If, For, Switch)** konusuna geçebiliriz. Özellikle `if` içinde değişken tanımlama ve `while` döngüsünün olmaması ilginçtir.
# stringer ınterface (early)

Created: January 26, 2026 4:15 AM

Gayet haklısın, bu Go'nun C# geliştiricilerine en ters gelen kısmıdır (Implicit Interface). Gel bu konuyu teoriden çıkarıp, doğrudan **C# kodunun Go karşılığına** çevirerek inceleyelim.

Sorun şu: C#'ta bir "sözleşme" imzalarsın (`: IInterface`), Go'da ise sadece "işi yaparsın".

### 1. C# Tarafı (Senin Bildiğin Dünya)

C#'ta `User` sınıfının ekrana nasıl yazılacağını değiştirmek için `Object` sınıfından gelen `ToString` metodunu `override` edersin.

C#

```go
// C# KODU
public class User 
{
    public int Id { get; set; }
    public string Name { get; set; }

    // BURAYA DİKKAT: Açıkça "override" yazarak sistemi uyarıyorsun.
    public override string ToString() 
    {
        return Id + " - " + Name;
    }
}

// Kullanım
var u = new User { Id = 1, Name = "Ahmet" };
Console.WriteLine(u); // Çıktı: 1 - Ahmet
```

*Buradaki mantık:* Derleyiciye "Ben `ToString`'i değiştirdim, benimkini kullan" dersin.

---

### 2. Go Tarafı (Yeni Dünya)

Go'da `override` kelimesi yoktur. `implements` kelimesi yoktur.

Go'nun mantığı şudur: **"Eğer bir struct, `String()` adında, parametre almayan ve geriye `string` döndüren bir metoda sahipse, o struct benim için `Stringer`dır."**

Aynı kodun Go versiyonuna bakalım:

Go

```go
// GO KODU

// 1. Önce Veri Yapısı (Struct)
type User struct {
    ID   int
    Name string
}

// 2. Metot Tanımı (Override YOK, sadece tanımlama var)
// Bu fonksiyon User tipine "yapışık" bir fonksiyondur.
func (u User) String() string {
    return fmt.Sprintf("%d - %s", u.ID, u.Name)
}

func main() {
    u := User{ID: 1, Name: "Ahmet"}
    
    // 3. Büyü burada gerçekleşiyor
    fmt.Println(u) 
}
```

### 3. Arka Planda Ne Oluyor? (Mimari Akış)

Sen `fmt.Println(u)` komutunu çalıştırdığında, Go'nun `fmt` kütüphanesi arka planda şu kontrolü yapar (Pseudo-code mantığıyla):

1. **Soru:** Bana gönderilen `u` değişkeni, `String()` isimli bir metoda sahip mi?
2. **Cevap:** Evet, yukarıda tanımladın.
3. **Karar:** O zaman ben varsayılan çıktıyı basmam, geliştiricinin yazdığı bu metodu çalıştırır ve onun sonucunu ekrana basarım.

Eğer o metodu yazmasaydın (silseydin), `fmt` paketi şöyle diyecekti:

1. **Soru:** `String()` metodu var mı?
2. **Cevap:** Hayır.
3. **Karar:** O zaman ben standart struct çıktısını (`{1 Ahmet}`) basarım.

### Özetle Fark

- **C#:** "Ben bu interface'i uyguluyorum" diye **beyan edersin** (Explicit).
- **Go:** Sadece metodu yazarsın, derleyici "Sen bu metodu yazdığına göre bu interface'i uyguluyorsun demektir" diye **varsayar** (Implicit).

Bu "beyan etmeme" durumu (Implicit), Go mimarisinin en temel taşıdır. Bu sayede 3. parti bir kütüphanedeki struct'a bile, kodunu değiştirmeden kendi interface'inmiş gibi davranabilirsin.
# Strings Review

Created: January 31, 2026 12:26 AM

[for - for range string](for%20-%20for%20range%20string%202f8ad6137d2480189b6ec1a535a9ff57.md)

Rob Pike'ın bu makalesi ("Strings, bytes, runes and characters in Go"), Go dünyasının **Anayasasıdır**. Her Go geliştiricisi, özellikle de senin gibi "kaputun altına" (under the hood) meraklı mühendisler bunu adeta bir manifesto gibi okur.

---

### Ders 1: "String Aslında Nedir?" (Tanım Mücadelesi)

Rob Pike söze balyozla giriyor:

> "In Go, a string is in effect a read-only slice of bytes."
> 
> 
> (Go'da string, aslında salt okunur bir byte dilimidir.)
> 

C# geliştiricisi olarak senin kafandaki String şudur: "Karakterler topluluğu".

Rob Pike diyor ki: **Hayır.** String, **Byte'lar topluluğudur.**

- **Mimari Gerçek:** String'in içinde ille de "Yazı" (Text) olmak zorunda değildir. İçinde rastgele binary veri (Resim dosyası header'ı, şifrelenmiş çöp veri vs.) de olabilir. String sadece bir konteynerdir.
- **Kanıt:** Makaledeki `\xbd\xb2\x3d...` örneği. Bu byte'lar geçerli bir yazı değil, tamamen çöp (garbage). Ama Go buna "Tamam, bu bir stringdir" der.

**Özet:** String = `[]byte` (ama değiştirilemez olanı). İçeriğinin anlamlı bir yazı olup olmadığı Go'nun umurunda değildir.

---

### Ders 2: Ekrana Basma Sanatı (Debugging)

Makalenin büyük kısmı `fmt.Printf` formatlarıyla ilgili. Neden? Çünkü stringler byte olduğu için, içinde "görünmez" veya "bozuk" karakterler olabilir. Debug yaparken bunları görmen lazım.

Pike bize şu taktikleri veriyor:

| **Format** | **Ne Yapar?** | **Ne Zaman Kullanılır?** |
| --- | --- | --- |
| `%s` | Normal yazı olarak basar. | Veri temizse. |
| `%x` | **Hexadecimal** (16'lık) olarak basar. | Veri bozuksa veya binary analiz yapıyorsan. |
| `%q` | Tırnak içinde basar (`"..."`). | Görünmez karakterleri (boşluk, tab) görmek için. |
| `%+q` | **ASCII Only.** Türkçe/Emoji karakterleri `\u` koduyla gösterir. | Unicode hatası arıyorsan (`ğ` yerine ne gidiyor?). |

**Makaledeki Örnek:**

İsveççe "Place of Interest" sembolü (⌘).

- `%s` -> `⌘`
- `%x` -> `e2 8c 98` (İşte bu 3 byte, o sembolün UTF-8 motorudur).

---

### Ders 3: UTF-8 ve "Literal" Kavramı

Burada çok ince bir mimari detay var.

Diyor ki: *"Go kaynak kodları (Source Code) her zaman UTF-8'dir."*

Sen kod editöründe (VS Code) şunu yazdığında:

Go

`const yer = "⌘"`

Editörün o dosyayı UTF-8 olarak kaydettiği için, Go derleyicisi o `⌘` sembolünü otomatik olarak byte'lara (`e2 8c 98`) çevirip string'in içine gömer.

**Kritik Ayrım:**

1. **String Değişkeni:** İçinde her türlü byte olabilir (UTF-8 olmak zorunda değil).
2. **String Literal (Kod içine elle yazılan):** Her zaman geçerli UTF-8'dir. Çünkü senin kaynak kodun UTF-8.

---

### Ders 4: Terminoloji Savaşı (Code Point vs Rune)

Rob Pike burada bilgisayar bilimindeki en büyük kafa karışıklığını çözüyor: **"Karakter (Character) nedir?"**

Diyor ki: *"Karakter kelimesi çok muğlak. Biz bunu kullanmayacağız."*

- **Code Point (Kod Noktası):** Unicode konsorsiyumunun her sembole verdiği tekil ID numarası.
    - Örn: `⌘` -> `U+2318`
- **Rune:** Go dilinin "Code Point"e verdiği isim. `int32` tipindedir.

**Neden Rune?**

Çünkü bir "Karakter" bazen birden fazla Code Point'ten oluşabilir!

- `à` harfi iki şekilde yazılabilir:
    1. Tek Code Point: `U+00E0` (Hazır 'à')
    2. İki Code Point: `U+0061` ('a') + `U+0300` (Üstüne aksan koy)

Go bu karmaşaya girmemek için "Biz Rune (Code Point) sayarız, gerisine karışmayız" diyor.

---

### Ders 5: `for range` Döngüsünün Büyüsü

Makalenin finali, bizim seninle yaptığımız **Egzersiz 3 ve 4'ün teorik ispatıdır.**

Pike diyor ki:

> "A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration."
> 

Yani;

- Normal `for` döngüsü **Makine gibidir**, byte byte gider.
- `range` döngüsü **İnsan gibidir**, karakter karakter (rune rune) gider.

Ve son olarak `unicode/utf8` paketini tanıtıyor. Eğer döngü kullanmadan bir stringin içindeki Rune sayısını bulmak istersen (bizim `len()` tuzağı):

`utf8.DecodeRuneInString` fonksiyonunu kullanın diyor.

---

### Makalenin Özeti (Senin İçin Tek Cümle)

**"String dediğin şey aslında sadece byte dizisidir; onu anlamlı bir yazı (Rune) olarak görmek senin (veya `range` döngüsünün) yorumuna kalmış bir şeydir."**

Bu makale, Go'nun neden C# gibi "String = Nesne" yaklaşımını reddedip, "String = Veri" yaklaşımını seçtiğinin manifestosudur.

Artık Rob Pike'ın ne demek istediğini tamamen biliyorsun. Bu teorik altyapıyla **Maps** konusuna geçtiğimizde, map key'lerinin nasıl çalıştığını çok daha iyi anlayacaksın. Geçelim mi Maps'e? 🚀
# Struct

Created: January 31, 2026 3:20 AM

Şu an tek odak noktamız: **"Go'da veri modelleme nasıl yapılır?"**
C#'taki `Class` mantığını Go'nun `Struct` mantığına tercüme edeceğiz.

Adım adım, en temelden başlayalım.

---

### Adım 1: Struct Nedir? (Tasarım)

C#'ta bir **Class** (Sınıf), hem veriyi (Properties) hem de davranışları (Methods) içinde tutan bir kutudur.
Go'da **Struct**, **sadece veriyi** tutan bir şablondur. Davranışlar (Metotlar) sonradan eklenir.

**Sözdizimi (Syntax):**`type` anahtar kelimesiyle başlar, `struct` ile biter.

Go

```go
// İsimlendirme Kuralı: PascalCase (Büyük harfle başlar) ise Public'tir.
// Eğer 'user' (küçük harf) yazsaydık, sadece bu pakette (dosyada) görünürdü (Private).
type User struct {
    ID       int
    Name     string
    Email    string
    IsActive bool
}
```

Bu tanım bellekte yer kaplamaz. Bu sadece bir **plandır** (Blueprint).

---

### Adım 2: Nesne Oluşturma (Instancing)

Bu plandan nasıl gerçek bir veri oluştururuz? 3 yaygın yöntem vardır.

### Yöntem A: İçi Dolu Oluşturma (En Yaygın - Best Practice) ✅

C#'taki `new User { Name = "..." }` mantığı.

Go

```go
u1 := User{
    ID:       1,
    Name:     "Enes",
    Email:    "enes@gmail.com",
    IsActive: true, // Son elemandan sonra virgül koymak ZORUNLUDUR.
}
```

### Yöntem B: Boş Oluşturma (Zero Value)

Veritabanından dolduracaksan bunu kullanırsın.

Go

```go
var u2 User 
// Şu an u2.ID = 0, u2.Name = "", u2.IsActive = false
```

### Yöntem C: Sıralı Atama (Tavsiye Edilmez ❌)

Field isimlerini yazmadan sırayla değer vermek. Okuması zordur, yapı değişirse kod patlar.

Go

```go
u3 := User{1, "Enes", "x@x.com", true} // Hangi değer hangisiydi? Karışık.
```

---

### Adım 3: Embedding (C#'taki Base Entity Mantığı)

Geldik senin için en önemli kısma.
Go'da `Inheritance` (Kalıtım) yoktur. Yani `User : BaseEntity` diyemezsin.
Bunun yerine **Composition (İçerme/Gömme)** vardır.

Mantık şudur: Bir struct'ı diğerinin içine **isimsiz** olarak koyarsan, Go o struct'ın özelliklerini ana struct'a kopyalar (Promote eder).

**Kod Üzerinde Görelim:**

Go

```go
// 1. Ortak Özellikler (Base)
type BaseEntity struct {
    ID        int
    CreatedAt string
}

// 2. Asıl Yapı (User)
type User struct {
    BaseEntity // <-- İsim vermedik! Sadece tipini yazdık.
    
    Username string
    Password string
}

func main() {
    // KULLANIMI:
    u := User{}
    
    // Normalde şöyle olması gerekirdi: u.BaseEntity.ID
    // Ama Go diyor ki: "Direkt erişmene izin veriyorum."
    u.ID = 100 
    u.CreatedAt = "2025-01-31"
    u.Username = "eneskus"
    
    fmt.Println(u.ID, u.Username)
}
```

**Fark:** C#'ta `User` **bir** `BaseEntity`'dir (is-a). Go'da `User` `BaseEntity`'ye **sahiptir** (has-a), ama sanki kendisininmiş gibi kullanır.

---

### Adım 4: İç İçe Structlar (Nested)

Bazen struct içinde başka bir struct'ı isim vererek kullanırız. Örneğin bir Adres bilgisi.

Go

```go
type Address struct {
    City    string
    Country string
}

type Customer struct {
    Name    string
    Address Address // <-- Burada 'Address' ismini verdik.
}

func main() {
    c := Customer{}
    c.Name = "Pixelance"
    
    // c.City DİYEMEYİZ! Çünkü isim verdik (Embedding değil).
    // Mecburen ismini kullanacağız:
    c.Address.City = "Konya" 
}
```

---

### ÖZET (Cheat Sheet) 📝

1. **Tanımlama:** `type Isim struct { ... }`
2. **Kalıtım Yok:** `BaseEntity`'yi struct'ın içine isimsiz yazarsın (`Embedding`).
3. **Erişim:** Nokta (`.`) ile erişirsin. `u.Name`.
4. **Public/Private:** Büyük harfle başlarsa (`Name`) Public, küçük (`name`) ise Private.
# Struct JSON

Created: January 31, 2026 3:58 AM

C#'ta (ASP.NET Core) bir DTO yazarken, property isimlerini JSON'da küçük harfle göndermek veya bazı alanları gizlemek için **Attribute** kullanırsın:

C#

```go
// C# Yaklaşımı
public class User {
    [JsonProperty("user_id")]  // İsim değiştirme
    public int ID { get; set; }

    [JsonIgnore]               // Gizleme
    public string Password { get; set; }
}
```

Go'da **Attribute** yoktur. Bunun yerine **Struct Tags (Yapı Etiketleri)** vardır.
Mantık aynıdır: **Meta-Data.**

Gel bu konuyu senin sevdiğin gibi "Under the hood" (Kaputun altı) mantığıyla inceleyelim.

---

### 1. Mekanizma: Backtick ve Reflection 🧠

Go'da Struct Tag'ler, struct alanının hemen yanına **Backtick (```)** içine yazılır.

Go

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
}
```

**Nasıl Çalışır? (Under the Hood)**

1. **Derleme Zamanı:** Go derleyicisi bu string ifadeyi (`json:"id"`) alır ve struct'ın tip tanımının içine gömer.
2. **Çalışma Zamanı (Runtime):** `encoding/json` paketi (veya ORM kütüphaneleri), **Reflection** kullanarak bu string'i okur ve *"Haa, geliştirici ID alanını JSON'a çevirirken 'id' adını kullanmamı istemiş"* der.

---

### 2. JSON Etiketlerinin Süper Güçleri 🦸‍♂️

Bir backend mühendisi olarak bilmen gereken 4 temel özellik vardır.

### A. İsim Değiştirme (Renaming)

Go'da alanlar Public olsun diye `PascalCase` (Büyük Harf) yazarız. Ama JSON dünyası `camelCase` veya `snake_case` sever.

Go

```go
Field string `json:"firstName"` // JSON çıktısı: {"firstName": "..."}
```

### B. Veri Yoksa Gönderme (`omitempty`) ⚠️

Bu özellik veri trafiğini azaltmak için harikadır ama **Zero Value** tuzağı vardır.
Eğer alanın değeri **Zero Value** (0, "", false, nil) ise, o alan JSON'a hiç eklenmez.

Go

```go
// Email boşsa JSON'da "email" alanı hiç görünmez.
Email string `json:"email,omitempty"`
```

**Tuzak:** `IsActive bool` json:"active,omitempty"`` dersen ve kullanıcı `false` (Pasif) ise, JSON'da `active` alanı hiç gitmez! Frontend geliştirici "Nerede bu alan?" diye ağlar.

### C. Alanı Gizleme / Yoksayma ()

Şifre, Token gibi alanları struct içinde tutman gerekebilir ama dışarıya (JSON) basarken gizlemelisin.

Go

```go
Password string `json:"-"` // Asla JSON'a çıkmaz.
```

### D. Tip Dönüşümü (`string`)

Bazen Frontend (JavaScript), 64-bit integer sayıları (Long) tutarken yuvarlama hatası yapar. Bu yüzden ID'leri sayı olsa bile string olarak göndermek istersin.

Go

```go
ID int64 `json:"id,string"` 
// Go'da sayı (12345) -> JSON'da String ("12345")
```

---

### 3. C# vs Go Karşılaştırması

| Özellik | C# (`Newtonsoft` / `System.Text.Json`) | Go (`encoding/json`) |
| --- | --- | --- |
| **İsim Değiştirme** | `[JsonProperty("name")]` | ``json:"name"`` |
| **Yoksayma** | `[JsonIgnore]` | ``json:"-"`` |
| **Boşsa Atla** | `NullValueHandling.Ignore` | ``json:",omitempty"`` |
| **Sözdizimi** | Attribute (Sınıfın tepesinde) | Tag (Satırın sonunda) |

E-Tablolar'a aktar

---

### 💻 Pratik Zamanı: "API Response Modeli"

Hadi öğrendiklerimizi birleştiren gerçekçi bir **API Response** senaryosu kodlayalım.

Bu kodu çalıştırıp JSON çıktısını analiz etmeni istiyorum. Özellikle `omitempty` ve `Password` alanına dikkat et.

Go

```go
package main

import (
	"encoding/json" // JSON işlemleri için standart paket
	"fmt"
)

// DTO (Data Transfer Object) Tasarımı
type UserDTO struct {
	// 1. İsim Değiştirme: Go'da ID büyük, JSON'da "user_id" küçük
	ID int `json:"user_id"`

	// 2. Standart Kullanım
	Name string `json:"full_name"`

	// 3. Gizli Alan: Şifre struct'ta var ama JSON'da OLMAYACAK
	Password string `json:"-"`

	// 4. OmitEmpty: Eğer email boşsa, JSON'da yer kaplamasın
	Email string `json:"email,omitempty"`

	// 5. Type Coercion: Sayıyı string olarak gönder (Frontend dostu)
	Balance float64 `json:"balance,string"`
}

func main() {
	// Senaryo 1: Dolu Kullanıcı
	user1 := UserDTO{
		ID:       101,
		Name:     "Enes Furkan",
		Password: "super_secret_pass", // Bu gizlenecek
		Email:    "enes@sunstech.com",
		Balance:  450.50, // String olarak gidecek "450.5"
	}

	// Marshaling (Struct -> JSON)
	// C#'taki JsonConvert.SerializeObject()
	jsonData1, _ := json.Marshal(user1)
	fmt.Println("--- Dolu Kullanıcı JSON ---")
	fmt.Println(string(jsonData1))

	// Senaryo 2: Eksik Kullanıcı (Email Yok)
	user2 := UserDTO{
		ID:   102,
		Name: "Mehmet",
		// Email vermedik (Zero Value: "") -> JSON'da görünmeyecek
		// Balance vermedik (Zero Value: 0) -> "0" olarak gidecek
	}

	jsonData2, _ := json.Marshal(user2)
	fmt.Println("\n--- Eksik Kullanıcı JSON (Email Yok) ---")
	fmt.Println(string(jsonData2))
}
```

### 🕵️ Görev: Çıktı Analizi

Kodu çalıştır ve şu sorulara cevap ver (Kendi kendine veya bana yazabilirsin):

1. `Password` alanı JSON çıktısında var mı?
2. `user2` (Mehmet) çıktısında `email` alanı var mı?
3. `balance` alanı sayı olarak mı (`450.5`) yoksa string olarak mı (`"450.5"`) gitti?
# Struct Tags

Created: January 31, 2026 4:01 AM

Bu gönderdiğin liste, Go dünyasının **"Sarı Sayfalar"ı (Yellow Pages)** gibidir. 📒

Go derleyicisi `json` veya `xml` haricindeki etiketleri (tags) kendiliğinden bilmez. Bunların hepsi üçüncü parti kütüphanelerdir. Bu liste, kütüphane yazarları "Ben hangi kelimeyi rezerve etsem?" diye baktığında çakışma olmasın diye tutulan bir kayıttır.

Bir .NET Core Backend geliştiricisi olarak senin için bu listedeki **en kritik 4 etiketi** ve ne işe yaradıklarını, C# karşılıklarıyla analiz edelim.

---

### 1. `json` & `xml` & `yaml` (Serileştirme Ailesi)

Listenin en başındakiler. Veriyi bir formattan diğerine çevirmek için kullanılır.

- **C# Karşılığı:** `[JsonProperty("name")]`, `[XmlElement("Name")]`
- **Kullanım:** API request/response ve Config dosyaları.

```go
type Config struct {
    // Hem JSON hem YAML (Kubernetes/Docker configleri) için etiket
    Port int `json:"port" yaml:"port"` 
}
```

### 2. `db` & `gorm` (Veritabanı / ORM Ailesi) 🗄️

Listede `gorm` ve `db` (sqlx) göreceksin. Bunlar Entity Framework Core veya Dapper mantığıdır.

- **C# Karşılığı:** `[Column("user_id")]`, `[Key]`, `[NotMapped]`
- **Mantık:** Struct alanını veritabanı tablosundaki sütunla eşleştirir.

```go
type User struct {
    // Gorm'a diyoruz ki: Bu Primary Key'dir ve veritabanında adı 'id'dir.
    ID uint `gorm:"primaryKey;column:id"` 
    
    // Bu alan veritabanında yok, sadece kodda var (NotMapped)
    Token string `gorm:"-"`
}
```

### 3. `validate` (Doğrulama Ailesi) ✅

Listede `github.com/go-playground/validator` adresini göreceksin. Go'da en çok kullanılan validasyon kütüphanesidir.

ASP.NET Core'daki **Data Annotations**'ın birebir karşılığıdır.

- **C# Karşılığı:** `[Required]`, `[EmailAddress]`, `[MinLength(5)]`

```go
type LoginRequest struct {
    // Zorunlu alan ve geçerli bir email formatı olmalı
    Email string `validate:"required,email"` 
    
    // Zorunlu ve en az 6 karakter
    Password string `validate:"required,min=6"`
}
```

### 4. `mapstructure` (Generic Dönüşüm) 🔄

Bunu sık göreceksin. Genelde `config.yaml` gibi dosyalardan okunan `map[string]interface{}` tipindeki veriyi, senin tanımladığın `struct` yapısına oturtmak için kullanılır.

---

### 🚀 Uygulamalı Lab: "The Mega Struct"

Şimdi seninle gerçek bir Backend senaryosu kuralım.

Bir **Kullanıcı Kayıt (Sign Up)** DTO'su tasarlayacağız. Ama bu struct öyle yetenekli olacak ki:

1. **JSON** olarak Frontend'den gelecek.
2. **Validator** ile verisi kontrol edilecek.
3. **ORM (Gorm)** ile veritabanına yazılacak.

Hepsini **TEK SATIRDA** nasıl yapıyoruz? İşte Go'nun pratikliği.

```go
package main

import (
	"encoding/json"
	"fmt"
	"reflect" // Tag'leri okumak için
)

// SENARYO: User hem API'den geliyor, hem validasyondan geçiyor, hem DB'ye gidiyor.
type User struct {
	// 1. JSON: "user_id" olarak gelir/gider.
	// 2. GORM: Veritabanında Primary Key'dir.
	ID int `json:"user_id" gorm:"primaryKey"`

	// 1. JSON: "username" ismini kullanır.
	// 2. VALIDATE: Zorunludur.
	// 3. GORM: Veritabanında sütun adı "user_name" olsun.
	Username string `json:"username" validate:"required" gorm:"column:user_name"`

	// 1. JSON: ASLA görünmez (Şifre dışarı sızmaz).
	// 2. VALIDATE: En az 6 karakter olmalı.
	Password string `json:"-" validate:"min=6"`
	
	// Sadece bellek içi bir alan, DB'ye gitmez, JSON'a gitmez.
	SessionToken string `json:"-" gorm:"-"`
}

func main() {
	u := User{
		ID:       101,
		Username: "eneskus",
		Password: "gizlisifre123",
	}

	// --- 1. Reflection ile Tag Okuma (Under the hood) ---
	// Go kütüphaneleri arka planda bunu yapar:
	t := reflect.TypeOf(u)
	field, _ := t.FieldByName("Username")

	fmt.Println("--- TAG ANALİZİ ---")
	fmt.Println("Ham Tag Verisi:", field.Tag)
	fmt.Println("JSON Değeri   :", field.Tag.Get("json"))
	fmt.Println("GORM Değeri   :", field.Tag.Get("gorm"))
	fmt.Println("Validate Değeri:", field.Tag.Get("validate"))

	// --- 2. JSON Çıktısı Testi ---
	fmt.Println("\n--- JSON ÇIKTISI ---")
	jsonData, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(jsonData))
	// Dikkat et: Password yok, ID "user_id" oldu.
}
```

### 🧠 Mühendislik Notu: Boşluk Kuralı (Syntax Trap)

Burada çok önemli bir sözdizimi kuralı var. C# geliştiricilerinin en çok hata yaptığı yerdir.

**Doğru:**

`json:"name" gorm:"column:name"` (Anahtarlar arasında **boşluk** var).

**Yanlış:**

`json:"name",gorm:"column:name"` (Virgül koyma! Go anlamaz).

`json: "name"` (İki noktadan sonra boşluk koyma! Tag bozulur).

---

### Özet

O listede gördüğün her şey, bir kütüphanenin "Ben bu struct alanını işlerken bu etikete bakacağım" dediği bir sözleşmedir.

Senin şu an için bilmen gerekenler:

1. **`json`**: API için.
2. **`xml`**: XML API'ler için.
3. Diğerleri (`gorm`, `validate`, `yaml`) ilgili kütüphaneyi projene dahil ettiğinde anlam kazanır.

Bu mantık oturduysa, artık veriyi JSON'a çevirmeyi de öğrendik. Şimdi bu verilerle iş yapan **Fonksiyonlara (Functions)** ve **Metotlara (Methods)** geçebiliriz. C#'tan farklı olarak Go'da fonksiyonlar "Birinci Sınıf Vatandaş"tır (First Class Citizen). Hazır mısın? 🚀
# Type Conversion

Created: January 27, 2026 12:14 AM

C# ve ASP.NET Core geliştiricisi olarak en çok zorlanacağın ama zamanla "İyi ki böyle yapmışlar" diyeceğin bir konuya geldik: **Type Conversion (Tip Dönüşümü).**

C# tarafında derleyici (compiler) senin adına birçok şeyi sessizce halleder.

- `int` $\rightarrow$ `long` (Implicit Conversion)
- `float` $\rightarrow$ `double`

**Go'da "Implicit Conversion" (Gizli Dönüşüm) YOKTUR.**

Go mimarisi der ki: "Eğer bir veri tipini değiştireceksen, bunu açıkça (explicitly) yazmalısın. Ben senin yerine niyet okumam."

Bu yaklaşım, özellikle büyük backend projelerinde veri kaybını ve taşma (overflow) hatalarını derleme zamanında yakalamanı sağlar.

### 1. Sayısal Dönüşümler: Katı Sınırlar

En yakın akraba olan tipler (`int` ve `int64`) bile birbirine doğrudan atanamaz.

Go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 3
	var y int = 4

	// HATA: math.Sqrt fonksiyonu parametre olarak 'float64' ister.
	// C# olsa 'int'i otomatik 'double'a çevirirdi. Go yapmaz!
	// f := math.Sqrt(x*x + y*y) 

	// DOĞRU: Açıkça dönüştürmelisin.
	f := math.Sqrt(float64(x*x + y*y))
	
	// f şu an float64 (5.0). Bunu uint'e çevirmek istiyoruz.
	// z := uint(f) // Yine açıkça yazmalısın.
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
```

**Mimari Not:** Bu katılık sayesinde, 32-bit bir sayıyı yanlışlıkla 8-bit bir değişkene atayıp veri kaybetme (overflow) riskin sıfıra iner. Derleyici seni durdurur.

### 2. String ve Byte Dönüşümü (Maliyetli Dönüşüm)

Daha önce konuştuğumuz "String aslında byte dizisidir" konusunu hatırlıyorsun. Bu iki tip arasında geçiş yapabilirsin ama bunun bir **maliyeti** vardır.

Go

`str := "Go"
bytes := []byte(str) // String -> Byte Slice
str2 := string(bytes) // Byte Slice -> String`

**Kritik Mimari Detay:**

Bu işlem C#'taki gibi sadece bir etiket değişimi (casting) değildir. Go burada **bellek tahsisi (memory allocation)** yapar ve veriyi kopyalar.

- Çünkü `string` **immutable** (değişmez).
- `[]byte` **mutable** (değişebilir).
    
    Eğer kopyalama yapmasaydı, byte dizisini değiştirerek string'i bozabilirdin. Bu yüzden büyük metinlerde sürekli dönüşüm yapmak performansı düşürür.
    

### 3. Custom Types (Domain Güvenliği)

Backend mimarisinde **Domain Driven Design (DDD)** uygularken bu özellik harikadır.

Diyelim ki `UserID` ve `OrderID` ikisi de `int` veritabanında. Ama kod içinde yanlışlıkla bir kullanıcının ID'sini sipariş arama fonksiyonuna göndermek istemiyorsun.

Go

```go
type UserID int
type OrderID int

func FindOrder(id OrderID) { ... }

func main() {
    var uID UserID = 10
    var oID OrderID = 20

    // FindOrder(uID) // DERLEME HATASI! (Cannot use UserID as OrderID)
    
    // Ancak sen "Ben ne yaptığımı biliyorum" dersen dönüşüm yapabilirsin:
    FindOrder(OrderID(uID)) 
}
```

Bu özellik, "Primitive Obsession" (İlkel tip takıntısı) anti-pattern'ini engellemek için mükemmeldir.

### 4. Interface Conversions (Type Assertion)

C#'taki `as` operatörü veya `(string)obj` casting işlemi Go'da **Type Assertion** olarak geçer. `interface{}` (herhangi bir şey) tutan bir değişkenden asıl tipi çıkarmak için kullanılır.

Go

```go
var data interface{} = "Merhaba"

// 1. Güvensiz Yöntem (Panic Riski)
s := data.(string) // Eğer data string değilse uygulama çöker (Panic)!

// 2. Güvenli Yöntem (Comma-ok idiom) -> C#'taki 'as' mantığı
s, ok := data.(string)
if ok {
    fmt.Println("Bu bir string:", s)
} else {
    fmt.Println("Bu string değil.")
}
```

### 5. Büyük Tuzak: `string(65)`

C# geliştiricilerinin en sık düştüğü hata budur. Bir sayıyı string'e çevirmek için sakın `string()` casting kullanma!

Go

```go
num := 65
str := string(num) 
// BEKLENTİ: "65"
// GERÇEK: "A" (Çünkü 65, 'A' harfinin Unicode karşılığıdır!)
```

**Doğrusu:** Sayıyı metne çevirmek (Parsing/Formatting) bir dönüştürme değil, bir işlemdir.

- `strconv.Itoa(65)` $\rightarrow$ `"65"`

---

### Özet: Dönüşüm Kuralları

1. **Sayılar:** `T(v)` sözdizimi zorunludur. `float64(i)`, `int(f)`. Gizli dönüşüm yoktur.
2. **String/Byte:** Dönüşüm mümkündür ama bellek kopyalaması (allocation) yapar.
3. **Interface:** Dinamik tipler için `val.(T)` (Type Assertion) kullanılır.
4. **Parsing:** Sayıyı yazıya çevirmek (`"123"` -> `123`) casting değildir, `strconv` paketi kullanılır.

**Sıradaki Adım:**

Veri tiplerini ve dönüşümlerini tamamen bitirdik. Artık kodun akışını yöneteceğimiz **Control Structures** (If, Else, Switch) konusuna geçebiliriz. Özellikle `if` bloğu içinde değişken tanımlama yeteneği ve `switch` yapısının C#'tan ne kadar daha güçlü olduğunu (break yazmaya gerek olmaması gibi) görmek ister misin?

[String Type conversion](String%20Type%20conversion%202f4ad6137d248066875edf171d801f35.md)
Name,Created,Tags
"1.1 Function Basics: Sözdizimi, Parametreler ve C# Farkları","February 1, 2026 2:37 AM",
"1.2 Multiple Return Values ","February 1, 2026 3:01 AM",
Name,Created,Tags
"1.1 Function Basics: Sözdizimi, Parametreler ve C# Farkları","February 1, 2026 2:37 AM",
"1.2 Multiple Return Values ","February 1, 2026 3:01 AM",
# var & :=

Created: January 26, 2026 3:24 AM

ASP.NET Core geçmişinle `var` (C#) kullanımına alışkınsın, ancak Go'da bu ayrım sadece "sözdizimsel şeker" (syntactic sugar) değil, aynı zamanda değişkenin **scope (kapsam)** ve **yaşam döngüsü** hakkında kod okuyan kişiye sinyal veren mimari bir araçtır.

İşte "Neyi, neden kullanıyoruz?" sorusunun cevabı:

### 1. `var` Anahtar Kelimesi: Explicit (Açık) ve Geniş Kapsam

`var`, C#'taki tip belirterek tanımlamaya (`int x;`) benzer. Go'da bunu kullanmanın temel mimari sebepleri şunlardır:

- **Paket Seviyesi (Package Scope):** Fonksiyonların dışında (global alanda) değişken tanımlamanın **tek** yoludur. Go, fonksiyon dışındaki her ifadenin bir anahtar kelime (`var`, `func`, `const` vb.) ile başlamasını zorunlu kılar.
- **Zero Value Niyeti:** Eğer bir değişkene hemen değer atamayacaksan ve onun "boş" (default) haliyle başlamasını istiyorsan `var` kullanırsın. Bu, okuyucuya şunu söyler: *"Bu değişken tanımlandı ama değeri daha sonra, muhtemelen bir logic sonucunda dolacak."*Go
    
    `var count int // Okuyucu anlar ki: Başlangıçta 0, sonra değişecek.`
    
- **Tip Belirginliği:** Bazen sağ taraftaki değerin tipi (inference) istediğimiz tip olmayabilir. Örneğin, varsayılan `int` yerine `int64` istiyorsan `var` ile bunu dikte edersin.

### 2. `:=` Operatörü: Short Declaration (Kısa Tanımlama)

Bu, C#'taki `var` (implicit typing) ile neredeyse aynıdır.

- **Sadece Fonksiyon İçi (Local Scope):** Sadece bir fonksiyonun bloğu içinde çalışır.
- **Init Zorunluluğu:** Hem tanımlama hem atama yapar. Başlangıç değeri vermek zorundasın.
- **Temiz Kod (Clean Code):** Yerel değişkenlerin tipini sürekli yazmak göz yorar. `:=` odağı tipe değil, iş mantığına (business logic) çeker.

### Kritik Mimari Fark: "Redeclaration" (Yeniden Tanımlama) Kuralı

Backend geliştirirken en sık karşılaşacağın ve C#'tan en farklı olan durum budur.

C#'ta aynı scope içinde `var x = 1; var x = 2;` dersen hata alırsın.

Go'da `:=` operatörü, **eğer sol tarafta en az bir tane YENİ değişken varsa** çalışmaya devam eder. Bu, özellikle hata yönetiminde (error handling) kritik bir kolaylık sağlar.

Go

```go
func main() {
    // f, os.Open'dan dönen bir dosya, err ise hata
    f, err := os.Open("file.txt") 
    
    // DİKKAT: 'err' burada zaten tanımlı. 
    // Ancak 'd' (data) yeni olduğu için Go buna izin verir.
    // 'err' değişkenine yeni değer atanır (assignment), 'd' ise tanımlanır (declaration).
    d, err := f.Stat() 
}
```

**Neden Önemli?** Eğer bu özellik olmasaydı, her işlem için `err1`, `err2`, `err3` gibi farklı hata değişkenleri tanımlamak veya sürekli `var err error` ile önceden tanımlama yapmak zorunda kalırdık. Bu tasarım, Go'nun akışkan kod yapısını destekler.

### Özet: Karar Matrisi

| **Durum** | **Kullanılması Gereken** | **Sebep** |
| --- | --- | --- |
| **Fonksiyon Dışı** | `var` | Zorunluluk. |
| **Fonksiyon İçi (Genel)** | `:=` | Okunabilirlik ve hız. |
| **Zero Value İhtiyacı** | `var` | Başlangıç değeri (0, nil) atamanın en temiz yolu. |
| **Interface Dönüşümü** | `var` | Spesifik bir interface tipiyle saklamak istediğinde. |

**Sıradaki Adım:**

Değişkenlerin tanımlanmasını hallettik. Şimdi bu değişkenlerin bellekte varsayılan olarak ne durumda olduğunu anlamak için **Zero Values** konusuna detaylıca bakalım mı? Yoksa Go'nun tiplerine (Basic Types) mi geçelim? Zero Value, "Null Reference" hatalarını mimari olarak nasıl engelliyor, orası ilginç olabilir.
# Variables and Constans

Created: January 26, 2026 3:02 AM

Go'nun değişken ve sabit (constant) yönetimindeki yaklaşımı, sadelik ve güvenliği birleştirir. ASP.NET Core geçmişin olduğu için C#'taki `var` keyword'üne aşinasın, ancak Go'da işler biraz daha farklı ve katı kurallara bağlıdır.

Konuyu derinlemesine kavramak için şu mimari başlıkları zihninde oturtman gerekir:

### Bu Konuyu Anlamak İçin Gereken Altyapı

<aside>
✅

[Here ](Here%202f4ad6137d2480b788cdc3a189ccff95.md)

</aside>

1. **Memory Allocation (Stack vs Heap):** Değişkenlerin nerede oluşturulduğu (Scope) performans için kritiktir.
2. **Static Typing vs Dynamic Typing:** Go statik tipli bir dildir, ancak `:=` operatörü ile dinamik dillerin esnekliğini hissettirir.
3. **Zero Value:** Go'da "uninitialized variable" (başlatılmamış değişken) diye bir kavram yoktur. Bu güvenlik mekanizmasını anlamak şarttır.

---

### 1. Değişken Tanımlama Yöntemleri ve Mimari Tercihler

Go'da değişken tanımlamanın iki ana yolu vardır ve hangisini seçeceğin **"niyetini" (intent)** belirtir.

### A. Uzun Deklarasyon (`var`)

Daha resmi ve explicit (açık) bir yöntemdir. Genelde iki durumda kullanılır:

1. **Paket Seviyesinde (Global Scope):** Fonksiyonların dışında değişken tanımlarken `:=` kullanamazsın, `var` zorunludur.
2. **Zero Value Kullanılacaksa:** Eğer değişkene hemen bir değer atamayacaksan ve varsayılan değeriyle (0, "", false) başlamasını istiyorsan kullanılır.

Go

```go
var count int // Değer atanmadı, otomatik olarak 0 oldu.
```

### B. Kısa Deklarasyon (`:=`)

C#'taki `var` kullanımıyla benzerdir ancak sadece **fonksiyon içinde** çalışır.

- Hem tanımlama (declaration) hem atama (assignment) yapar.
- Derleyici tipi sağdaki değere göre otomatik belirler (Type Inference).

Go

```go
func main() {
    name := "Go" // String olduğu derleme zamanında belirlenir.
}
```

> Best Practice: Fonksiyon içinde her zaman := kullanmaya çalış. Eğer bir değişkeni var ile tanımlayıp sonra değer atıyorsan, Go topluluğunda bu genellikle "fazla kod" (non-idiomatic) olarak görülür.
> 

### 2. Zero Value (Sıfır Değer Güvencesi)

C# tarafında `string` bir değişkeni initialize etmezsen `null` olabilir ve `NullReferenceException` alabilirsin. Go'da bu mimari risk **Zero Value** kavramı ile minimize edilmiştir.

Bir değişkeni sadece tanımladığında (`var i int`), bellekte ona yer ayrılır ve türüne uygun varsayılan değerle doldurulur:

- `int` $\rightarrow$ `0`
- `bool` $\rightarrow$ `false`
- `string` $\rightarrow$ `""` (boş string, null değil!)
- `pointer`, `interface`, `slice`, `channel` $\rightarrow$ `nil`

Bu sayede, Go programları "rastgele bellek çöpü" (garbage value) ile çalışmaz.

### 3. Sabitler (Constants) ve Derleme Zamanı

Go'daki `const` yapısı C#'tan çok daha kısıtlı ama performanslıdır.

- **Compile-time Only:** Sabitlerin değeri **derleme anında** bilinmelidir. Çalışma zamanında (runtime) hesaplanan bir fonksiyon sonucu `const` olamaz.
- **Typed vs Untyped Constants:** Go'da sabitler bir tipe sahip olmak zorunda değildir.Go
    
    ```go
    const Pi = 3.14 // Tipi yok (untyped float). Kullanıldığı yere göre float32 veya float64 gibi davranabilir.
    ```
    
    Bu esneklik, matematiksel işlemlerde tip dönüşümü (casting) karmaşasını azaltır.
    

### 4. Dikkat Edilmesi Gereken Tuzak: Variable Shadowing

İç içe bloklarda (scope) aynı isimli değişken tanımlarsan, içteki değişken dıştakini "gölgeler" (shadow).

Go

```go
func main() {
    x := 10
    if true {
        x := 5 // Bu yeni bir 'x' değişkenidir, dışarıdakini etkilemez!
        fmt.Println(x) // 5 yazar
    }
    fmt.Println(x) // 10 yazar
}
```

**Mimari Not:** Shadowing, özellikle hata yönetiminde (`err` değişkeni) sıkça yapılan ve buglara yol açan bir durumdur. Kodun okunabilirliğini düşürdüğü için dikkatli kullanılmalıdır.

---

### Özet Karşılaştırma

| **Özellik** | **C# (ASP.NET)** | **Go** | **Mimari Yorum** |
| --- | --- | --- | --- |
| **Global Değişken** | `public static int x;` | `var X int` | Go'da büyük harf `public` (exported) demektir. |
| **Local Değişken** | `var x = 10;` | `x := 10` | Go, local scope'ta kısalığı sever. |
| **Başlangıç Değeri** | `default(int)` (manuel gerekebilir) | Otomatik Zero Value | Go daha güvenli bir bellek başlangıcı sunar. |
| **Kullanılmayan Değişken** | Warning (Uyarı) verir. | **Error (Hata) verir.** | Go, kodda kullanılmayan değişken varsa **derlenmez**. Bu, kod kirliliğini (dead code) mimari düzeyde engeller. |
# Zero Values

Created: January 26, 2026 3:33 AM

Harika, şimdi Go'nun en güçlü ve güvenli özelliklerinden biri olan **Zero Values** kavramını, ASP.NET Core geçmişinle kıyaslayarak mimari açıdan inceleyelim. Ardından şu ana kadar öğrendiğimiz (Package yapısı, Import, `var` vs `:=` ve Zero Values) her şeyi kapsayan "Best Practice" odaklı bir örnek yapalım.

### Zero Values: Mimari Güvenlik Sigortası

C# tarafında (özellikle non-nullable reference types öncesinde) bir nesne oluşturduğunda, eğer initialize etmezsen değeri `null` olur. Bu da meşhur `NullReferenceException` hatasının ana kaynağıdır.

Go'da ise **"Uninitialized Variable" (Başlatılmamış Değişken) diye bir şey yoktur.** Bir değişken tanımladığın an, derleyici ona bellekte bir yer ayırır ve o tipin "Zero Value"su ile doldurur.

**Neden Önemli?**

1. **Bellek Güvenliği:** Rastgele bellek adreslerine (garbage values) erişimi engeller.
2. **Kod Sadelikleri:** Çoğu zaman `if x == nil` veya `if x == ""` kontrolü yapmana gerek kalmaz, çünkü varsayılan değer mantıklı bir başlangıç noktasıdır. Go'da yaygın bir söz vardır: **"Make the zero value useful."** (Sıfır değerini işe yarar hale getir).

### Zero Value Tablosu

| **Tip Ailesi** | **Zero Value** | **Mimari Not** |
| --- | --- | --- |
| **Numeric** (`int`, `float`, vb.) | `0` | Sayaçlar (counters) başlatılmadan direkt artırılabilir (`count++`). |
| **Boolean** | `false` | Feature flag'ler varsayılan olarak kapalı gelir. |
| **String** | `""` (Boş String) | **Dikkat:** Go'da string `nil` olamaz! Bu, string işlemlerinde null check yapma yükünü kaldırır. |
| **Pointers, Interfaces, Slices, Maps** | `nil` | Referans tipleri bellekte henüz bir yeri işaret etmez. |

---

### Konsolide Örnek: Teoriyi Pratiğe Dökmek

Şimdiye kadar öğrendiğimiz şu başlıkları tek bir kod dosyasında birleştirelim:

1. `package main` ve `import` yapısı.
2. `var` ile global değişken (State management simülasyonu).
3. `var` ile Zero Value kullanımı (Default değerlere güvenmek).
4. `:=` ile Short Declaration (Hızlı ve lokal tanımlama).
5. String'lerin `nil` olmaması.

Bu kodu bir API'nin request işleme mantığının çok basitleştirilmiş bir simülasyonu gibi düşünebilirsin.

Go

```go
package main

import "fmt"

// 1. GLOBAL SCOPE (Paket Seviyesi)
// Burada ':=' KULLANAMAYIZ. Zorunlu olarak 'var' kullanıyoruz.
// Bu değişkenler programın yaşam döngüsü boyunca erişilebilir (bu pakette).
var (
	appVersion  string = "1.0.0" // Explicit initialization
	totalRequests int             // Değer atamadık -> Zero Value: 0
)

func main() {
	// 2. LOCAL SCOPE (Fonksiyon İçi)
	
	// Senaryo: Bir API isteği geldiğini varsayalım.
	
	// A. 'var' Kullanımı (Zero Value İhtiyacı)
	// Henüz bir hata oluşmadı, varsayılan olarak 'boş string' olmasını istiyoruz.
	// C#'taki gibi 'string lastError = null' değil, "" olur.
	var lastError string 
	
	// Bir boolean flag. Varsayılanı 'false'tur.
	// "Henüz işlenmedi" anlamına gelir.
	var isProcessed bool

	fmt.Println("--- Başlangıç Durumu (Zero Values) ---")
	fmt.Printf("Toplam İstek: %d\n", totalRequests)
	fmt.Printf("Son Hata: '%s' (Burasi bos string)\n", lastError)
	fmt.Printf("İşlendi mi?: %v\n", isProcessed)

	// 3. ':=' Kullanımı (Short Declaration)
	// Fonksiyon içinde olduğumuz için ve hemen değer atadığımız için en idiomatic yol budur.
	// Tipini (string) belirtmemize gerek yok, Go anlar (Type Inference).
	requestID := "req-12345"
	clientIP := "192.168.1.1"

	// Mantıksal işlemler (Simülasyon)
	totalRequests++ // 0'dı, şimdi 1 oldu. (Init etmeye gerek kalmadı!)
	isProcessed = true
	
	fmt.Println("\n--- İşlem Sonrası ---")
	fmt.Printf("[%s] ID'li istek %s adresinden alındı.\n", requestID, clientIP)
	fmt.Printf("Güncel Toplam İstek: %d\n", totalRequests)
	fmt.Printf("İşlem Durumu: %v\n", isProcessed)
}
```

### Kodun Mimari Okuması:

1. **`totalRequests`**: Başlangıç değeri vermedik ama `totalRequests++` diyebildik. C#'ta local bir değişkende bunu yapsaydık (initialize etmeden) hata alırdık. Go'da `0` garantidir.
2. **`lastError`**: Ekrana yazdırırken hata vermedi, sadece boşluk bastı. Çünkü `string` bir değer tipidir (value semantic) ve varsayılanı boştur.
3. **Değişken Seçimi**: Global ayarlar için `var` bloğu, hızlı yerel tanımlar için `:=` kullandık. Kodun okunabilirliği ve niyeti açık.

---

Bu örnekle temel yapıyı kurmuş olduk. Şimdi veri tiplerinin derinliklerine inmeye hazırız.

**Sıradaki Adım:** Go'nun `int`, `float` gibi sayısal tiplerini inceleyip, özellikle backend tarafında veri tabanı ID'leri veya parasal işlemler için kritik olan **Type Conversion (Tip Dönüşümü)** konusuna geçelim mi? Go, C#'ın aksine otomatik (implicit) dönüşümleri hiç sevmez, sebebini anlamak önemlidir.
# 1.1 Function Basics: Sözdizimi, Parametreler ve C# Farkları

Created: February 1, 2026 2:37 AM

Harika bir yaklaşım. C# bilgin burada altın değerinde çünkü **"neyin ne olduğunu"** biliyorsun, sadece **"Go'da nasıl yapıldığını"** ve **"felsefi farkını"** öğreneceksin. Bu da öğrenme hızını 10 katına çıkarır.

Hadi **Bölüm 1.1: Function Basics** ile başlayalım.

---

# 1.1 Function Basics: Sözdizimi, Parametreler ve C# Farkları

C#'ta bir metot yazarken beynin şöyle çalışır:

`[Erişim Belirteci] [Dönüş Tipi] [İsim] ([Tip] [Parametre])`

Go'da ise bu sıralama tamamen değişir. Go'nun mottosu şudur: **"İsimler tiplerden daha önemlidir, o yüzden önce isim gelir."**

### 1. Anatomi ve Sözdizimi Karşılaştırması 🆚

Gel aynı toplama fonksiyonunu iki dilde yazalım.

**C#:**

C#

```go
// Erişim: public
// Dönüş: int
// Parametre: int a
public int Add(int a, int b) 
{
    return a + b;
}
```

**Go:**

Go

```go
// Keyword: func
// İsim: Add
// Parametre: a int (Önce isim, sonra tip!)
// Dönüş: int (En sonda)
func Add(a int, b int) int {
    return a + b
}
```

### 2. Kritik Farklar (Not Defterine Eklenecekler) 📝

### A. Erişim Belirteçleri (Public/Private)

C#'ta `public`, `private`, `protected` kelimeleri havada uçuşur.

Go'da **erişim belirteci (keyword) yoktur.** Her şey **İSİMDE** biter.

- **Büyük Harf (PascalCase):** `Add` -> **Public** (Diğer paketlerden erişilebilir. Buna "Exported" denir).
- **Küçük Harf (camelCase):** `add` -> **Private** (Sadece tanımlandığı paketin içinde kullanılabilir).

### B. Parametre Gruplama (Parameter Shortening) ✨

C#'ta parametrelerin tipleri aynı olsa bile tekrar tekrar yazmak zorundasın: `(int a, int b, int c)`.

Go'da ardışık parametreler aynı tipteyse, sadece sonuncusuna tipi yazman yeterlidir.

Go

```go
// Uzun Hali
func Hesapla(a int, b int, c int) int { ... }

// Go Hali (Kısa ve Temiz)
func Hesapla(a, b, c int) int { ... }
```

### C. Method Overloading (Aşırı Yükleme) YOKTUR! 🛑

C# geliştiricisi olarak en büyük şoku burada yaşayacaksın.

C#'ta:

- `Add(int a, int b)`
- `Add(double a, double b)`
    
    Aynı isimle iki metot yazabilirsin.
    

**Go'da bu yasaktır.** Bir pakette fonksiyon ismi benzersiz (unique) olmalıdır.

- `AddInt(a, b int)`
- `AddFloat(a, b float64)`
    
    şeklinde açıkça isimlendirmelisin. Go, "Büyü yapma, açık ol" der.
    

### D. Süslü Parantez Kuralı `{`

C#'ta `{` işaretini alt satıra koymayı sevenlerdensen (Allman style), Go seni üzer.

Go'da açılış parantezi `{` fonksiyonla **aynı satırda olmak zorundadır.** Yoksa kod derlenmez (Syntax Error).

---

### 🚀 Görev 1.1: "Selamlayıcı"

Şimdi C# kaslarını Go'ya çevirelim. Aşağıdaki C# metodunun Go karşılığını yazmanı istiyorum.

**C# Kodu:**

C#

```go
namespace Core 
{
    public class Greeter 
    {
        // 1. Dışarıya açık olsun (Public)
        // 2. Parametreler string olsun
        public string SayHello(string name, string surname) 
        {
            return "Merhaba " + name + " " + surname;
        }
    }
}
```

**İstenen Go Kodu:**

1. Fonksiyon ismini "Public" mantığına göre ayarla.
2. Parametreleri "Shortening" (Gruplama) tekniğiyle yaz.
3. `main` paketi altında basitçe tanımla (Class içine almana gerek yok, Go'da fonksiyonlar özgürdür).

Kodunu yaz, syntax'ın oturduğunu görelim, sonra 1.2'ye (Multiple Return Values) geçelim. Bekliyorum! ⌨️

---

### 

C#'ta metotların başına `public` yazmazsan varsayılan olarak `private` olur ya?

Go'da da fonksiyonun baş harfini **küçük** yazarsan (`greeting`), o fonksiyon `private` (unexported) olur.
# 1.2 Multiple Return Values

Created: February 1, 2026 3:01 AM

### 

---

# 1.2 Multiple Return Values (Çoklu Dönüş)

C# geliştiricisi olarak şu senaryoyu binlerce kez yaşamışsındır: Bir metottan hem **sonucu** hem de **başarılı mı?** bilgisini dönmek istersin.

**C#'taki Çile (Eski Yöntemler):**

1. **`out` parametresi:** `public bool TryParse(string s, out int result)` -> Okuması zordur, akışı bozar.
2. **`Tuple`:** `public (int, bool) Hesapla()` -> C# 7.0 ile geldi ama syntax hala biraz kalabalık.
3. **Class/Struct Wrapper:** `Result` objesi oluşturmak.

**Go'daki Çözüm:**

Go der ki: *"Neden fonksiyonlar sadece bir şey dönmek zorunda olsun ki?"*

Go'da fonksiyonlar, virgülle ayrılmış birden fazla değeri doğal olarak dönebilir.

### Sözdizimi (Syntax)

```go
package main

import "fmt"

// Dönüş tipi parantez içinde: (int, bool)
// Anlamı: "Ben sana bir int, bir de bool fırlatacağım."
func Hesapla(a int) (int, bool) {
    if a > 10 {
        // İki değer dönüyoruz
        return a * 2, true 
    }
    // Burada da iki değer dönmek ZORUNDAYIZ
    return 0, false 
}

func main() {
    // Dönüşleri iki değişkene yakalıyoruz
    sayi, basariliMi := Hesapla(15)

    if basariliMi {
        fmt.Println("Sonuç:", sayi)
    } else {
        fmt.Println("İşlem başarısız")
    }
}
```

### 🧠 C# Geliştiricisi İçin Kritik Notlar

1. **Sıra Önemlidir:** İlk dönülen değer ilk değişkene, ikinci ikinciye gider.
2. **Zorunluluk:** Fonksiyon `(int, bool)` dönüyorsa, `return` satırında mutlaka iki değer olmalıdır. `return 5` dersen hata alırsın.
3. **Blank Identifier (`_`):** Eğer dönen değerlerden biri işine yaramıyorsa, onu `_` ile çöpe atabilirsin.
    - `sonuc, _ := Hesapla(20)` (Başarılı mı kısmını yuttuk).

---

### 🚀 Görev 1.2: "Tam Sayı Bölme Operatörü"

Şimdi bu özelliği pekiştirmek için bir fonksiyon yazmanı istiyorum.

**Senaryo:** İki tam sayıyı bölen güvenli bir fonksiyon yazacağız.

1. Fonksiyonun adı `Divide` olsun (Public/Exported kuralına dikkat! 😉).
2. İki parametre alsın: `bolunen` (int), `bolen` (int).
3. **İki değer dönsün:**
    - `sonuc` (int): Bölme işleminin sonucu.
    - `durum` (string): Eğer işlem başarılıysa `"Başarılı"`, başarısızsa `"Sıfıra Bölünemez"` stringi dönsün.
4. **Logic:**
    - Eğer `bolen` 0 ise -> Sonuç `0`, Durum `"Sıfıra Bölünemez"` dön.
    - Değilse -> Bölme sonucunu ve `"Başarılı"` dön.
5. `main` içinde bunu hem `10/2` hem de `10/0` için çağırıp ekrana bas.

Bu görevle hem **Control Flow (if/else)** hem de **Multiple Return** kaslarını çalıştıracağız. Kodunu bekliyorum!
# Arrays

Created: January 27, 2026 1:27 AM

Harika bir tespit. C# dünyasındaki `List<T>` yapısının Go'da adı **Slice** (Dilim)'dır. Ama Go mimarisinde Slice'ları anlayabilmek için önce onların üzerine inşa edildiği **Array (Dizi)** kavramını, "bellek yönetimi" (memory layout) açısından anlamamız şart.

C# geliştiricilerini en çok şaşırtan mimari farkla başlayalım:

**C#'ta diziler (`int[]`) Referans Tipidir (Heap'te yaşar).**

**Go'da diziler (`[n]int`) Değer Tipidir (Value Type - Stack'te yaşayabilir).**

Bu farkın mimari sonuçlarını inceleyelim.

### 1. Array (Dizi): Katı ve Sabit Yapı

Go'da bir dizi tanımladığında, bellekte o boyut kadar yer **garanti** edilir ve bu boyut **değiştirilemez**.

### A. Boyut, Tipin Kendisidir (Type Identity)

C#'ta `int[3]` ve `int[4]` aynı tiptedir (`int[]`). Sadece uzunlukları (Length) farklıdır.

Go'da ise bunlar **tamamen farklı tiplerdir.** Birbirlerine atanamazlar.

Go

```go
var a [3]int // Tipi: [3]int
var b [4]int // Tipi: [4]int

// a = b // DERLEME HATASI! (Tipler uyuşmuyor)
```

**Mimari Neden:** Bu katılık, derleyicinin bellekte ne kadar yer ayıracağını milimetrik olarak bilmesini sağlar. Memory Allocation optimizasyonu için kritiktir.

### B. Value Semantics (Kopyalama Maliyeti) ⚠️

Backend performansını etkileyen en büyük fark budur.

- **C#:** Bir diziyi fonksiyona parametre geçersen, sadece bellekteki adres (referans) kopyalanır. Veri kopyalanmaz.
- **Go:** Bir diziyi fonksiyona parametre geçersen veya başka değişkene atarsan, **tüm dizi elemanları tek tek kopyalanır.**

Go

```go
func main() {
    original := [3]int{1, 2, 3}
    
    // Kopyası oluşur (Pass by Value)
    kopya := original 
    
    kopya[0] = 999
    
    fmt.Println(original[0]) // Çıktı: 1 (Değişmedi!)
    fmt.Println(kopya[0])    // Çıktı: 999
}
```

**Mimari Ders:** Eğer `[1000000]int` diye devasa bir dizin varsa ve bunu bir fonksiyona düz parametre olarak gönderirsen, RAM'de anlık olarak o kadar veri kopyalanır ve CPU kopyalama işlemiyle meşgul olur. Bu yüzden Go'da ham Array'ler nadiren fonksiyona gönderilir.

### 2. Tanımlama Yöntemleri

Go

```go
// Yöntem 1: Zero Value (Hepsini 0 yapar)
var arr [5]int 

// Yöntem 2: Değer Atayarak
arr2 := [3]int{10, 20, 30}

// Yöntem 3: Derleyici Saysın (...)
// Derleyici eleman sayısına bakar (3 tane) ve tipi [3]int yapar.
arr3 := [...]int{10, 20, 30}
```

### 3. Neden Array Var? (List varken?)

"Madem kopyalaması pahalı, boyutu sabit, neden kullanıyoruz?" diyeceksin.

1. **Slice'ların Temeli:** Birazdan göreceğimiz esnek `Slice` (List benzeri) yapısı, arka planda gizlice bir Array kullanır. Array olmadan Slice olamaz.
2. **Stack Allocation:** Array'ler boyutu bilindiği için Heap yerine Stack'te oluşturulabilir. Bu da Garbage Collector'a yük bindirmez (Zero Allocation).
3. **Predictability:** Bellek düzeni (Memory Layout) %100 tahmin edilebilirdir. CPU cache dostudur.

### Özet: C# vs Go Array

| **Özellik** | **C# (int[])** | **Go ([n]int)** | **Mimari Etki** |
| --- | --- | --- | --- |
| **Tip Türü** | Reference Type | **Value Type** | Go'da atama yapınca veri kopyalanır. |
| **Boyut** | Dinamik (Heap'te) | **Statik (Tipin parçası)** | `[3]int` ile `[4]int` farklı dünyalardır. |
| **Yerleşim** | Heap | Stack veya Struct içi | Go'da küçük diziler GC'yi yormaz.a |

---

# Ek Olarak

Kesinlikle çok doğru bir sezgi. Go'da Slice mekanizmasının "sihrini" çözmek için, arkasında dönen **Array (Dizi)** gerçeğini, özellikle bellek yönetimi (Memory Layout) açısından anlamak şarttır. Çünkü Slice, aslında kendi başına veri tutmaz; sadece alttaki bir Array'e açılan **penceredir**.

C# ve ASP.NET Core mimarisinden gelen biri için Array'leri anlamanın en iyi yolu, onları **C# Struct** gibi düşünmektir.

Gel, Array konusunu bir backend mühendisinin bilmesi gereken "Low-Level" detaylarla inceleyelim.

### 1. Bellek Düzeni: "Bitişik Bloklar" (Contiguous Memory)

Go'da `var arr [4]int` dediğin an, işletim sistemi RAM'de yan yana duran 4 tane tam sayı (4 x 8 byte = 32 byte) kadar yer ayırır.

- C#'ta diziler birer **Nesnedir** (Object Header, Type Pointer vb. meta verileri vardır).
- Go'da diziler **saf veridir** (Flat Memory). Başında veya sonunda gizli bir header yoktur.

Bu sayede CPU'nun **Cache Line** mekanizması çok verimli çalışır. Veriler bitişik olduğu için CPU hepsini tek seferde cache'e çekebilir.

### 2. Value Semantics: "Struct" Gibi Davranır

Bu madde, C# geliştiricileri için en büyük tuzaktır.

C#'ta: `int[] a = new int[3]; int[] b = a;`

Burada `b`, `a`'nın gösterdiği yere referans verir (Pointer kopyalanır). `b[0]` değişirse `a[0]` da değişir.

Go'da:

Array'ler **Değer Tipidir (Value Type).** Tıpkı C#'taki `struct` veya `int` gibi davranırlar.

Go

```go
package main

import "fmt"

func main() {
    // 1. Orijinal Dizi
    original := [3]int{1, 2, 3}

    // 2. Atama Yapma (Assignment)
    // DİKKAT: Burada tüm dizi (3 eleman da) bellekte yeni bir yere KOPYALANIR.
    copyArr := original 

    // 3. Kopyayı Değiştirme
    copyArr[0] = 999

    fmt.Println("Orijinal:", original) // [1 2 3] -> Değişmedi!
    fmt.Println("Kopya:   ", copyArr)  // [999 2 3]
}
```

**Mimari Sonuç:**

Eğer 1 MB'lık bir dizin varsa ve bunu `x := y` dersen veya bir fonksiyona gönderirsen, RAM'de 1 MB'lık yeni yer açılır ve kopyalanır. Bu yüzden Go'da büyük Array'ler asla doğrudan taşınmaz.

### 3. Pointer Kullanımı: C# Referans Davranışını Simüle Etmek

"Peki ben kopyalamak istemiyorum, C#'taki gibi orijinali değiştirmek istiyorum" dersen, Array'in kendisini değil, **adresini (Pointer)** göndermelisin.

Backend servislerinde performans için genelde bu yöntem kullanılır.

Go

```go
// Parametre olarak dizinin kendisini değil, Pointer'ını (*[3]int) alıyoruz.
// Böylece sadece 8 byte'lık adres kopyalanır, verinin kendisi değil.
func updateFirst(arr *[3]int) {
    // (*arr)[0] = 100 yazmak yerine Go bize sözdizimsel şeker sunar:
    arr[0] = 100 
}

func main() {
    data := [3]int{1, 2, 3}
    
    // '&' operatörü ile adresini gönderiyoruz
    updateFirst(&data) 
    
    fmt.Println(data) // [100 2 3] -> Orijinal veri değişti!
}
```

### 4. Çok Boyutlu Diziler (Multidimensional Arrays)

Matris işlemleri veya oyun haritaları gibi yapılar için kullanılır. C#'taki `int[,]` (Rectangular Array) mantığına benzer.

Bellekte yine hepsi peş peşe (linear) dizilir.

Go

```go
func main() {
    // 2 Satır, 3 Sütunlu bir matris
    var matrix [2][3]int
    
    matrix[0] = [3]int{10, 20, 30}
    matrix[1] = [3]int{40, 50, 60}

    // Veya kısa tanımlama:
    grid := [2][2]string{
        {"A", "B"},
        {"C", "D"},
    }
    
    fmt.Println(grid[1][0]) // "C"
}
```

### 5. Array Nerede Kullanılır? (Slice varken neden var?)

Slice varken ham Array kullanacağın yerler çok sınırlıdır ama mimari açıdan kritiktir:

1. **Sabit Boyutlu Veriler:**
    - **UUID:** Her zaman 16 byte'tır. `[16]byte` kullanmak Slice kullanmaktan çok daha performanslıdır (Heap allocation yapmaz).
    - **IPv4 Adresi:** Her zaman 4 byte'tır. `[4]byte`.
    - **Kriptografik Hash (SHA-256):** Her zaman 32 byte'tır. `[32]byte`.
2. **Zero Allocation (GC Dostu):**
    
    Array'ler fonksiyon içinde tanımlanırsa **Stack** üzerinde oluşturulabilir. Fonksiyon bitince Stack temizlenir, Garbage Collector devreye girmez. Slice ise genelde Heap'e çıkar ve GC yükü oluşturur. Yüksek performanslı sistemlerde (High Frequency Trading) bu yüzden Array tercih edilir.
    

### Özet: Array Mimarisi

| **Özellik** | **Açıklama** |
| --- | --- |
| **Kapasite** | Derleme zamanında (Compile-time) sabittir. Çalışırken büyüyüp küçülemez. |
| **Bellek** | Değer tipidir (Value Type). Atamalarda kopyalanır. |
| **Performans** | Stack'te yaşayabilir, CPU Cache dostudur. |
| **Tip Güvenliği** | `[5]int` ile `[10]int` farklı tiplerdir, birbirine dönüşmez. |

Array mantığı tam olarak oturduysa, Go'nun bu katı yapıyı esnetmek için icat ettiği ve Array'lerin üzerine bir katman olarak inşa ettiği **Slices** konusuna geçebiliriz. Slice'lar, Array'lerin pointer gücünü kullanarak C#'taki `List<T>` konforunu sunar.
# basically

Haklısın, teknik detaylarda boğulmadan doğrudan pratiğe ve sözdizimine (syntax) odaklanalım. İstediğin "öğrenmen gerekenler" formatında, en sade haliyle dönüşümleri anlatıyorum.

Bu dönüşümleri anlamak için şu 2 temel kuralı bilmen yeterli:

1. **Array (Dizi):** Boyutu sabittir, değişmez (Örn: `[3]int`).
2. **Slice (Dilim):** Boyutu esnektir, değişebilir (Örn: `[]int`).

İşte C# bildiğin için rahat kavrayacağın o dönüşümler:

---

### 1. Array'den Slice'a Dönüşüm (En Sık Kullanılan)

Go'da bir Array'i Slice'a çevirmek çok kolaydır. Genellikle API'lara veri gönderirken bunu yaparsın.

**Öğrenmen Gereken Operatör:** `[:]` (Tümünü seç)

Go

```go
func main() {
    // 1. Sabit boyutlu bir Array tanımlayalım
    var myArr [3]int = [3]int{10, 20, 30}

    // 2. Bunu Slice'a dönüştürelim
    // Dizi isminin sonuna [:] koyman yeterli.
    var mySlice []int = myArr[:]

    // Artık mySlice, myArr'a bakan bir penceredir.
}
```

- **Püf Nokta:** Bu işlem veriyi kopyalamaz. `mySlice` üzerinde bir değişiklik yaparsan, `myArr` da değişir (Referans mantığı).

---

### 2. Slice'tan Array'e Dönüşüm (Go 1.20 ve Sonrası)

Slice esnek olduğu için sabit bir Array'e sığdırmak eskiden zordu. Go 1.20 sürümüyle bu iş standartlaştı.

**Öğrenmen Gereken Kural:** Slice'ın uzunluğu, hedef Array'in uzunluğuna **tam eşit** olmalıdır.

Go

```go
func main() {
    // 1. Esnek bir Slice'ımız olsun
    var mySlice []int = []int{10, 20, 30}

    // 2. Bunu sabit boyutlu Array'e dönüştürelim
    // Doğrudan tip dönüşümü (casting) yapıyoruz.
    myArr := [3]int(mySlice)

    // Not: Eğer mySlice'ın eleman sayısı 3 değilse (eksik veya fazla),
    // program "panic" (hata) verir ve çalışmayı durdurur.
}
```

- **Püf Nokta:** Bu işlem (Go 1.20 yöntemi) veriyi **kopyalar**. Yani `myArr` üzerinde yaptığın değişiklik, `mySlice`'ı etkilemez. Bağ kopar.

---

### Özet: Ne Zaman Hangisi?

| **Dönüşüm Yönü** | **Kod Örneği** | **Ne Yapar?** |
| --- | --- | --- |
| **Array -> Slice** | `arr[:]` | Referans oluşturur (Hızlıdır, veri kopyalamaz). |
| **Slice -> Array** | `[3]int(slice)` | Veriyi kopyalar (Yepyeni, bağımsız bir dizi oluşturur). |

Şu an bu basit dönüşüm mantığı oturdu mu? Yoksa "neden Array yerine hep Slice kullanmıyoruz?" kısmına mı geçelim?
# Bool

Harika, Boolean konusunu ASP.NET Core geçmişine ve mimari bakış açına uygun, not alabileceğin temiz bir formatta toparladım.

Buradaki en kritik not; **Strictness (Katılık)** ve **JSON Serialization (Pointer kullanımı)** başlıklarıdır.

---

# 📌 Go Notları: Booleans (`bool`)

### 1. Temel Tanım ve Mimari Davranış

- **Tip:** `bool`
- **Değerler:** Sadece `true` veya `false`.
- **Zero Value:** `false` (Başlatılmamış bir boolean her zaman `false` kabul edilir).
- **Boyut:** Genellikle 1 byte (Derleyiciye ve mimariye göre değişebilir).

### 2. Tip Güvenliği ve Katılık (No Implicit Conversion)

C++ veya JavaScript'in aksine, Go'da **hiçbir şey** otomatik olarak boolean'a dönüşmez. Karşılaştırma sonucu kesinlikle `bool` olmalıdır.

Go

```go
// YANLIŞ (Derlenmez)
if 1 { ... } 
if "text" { ... }

// DOĞRU
if count > 0 { ... }
if name != "" { ... }
```

### 3. Logic: Short-Circuit Evaluation (Kısa Devre)

C#'taki gibi çalışır. Performans ve "Side Effect" yönetimi için kritiktir.

- **AND (`&&`):** İlk ifade `false` ise, ikinciye **bakılmaz** (çalıştırılmaz).
- **OR (`||`):** İlk ifade `true` ise, ikinciye **bakılmaz**.

Go

```go
// Eğer user nil ise, user.IsActive() çağrılmaz ve Panic (Null Ref) oluşmaz.
if user != nil && user.IsActive() { ... }
```

### 4. JSON Serialization ve `omitempty` Tuzağı (Kritik)

REST API tasarlarken en önemli mimari karar noktasıdır.

- **Sorun:** Struct tag'inde `omitempty` kullanılırsa, Go `Zero Value` (yani `false`) olan alanları JSON çıktısına **yazmaz**.
- **Risk:** `IsActive: false` gönderdiğinde API bunu hiç göndermemiş gibi davranır (`{}`). Alıcı taraf (Frontend) varsayılanı `true` sanabilir.

**Çözüm Tablosu:**

| **İhtiyaç** | **Kullanım Tipi** | **JSON Etiketi** | **Davranış** |
| --- | --- | --- | --- |
| `false` değeri hep gitsin | `bool` | `json:"isActive"` | `false` ise JSON'da görünür: `false` |
| Sadece `true` ise gitsin | `bool` | `json:"isActive,omitempty"` | `false` ise JSON'dan silinir. (Tuzak!) |
| **Tristate (True/False/Null)** | **`*bool`** | `json:"isActive,omitempty"` | `false` ise görünür (`false`), `nil` ise silinir. |

### 5. String Dönüşümü ve Parsing

Backend'de Query String veya Form Data okurken kullanılır.

- **String to Bool:** `strconv.ParseBool(str)`
    - Sadece "true"/"false" değil; `"1"`, `"t"`, `"T"`, `"TRUE"` hepsini kabul eder. API esnekliği sağlar.
- **Bool to Int:** Go'da `int(myBool)` diye bir cast **yoktur**. Elle yazmalısın:Go
    
    ```go
    func boolToInt(b bool) int {
        if b { return 1 }
        return 0
    }
    ```
    

### 6. İsimlendirme (Naming Convention)

- **Struct Fields:** `IsActive`, `HasPermission` (Okunabilirlik için C# stili kabul görür).
- **Methods (Getters):** `Is` eki **atılır**.
    - C#: `user.IsActive()`
    - Go: `user.Active()` (Daha minimalist).

---

Bu notları kaydettiysen, Go'nun C#'tan en keskin ayrıldığı ve mülakatların vazgeçilmezi olan **Strings** (Aslında String Değildirler) konusuna geçebiliriz. Hazır mısın?
# Comma-ok Idiom

Created: February 1, 2026 12:29 AM

Harika, en temelden ve "Under the Hood" (Kaputun altı) mantığıyla gidelim. Bu konu, Go'da hata yönetimi ve akış kontrolünün omurgasıdır.

C# geliştiricisi olarak senin kafanda şu soru var: **"Neden tek bir değer dönmek varken iki değer dönüyoruz?"**

Hadi bunu bir senaryo ile parçalayalım.

---

### 1. Problem: "Sıfırın Laneti" (The Curse of Zero Value)

Go'nun felsefesini hatırlayalım: **Go'da Exception (Hata Fırlatma) yoktur.**

C#'ta `dictionary["olmayan_key"]` dersen program patlar (Exception). Go'da ise patlamaz, o tipin **Sıfır Değerini (Zero Value)** döner.

**Senaryo:** Bir sınıfın not defterini tutuyoruz.

- Ahmet: 90 aldı.
- Mehmet: Sınavdan 0 aldı (Kötü geçti).
- **Veli:** Sınıfta yok (Kaydı yok).

Şimdi Veli'nin notunu soralım:

```go
notlar := map[string]int{
    "Ahmet": 90,
    "Mehmet": 0,
}

veliNotu := notlar["Veli"] // Veli yok. int'in sıfır değeri nedir? 0.
mehmetNotu := notlar["Mehmet"] // Mehmet var. Notu kaç? 0.
```

🛑 **KRİTİK SORUN:**

Hem Mehmet (Var ama 0 aldı) için, hem de Veli (Yok) için elimizde **0** sonucu var.

Ben sadece `0`'a bakarak **"Öğrenci mi başarısız?"** yoksa **"Öğrenci mi yok?"** sorusunun cevabını veremem. Bu bir **Belirsizlik (Ambiguity)** yaratır.

---

### 2. Çözüm: "Comma-ok" İdiyomu (Dedektif) 🕵️‍♂️

Go mühendisleri demiş ki: *"Haritadan bir şey istendiğinde, sadece değeri vermeyelim. Yanına bir de 'Bulundu mu?' bilgisini iliştirelim."*

Sözdizimi şöyledir:

```go
deger, ok := map[anahtar]
```

- **deger:** Aradığın verinin kendisi (Bulunamazsa Zero Value).
- **ok (bool):** Veri gerçekten orada mı? (`true` = Var, `false` = Yok).

*(Not: `ok` ismi bir gelenektir. İstersen `varMi`, `found`, `exists` de diyebilirsin ama %99 `ok` kullanılır).*

---

### 3. Kullanım Alanı 1: MAPS (Haritalar) 🗺️

En yaygın kullanım burasıdır. Az önceki örneği düzeltelim.

Go

```go
package main

import "fmt"

func main() {
    notlar := map[string]int{
        "Ahmet":  90,
        "Mehmet": 0, // Mehmet var ama notu sıfır
    }

    // Veli'yi soruyoruz
    // Go burada Exception fırlatmaz. İki değer döner.
    puan, kayitVarMi := notlar["Veli"]

    fmt.Println("Puan:", puan)       // Çıktı: 0
    fmt.Println("Kayıt Var mı:", kayitVarMi) // Çıktı: false

    // ANALİZ
    if kayitVarMi {
        fmt.Println("Veli'nin notu:", puan)
    } else {
        fmt.Println("Veli diye bir öğrenci yok! O '0' değeri yalan.")
    }
    
    // Mehmet'i soralım
    puan2, ok2 := notlar["Mehmet"]
    // puan2 -> 0 gelir.
    // ok2    -> true gelir. (Çünkü anahtar haritada var!)
}
```

---

### 4. Kullanım Alanı 2: TYPE ASSERTION (Tip Dönüşümü) 🎭

Burası C#'taki `boxing/unboxing` veya `as` operatörü gibidir.

Go'da `interface{}` tipi, her şeyi tutabilir (C#'taki `object` gibi). Ama içindeki veriyi kullanmak için onu orijinal haline çevirmemiz gerekir.

**Sorun:** Eğer içindeki string olan bir kutuyu, zorla int yapmaya çalışırsan program **PANIC** (Crash) olur.

Go

```go
var kutu interface{} = "Merhaba"

// Hatalı Yöntem (Unsafe Cast)
// sayi := kutu.(int) // PANIC! "String'i int yapamazsın" diye program çöker.
```

**Çözüm (Comma-ok):**

"Bunu int yapmayı dene, yapamazsan programı patlatma, bana `false` dön."

Go

```go
var kutu interface{} = "Merhaba"

// Güvenli Yöntem
// deger, ok := degisken.(HedefTip)
sayi, ok := kutu.(int)

if ok {
    fmt.Println("Başarılı! Sayı:", sayi)
} else {
    fmt.Println("Hata! Bu kutunun içinde int yok.") 
    // Program çökmedi, 'sayi' değişkeni 0 oldu, 'ok' false oldu.
}
```

---

### 5. BEST PRACTICE: `If` ile Birleştirme (Pro Seviye) 🚀

Az önce öğrendiğin **"If Statement Initialization"** (noktalı virgüllü if) konusuyla `comma-ok` birleşince ortaya Go'nun en ikonik kodu çıkar.

Backend kodlarında bunu binlerce kez yazacaksın.

**Uzun (Amatör) Yol:**

```go
deger, ok := configs["db_port"] // Değişken scope dışına taştı
if ok {
    Connect(deger)
}
```

**Kısa (Pro) Yol:**

```go
// "db_port" anahtarını iste; eğer (;) ok true ise bloğa gir.
if port, ok := configs["db_port"]; ok {
    // 'port' sadece burada yaşıyor.
    fmt.Println("Veritabanına bağlanılıyor:", port)
} else {
    fmt.Println("Config dosyasında port eksik!")
}
```

---

### Özet: Comma-ok Nedir?

1. **Dedektiftir:** "Sadece sonuca bakma, sonucun *gerçek* olup olmadığına bak" der.
2. **Hata Önleyicidir:** Programın patlamasını (Panic) veya yanlış veriyle (Zero Value) işlem yapmasını engeller.
3. **İki Yer:** %90 **Map**'lerde, %10 **Interface** dönüşümlerinde (Type Assertion) kullanılır.

Bu mantık oturduysa, seninle **Döngüler (Loops)** ve **Retry Mechanism** görevine dönebiliriz. Durum nedir? 🕵️‍♂️
