CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    body TEXT, 
    author_id INTEGER
);