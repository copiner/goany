package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}


func main() {
	// Capture connection properties.
	
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
	    DBName: "recordings",
	    AllowNativePasswords: true,
	}
	
	// Get a database handle.
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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
	/*
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
	*/
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
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

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
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


func canPurchase(id int, quantity int) (bool, error) {
    var enough bool
    // Query for a value based on a single row.
    if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
        quantity, id).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return false, fmt.Errorf("canPurchase %d: unknown album", id)
        }
        return false, fmt.Errorf("canPurchase %d: %v", id)
    }
    return enough, nil
}



// CreateOrder creates an order for an album and returns the new order ID.
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {

    // Create a helper function for preparing failure results.
    fail := func(err error) (int64, error) {
        return fmt.Errorf("CreateOrder: %v", err)
    }

    // Get a Tx for making transaction requests.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }
    // Defer a rollback in case anything fails.
    defer tx.Rollback()

    // Confirm that album inventory is enough for the order.
    var enough bool
    if err = tx.QueryRowContext(ctx, "SELECT (quantity >= ?) from album where id = ?",
        quantity, albumID).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return fail(fmt.Errorf("no such album"))
        }
        return fail(err)
    }
    if !enough {
        return fail(fmt.Errorf("not enough inventory"))
    }

    // Update the album inventory to remove the quantity in the order.
    _, err = tx.ExecContext(ctx, "UPDATE album SET quantity = quantity - ? WHERE id = ?",
        quantity, albumID)
    if err != nil {
        return fail(err)
    }

    // Create a new row in the album_order table.
    result, err := tx.ExecContext(ctx, "INSERT INTO album_order (album_id, cust_id, quantity, date) VALUES (?, ?, ?, ?)",
        albumID, custID, quantity, time.Now())
    if err != nil {
        return fail(err)
    }
    // Get the ID of the order item just created.
    orderID, err := result.LastInsertId()
    if err != nil {
        return fail(err)
    }

    // Commit the transaction.
    if err = tx.Commit(); err != nil {
        return fail(err)
    }

    // Return the order ID.
    return orderID, nil
}


func QueryWithTimeout(ctx context.Context) {
    // Create a Context with a timeout.
    queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Pass the timeout Context with a query.
    rows, err := db.QueryContext(queryCtx, "SELECT * FROM album")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Handle returned rows.
}
