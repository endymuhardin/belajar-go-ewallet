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

## Setup Database ##

1. Jalankan database dengan docker compose

    ```
    docker compose up
    ```

2. Download database driver ke dalam folder `flyway/drivers`

    ```
    wget -c "https://jdbc.postgresql.org/download/postgresql-42.7.4.jar" -P flyway/drivers/
    ```


3. Jalankan migrasi database untuk membuat tabel-tabel

    ```
    docker run --rm --network host -v ${PWD}/flyway:/flyway/project flyway/flyway migrate -workingDirectory="project"
    ```

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
* [How to structure API project](https://www.youtube.com/watch?v=EqniGcAijDI)