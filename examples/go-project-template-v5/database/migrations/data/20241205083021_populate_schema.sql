-- +goose Up

-- Users
INSERT INTO users (email) VALUES ('brent58@yahoo.com');
INSERT INTO users (email) VALUES ('vrodriguez@perez.org');
INSERT INTO users (email) VALUES ('mikayla87@yahoo.com');
INSERT INTO users (email) VALUES ('millerwillie@yahoo.com');
INSERT INTO users (email) VALUES ('lwright@yahoo.com');
INSERT INTO users (email) VALUES ('knightkatherine@gmail.com');
INSERT INTO users (email) VALUES ('brooksanthony@gmail.com');
INSERT INTO users (email) VALUES ('jamesmichael@gmail.com');
INSERT INTO users (email) VALUES ('scottkeith@yahoo.com');
INSERT INTO users (email) VALUES ('cdavidson@gmail.com');

-- Customers
INSERT INTO customers (email) VALUES ('brent58@yahoo.com');
INSERT INTO customers (email) VALUES ('vrodriguez@perez.org');
INSERT INTO customers (email) VALUES ('mikayla87@yahoo.com');
INSERT INTO customers (email) VALUES ('millerwillie@yahoo.com');
INSERT INTO customers (email) VALUES ('lwright@yahoo.com');
INSERT INTO customers (email) VALUES ('knightkatherine@gmail.com');
INSERT INTO customers (email) VALUES ('brooksanthony@gmail.com');
INSERT INTO customers (email) VALUES ('jamesmichael@gmail.com');
INSERT INTO customers (email) VALUES ('scottkeith@yahoo.com');
INSERT INTO customers (email) VALUES ('cdavidson@gmail.com');

-- Parent Categories
INSERT INTO product_categories (name, parent_id) VALUES ('Electronics', NULL);
INSERT INTO product_categories (name, parent_id) VALUES ('Furniture', NULL);
INSERT INTO product_categories (name, parent_id) VALUES ('Clothing', NULL);
INSERT INTO product_categories (name, parent_id) VALUES ('Kitchenware', NULL);
INSERT INTO product_categories (name, parent_id) VALUES ('Sports Equipment', NULL);
INSERT INTO product_categories (name, parent_id) VALUES ('Smartphones', 1); -- Electronics
INSERT INTO product_categories (name, parent_id) VALUES ('Laptops', 1); -- Electronics
INSERT INTO product_categories (name, parent_id) VALUES ('Tables', 2); -- Furniture
INSERT INTO product_categories (name, parent_id) VALUES ('Chairs', 2); -- Furniture
INSERT INTO product_categories (name, parent_id) VALUES ('Men''s Clothing', 3); -- Clothing
INSERT INTO product_categories (name, parent_id) VALUES ('Women''s Clothing', 3); -- Clothing
INSERT INTO product_categories (name, parent_id) VALUES ('Cookware', 4); -- Kitchenware
INSERT INTO product_categories (name, parent_id) VALUES ('Utensils', 4); -- Kitchenware
INSERT INTO product_categories (name, parent_id) VALUES ('Fitness Equipment', 5); -- Sports Equipment
INSERT INTO product_categories (name, parent_id) VALUES ('Outdoor Gear', 5); -- Sports Equipment

INSERT INTO products (category_id, name, description) VALUES (6, 'iPhone 14', 'The latest smartphone with advanced features.'); -- Smartphones
INSERT INTO products (category_id, name, description) VALUES (6, 'Samsung Galaxy S22', 'A powerful Android smartphone.'); -- Smartphones
INSERT INTO products (category_id, name, description) VALUES (7, 'MacBook Pro', 'A high-performance laptop for professionals.'); -- Laptops
INSERT INTO products (category_id, name, description) VALUES (7, 'Dell XPS 15', 'A sleek and powerful laptop.'); -- Laptops
INSERT INTO products (category_id, name, description) VALUES (8, 'Dining Table', 'A sturdy wooden dining table for six people.'); -- Tables
INSERT INTO products (category_id, name, description) VALUES (8, 'Coffee Table', 'A modern glass coffee table.'); -- Tables
INSERT INTO products (category_id, name, description) VALUES (9, 'Office Chair', 'An ergonomic office chair for maximum comfort.'); -- Chairs
INSERT INTO products (category_id, name, description) VALUES (9, 'Recliner', 'A luxurious recliner for your living room.'); -- Chairs
INSERT INTO products (category_id, name, description) VALUES (10, 'Men''s T-Shirt', 'A comfortable cotton t-shirt available in various sizes.'); -- Men's Clothing
INSERT INTO products (category_id, name, description) VALUES (10, 'Men''s Jeans', 'Classic blue denim jeans.'); -- Men's Clothing
INSERT INTO products (category_id, name, description) VALUES (11, 'Women''s Dress', 'A stylish evening dress.'); -- Women's Clothing
INSERT INTO products (category_id, name, description) VALUES (11, 'Women''s Jacket', 'A warm winter jacket.'); -- Women's Clothing
INSERT INTO products (category_id, name, description) VALUES (12, 'Non-Stick Frying Pan', 'A durable non-stick frying pan for everyday cooking.'); -- Cookware
INSERT INTO products (category_id, name, description) VALUES (12, 'Stainless Steel Pot', 'A large stainless steel pot for soups and stews.'); -- Cookware
INSERT INTO products (category_id, name, description) VALUES (13, 'Cutlery Set', 'A 16-piece stainless steel cutlery set.'); -- Utensils
INSERT INTO products (category_id, name, description) VALUES (13, 'Spatula', 'A heat-resistant silicone spatula.'); -- Utensils
INSERT INTO products (category_id, name, description) VALUES (14, 'Treadmill', 'A high-quality treadmill for indoor running.'); -- Fitness Equipment
INSERT INTO products (category_id, name, description) VALUES (14, 'Dumbbell Set', 'An adjustable dumbbell set for strength training.'); -- Fitness Equipment
INSERT INTO products (category_id, name, description) VALUES (15, 'Tent', 'A waterproof camping tent for 4 people.'); -- Outdoor Gear
INSERT INTO products (category_id, name, description) VALUES (15, 'Sleeping Bag', 'A lightweight sleeping bag for outdoor adventures.'); -- Outdoor Gear

