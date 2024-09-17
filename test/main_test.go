package test

import (
	"testing"

	"github.com/khairulharu/gojwt/internal/components"
	"github.com/khairulharu/gojwt/internal/config"
	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func TestMain(m *testing.M) {

	conf := config.New()

	dbGorm = components.NewDatabaseConnection(conf)

	m.Run()
}
