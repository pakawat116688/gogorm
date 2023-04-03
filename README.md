# gogorm
    สามารถ go run ได้ตามปกติ หรือทำการ Deploy ลงบน kubenetes โดยให้ทำการสร้าง Docker Image ขึ้นมาก่อนและนำไปใส่ใร kubernetes

# Postman
- localhost:30003/signup + Body {"Username": "","Password": ""} Method POST
- localhost:30003/signin + Body {"Username": "","Password": ""} Method POST
- localhost:30003/user/users Method GET
- localhost:30003/userone/:username Method GET


# library
- go get gorm.io/gorm
- go get gorm.io/driver/sqlite
- go get github.com/spf13/viper
- go get golang.org/x/crypto/bcrypt
- go get github.com/golang-jwt/jwt/v4 -> Not use
- go get github.com/LdDl/fiber-jwt/v2
- go get github.com/dgrijalva/jwt-go/v4
- go get github.com/gofiber/fiber/v2/middleware/requestid
- go get github.com/gofiber/jwt/v3 -> search gofiber jwt = jwtware (github.com/gofiber/jwt/v3)