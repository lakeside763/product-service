CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  sku VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_price ON products(price);

CREATE TABLE discounts (
  id SERIAL PRIMARY KEY,
  sku VARCHAR(50),
  category VARCHAR(255),
  discount_percentage INT CHECK(discount_percentage BETWEEN 0 AND 100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_discounts_category ON discounts(category);
CREATE INDEX idx_discounts_sku ON discounts(sku);

INSERT INTO products (sku, name, category, price)
VALUES
    ('000001', 'BV Lean leather ankle boots', 'boots', 89000),
    ('000002', 'BV Lean leather ankle boots', 'boots', 99000),
    ('000003', 'Ashlington leather ankle boots', 'boots', 71000),
    ('000004', 'Naima embellished suede sandals', 'sandals', 79500),
    ('000005', 'Nathane leather sneakers', 'sneakers', 59000);

INSERT INTO discounts (sku, category, discount_percentage)
VALUES
    ('000001', NULL, 15), -- SKU-specific discount
    (NULL, 'boots', 30),  -- Category-specific discount
    (NULL, 'sandals', 20),
    (NULL, 'sneakers', 25),
    ('000003', NULL, 35);
