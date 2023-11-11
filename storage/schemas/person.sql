CREATE TABLE persons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE,
    firstname TEXT,
    lastname TEXT,
    verified INTEGER,
    joined INTEGER
)