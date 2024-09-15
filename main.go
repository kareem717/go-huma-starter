package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	server "proj/internal/server"
	"proj/internal/service/domain"
	"proj/internal/storage/postgres"

	"github.com/supabase-community/supabase-go"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Port               int    `help:"Port to listen on" short:"p" default:"8080"`
	DatabaseURL        string `help:"Database URL" short:"d"`
	APIName            string `help:"API Name" short:"n"`
	APIVersion         string `help:"API Version" short:"v"`
	BaseURL            string `help:"Base API URL" short:"B"`
	SupabaseHost       string `help:"Supabase Host" short:"s"`
	SupabaseServiceKey string `help:"Supabase Service Key" short:"k"`
}

func (o *Options) config() {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		o.Port = port
	}

	o.DatabaseURL = os.Getenv("DATABASE_URL")
	o.APIName = os.Getenv("API_NAME")
	o.APIVersion = os.Getenv("API_VERSION")
	o.BaseURL = os.Getenv("BASE_API_URL")
	o.SupabaseHost = os.Getenv("SUPABASE_HOST")
	o.SupabaseServiceKey = os.Getenv("SUPABASE_SERVICE_KEY")
}

func main() {
	// Load environment variables from .env.local
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Println("Error loading .env.local file")
	}

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		options.config()

		ctx := context.Background()
		logger := zap.New(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
				zapcore.AddSync(os.Stdout), zap.InfoLevel))

		postgresConfig := postgres.NewConfig(options.DatabaseURL)
		repositories := postgres.NewRepository(postgresConfig, ctx, logger)

		services := domain.NewService(repositories)

		supabaseClient, err := supabase.NewClient(
			options.SupabaseHost,
			options.SupabaseServiceKey,
			&supabase.ClientOptions{},
		)
		if err != nil {
			logger.Fatal("Failed to create supabase client", zap.Error(err))
		}

		server := server.NewServer(
			services,
			options.APIName,
			options.APIVersion,
			logger,
			supabaseClient,
		)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			server.Serve(fmt.Sprintf(":%d", options.Port))
		})
	})

	cli.Run()
}
