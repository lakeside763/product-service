INSERT INTO products (name, category, price)
VALUES
    ('BV Lean leather ankle boots', 'boots', 89000),
    ('BV Lean leather ankle boots', 'boots', 99000),
    ('Ashlington leather ankle boots', 'boots', 71000),
    ('Naima embellished suede sandals', 'sandals', 79500),
    ('Nathane leather sneakers', 'sneakers', 59000);

INSERT INTO discounts (sku, category, discount_percentage)
VALUES
    ('0000000001', NULL, 15), -- SKU-specific discount
    (NULL, 'boots', 30),  -- Category-specific discount
    (NULL, 'sandals', 20),
    (NULL, 'sneakers', 25);

