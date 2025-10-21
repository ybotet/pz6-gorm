package main

import (
	"log"
	"net/http"

	"example.com/pz6-gorm/internal/db"
	httpapi "example.com/pz6-gorm/internal/http"
	"example.com/pz6-gorm/internal/models"
)

func main() {
	d := db.Connect()

	// Автоматически создаст (или обновит) таблицы под наши модели
	if err := d.AutoMigrate(&models.User{}, &models.Note{}, &models.Tag{}); err != nil {
		log.Fatal("migrate:", err)
	}

	r := httpapi.BuildRouter(d)

	log.Println("listening on :8087")
	log.Fatal(http.ListenAndServe(":8087", r))
}
