package repository

import (
	"context"
	"log"
	"fmt"
	"github.com/PyMarcus/codes_download/tools"
    "github.com/jackc/pgx/v4/pgxpool"
)

func psqlConnect() *pgxpool.Pool{
	credentials := tools.GetDataBaseCredentials()
	conn, _ := pgxpool.Connect(context.Background(), fmt.Sprintf("user=%s host=%s dbname=%s password=%s sslmode=disable", credentials["user"], credentials["host"], credentials["database"], credentials["password"]))
	
	log.Println("Success to connect into PostgreSQL")
    return conn
}