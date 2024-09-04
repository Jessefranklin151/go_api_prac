package di

import (
	"log"
	"os"

	"github.com/Jessefranklin151/go_api_prac/go_api_prac/config"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/controllers"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/data"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/handler"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/repositories"
	"github.com/joho/godotenv"
)

type DI struct {
	PlayerHandler *handler.PlayerHandler
}

var Dependencies *DI

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseType := os.Getenv("DATABASE_TYPE")

	var dataLayer data.IPlayerData

	switch databaseType {
	case "memory":
		dataLayer = &data.MemoryPlayerData{}
	case "postgres":
		db, err := config.InitializeDB()
		if err != nil {
			panic(err)
		}
		dataLayer = &data.DatabasePlayerData{
			DB: db,
		}
	default:
		panic("Unsupported database type")
	}

	Dependencies = &DI{
		PlayerHandler: initPlayerHandler(dataLayer),
	}
}

func initPlayerHandler(dataLayer data.IPlayerData) *handler.PlayerHandler {
	playerRepository := repositories.CreatePlayerRepository(dataLayer)
	playerController := controllers.CratePlayerController(playerRepository)
	return handler.InitPlayerHandler(playerController)
}
