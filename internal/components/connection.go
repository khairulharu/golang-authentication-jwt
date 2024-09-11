package components

import (
	"fmt"
	"log"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s "+"port=%s "+"user=%s "+"password=%s "+"dbname=%s "+"sslmode=disable",
		cnf.DB.Host, cnf.DB.Port, cnf.DB.User, cnf.DB.Pass, cnf.DB.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("err when connect to database: %s", err.Error())
	}

	&gorm.Migrator(domain.User)

	fmt.Println("database Connnet")

	return db
}
