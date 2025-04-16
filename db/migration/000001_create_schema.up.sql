CREATE TABLE category(
    name text PRIMARY KEY
);

CREATE TABLE txn(
    id integer PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    amount integer NOT NULL,
    category text NOT NULL REFERENCES category(name),
    description text
);

INSERT INTO category
    VALUES ('Grocery'),
('Fruits and Vegetables'),
('Apparels'),
('Home Needs'),
('Toiletries'),
('Pet Food'),
('Doctor'),
('Veterinarian'), 
('Electricity'),
('Maintenance'),
('Maid'),
('Carwash'),
('Travel'),
('Transportation'),
('Petrol/Diesel'),
('Gas Cylinder'),
('Flowers'),
('Miscellaneous'),
('Donation'),
('Medicines');

