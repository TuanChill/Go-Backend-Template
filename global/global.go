package global

import (
	"database/sql"
	"fmt"

	"go_template/configs"
	"go_template/internal/controllers/initialization"
	"go_template/internal/models"
	pkg "go_template/pkg/setting"

	firebase "firebase.google.com/go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	Cfg          models.Config
	DB           *sql.DB
	Cache        *redis.Client
	AdminSdk     *firebase.App
	MessageQueue *amqp.Connection
)

func init() {
	//* CONFIG
	var err error
	Cfg, err = configs.LoadConfig("configs/yaml")
	if err != nil {
		panic(err)
	}

	//* DATABASE
	DB, err = initialization.ConnectPG(Cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}

	//* CACHE
	Cache, err = initialization.ConnectRedis(Cfg)
	if err != nil {
		fmt.Printf("Error connecting to Redis: %v\n", err)
		panic(err)
	}

	//* FIREBASE
	AdminSdk, err = pkg.InitializeApp()
	if err != nil {
		fmt.Printf("Error connecting to firebase: %v\n", err)
		panic(err)
	}

	//* RABBITMQ
	MessageQueue, err = initialization.ConnectRabbitMQ(Cfg.RabbitMQ.URL)
	if err != nil {
		fmt.Printf("Error connecting to RabbitMq: %v\n", err)
		panic(err)
	}
}
