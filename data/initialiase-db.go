package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Initialize(db *sql.DB) {

	statement, err := db.Prepare(planet)
	if err != nil {
		log.Println(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println("Planets table already exists")
	}

	statement, err = db.Prepare(moon)
	if err != nil {
		log.Println(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println("Moon table already exists")
	}

	PopulatePlanets(db)
	PopluateMoons(db)
}

func PopulatePlanets(db *sql.DB) {
	rows := rowCount(db, "SELECT COUNT(*) FROM planets")

	if rows > 0 {
		log.Println("Planet table already populated", rows)
		return
	}
	statement, _ := db.Prepare(`INSERT INTO planets (
		name,
		description,
		density,
		tilt,
		imageUrl,
		rotationperiod,
		period,
		radius,
		moons,
		au,
		eccentricity,
		velocity,
		mass,
		inclination,
		ordinal
		)
		VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)

	statement.Exec("Mercury", "Mercury—the smallest planet in our solar system and closest to the Sun—is only slightly larger than Earth's Moon. Mercury is the fastest planet, zipping around the Sun every 88 Earth days.", 5.43, 2, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/2_mercury_480x320_new.jpg", 1408, 0.24, 2439, 0, 0.3871, 0.206, 47.89, 0.06, 7.0, 1)
	statement.Exec("Venus", "Venus spins slowly in the opposite direction from most planets. A thick atmosphere traps heat in a runaway greenhouse effect, making it the hottest planet in our solar system.", 5.25, 177.3, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/3_480x320_venus.png", 5832, 0.62, 6052, 0, 0.7233, 0.007, 35.04, 0.82, 3.4, 2)
	statement.Exec("Earth", "Earth—our home planet—is the only place we know of so far that’s inhabited by living things. It's also the only planet in our solar system with liquid water on the surface.", 5.52, 23.45, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/4_earth_480x320.jpg", 23.93, 1, 6378, 1, 1, 0.017, 29.79, 1, 0, 3)
	statement.Exec("Mars", "Mars is a dusty, cold, desert world with a very thin atmosphere. There is strong evidence Mars was—billions of years ago—wetter and warmer, with a thicker atmosphere.", 3.95, 25.19, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/6_mars_480x320.jpg", 24.62, 1.88, 3397, 2, 1.524, 0.093, 24.14, 0.11, 1.85, 4)
	statement.Exec("Jupiter", "Jupiter is more than twice as massive than the other planets of our solar system combined. The giant planet's Great Red spot is a centuries-old storm bigger than Earth.", 1.33, 3.12, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/9_jupiter_480x320_new.jpg", 9.92, 11.86, 71490, 28, 5.203, 0.048, 13.06, 317.89, 1.3, 5)
	statement.Exec("Saturn", "Adorned with a dazzling, complex system of icy rings, Saturn is unique in our solar system. The other giant planets have rings, but none are as spectacular as Saturn's.", 0.69, 26.73, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/38_saturn_480x320.jpg", 10.66, 29.46, 60268, 30, 9.539, 0.056, 9.64, 95.18, 2.49, 6)
	statement.Exec("Uranus", "Uranus—seventh planet from the Sun—rotates at a nearly 90-degree angle from the plane of its orbit. This unique tilt makes Uranus appear to spin on its side.", 1.29, 97.86, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/69_uranus_480x320.jpg", 17.24, 84.01, 25559, 24, 19.19, 0.046, 6.81, 14.53, 0.77, 7)
	statement.Exec("Neptune", "Neptune—the eighth and most distant major planet orbiting our Sun—is dark, cold and whipped by supersonic winds. It was the first planet located through mathematical calculations.", 1.64, 29.6, "https://solarsystem.nasa.gov/system/stellar_items/list_view_images/90_neptune_480x320.jpg", 16.11, 164.79, 25269, 8, 30.06, 0.01, 6.81, 17.14, 1.77, 8)
	log.Println("Planet table populated")
}

func PopluateMoons(db *sql.DB) {
	rows := rowCount(db, "SELECT COUNT(*) FROM moons")

	if rows > 0 {
		log.Println("Moon table already populated", rows)
		return
	}
	statement, _ := db.Prepare(`INSERT INTO moons (
		name,
		planetid,
		ordinal
		)
		VALUES
		(?,?,?)`)

	statement.Exec("Moon", 3, 1)
	statement.Exec("Deimos", 4, 1)
	statement.Exec("Phobos", 4, 2)
	statement.Exec("Io", 5, 1)
	statement.Exec("Europa", 5, 2)
	statement.Exec("Ganymede", 5, 3)
	statement.Exec("Enceladus", 6, 1)
	statement.Exec("Titan", 6, 2)
	statement.Exec("Iapetus", 6, 3)
	statement.Exec("Mimas", 6, 4)
	statement.Exec("Prometheus", 6, 5)
	statement.Exec("Miranda", 7, 1)
	statement.Exec("Ariel", 7, 2)
	statement.Exec("Umbriel", 7, 3)
	statement.Exec("Titania", 7, 4)
	statement.Exec("Oberon", 7, 5)
	statement.Exec("Nereid", 8, 1)
	statement.Exec("Larissa", 8, 2)
	statement.Exec("Triton", 8, 3)
	log.Println("Moon table populated")
}

func rowCount(db *sql.DB, query string) int {
	count := 0
	row := db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}
