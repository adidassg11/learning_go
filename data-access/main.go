// tutorial from https://go.dev/doc/tutorial/database-access

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB // in prod we'd avoid a global var for the handle

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// capture connection props
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a db handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	//var artist string = "John Coltrane" // This works as well
	var artist = "John Coltrane"
	albums, err := albumsByArtist(artist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found for %v: %v\n", artist, albums)

	var albumId = 2 // can also do `var albumId int64 = 2`
	alb, err := albumByID(int64(albumId))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found for album id {%d}: %v\n", albumId, alb)

	// the := must be declaring at least one new variable.
	// It automatically acts as declaration+reassignment is something exists already
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of newly added album: %d\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name
func albumsByArtist(name string) ([]Album, error) {
	// an albums slice to hold data from returned rows
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	// loop thru rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsByID %d: no such album", id)
		}
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
