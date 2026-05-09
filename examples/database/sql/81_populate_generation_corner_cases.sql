-- +goose Up

-- Small seed dataset for the generator corner-case schemas.

insert into bookstore_catalog.publishers (code, name, country_code, website, founded_on)
values
    ('no_starch', 'No Starch Press', 'US', 'https://nostarch.com', '1994-01-01'),
    ('manning', 'Manning Publications', 'US', 'https://www.manning.com', '1990-01-01'),
    ('pragprog', 'The Pragmatic Bookshelf', 'US', 'https://pragprog.com', '1999-01-01');

insert into bookstore_catalog.authors (author_id, display_name, legal_name, biography, metadata, born_on)
values
    ('00000000-0000-0000-0000-000000000101', 'Jon Bodner', 'Jon Bodner', 'Writes about Go.', '{"topics":["go","backend"]}', '1975-01-01'),
    ('00000000-0000-0000-0000-000000000102', 'Alex Edwards', 'Alex Edwards', 'Writes about Go web development.', '{"topics":["go","web"]}', '1985-01-01'),
    ('00000000-0000-0000-0000-000000000103', 'Thorsten Ball', 'Thorsten Ball', 'Writes about interpreters and compilers.', '{"topics":["compilers","go"]}', '1987-01-01');

insert into bookstore_catalog.books
    (book_id, publisher_code, isbn13, title, subtitle, description, price, weight_grams, rating, published_on, tags, attrs)
values
    (1001, 'no_starch', '9781718502895', 'Learning Go', 'An Idiomatic Approach to Real-World Go Programming', 'A practical Go book.', 49.99, 900, 4.80, '2024-01-10', array['go','backend'], '{"level":"intermediate"}'),
    (1002, 'manning', '9781617299759', 'Let us Go Further', null, 'A Go web development book.', 44.99, 700, 4.70, '2021-11-01', array['go','web'], '{"level":"intermediate"}'),
    (1003, 'pragprog', '9780984782857', 'Writing An Interpreter In Go', null, 'Interpreter implementation in Go.', 39.99, 600, 4.90, '2018-06-01', array['go','compilers'], '{"level":"advanced"}');

insert into bookstore_catalog.book_authors (book_id, author_id, contribution_order, role)
values
    (1001, '00000000-0000-0000-0000-000000000101', 1, 'author'),
    (1002, '00000000-0000-0000-0000-000000000102', 1, 'author'),
    (1003, '00000000-0000-0000-0000-000000000103', 1, 'author');

insert into bookstore_catalog.book_translations (book_id, language_code, translated_title, translated_by, released_on)
values
    (1001, 'kk', 'Go тілін үйрену', 'Localizer Team KZ', '2025-01-01'),
    (1001, 'ru', 'Изучаем Go', 'Localizer Team RU', '2025-02-01');

insert into bookstore_sales.customers (customer_id, email, full_name, marketing_opt_in)
values
    ('00000000-0000-0000-0000-000000000201', 'ada@example.test', 'Ada Lovelace', true),
    ('00000000-0000-0000-0000-000000000202', 'grace@example.test', 'Grace Hopper', false);

insert into bookstore_sales.orders (order_id, customer_id, status, placed_at, paid_at, comment)
values
    (2001, '00000000-0000-0000-0000-000000000201', 'paid', current_timestamp, current_timestamp, 'First generated-schema test order'),
    (2002, '00000000-0000-0000-0000-000000000202', 'placed', current_timestamp, null, 'Waiting for payment');

insert into bookstore_sales.order_lines (order_id, line_no, book_id, quantity, unit_price, discount_amount)
values
    (2001, 1, 1001, 1, 49.99, 0),
    (2001, 2, 1003, 1, 39.99, 5.00),
    (2002, 1, 1002, 2, 44.99, 0);

insert into bookstore_sales.discount_codes (code, description, percent_off, valid_period, max_uses)
values
    ('GO10', 'Ten percent off Go books', 10.00, '[2026-01-01,2027-01-01)'::daterange, 100),
    ('BOOKWORM', 'Small loyal-reader discount', 5.00, '[2026-01-01,)'::daterange, null);

insert into bookstore_inventory.warehouses (code, name, address, timezone)
values
    ('ALA_MAIN', 'Almaty Main Warehouse', '{"city":"Almaty","country":"KZ"}', 'Asia/Almaty'),
    ('AST_BACKUP', 'Astana Backup Warehouse', '{"city":"Astana","country":"KZ"}', 'Asia/Almaty');

insert into bookstore_inventory.stock_levels
    (warehouse_code, book_id, available_qty, reserved_qty, reorder_threshold, last_counted_at)
values
    ('ALA_MAIN', 1001, 12, 1, 3, current_timestamp),
    ('ALA_MAIN', 1002, 8, 2, 3, current_timestamp),
    ('AST_BACKUP', 1003, 5, 0, 2, current_timestamp);

insert into bookstore_inventory.stock_events (warehouse_code, book_id, delta_qty, reason, payload)
values
    ('ALA_MAIN', 1001, 10, 'initial_stock', '{"source":"manual"}'),
    ('ALA_MAIN', 1001, -1, 'sale', '{"order_id":2001}');

insert into bookstore_import.import_batches (source_name, batch_no, file_name, row_count, metadata)
values
    ('catalog_csv', 1, 'catalog-2026-01.csv', 3, '{"encoding":"utf-8"}');

insert into bookstore_import.import_errors (source_name, batch_no, row_no, column_name, message, raw_payload)
values
    ('catalog_csv', 1, 42, 'isbn13', 'invalid ISBN-13 checksum', '{"isbn13":"bad"}');

insert into bookstore_events.book_events (happened_at, book_id, event_type, payload)
values
    ('2026-01-15 12:00:00+00', 1001, 'viewed', '{"source":"homepage"}'),
    ('2026-02-10 12:00:00+00', 1002, 'purchased', '{"order_id":2002}');
