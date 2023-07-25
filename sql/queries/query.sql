-- name: CreateTestimonial :exec
INSERT INTO testimonials (id, name, testimonial) VALUES (?, ?, ?);

-- name: GetTestimonial :many
SELECT * FROM testimonials ORDER BY name;

-- name UpdateTestimonial :exec
UPDATE testimonials SET name = ?, testimonial = ? WHERE id = ?;

-- name: DeleteTestimonial :exec
DELETE FROM testimonials WHERE id = ?;
