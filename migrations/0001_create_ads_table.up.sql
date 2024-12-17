CREATE TABLE IF NOT EXISTS ads (
                                   id SERIAL PRIMARY KEY,
                                   title TEXT NOT NULL,
                                   description TEXT NOT NULL,
                                   price REAL NOT NULL
);