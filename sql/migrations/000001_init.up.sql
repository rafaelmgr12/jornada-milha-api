START TRANSACTION;
CREATE TABLE IF NOT EXISTS testimonials (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    testimonial TEXT NOT NULL
);
COMMIT;