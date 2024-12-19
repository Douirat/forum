package models

import (
    "database/sql"
    "log"
    _"github.com/mattn/go-sqlite3"
)

// GetDataFromDB retrieves data from the SQLite database
func GetDataFromDB() []string {
    db, err := sql.Open("sqlite3", "./database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT name FROM items")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var items []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        items = append(items, name)
    }

    return items
}
