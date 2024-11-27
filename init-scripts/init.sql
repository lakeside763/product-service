CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create a sequence explicitly for serial_id
CREATE SEQUENCE IF NOT EXISTS product_seq START 1;

CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  serial_id INT NOT NULL DEFAULT nextval('product_seq'),
  sku TEXT GENERATED ALWAYS AS (LPAD(serial_id::TEXT, 10, '0')) STORED, -- Auto-generated
  name VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  price INT NOT NULL, -- Stored in cents (or whole number)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_serial_id ON products(serial_id);
CREATE INDEX idx_products_category ON products(category);
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
