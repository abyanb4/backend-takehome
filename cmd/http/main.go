package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	commentDmn "backend-takehome/internal/domain/comment"
	postDmn "backend-takehome/internal/domain/post"
	userDmn "backend-takehome/internal/domain/user"
	"backend-takehome/internal/server/http"
	commentUc "backend-takehome/internal/usecase/comment"
	postUc "backend-takehome/internal/usecase/post"
	userUc "backend-takehome/internal/usecase/user"

	config "backend-takehome/internal/common/config"

	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB
var cacheRsc *cache.Cache
var cfg *config.Config

func main() {
	// Init Config
	cfg = &config.Config{
		Cache: &config.CacheConfig{
			CleanupInterval:   1 * time.Minute,
			DefaultExpiration: 30 * time.Minute,
			TokenExpiration:   120 * time.Minute,
		},
	}

	// Init DB
	initDB()

	// init cache
	cacheRsc = cache.New(cfg.Cache.DefaultExpiration, cfg.Cache.CleanupInterval)

	// init Domain
	postDomain := postDmn.NewPostDomain(db)
	userDomain := userDmn.NewUserDomain(userDmn.UserDomain{
		DB: db,
	})
	commentDomain := commentDmn.NewCommentDomain(db)

	// init Usecase
	ucHandler := http.HTTPHandler{
		PostUsecase: postUc.NewPostUsecase(postDomain, commentDomain),
		UserUsecase: userUc.NewUserUsecase(userUc.UserUsecase{
			UserDomain: userDomain,
			Cache:      cacheRsc,
			Config:     cfg,
		}),
		CommentUsecase: commentUc.NewCommentUsecase(commentDomain),
	}

	// Start http server
	r := gin.Default()
	http.RegisterHandler(r, &ucHandler)

	r.Run(":" + os.Getenv("APP_PORT"))
}

func initDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		log.Fatal("One or more environment variables are not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		user,
		password,
		host,
		port,
		name,
	)

	var err error
	for i := 0; i < 10; i++ { // Retry up to 10 times
		db, err = sqlx.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			fmt.Println("Database connection is successful!")
			return
		}

		log.Printf("Failed to connect to database: %v. Retrying in 2 seconds...", err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Failed to connect to database after several retries: %v", err)
}
