package main

import (
	"flag"
	"log"
	"musictracks/internal/models"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type application struct {
	tracks  *models.TrackModel
	artists *models.ArtistModel
	images  *models.ImageModel
}

func main() {
	addr := flag.String("port", ":8000", "Http network address")
	dsn := flag.String("dsn", "", "Mysql data source name using GORM")
	flag.StringVar(&clientID, "cid", "", "Client ID for spotify api")
	flag.StringVar(&clientSecret, "cs", "", "CLient secret required for spotify api")

	flag.Parse()

	db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{})
	checkError(err)

	app := &application{
		tracks:  &models.TrackModel{DB: db},
		artists: &models.ArtistModel{DB: db},
		images:  &models.ImageModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Fatal(srv.ListenAndServe())
}
