module github.com/one-piece-team1/one-piece-recovery

go 1.14

replace (
	github.com/one-piece-team1/one-piece-recovery/src/apps => ../one-piece-recovery/src/apps
	github.com/one-piece-team1/one-piece-recovery/src/config => ../one-piece-recovery/src/config
	github.com/one-piece-team1/one-piece-recovery/src/controllers/healths => ../one-piece-recovery/src/controllers/healths
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg/v10 v10.8.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)
