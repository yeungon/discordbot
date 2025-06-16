-- INSERT
-- name: CreateStudent :one
INSERT INTO students (
    name, student_code, gender, dob, dob_format, class, class_code,
    ethnic, national_id, phone, email, province, address, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14
) RETURNING *;


-- SELECT
-- name: GetStudentByID :one
SELECT * FROM students WHERE id = $1;

-- name: GetStudentByStudentCode :one
SELECT * FROM students WHERE student_code = $1;

-- name: GetStudentByName :one
SELECT * FROM students WHERE name = $1;

-- name: ListStudents :many
SELECT * FROM students ORDER BY id;

-- name: ListStudentsPaginated :many
SELECT * FROM students ORDER BY id LIMIT $1 OFFSET $2;

-- name: SearchStudents :many
SELECT * FROM students WHERE search_vector @@ plainto_tsquery('simple', $1) ORDER BY id;

-- name: SearchStudentsByPhrase :many
SELECT * FROM students
WHERE search_vector @@ phraseto_tsquery('simple', $1) ORDER BY id;

-- UPDATE
-- name: UpdateStudent :exec
UPDATE students
SET
    name = $2,
    student_code = $3,
    gender = $4,
    dob = $5,
    dob_format = $6,
    class = $7,
    class_code = $8,
    ethnic = $9,
    national_id = $10,
    phone = $11,
    email = $12,
    province = $13,
    address = $14,
    notes = $15
WHERE id = $1;

-- DELETE
-- name: DeleteStudent :exec
DELETE FROM students WHERE id = $1;

-- name: SearchStudentsFilteredPaginated :many
SELECT *
FROM students
WHERE
  ($1::text IS NULL OR class = $1) AND
  ($2::text IS NULL OR gender = $2) AND
  ($3::text IS NULL OR search_vector @@ plainto_tsquery('simple', $3))
ORDER BY id
LIMIT $4 OFFSET $5;
