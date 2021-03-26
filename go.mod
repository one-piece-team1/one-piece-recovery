module github.com/one-piece-team1/one-piece-recovery

go 1.14

replace (
	github.com/one-piece-team1/one-piece-recovery/src/apps => ../one-piece-recovery/src/apps
	github.com/one-piece-team1/one-piece-recovery/src/config => ../one-piece-recovery/src/config
	github.com/one-piece-team1/one-piece-recovery/src/controllers/healths => ../one-piece-recovery/src/controllers/healths
	github.com/one-piece-team1/one-piece-recovery/src/domains/enums => ../one-piece-recovery/src/domains/enums
	github.com/one-piece-team1/one-piece-recovery/src/domains/recoveries => ../one-piece-recovery/src/domains/recoveries
	github.com/one-piece-team1/one-piece-recovery/src/logger => ../one-piece-recovery/src/logger
	github.com/one-piece-team1/one-piece-recovery/src/pkg => ../one-piece-recovery/src/pkg
	github.com/one-piece-team1/one-piece-recovery/src/pkg/postgres => ../one-piece-recovery/src/pkg/postgres
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg/v10 v10.8.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/postgres v1.0.8 // indirect
	gorm.io/gorm v1.21.4 // indirect
)
