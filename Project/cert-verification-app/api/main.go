package main

import (
    "database/sql"
    "encoding/json"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type Certificate struct {
    ID     int    `json:"id"`
    Serial string `json:"serial"`
    Status string `json:"status"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/certdb")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    http.HandleFunc("/api/certificates", getCertificates)
    http.ListenAndServe(":8080", nil)
}

func getCertificates(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    rows, err := db.Query("SELECT id, serial, status FROM certificates")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var certificates []Certificate
    for rows.Next() {
        var cert Certificate
        if err := rows.Scan(&cert.ID, &cert.Serial, &cert.Status); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        certificates = append(certificates, cert)
    }

    json.NewEncoder(w).Encode(certificates)
}