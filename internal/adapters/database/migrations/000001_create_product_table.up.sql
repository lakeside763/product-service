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