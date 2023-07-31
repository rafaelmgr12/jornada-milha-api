START TRANSACTION;
CREATE TABLE IF NOT EXISTS destinations (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    photo TEXT NOT NULL
);
COMMIT;