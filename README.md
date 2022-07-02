# Go-Network-P2P-API

noise , merkezi olmayan uygulamalar ve Go'da yazılmış kriptografik protokoller için üzerinde düşünülmüş, kullanımı kolay bir P2P ağ yığınıdır.

Gürültü , az miktarda iyi test edilmiş, üretim düzeyinde bağımlılıklar kullanılarak çok sayıda cihazda minimum, sağlam, geliştirici dostu, performanslı, güvenli ve platformlar arası olacak şekilde yapılmıştır.


# Ornek - 1

User-2 Postman
http://localhost:8000/Al

body, x-www-from-urlencoded 

  Key      Value
mesaj      Selam

-------------------------------------------------

User-1 Postman
http://localhost:8080/Gonder

body, x-www-from-urlencoded 

  Key      Value
mesaj      Selam
adres      :50561
# ornek - 2

User-2 Postman
http://localhost:8080/Al

body, x-www-from-urlencoded 

  Key      Value
mesaj      Selam

-------------------------------------------------

User-1 Postman
http://localhost:8000/Gonder

body, x-www-from-urlencoded 

  Key      Value
mesaj      Selam
adres      :50561
