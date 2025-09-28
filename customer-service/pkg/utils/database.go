package utils

import (
    "database/sql"
    "fmt"
    "log"
    "customer-service/pkg/config"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDBFromConfig() {
    user := config.GetString("spring.datasource.username")
    pass := config.GetString("spring.datasource.password")
    url := config.GetString("spring.datasource.url") // ex: localhost:3308/customerdb

    // Séparer host:port et database
    var hostPort, dbName string
    n := len(url)
    sep := -1
    for i := 0; i < n; i++ {
        if url[i] == '/' {
            sep = i
            break
        }
    }
    if sep == -1 {
        log.Fatal("❌ URL MySQL invalide, format attendu host:port/database")
    }
    hostPort = url[:sep]
    dbName = url[sep+1:]

    // DSN correct
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, hostPort, dbName)

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("❌ Erreur connexion DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("❌ DB inaccessible:", err)
    }

    log.Println("✅ Connexion DB réussie avec Config Server")
}
