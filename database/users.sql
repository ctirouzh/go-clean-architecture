CREATE TABLE users (
	id VARCHAR(36) NOT NULL,
	username VARCHAR(20) NOT NULL,
	email VARCHAR(80) NOT NULL,
	role_type TINYINT,
	password_hash VARCHAR(128),
	created_at DATETIME,
	updated_at DATETIME,
);

ALTER TABLE users
ADD CONSTRAINT PK_user PRIMARY KEY (id)
ADD CONSTRAINT UC_user UNIQUE (username, email);