-- name: CreateTestimonial :exec
INSERT INTO testimonials (id, name, testimonial) VALUES (?, ?, ?);

-- name: GetTestimonial :many
SELECT * FROM testimonials ORDER BY name;

-- name: UpdateTestimonial :exec
UPDATE testimonials SET name = ?, testimonial = ? WHERE id = ?;

-- name: DeleteTestimonial :exec
DELETE FROM testimonials WHERE id = ?;

-- name: CreateDestination :exec
INSERT INTO destinations (id, photo1, photo2, name, meta, text_description, price) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetDestination :many
SELECT id, photo1, photo2, name, meta, text_description, price FROM destinations ORDER BY name;

-- name: UpdateDestination :exec
UPDATE destinations SET photo1 = ?, photo2 = ?, name = ?, meta = ?, text_description = ?, price = ? WHERE id = ?;

-- name: DeleteDestination :exec
DELETE FROM destinations WHERE id = ?;

-- name: GetDestinationsByName :many
SELECT id, photo1, photo2, name, meta, text_description, price
FROM destinations
WHERE name LIKE CONCAT('%', ?, '%');

-- name: GetDestinationById :one
SELECT id, photo1, photo2, name, meta, text_description, price FROM destinations WHERE id = ?;