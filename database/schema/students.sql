-- name: CreateStudentsTable :exec
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name TEXT,
    student_code TEXT,
    gender TEXT,
    dob TEXT,
    dob_format TEXT,
    class TEXT,
    class_code TEXT,
    ethnic TEXT,
    national_id TEXT,
    phone TEXT,
    email TEXT,
    province TEXT,
    address TEXT,
    notes TEXT,
    search_vector TSVECTOR
);
