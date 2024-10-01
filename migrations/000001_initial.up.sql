BEGIN;

CREATE TABLE IF NOT EXISTS song (
    group VARCHAR(300) NOT NULL,
    song_name VARCHAR(200) NOT NULL,
    release_date DATE,
    text TEXT,
    link, VARCHAR(300),
    PRIMARY KEY(group, song_name)
)

END;