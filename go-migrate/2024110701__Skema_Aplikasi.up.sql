begin;

create table customer (
    id varchar(36),
    customer_name varchar(200) not null,
    email varchar(100) not null,
    mobile_phone varchar(50) not null,
    primary key (id)
);

create table wallet (
    id varchar(36),
    id_customer varchar(36) not null,
    wallet_name varchar(50) not null,
    amount decimal(19,2) not null,
    primary key (id),
    foreign key (id_customer) references customer(id)
);

create table wallet_transaction (
    id varchar(36),
    id_wallet varchar(36) not null,
    transaction_time timestamp not null,
    remarks varchar(255) not null,
    transaction_type varchar(30) not null,
    amount decimal(19,2) not null,
    primary key (id),
    foreign key (id_wallet) references wallet(id)
);

commit;