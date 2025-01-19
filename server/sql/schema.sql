CREATE TABLE exchange_rate (
    id INTEGER PRIMARY KEY,
    code TEXT NOT NULL,
    codein TEXT NOT NULL,
    name TEXT NOT NULL,
    high REAL NOT NULL,
    low REAL NOT NULL,
    varBid REAL NOT NULL,
    pctChange REAL NOT NULL,
    bid REAL NOT NULL,
    ask REAL NOT NULL,
    timestamp INTEGER NOT NULL,
    create_date TEXT NOT NULL
);