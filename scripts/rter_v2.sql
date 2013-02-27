-- MySQL rter v2
-- ===========
-- Run these commands to setup the MySQL databases for the rter v2 project

SET foreign_key_checks = 0;

DROP TABLE IF EXISTS roles;
CREATE TABLE IF NOT EXISTS roles (
	role VARCHAR(64) NOT NULL,
	permissions INT NOT NULL DEFAULT 0,
	PRIMARY KEY(role)
);

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(64) NOT NULL,
	password CHAR(128) NOT NULL,
	salt CHAR(16) NOT NULL,
	role VARCHAR(64),
	trust_level INT,
	create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id),
	FOREIGN KEY(role) REFERENCES roles (role) ON UPDATE CASCADE
);

DROP TABLE IF EXISTS user_directions;
CREATE TABLE IF NOT EXISTS user_directions (
	user_id INT NOT NULL,
	lock_user_id INT,
	command VARCHAR(64),
	heading DECIMAL(9, 6),
	lat DECIMAL(9, 6),
	lng DECIMAL(9, 6),
	update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY(user_id),
	FOREIGN KEY(user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

DROP TABLE IF EXISTS items;
CREATE TABLE IF NOT EXISTS items (
	id INT NOT NULL AUTO_INCREMENT,
	type VARCHAR(64) NOT NULL,
	start_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	stop_time TIMESTAMP,
	heading DECIMAL(9, 6),
	lat DECIMAL(9, 6),
	lng DECIMAL(9, 6),
	author_id INT NOT NULL,
	thumbnail_URI VARCHAR(2048),
	content_URI VARCHAR(2048),
	upload_URI VARCHAR(2048),
	PRIMARY KEY(id),
	FOREIGN KEY(author_id) REFERENCES users (id) ON UPDATE CASCADE
);

DROP TABLE IF EXISTS item_comments;
CREATE TABLE IF NOT EXISTS item_comments (
	id INT NOT NULL AUTO_INCREMENT,
	item_id INT NOT NULL,
	author_id INT NOT NULL,
	body TEXT NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(item_id) REFERENCES items (id) ON UPDATE CASCADE,
	FOREIGN KEY(author_id) REFERENCES users (id) ON UPDATE CASCADE
);

DROP TABLE IF EXISTS taxonomy;
CREATE TABLE IF NOT EXISTS taxonomy (
	id INT NOT NULL,
	author_id INT NOT NULL,
	create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	automated TINYINT(1) NOT NULL DEFAULT 0,
	term VARCHAR(256) NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(author_id) REFERENCES users (id) ON UPDATE CASCADE
);

DROP TABLE IF EXISTS taxonomy_rankings;
CREATE TABLE IF NOT EXISTS taxonomy_rankings (
	id INT NOT NULL,
	taxonomy_id INT NOT NULL,
	update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	ranking TEXT NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(taxonomy_id) REFERENCES taxonomy (id) ON UPDATE CASCADE
);

-- DROP TABLE IF EXISTS taxonomy_rankings_archive;
-- CREATE TABLE IF NOT EXISTS taxonomy_rankings_archive (
-- 	id INT NOT NULL,
-- 	taxonomy_ranking_id INT NOT NULL,
-- 	taxonomy_id INT NOT NULL,
-- 	update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
-- 	ranking TEXT NOT NULL,
-- 	PRIMARY KEY(id),
-- 	FOREIGN KEY(taxonomy_ranking_id) REFERENCES taxonomy_rankings (id) ON UPDATE CASCADE,
-- 	FOREIGN KEY(taxonomy_id) REFERENCES taxonomy (id) ON UPDATE CASCADE
-- );

SET foreign_key_checks = 1;