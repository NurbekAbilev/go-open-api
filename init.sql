CREATE TABLE position (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    salary NUMERIC(10, 2) NOT NULL
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    position_id INT REFERENCES position(id)
);
