package main

import (
	"bail/config"
	"bail/delivery/core"
	"bail/infrastructure/hash"
	"bail/infrastructure/jwt"
	userrepo "bail/infrastructure/repo/user"
	"time"

	usercontroller "bail/delivery/controller/user"
	"bail/delivery/router"
	usercmd "bail/usecases/user/command"
	userqry "bail/usecases/user/query"
	"fmt"
	"log"

	db "bail/infrastructure/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	cfg := config.Envs

	// Initialize MongoDB client and perform migrations
	mongoClient := initDB(cfg)

	// Initialize services
	userRepo := initRepos(&cfg, mongoClient)

	jwtService := jwt.New(
		jwt.Config{
			SecretKey: config.Envs.JWTSecret,
			Issuer:    config.Envs.ServerHost,
			ExpTime:   time.Duration(config.Envs.JWTExpirationInSeconds) * time.Second,
		})

	hashService := &hash.Service{}
	userController := initUserController(userRepo, hashService, jwtService)

	routerConfig := router.Config{
		Addr:        fmt.Sprintf(":%s", cfg.ServerPort),
		BaseURL:     "/api",
		Controllers: []core.IController{userController},
		JwtService:  jwtService,
	}

	r := router.NewRouter(routerConfig)

	// Start the HTTP server
	log.Printf("Starting server on %s", cfg.ServerPort)
	if err := r.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initDB(cfg config.Config) *mongo.Client {
	mongoClient := db.Connect(db.Config{
		ConnectString: cfg.DBConnectionString,
	})

	db.Migrate(mongoClient, cfg.DBName)

	return mongoClient
}

func initRepos(cfg *config.Config, mongoClient *mongo.Client) *userrepo.Repo {
	userRepo := userrepo.New(mongoClient, cfg.DBName, "user")
	return userRepo

}

func initUserController(userRepo *userrepo.Repo, hashService *hash.Service, jwtService *jwt.Service) *usercontroller.UserController {

	loginHandler := usercmd.NewLoginHandler(usercmd.LoginConfig{
		UserRepo:     userRepo,
		JwtService:   jwtService,
		HashService:  hashService,
	},
	)

	signupHandler := usercmd.NewSignUpHandler(usercmd.SignUpConfig{
		UserRepo:    userRepo,
		JwtService:  jwtService,
		HashService: hashService,
	},
	)

	updateHandler := usercmd.NewUpdateUserHandler(usercmd.UpdateUserConfig{
		UserRepo: userRepo,
	},
	)

	return usercontroller.New(usercontroller.Config{
		SignupUserHandler: signupHandler,
		UpdateUserHandler: updateHandler,
		LoginUserHandler:     loginHandler,
		GetEmployeeHandler:    userqry.NewGetusersHandler(userRepo),
		GetUserHandler:        userqry.NewGetHandler(userRepo),
		DeleteEmployeeHandler: usercmd.NewDeleteHandler(userRepo),
		PromoteUserHandler:    usercmd.NewPromoteHandler(userRepo),
	},
	)

}
