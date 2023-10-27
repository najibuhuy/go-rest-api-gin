install modul
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlserver
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
go get github.com/githubnemo/compileDaemon
go install github.com/githubnemo/compileDaemon

Run the program in local comp:
compiledaemon --command="./go-rest-api-gin"