-- +++ UP Migration
CREATE TABLE farmers (
    id SERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    national_id VARCHAR(100) NOT NULL,
    address TEXT,
    birth_date DATE,
    gender ENUM('MALE', 'FEMALE', 'OTHER'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- --- DOWN Migration
DROP TABLE IF EXISTS farmers;