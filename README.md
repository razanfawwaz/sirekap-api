# Data Pemilu 2024 - API

## Mengenai Projek Ini
API ini menyediakan data Pemilu 2024 yang dikumpulkan oleh [Data Pemilu](https://data-pemilu.vercel.app) dari Situs Resmi [Pemilu 2024 KPU](https://pemilu2024.kpu.go.id).

## Endpoint
Untuk menggunakan API ini dapat menggunakan endpoint berikut:

```
GET /api/v1/ppwp?limit=xx&offset=xx
```
Endpoint di atas akan menampilkan data dari TPS yang sudah memiliki data perhitungan suara. 

Agar dapat menggunakan data, jangan lupa untuk memakai query parameter `limit` dan pastikan untuk mengganti `xx` dengan jumlah data yang ingin ditampilkan.

## Cara Menjalankan
Clone repository ini dengan cara:
```
git clone https://github.com/razanfawwaz/sirekap-api
```

Kemudian install dependencies dengan cara:
```
go mod tidy
```

Setelah itu, jalankan aplikasi dengan cara:
```
go run main.go
```

## Environment Variable
Untuk menjalankan aplikasi ini, dibutuhkan environment variable berikut:

```
PORT=
DB_HOST =
DB_PORT =
DB_NAME =
DB_USERNAME =
DB_PASSWORD =
GIN_MODE
```

## Informasi Tambahan
API ini masih dalam tahap pengembangan dan bersifat open-source. Jika ingin berkontribusi, silahkan buat pull request.


