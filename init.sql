CREATE TABLE postgres.public.Position (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Salary NUMERIC(10, 2) NOT NULL
);

CREATE TABLE postgres.public.Employee (
    ID SERIAL PRIMARY KEY,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    PositionID INT REFERENCES Position(ID)
);
