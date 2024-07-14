package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/DanilMankiev/SofiaApplication/config"
	v1 "github.com/DanilMankiev/SofiaApplication/internal/controllers/http/v1"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
	"github.com/DanilMankiev/SofiaApplication/internal/service"
	server "github.com/DanilMankiev/SofiaApplication/pkg/httpserver"
	"github.com/DanilMankiev/SofiaApplication/pkg/otp"
	"github.com/DanilMankiev/SofiaApplication/pkg/postgres"
	"github.com/DanilMankiev/SofiaApplication/pkg/rabbitmq"
	"github.com/gin-gonic/gin"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)



func Run(cfg *config.Config){
	logrus.SetFormatter(new(logrus.JSONFormatter))
	
	//Postgres
	pg,err:=postgres.New(postgres.Postgres{
		User: cfg.Postgres.User,
		Port: cfg.Postgres.Port,
		Dbname: cfg.Postgres.Dbname,
		Host: cfg.Postgres.Host,
		Sslmode: cfg.Postgres.Sslmode,
		Password: cfg.Postgres.Password,
	})
	if err!=nil{
		logrus.Fatalf("Failed to connect database:%s",err.Error())
	}
	
	// RabbitMQ
	rabbit,err:=rabbitmq.NewRabbitMQ(cfg.RabbitMQ.Url)
	if err!=nil{
		logrus.Fatalf("Failed to connect rabbitmq:%s",err.Error())
	}

	//OTP
	otpGenerator:=otp.NewOTPGenerator()

	//Repository
	repo:=repository.New(pg)
	
	//Service
	service:=service.New(repo,rabbit,otpGenerator,cfg.Authorization.VerificationCodeLength)

	//HTTP server
	router:=gin.New()
	handler:=v1.NewHandler(service)
	handler.NewRouter(router)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.PortHTTP, router); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	//Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := pg.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}