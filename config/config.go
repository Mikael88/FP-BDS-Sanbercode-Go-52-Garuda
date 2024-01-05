package config

import "os"

var (
	MongoURI  = os.Getenv("MONGO_URI")
	SecretKey = os.Getenv("SECRET_KEY")
)
