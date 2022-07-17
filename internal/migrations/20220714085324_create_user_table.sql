-- +goose Up

-- TODO: Make age an integer
CREATE TABLE IF NOT EXISTS users 
(
	id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),	
	username  	VARCHAR(50) NOT NULL,
	age 			INTEGER 	NOT NULL,
	email 		VARCHAR (40)	NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users ;

-- goose postgres "user=postgres password=postgres dbname=gamedev sslmode=disable" up  
-- goose postgres "user=postgres password=postgres dbname=gamedev_test sslmode=disable" up  