package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputados_all"
	grpc_service "github.com/dariocasas/deputados3/src/internal/infra/grpc"
	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"
	"github.com/dariocasas/deputados3/src/internal/usecase/db_exists_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/drop_db_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/fetch_deputados_ids_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/get_deputado_from_api_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/get_fotos_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/populate_db_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/populate_index_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/scan_page"
	"github.com/dariocasas/deputados3/src/internal/usecase/shutdown_usecase"
	"github.com/dariocasas/deputados3/src/pkg/class_finder"
	"github.com/dariocasas/deputados3/src/pkg/http_client"
	"github.com/dariocasas/deputados3/src/pkg/url_open"
	"google.golang.org/grpc"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	config, err := configs.NewConfig(".")
	if err != nil {
		panic(err)
	}

	log.Println("Connecting to mongoDB...")
	db, err := mongodb.ConnectMongoDB(config)
	if err != nil {
		panic(err)
	}
	defer mongodb.DisconnectMongoDB(db)

	deputadoIndexRepository := repo_deputado_index.NewDeputadoIndexRepository(db, config)
	deputadoRepository := repo_deputado.NewDeputadoRepository(db, config)

	httpClient := http_client.NewHttpClient(config)
	apiDeputadosAll := api_deputados_all.NewApiDeputadosAll(config, httpClient)
	apiDeputado := api_deputado.NewApiDeputado(config, httpClient)

	classFinder := class_finder.NewClassFinder()
	scanPageDeputado := scan_page.NewScanPageDeputado(classFinder)

	// UseCases ------------

	fetchDeputadosIdsUseCase := fetch_deputados_ids_usecase.NewFetchDeputadosIdsUseCase(deputadoIndexRepository, apiDeputadosAll)
	dbStatusUseCase := db_exists_usecase.NewDbStatusUseCase(db, config, deputadoIndexRepository)
	getDeputadoFromApiUseCase := get_deputado_from_api_usecase.NewGetDeputaFromApiUseCase(apiDeputado)
	populateDbUseCase := populate_db_usecase.NewPopulateDbUseCase(
		config, deputadoIndexRepository, deputadoRepository, apiDeputado,
		scanPageDeputado, httpClient, getDeputadoFromApiUseCase,
	)
	populateIndexUseCase := populate_index_usecase.NewPopulateIndexUseCase(
		config, deputadoIndexRepository, apiDeputadosAll,
	)
	dropDbUsecase := drop_db_usecase.NewDropDbUseCase(deputadoIndexRepository, deputadoRepository)
	fotosUsecase := get_fotos_usecase.NewFotosUseCase(deputadoRepository)

	// GRPC ------------

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	dbService := grpc_service.NewDbService(config,
		*dropDbUsecase, *dbStatusUseCase, *fetchDeputadosIdsUseCase,
		*populateDbUseCase, *populateIndexUseCase,
	)
	pb.RegisterDbServiceServer(grpcServer, dbService)
	fotosService := grpc_service.NewFotosService(*fotosUsecase)
	pb.RegisterFotosServiceServer(grpcServer, fotosService)

	shutdownUsecase := shutdown_usecase.NewShutdownUseCase(grpcServer, db, config)
	serverService := grpc_service.NewServerService(*shutdownUsecase)
	pb.RegisterServerServiceServer(grpcServer, serverService)

	log.Printf("Starting gRPC server on %s:%s", config.GRPCServerAddr, config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GRPCServerAddr, config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	url_open.OpenUrl("deputados.exe")

	go func() {
		<-c
		os.Exit(1)
	}()

	for {
		time.Sleep(10 * time.Second)
	}

}
