# Query Test

## 1. Query create table dan seed untuk table "transaksi"

```
CREATE TABLE public.transaksi (
	id int8 NOT NULL DEFAULT nextval('order_id_seq'::regclass),
	tanggal_order timestamp NULL,
	status_pelunasan varchar NULL,
	tanggal_pembayaran timestamp NULL,
	CONSTRAINT order_pk PRIMARY KEY (id)
);

INSERT INTO public.transaksi (tanggal_order,status_pelunasan,tanggal_pembayaran) VALUES
	 ('2020-12-01 11:30:00.000','lunas','2020-12-01 12:00:00.000'),
	 ('2020-12-02 10:30:00.000','pending',NULL);
```

## 2. Query create table dan seed untuk table "detail_transaksi"

```
CREATE TABLE public.detail_transaksi (
	id bigserial NOT NULL,
	id_transaksi bigserial NOT NULL,
	harga float8 NULL,
	jumlah int4 NULL,
	subtotal float8 NULL,
	CONSTRAINT transaksi_pk PRIMARY KEY (id),
	CONSTRAINT detail_transaksi_fk FOREIGN KEY (id_transaksi) REFERENCES public.transaksi(id)
);

INSERT INTO public.detail_transaksi (harga,jumlah,subtotal) VALUES
	 (10000.0,2,20000.0),
	 (5000.0,3,15000.0),
	 (10000.0,1,10000.0);
```

## 3. Query menampilkan data transaksi dengan total dan jumlah detail transaksi

```
SELECT a.*, sum(b.subtotal) as total, sum(b.jumlah) as jumlah_barang 
from transaksi a
left join detail_transaksi b on a.id = b.id_transaksi
group by a.id
```