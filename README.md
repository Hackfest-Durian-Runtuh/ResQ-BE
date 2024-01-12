# ResQ-BE

## Description

Sistem Pemanggilan Layanan Darurat Terpadu ini dirancang untuk mempercepat dan meningkatkan efisiensi penanganan gawat darurat di berbagai negara berkembang seperti Indonesia. Sistem ini memungkinkan pengguna untuk dapat mengakses nomor berbagai unit layanan darurat (ambulans, pemadam kebakaran, bantuan SAR, dan kepolisian) di setiap daerah.

Aplikasi juga memungkinkan pengguna menyimpan informasi pribadi serta rekam medis yang telah diisi sebelumnya Pre-Filled Information. Informasi tersebut akan secara otomatis terkirim saat panggilan darurat dilakukan tanpa perlu penyampaian secara berulang.

Selain itu, terdapat fitur Live Tracking sehingga pengguna dan petugas lapangan dapat saling terhubung dan mendapatkan informasi update secara berkala terkait lokasi dan status penanganan.

## Requirements

```
- Golang
- MySQL
- WebSMS API
- Firebase Cloud Messaging
```

## Getting Started

### Install dependencies
```bash
$ go mod download
```

### Complete the .env file
```bash
$ cp .env.example .env
```

### Run the server
```bash
$ go run main.go
```

### Build the server
```bash
$ go build
```

### Run the server
```bash
$ ./resq-be
```

## API Documentation

Please refer to the [docs](https://github.com/Hackfest-Durian-Runtuh/ResQ-BE/tree/main/docs) folder for the API documentation.


## Authors

- [Aditya Rizky Ramadhan](https://github.com/adityarizkyramadhan)


