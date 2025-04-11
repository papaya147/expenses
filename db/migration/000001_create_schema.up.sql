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
    VALUES ('Food'),
('Transport'),
('Entertainment'),
('Clothes'),
('Other');

