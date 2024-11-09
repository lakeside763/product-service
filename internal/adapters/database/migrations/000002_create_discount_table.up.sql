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
