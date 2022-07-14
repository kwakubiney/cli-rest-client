-- +goose Up
CREATE TABLE IF NOT EXISTS users 
(
	id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),	
	username  	VARCHAR(50) NOT NULL,
	age 			VARCHAR 	NOT NULL,
	email 		VARCHAR (40)	NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users ;