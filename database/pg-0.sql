DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS coordinate;

CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	email VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS coordinate (
	user_id INTEGER NOT NULL REFERENCES users(user_id),
	id SERIAL PRIMARY KEY,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL,
	altitude REAL NOT NULL,
	created_at TIMESTAMPTZ NOT NULL default current_timestamp
);

CREATE INDEX coord_user_index ON coordinate(user_id);
