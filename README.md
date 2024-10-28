# Aplikasi e-Wallet dengan Go Lang #

## Fitur aplikasi ##

* Melihat posisi saldo (via web dan API)
* Topup saldo (via API)
* Transaksi purchase/payment (via API)

## Tech Stack ##

* Go 1.23.2

    * [Package net/http](https://pkg.go.dev/net/http)
    * [Package html/template](https://pkg.go.dev/html/template)
    * [Package database/sql](https://pkg.go.dev/database/sql)

* PostgreSQL 16

    * [PostgreSQL Driver](https://github.com/jackc/pgx)

## Menjalankan Aplikasi ##

```
go run .
```

## Membuat Aplikasi e-wallet ##

1. Inisialiasi go module

    ```
    go mod init github.com/endymuhardin/belajar-go-ewallet
    ```

2. Menyediakan http route untuk url berikut dengan dummy code :

    * `GET /api/wallet` : saldo terkini dalam format json
    * `POST /api/wallet/deposit` : untuk topup saldo
    * `POST /api/wallet/purchase` : untuk melakukan purchase
    * `POST /api/wallet/payment` : untuk melakukan payment
    * `GET /wallet` : info wallet berisi saldo dan last transaction dalam format HTML
    * `GET /wallet/purchase` : melihat riwayat pembelian dalam format HTML
    * `GET /wallet/payment` : melihat riwayat pembelian dalam format HTML

3. Membuat template HTML layout

4. Membuat screen html

5. Membuat kode program untuk akses database


## Referensi ##

* [Cara Membuat HTTP Service di Golang](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
* [Video : Cara Membuat HTTP Service di Golang](https://www.youtube.com/watch?v=0bmiwsv6LaA)
* [Go HTTP Standard Library](https://www.youtube.com/watch?v=H7tbjKFSg58)