# GO Blog API

Merhaba bu proje Go lang ile geliştirdiğim basit bir blog API'sidir.

Bu projeyi geliştirmekteki temel amacım projede Go Lang'ın kendi http kütüphanesini kullanarak ve 3. taraf kütüphanelere olan bağımlılığı en aza indirmeye çalışarak bir Blog API'si geliştirmekti

## Özellikler

-   Kulanıcı kayıt ve giriş
-   Blog gönderisi oluşturma, okuma, güncelleme, silme...
-   Yorum Yapma
-   JWT ile kimlik doğrulama ve yetkilendirme

## Proje Gereksinimleri ve Kurulum

### Gereksinimler

-   Go 1.20 ve üzeri

### Kurulum

1. Bu repoyu kendi bilgisayarınıza indirin:

```
git clone https://github.com/ahmetilboga2004/go-blog.git
```

2. Proje klasörüne gidin:

```
cd go-blog
```

3. Gerekli paketleri yükleyin:

```
go mod tidy
```

4. Projeyi çalıştırın:

```
go run cmd/main.go
```

Ve herhangi bir problem olmazsa proje başarılı bir şekilde çalışacaktır

## Yapılacaklar

-   Test kodları yazılacak
