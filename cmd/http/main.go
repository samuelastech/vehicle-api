package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/samuelastech/vehicle-api/docs/seed"
	"github.com/samuelastech/vehicle-api/internal/vehicle"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	seed := seed.VehicleJSONFile()
	rp := vehicle.NewRepository(seed)
	sv := vehicle.NewService(rp)
	hr := vehicle.NewHandler(sv)

	vehicle.NewRouter(r, hr)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
}
