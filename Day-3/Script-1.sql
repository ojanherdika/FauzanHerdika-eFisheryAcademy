--create table products
create table if not exists products(
	id_product serial not null,
	name char(50),
	primary key (id_product)
);
--insert values products
insert into products (name) values('pakan');
insert into products (name) values('vitamin');
insert into products (name) values('obat');
--delete values products
delete from products where id_product=3;
--update values products
update products set name ='makanan' where id_product=1;

--create table customers
create table if not exists customers(
	id_customer serial not null,
	customer_name char(50),
	primary key (id_customer)
);
--insert values customers
insert into customers (customer_name) values('ojan');
insert into customers (customer_name) values('hishna');
insert into customers (customer_name) values('haikal');
--delete values customers
delete from customers where id_customer=3;
--update values products
update customers set customer_name ='fauzan' where id_customer=1;

--create table orders
create table if not exists orders(
	id_order serial not null,
	order_date date not null,
	total float(50) not null,
	PRIMARY KEY(id_order)
);
--add foreign key to table orders
ALTER TABLE orders ADD COLUMN id_customer serial;
ALTER TABLE orders ADD COLUMN id_product serial;

ALTER TABLE orders 
   ADD CONSTRAINT fk_customer
   FOREIGN KEY (id_customer) 
   REFERENCES customers(id_customer);
ALTER TABLE orders 
   ADD CONSTRAINT fk_product
   FOREIGN KEY (id_product) 
   REFERENCES products(id_product);
--insert values orders
insert into orders (order_date,total,id_customer,id_product) values('12-12-22',1000,1,6);
insert into orders (order_date,total,id_customer,id_product) values('12-12-20',2000,2,5);
--delete values orders
delete from orders where id_order=2;
--update values orders
update orders set total =25000 where id_order=1;


