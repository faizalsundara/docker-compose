package main

import (
	// "banking/client"
	"banking/app"
	"log"
	"os"

	"github.com/joho/godotenv"

	// "log"

	"banking/routes"

	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func main() {

	cfg := jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831", // replace host
		},
	}

	closer, err := cfg.InitGlobalTracer(
		"main-service",
	)
	defer closer.Close()

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	errGet := godotenv.Load()
	if errGet != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	presenter := app.InitFactory()
	e := routes.New(presenter)
	e.Logger.Fatal(e.Start(":" + port))
}
