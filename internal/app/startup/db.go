package startup

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func DatabaseConnect(config DatabaseConfig) (*gorm.DB, error) {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//err = driver.DB().Ping()

	//sqlDB, _ := db.DB()
	//
	//err = sqlDB.Close()
	//if err != nil {
	//	return nil, err
	//}

	//if err != nil {
	//	return nil, err
	//}
	return db, nil
}

//func DatabaseConnect(config DatabaseConfig) (*ent.Client, error) {
//	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
//		config.Host, config.Port, config.User, config.DBName, config.Password, config.SslMode)
//	driver, err := sql.Open(dialect.Postgres, dsn)
//	if err != nil {
//		return nil, err
//	}
//
//	err = driver.DB().Ping()
//	if err != nil {
//		return nil, err
//	}
//
//	return ent.NewClient(ent.Driver(driver)), nil
//}
