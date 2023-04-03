# User Management API
Proyek ini dibuat untuk menjawab Soal Coding Test.

Proyek ini dibuat menggunakan salah satu contoh dari [repository `awesome-compose`](https://github.com/docker/awesome-compose/) dengan stack `nginx-golang-mariadb` yang bisa anda lihat [di sini](https://github.com/docker/awesome-compose/tree/master/nginx-golang-mysql).

Teknologi yang digunakan pada Proyek ini adalah :
- Nginx : Web Server (proxy)
- MariaDB : SQL Database Server (MySQL)
- GoLang : Bahasa Pemrograman Backend (IDE)
- Docker : Container
- Git : Sistem Pengontrol Versi (Kode)
- Postman : Dokumentasi API

## Bagaimana Cara Menjalankan Proyek Ini?
Ada beberapa langkah mudah untuk menjalankan Proyek ini namun pastikan bahwa Komputer anda telah terinstall Docker dan Git.

### 1. Clone Proyek
```
$ mkdir PROJECT && cd PROJECT
$ git clone https://github.com/iFadli/test-crud-user-management-api
... # tunggu hingga selesai # ...
Cloning into 'test-user-management-api'...
remote: Enumerating objects: xxx, done.
remote: Counting objects: 100% (xxx/xxx), done.
remote: Compressing objects: 100% (xxx/xxx), done.
remote: Total xxx (delta xx), reused xxx (delta xx), pack-reused x
Receiving objects: 100% (xxx/xxx), xx.xx KiB | xxx.xx KiB/s, done.
Resolving deltas: 100% (xx/xx), done.
$ cd test-user-management-api
```


### 3. Jalankan Proyek ini dengan docker-compose
```
$ docker-compose up -d
... # tunggu hingga selesai # ...
[+] Running 3/3
 ⠿ Container test-user-management-api-db-1       Healthy                                                                                                                        4.2s
 ⠿ Container test-user-management-api-backend-1  Started                                                                                                                        4.4s
 ⠿ Container test-user-management-api-proxy-1    Started                                                                                                                        4.6s
```

### 4. Akses Proyek sesuai Kebutuhan
Pada pengaturan Docker Proyek ini, secara default akan meng-expose Port 2 Service yang digunakan; Yakni, Database (MariaDB : 3306) dan Web Server (Nginx : 8080).

Jika ingin menghubungkan Database dengan Tool Database Manager seperti DBeaver, anda dapat menyesuaikan konfirugasi dengan File `.env`.
#### !! DBeaver
Berikut langkah-langkah konfigurasi DBeaver :

>1. Buka `DBeaver`.
>2. Klik ikon `Connect to a Database` di Pojok-Kiri-Atas.
>3. Pada Kategori `Popular`, pilih `MariaDB` lalu klik Next.
>4. Isikan Konfigurasi sesuai dengan File `config.go` pada Proyek ini.
>5. Jika berhasil, akan ada Ikon centang hijau pada daftar Database di sebelah kiri.
#### !! Postman Collection
Selain pengaturan Database Manager dari luar Docker yang dapat mengakses ke service Database, di Proyek ini juga disematkan Collection Postman yang dapat anda Gunakan untuk mencoba API.

## Daftar API

Daftar API dapat dilihat dari Collection API Postman pada Repository path Backend.

Lakukan perintah berikut di Terminal untuk Menonaktifkan Proyek :
```
$ docker-compose down --volumes
```
