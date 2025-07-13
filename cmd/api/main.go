package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"xyz-multifinance/transaction-service/pkg/database"
	httpHandler "xyz-multifinance/transaction-service/internal/handler/http"
	"xyz-multifinance/transaction-service/internal/repository/mysql"
	"xyz-multifinance/transaction-service/internal/usecase"
)

func main() {
	dbConfig := database.Config{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "3306"),
		DBName:   getEnv("DB_NAME", "xyz_db"),
	}

	db, err := database.NewConnection(dbConfig)
	if err != nil { log.Fatalf("FATAL: Could not start database: %v", err) }
	defer db.Close()
	log.Println("Database connection established successfully.")

	transactionRepo := mysql.NewMysqlTransactionRepository(db)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo)
	transactionHandler := httpHandler.NewTransactionHandler(transactionUC)

	r := mux.NewRouter()
	r.Handle("/v1/transactions", httpHandler.AuthMiddleware(http.HandlerFunc(transactionHandler.CreateTransaction))).Methods("POST")

	port := getEnv("PORT", "8080")
	log.Printf("Server is starting on port :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("FATAL: Could not start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok { return value }
	return fallback
}