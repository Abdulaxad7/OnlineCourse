package configuration

import (
	"fmt"
	"os"
)

var Configuration = con()

func con() conDB {
	return conDB{
		PublicHost: GETEnv("PUBLIC_HOST", "https://localhost"),
		Port:       GETEnv("PORT", "8080"),
		DBUser:     GETEnv("DBUSER", "root"),
		DBAddr:     fmt.Sprintf("%s:%s", GETEnv("DB_HOST", "127.0.0.1"), GETEnv("DB_POST", "3306")),
		DBPassword: GETEnv("DBPASSWORD", "0987poiulkjh"),
		DBName:     GETEnv("DBNAME", "data12"),
	}
}
func GETEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
