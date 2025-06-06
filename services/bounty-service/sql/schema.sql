CREATE TABLE bounties (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    points INT NOT NULL
);