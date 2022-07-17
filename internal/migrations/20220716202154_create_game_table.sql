-- +goose Up
CREATE TABLE IF NOT EXISTS games
(
	id 			UUID 		    PRIMARY KEY DEFAULT gen_random_uuid(),	
	title 	    VARCHAR(50)     NOT NULL,
	age_rating	VARCHAR(4)      NOT NULL,
    description VARCHAR         NOT NULL,
    publisher   VARCHAR         NOT NULL,
    url         VARCHAR         NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS games;


