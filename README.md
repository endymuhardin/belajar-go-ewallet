# Aplikasi e-Wallet dengan Go Lang #

## Fitur aplikasi ##

* Melihat posisi saldo (via web dan API)
* Topup saldo (via API)
* Transaksi purchase/payment (via API)

## Tech Stack ##

* Go 1.23.2

    * [Package net/http](https://pkg.go.dev/net/http)
    * [Package html/template](https://pkg.go.dev/html/template)

* PostgreSQL 16

    * [PostgreSQL Driver](https://github.com/jackc/pgx)

## Setup Database ##

Jalankan database dengan docker compose

```
docker compose up
```

## Migrasi Skema dengan Flyway ##

1. Download database driver ke dalam folder `flyway/drivers`

    ```
    wget -c "https://jdbc.postgresql.org/download/postgresql-42.7.4.jar" -P flyway/drivers/
    ```


2. Jalankan migrasi database untuk membuat tabel-tabel

    ```
    docker run --rm --network host -v ${PWD}/flyway:/flyway/project flyway/flyway migrate -workingDirectory="project"
    ```


## Migrasi dengan go-migrate ##

```
docker run -v ${PWD}/go-migrate:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://ewallet:ewallet123@localhost:5432/ewallet-db?sslmode=disable" up
```

## Akses ke database ##

1. Connect ke database

    ```
    docker exec -it belajar-go-ewallet-db-ewallet-1 psql -U ewallet -d ewallet-db
    ```

2. Melihat skema database

    ```
    ewallet-db=# \d
                List of relations
    Schema |        Name        | Type  |  Owner  
    --------+--------------------+-------+---------
    public | customer           | table | ewallet
    public | schema_migrations  | table | ewallet
    public | wallet             | table | ewallet
    public | wallet_transaction | table | ewallet
    (4 rows)
    ```

3. Melihat isi tabel

    ```
    ewallet-db=# select * from schema_migrations ;
    version   | dirty 
    ------------+-------
    2024110701 | f
    (1 row)
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

## To Do List ##

* Loading config dari external file dan OS Env
* Log debug : supaya logging statement bisa dimatikan via config
* Proper error handling

## Referensi ##

* [Cara Membuat HTTP Service di Golang](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
* [Video : Cara Membuat HTTP Service di Golang](https://www.youtube.com/watch?v=0bmiwsv6LaA)
* [Go HTTP Standard Library](https://www.youtube.com/watch?v=H7tbjKFSg58)
* [How to structure API project](https://www.youtube.com/watch?v=EqniGcAijDI)