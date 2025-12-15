package config

// var (
// 	db     *gorm.DB
// 	logger *Logger
// )

// func Init() error {
// 	// return errors.New("fake error")
// 	return nil
// }

// func GetLogger(p string) *Logger {
// 	logger = NewLogger(p)
// 	return logger
// }

// type Config struct {
// 	DBDriver string
// 	DBSource string
// }

// func LoadConfig() *Config {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("⚠️  Aviso: .env não encontrado, usando variáveis do sistema")
// 	}

// 	cfg := &Config{
// 		DBDriver: getEnv("DB_DRIVER", "postgres"),
// 		DBSource: getEnv("DB_SOURCE", "postgres://postgres:123@localhost:5432/meubanco?sslmode=disable"),
// 	}

// 	return cfg
// }

// func getEnv(key, fallback string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		return fallback
// 	}
// 	return value
// }
