package database

import (
	"fmt"
	// "log"

	"smartpatrol/models"
	"smartpatrol/pkg/util/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

type credential struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
	SSLMode  string
	Timezone string
}

func Init(driver string) {
	credential := &credential{
		Host: environment.Get("DB_HOST"),
		Port: environment.Get("DB_PORT"),
		User: environment.Get("DB_USER"),
		Pass: environment.Get("DB_PASSWORD"),
		Name: environment.Get("DB_NAME_SECURITY"),
		// SSLMode:  enviroment.Get("DB_SSLMODE"),
		Timezone: environment.Get("TIMEZONE_TZ"),
	}

	switch driver {
	case "postgres":
		credential.getPostgres()
	default:
		credential.getPostgres()
	}
}



func DBManager() *gorm.DB {
	return db
}

func (c *credential) getPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s Timezone=%s", c.Host, c.Port, c.User, c.Pass, c.Name, c.Timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	db.AutoMigrate(
		models.SecurityUser{},
		models.SecurityAccount{},
		models.SecurityRole{},
		models.SecurityRolePermission{},
		models.SecurityPermission{},
		models.SecuritySessionLog{},
		models.SecuritySessionLog{},
		models.CompanyCompany{},
		models.CompanyShift{},
		models.CompanyArea{},
	)
}