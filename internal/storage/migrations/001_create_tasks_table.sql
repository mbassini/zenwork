CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT DEFAULT '',
	priority TEXT DEFAULT 'unset',
	status TEXT DEFAULT 'pending',
	deadline DATETIME,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	time_spent INTEGER DEFAULT 0,
	is_tracking BOOLEAN DEFAULT 0,
	last_started DATETIME,
	finished_at DATETIME
);
