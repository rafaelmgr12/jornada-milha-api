-- name: CreateTestimonial :exec
INSERT INTO testimonials (id, name, testimonial) VALUES (?, ?, ?);

-- name: GetTestimonial :many
SELECT * FROM testimonials ORDER BY name;

-- name: UpdateTestimonial :exec
UPDATE testimonials SET name = ?, testimonial = ? WHERE id = ?;

-- name: DeleteTestimonial :exec
DELETE FROM testimonials WHERE id = ?;

-- name: CreateDestination :exec
INSERT INTO destinations (id, name, price, photo) VALUES (?, ?, ?, ?);

-- name: GetDestination :many
SELECT * FROM destinations ORDER BY name;

-- name: UpdateDestination :exec
UPDATE destinations SET name = ?, price = ?, photo = ? WHERE id = ?;

-- name: DeleteDestination :exec
DELETE FROM destinations WHERE id = ?;

-- name: GetDestinationsByName :many
SELECT id, name, price, photo
FROM destinations
WHERE name LIKE CONCAT('%', ?, '%');
