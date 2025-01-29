-- +goose Up

-- Clients
INSERT INTO client (email) VALUES ('brent58@yahoo.com');
INSERT INTO client (email) VALUES ('vrodriguez@perez.org');
INSERT INTO client (email) VALUES ('mikayla87@yahoo.com');
INSERT INTO client (email) VALUES ('millerwillie@yahoo.com');
INSERT INTO client (email) VALUES ('lwright@yahoo.com');
INSERT INTO client (email) VALUES ('knightkatherine@gmail.com');
INSERT INTO client (email) VALUES ('brooksanthony@gmail.com');
INSERT INTO client (email) VALUES ('jamesmichael@gmail.com');
INSERT INTO client (email) VALUES ('scottkeith@yahoo.com');
INSERT INTO client (email) VALUES ('cdavidson@gmail.com');

-- Parent Categories
INSERT INTO category (name, parent_id) VALUES ('Electronics', NULL);
INSERT INTO category (name, parent_id) VALUES ('Furniture', NULL);
INSERT INTO category (name, parent_id) VALUES ('Clothing', NULL);
INSERT INTO category (name, parent_id) VALUES ('Kitchenware', NULL);
INSERT INTO category (name, parent_id) VALUES ('Sports Equipment', NULL);
-- Subcategories with Parents
INSERT INTO category (name, parent_id) VALUES ('Smartphones', 1); -- Electronics
INSERT INTO category (name, parent_id) VALUES ('Laptops', 1); -- Electronics
INSERT INTO category (name, parent_id) VALUES ('Tables', 2); -- Furniture
INSERT INTO category (name, parent_id) VALUES ('Chairs', 2); -- Furniture
INSERT INTO category (name, parent_id) VALUES ('Men''s Clothing', 3); -- Clothing
INSERT INTO category (name, parent_id) VALUES ('Women''s Clothing', 3); -- Clothing
INSERT INTO category (name, parent_id) VALUES ('Cookware', 4); -- Kitchenware
INSERT INTO category (name, parent_id) VALUES ('Utensils', 4); -- Kitchenware
INSERT INTO category (name, parent_id) VALUES ('Fitness Equipment', 5); -- Sports Equipment
INSERT INTO category (name, parent_id) VALUES ('Outdoor Gear', 5); -- Sports Equipment

-- Electronics
INSERT INTO product (category_id, name, description) VALUES (6, 'iPhone 14', 'The latest smartphone with advanced features.'); -- Smartphones
INSERT INTO product (category_id, name, description) VALUES (6, 'Samsung Galaxy S22', 'A powerful Android smartphone.'); -- Smartphones
INSERT INTO product (category_id, name, description) VALUES (7, 'MacBook Pro', 'A high-performance laptop for professionals.'); -- Laptops
INSERT INTO product (category_id, name, description) VALUES (7, 'Dell XPS 15', 'A sleek and powerful laptop.'); -- Laptops
-- Furniture
INSERT INTO product (category_id, name, description) VALUES (8, 'Dining Table', 'A sturdy wooden dining table for six people.'); -- Tables
INSERT INTO product (category_id, name, description) VALUES (8, 'Coffee Table', 'A modern glass coffee table.'); -- Tables
INSERT INTO product (category_id, name, description) VALUES (9, 'Office Chair', 'An ergonomic office chair for maximum comfort.'); -- Chairs
INSERT INTO product (category_id, name, description) VALUES (9, 'Recliner', 'A luxurious recliner for your living room.'); -- Chairs
-- Clothing
INSERT INTO product (category_id, name, description) VALUES (10, 'Men''s T-Shirt', 'A comfortable cotton t-shirt available in various sizes.'); -- Men's Clothing
INSERT INTO product (category_id, name, description) VALUES (10, 'Men''s Jeans', 'Classic blue denim jeans.'); -- Men's Clothing
INSERT INTO product (category_id, name, description) VALUES (11, 'Women''s Dress', 'A stylish evening dress.'); -- Women's Clothing
INSERT INTO product (category_id, name, description) VALUES (11, 'Women''s Jacket', 'A warm winter jacket.'); -- Women's Clothing
-- Kitchenware
INSERT INTO product (category_id, name, description) VALUES (12, 'Non-Stick Frying Pan', 'A durable non-stick frying pan for everyday cooking.'); -- Cookware
INSERT INTO product (category_id, name, description) VALUES (12, 'Stainless Steel Pot', 'A large stainless steel pot for soups and stews.'); -- Cookware
INSERT INTO product (category_id, name, description) VALUES (13, 'Cutlery Set', 'A 16-piece stainless steel cutlery set.'); -- Utensils
INSERT INTO product (category_id, name, description) VALUES (13, 'Spatula', 'A heat-resistant silicone spatula.'); -- Utensils
-- Sports Equipment
INSERT INTO product (category_id, name, description) VALUES (14, 'Treadmill', 'A high-quality treadmill for indoor running.'); -- Fitness Equipment
INSERT INTO product (category_id, name, description) VALUES (14, 'Dumbbell Set', 'An adjustable dumbbell set for strength training.'); -- Fitness Equipment
INSERT INTO product (category_id, name, description) VALUES (15, 'Tent', 'A waterproof camping tent for 4 people.'); -- Outdoor Gear
INSERT INTO product (category_id, name, description) VALUES (15, 'Sleeping Bag', 'A lightweight sleeping bag for outdoor adventures.'); -- Outdoor Gear

-- Buys
INSERT INTO buy (client_id, description) VALUES (3, 'Bought office supplies.');
INSERT INTO buy (client_id, description) VALUES (1, 'Purchased electronics.');
INSERT INTO buy (client_id, description) VALUES (5, 'Ordered home furniture.');
INSERT INTO buy (client_id, description) VALUES (8, 'Shopping for a party.');
INSERT INTO buy (client_id, description) VALUES (2, 'Bought a gift for a friend.');
INSERT INTO buy (client_id, description) VALUES (6, 'Bulk order for work.');
INSERT INTO buy (client_id, description) VALUES (4, 'Personal shopping.');
INSERT INTO buy (client_id, description) VALUES (9, 'Last-minute shopping.');
INSERT INTO buy (client_id, description) VALUES (7, 'Gift for a colleague.');
INSERT INTO buy (client_id, description) VALUES (10, 'Birthday shopping.');

-- Buy Items
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (1, 2, 2, 99.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (2, 3, 1, 49.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (3, 5, 3, 19.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (4, 7, 2, 9.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (5, 9, 1, 39.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (6, 1, 4, 12.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (7, 4, 2, 25.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (8, 6, 3, 15.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (9, 8, 2, 59.99);
INSERT INTO buy_item (buy_id, product_id, quantity, price) VALUES (10, 10, 1, 89.99);
