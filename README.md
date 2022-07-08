# Go-Network-P2P-API

<img width="1353" alt="Ekran Resmi 2022-07-08 17 50 31" src="https://user-images.githubusercontent.com/92402372/178016582-37cf75af-63ab-421f-8add-4d57eaa10fbd.png">


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
# Ornek - 2

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
