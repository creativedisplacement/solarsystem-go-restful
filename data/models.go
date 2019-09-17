package data

const planet = `CREATE TABLE IF NOT EXISTS planets (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
	name VARCHAR(64),
	description VARCHAR(500),
	density NUMERIC,
	tilt NUMERIC,
	imageUrl VARCHAR(100),
	rotationperiod NUMERIC,
	period NUMERIC,
	radius INTEGER,
	moons INTEGER,
	au NUMERIC,
	eccentricity NUMERIC,
	velocity NUMERIC,
	mass NUMERIC,
	inclination NUMERIC,
	ordinal INTEGER
)`

const moon = `CREATE TABLE IF NOT EXISTS moons (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
	name VARCHAR(64),
	planetid INTEGER NOT NULL,
	ordinal INTEGER,
	FOREIGN KEY(planetid) REFERENCES planets(id)
)`
