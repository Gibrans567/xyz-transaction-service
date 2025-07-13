# **Layanan Transaksi \- PT XYZ Multifinance**

Proyek ini adalah sebuah backend service yang dibuat untuk menangani semua proses transaksi pembiayaan di PT XYZ Multifinance. Tujuannya sederhana: mengganti sistem lama yang monolitik dengan arsitektur microservices yang modern, cepat, dan aman.

Layanan ini fokus pada satu hal: **membuat transaksi baru secara aman**, termasuk memastikan limit kredit konsumen berkurang dengan benar bahkan saat ada banyak transaksi bersamaan (*concurrent-safe*).

## **✨ Fitur Utama & Pendekatan**

* **Clean Architecture**: Kode diatur dengan rapi ke dalam beberapa lapisan (domain, usecase, repository, handler). Hasilnya? Kode jadi lebih mudah dibaca, diuji, dan dikembangkan di kemudian hari.  
* **Aman dari Race Condition**: Proses pengurangan limit kredit dijamin aman menggunakan FOR UPDATE di level database. Tidak akan ada lagi cerita limit minus karena transaksi ganda.  
* **Siap untuk Dikembangkan**: Didesain sebagai microservice yang siap diintegrasikan dengan layanan lain (seperti layanan konsumen atau otentikasi).  
* **Siap Dibungkus (Dockerized)**: Sudah dilengkapi dengan Dockerfile dan docker-compose.yml agar mudah dijalankan di mana saja.

## **🛠️ Teknologi & Kebutuhan**

* **Bahasa**: Go (Golang)  
* **Database**: MySQL  
* **Alat Bantu**: Docker & Docker Compose  
* **Kebutuhan**: Anda hanya perlu [Go](https://go.dev/doc/install) dan [Docker](https://docs.docker.com/get-docker/) terinstal di komputer Anda.

## **🚀 Cara Menjalankan (Lokal)**

Hanya butuh 3 langkah untuk menjalankan proyek ini:

1. **Clone Proyek Ini**  
   git clone \[URL\_GITHUB\_ANDA\]  
   cd xyz-transaction-service

2. **Nyalakan Database** 🐳  
   docker-compose up \-d

   Perintah ini akan menyiapkan database MySQL Anda di dalam Docker.  
3. **Jalankan Aplikasi** ▶️  
   go run ./cmd/api

   Server Anda kini aktif di http://localhost:8080.

## **📖 API Endpoint**

### **Buat Transaksi Baru**

* **POST** /v1/transactions  
* **Headers**: Content-Type: application/json  
* **Body**:  
  {  
      "contract\_number": "KONTRAK-001",  
      "otr": 50000.0,  
      "admin\_fee": 2000.0,  
      "installment\_amount": 18000.0,  
      "interest\_amount": 4000.0,  
      "asset\_name": "Magic Jar Miyako",  
      "tenor": 3  
  }

* **Hasil Sukses (201 Created)**:  
  { "message": "Transaksi berhasil dibuat" }

* **Hasil Gagal (422 Unprocessable Entity)**:  
  { "error": "limit tidak mencukupi (sisa: 500000.00, butuh: 600000.00)" }

## **✅ Verifikasi Data di Database (Opsional)**

Setelah berhasil melakukan POST transaksi, Anda bisa mengecek langsung ke database untuk memastikan data benar-benar berubah.

1. **Masuk ke dalam container database**:  
   docker exec \-it xyz\_mysql\_db mysql \-u root \-p

   Saat diminta password, ketik: password  
2. **Jalankan query SQL**: Setelah masuk, jalankan perintah-perintah ini untuk melihat datanya.  
   \-- Pilih database yang akan digunakan  
   USE xyz\_db;

   \-- Cek apakah transaksi baru sudah tercatat  
   SELECT \* FROM transactions WHERE contract\_number \= 'KONTRAK-001';

   \-- Cek apakah limit konsumen sudah berkurang  
   \-- (Contoh: limit Budi (id=1) untuk tenor 3 bulan, awalnya 500rb)  
   \-- Hasilnya harus menjadi 450000  
   SELECT \* FROM customer\_limits WHERE customer\_id \= 1 AND tenor\_in\_months \= 3;

## **🧪 Testing**

Untuk menjalankan semua unit test, cukup jalankan:

go test ./...  
