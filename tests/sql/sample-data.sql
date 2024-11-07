insert into customer (id, customer_name, email, mobile_phone) values 
('c001', 'Customer 001', 'customer.001@yopmail.com', '08123456789001'),
('c002', 'Customer 002', 'customer.002@yopmail.com', '08123456789002'),
('c003', 'Customer 003', 'customer.003@yopmail.com', '08123456789003'),
('c004', 'Customer 004', 'customer.004@yopmail.com', '08123456789004'),
('c005', 'Customer 005', 'customer.005@yopmail.com', '08123456789005'),
('c006', 'Customer 006', 'customer.006@yopmail.com', '08123456789006');

insert into wallet (id, id_customer, wallet_name, amount) values 
('w101', 'c001', 'Wallet Customer 001', 100001),
('w201', 'c002', 'Wallet Customer 002', 200001),
('w301', 'c003', 'Wallet Customer 003', 300001),
('w401', 'c004', 'Wallet Customer 004', 400001),
('w501', 'c005', 'Wallet Customer 005', 500001),
('w601', 'c006', 'Wallet Customer 006', 600001);

insert into wallet_transaction (id, id_wallet, transaction_time, remarks, transaction_type, amount) values 
('wt001', 'w101', '2024-11-01 13:14:01', 'Saldo awal', 'DEPOSIT', 100001),
('wt002', 'w201', '2024-11-01 13:14:02', 'Saldo awal', 'DEPOSIT', 200001),
('wt003', 'w301', '2024-11-01 13:14:03', 'Saldo awal', 'DEPOSIT', 300001),
('wt004', 'w401', '2024-11-01 13:14:04', 'Saldo awal', 'DEPOSIT', 400001),
('wt005', 'w501', '2024-11-01 13:14:05', 'Saldo awal', 'DEPOSIT', 500001),
('wt006', 'w601', '2024-11-01 13:14:06', 'Saldo awal', 'DEPOSIT', 600001);