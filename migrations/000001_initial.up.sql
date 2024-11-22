CREATE TABLE IF NOT EXISTS songs (
    id VARCHAR(255) PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date VARCHAR(10),
    text TEXT,
    link TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_group_name_song_name ON songs (group_name, song_name);