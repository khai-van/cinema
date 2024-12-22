-- Create table for storing cinema details
CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    rows INT NOT NULL,
    columns INT NOT NULL
);

-- Create table for storing reserved seats
CREATE TABLE reserved_seats (
    id SERIAL PRIMARY KEY,
    cinema_id INT REFERENCES cinemas (id) ON DELETE CASCADE,
    row INT NOT NULL,
    col INT NOT NULL,
    status VARCHAR(10) CHECK (
        status IN ('reserved', 'canceled')
    ),
    UNIQUE (cinema_id, row, col)
);

-- Create an index on cinema_id, row, col for faster querying
CREATE INDEX idx_reserved_seats_cinema_row_col ON reserved_seats (cinema_id, row, col);