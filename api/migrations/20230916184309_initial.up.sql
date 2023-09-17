CREATE TABLE items (
	id varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	path text NOT NULL,
	size bigint,
	files json,
	PRIMARY KEY (id)
);