-- Buys
INSERT INTO purchases (customer_id, description) VALUES (3, 'Bought office supplies.');
INSERT INTO purchases (customer_id, description) VALUES (1, 'Purchased electronics.');
INSERT INTO purchases (customer_id, description) VALUES (5, 'Ordered home furniture.');
INSERT INTO purchases (customer_id, description) VALUES (8, 'Shopping for a party.');
INSERT INTO purchases (customer_id, description) VALUES (2, 'Bought a gift for a friend.');
INSERT INTO purchases (customer_id, description) VALUES (6, 'Bulk order for work.');
INSERT INTO purchases (customer_id, description) VALUES (4, 'Personal shopping.');
INSERT INTO purchases (customer_id, description) VALUES (9, 'Last-minute shopping.');
INSERT INTO purchases (customer_id, description) VALUES (7, 'Gift for a colleague.');
INSERT INTO purchases (customer_id, description) VALUES (10, 'Birthday shopping.');

-- Buy Items
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (1, 2, 2, 99.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (2, 3, 1, 49.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (3, 5, 3, 19.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (4, 7, 2, 9.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (5, 9, 1, 39.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (6, 1, 4, 12.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (7, 4, 2, 25.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (8, 6, 3, 15.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (9, 8, 2, 59.99);
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES (10, 10, 1, 89.99);

-- Product prices
insert into product_prices (product_id, product_price, product_price_period)
values
    (1, 1200.00,  '[2024-01-01, 2024-06-30)'),
    (1, 1100.00,  '[2024-07-01,)');

-- Currency
insert into public.currencies (currency_code, currency_value) values ('AFN', 'Afghani');
insert into public.currencies (currency_code, currency_value) values ('ALL', 'Lek');
insert into public.currencies (currency_code, currency_value) values ('AMD', 'Dram');
insert into public.currencies (currency_code, currency_value) values ('ARS', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('AUD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('AZN', 'Manat');
insert into public.currencies (currency_code, currency_value) values ('BAM', 'Marka');
insert into public.currencies (currency_code, currency_value) values ('BGN', 'Lev');
insert into public.currencies (currency_code, currency_value) values ('BIF', 'Franc');
insert into public.currencies (currency_code, currency_value) values ('BOB', 'Boliviano');
insert into public.currencies (currency_code, currency_value) values ('BRL', 'Real');
insert into public.currencies (currency_code, currency_value) values ('BSD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('BTN', 'Ngultrum');
insert into public.currencies (currency_code, currency_value) values ('BWP', 'Pula');
insert into public.currencies (currency_code, currency_value) values ('BYR', 'Ruble');
insert into public.currencies (currency_code, currency_value) values ('CAD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('CLP', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('CNY', 'Yuan Renminbi');
insert into public.currencies (currency_code, currency_value) values ('COP', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('CRC', 'Colon');
insert into public.currencies (currency_code, currency_value) values ('CUP', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('CZK', 'Koruna');
insert into public.currencies (currency_code, currency_value) values ('DKK', 'Krone');
insert into public.currencies (currency_code, currency_value) values ('DOP', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('EGP', 'Pound');
insert into public.currencies (currency_code, currency_value) values ('ETB', 'Birr');
insert into public.currencies (currency_code, currency_value) values ('EUR', 'Euro');
insert into public.currencies (currency_code, currency_value) values ('GBP', 'Pound');
insert into public.currencies (currency_code, currency_value) values ('GEL', 'Lari');
insert into public.currencies (currency_code, currency_value) values ('GMD', 'Dalasi');
insert into public.currencies (currency_code, currency_value) values ('GTQ', 'Quetzal');
insert into public.currencies (currency_code, currency_value) values ('HNL', 'Lempira');
insert into public.currencies (currency_code, currency_value) values ('HRK', 'Kuna');
insert into public.currencies (currency_code, currency_value) values ('HTG', 'Gourde');
insert into public.currencies (currency_code, currency_value) values ('HUF', 'Forint');
insert into public.currencies (currency_code, currency_value) values ('IDR', 'Rupiah');
insert into public.currencies (currency_code, currency_value) values ('ILS', 'Shekel');
insert into public.currencies (currency_code, currency_value) values ('IRR', 'Rial');
insert into public.currencies (currency_code, currency_value) values ('JMD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('JOD', 'Dinar');
insert into public.currencies (currency_code, currency_value) values ('JPY', 'Yen');
insert into public.currencies (currency_code, currency_value) values ('KES', 'Shilling');
insert into public.currencies (currency_code, currency_value) values ('KGS', 'Som');
insert into public.currencies (currency_code, currency_value) values ('KMF', 'Franc');
insert into public.currencies (currency_code, currency_value) values ('KRW', 'Won');
insert into public.currencies (currency_code, currency_value) values ('KZT', 'Tenge');
insert into public.currencies (currency_code, currency_value) values ('LKR', 'Rupee');
insert into public.currencies (currency_code, currency_value) values ('LRD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('LTL', 'Litas');
insert into public.currencies (currency_code, currency_value) values ('LYD', 'Dinar');
insert into public.currencies (currency_code, currency_value) values ('MAD', 'Dirham');
insert into public.currencies (currency_code, currency_value) values ('MDL', 'Leu');
insert into public.currencies (currency_code, currency_value) values ('MGA', 'Ariary');
insert into public.currencies (currency_code, currency_value) values ('MKD', 'Denar');
insert into public.currencies (currency_code, currency_value) values ('MMK', 'Kyat');
insert into public.currencies (currency_code, currency_value) values ('MNT', 'Tugrik');
insert into public.currencies (currency_code, currency_value) values ('MXN', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('MYR', 'Ringgit');
insert into public.currencies (currency_code, currency_value) values ('MZN', 'Metical');
insert into public.currencies (currency_code, currency_value) values ('NAD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('NGN', 'Naira');
insert into public.currencies (currency_code, currency_value) values ('NIO', 'Cordoba');
insert into public.currencies (currency_code, currency_value) values ('NOK', 'Krone');
insert into public.currencies (currency_code, currency_value) values ('NPR', 'Rupee');
insert into public.currencies (currency_code, currency_value) values ('NZD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('PAB', 'Balboa');
insert into public.currencies (currency_code, currency_value) values ('PEN', 'Sol');
insert into public.currencies (currency_code, currency_value) values ('PGK', 'Kina');
insert into public.currencies (currency_code, currency_value) values ('PHP', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('PKR', 'Rupee');
insert into public.currencies (currency_code, currency_value) values ('PLN', 'Zloty');
insert into public.currencies (currency_code, currency_value) values ('PYG', 'Guarani');
insert into public.currencies (currency_code, currency_value) values ('RSD', 'Dinar');
insert into public.currencies (currency_code, currency_value) values ('RUB', 'Ruble');
insert into public.currencies (currency_code, currency_value) values ('SAR', 'Rial');
insert into public.currencies (currency_code, currency_value) values ('SEK', 'Krona');
insert into public.currencies (currency_code, currency_value) values ('SLL', 'Leone');
insert into public.currencies (currency_code, currency_value) values ('SOS', 'Shilling');
insert into public.currencies (currency_code, currency_value) values ('SYP', 'Pound');
insert into public.currencies (currency_code, currency_value) values ('THB', 'Baht');
insert into public.currencies (currency_code, currency_value) values ('TJS', 'Somoni');
insert into public.currencies (currency_code, currency_value) values ('TMT', 'Manat');
insert into public.currencies (currency_code, currency_value) values ('TND', 'Dinar');
insert into public.currencies (currency_code, currency_value) values ('TZS', 'Shilling');
insert into public.currencies (currency_code, currency_value) values ('UAH', 'Hryvnia');
insert into public.currencies (currency_code, currency_value) values ('UGX', 'Shilling');
insert into public.currencies (currency_code, currency_value) values ('USD', 'Dollar');
insert into public.currencies (currency_code, currency_value) values ('UYU', 'Peso');
insert into public.currencies (currency_code, currency_value) values ('UZS', 'Som');
insert into public.currencies (currency_code, currency_value) values ('VEF', 'Bolivar');
insert into public.currencies (currency_code, currency_value) values ('VND', 'Dong');
insert into public.currencies (currency_code, currency_value) values ('XAF', 'Franc');
insert into public.currencies (currency_code, currency_value) values ('XOF', 'Franc');
insert into public.currencies (currency_code, currency_value) values ('YER', 'Rial');
insert into public.currencies (currency_code, currency_value) values ('ZAR', 'Rand');
insert into public.currencies (currency_code, currency_value) values ('ZWL', 'Dollar');
