package controllers

import (
	database "api/database"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)
var secret = []byte("12345")
var apiKey = "12345"
func PostUser(context *gin.Context) {
	
	var newUser database.User
	if err := context.BindJSON(&newUser); err != nil{
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err !=  nil{
		fmt.Println(err.Error())
	}
	newUser.Password = string(bytes)
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/godev")
	if err !=  nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	if userRow, userExists := database.FindUser(db,newUser.Login); !userExists{
		fmt.Println(userRow)
		insert,err := db.Query(fmt.Sprintf(`INSERT INTO user(login,password) VALUES("%v", "%v")`,newUser.Login,newUser.Password))
		errHandler(err)
		context.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "Não foi possivel inserir!"})
		defer insert.Close()
		
	}else{
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Usuario já existe!"})
	}
}

func Authentication(context *gin.Context){
	var user database.User
	if err := context.BindJSON(&user); err != nil{
		return
	}
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/godev")
	if err !=  nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var userFromDB database.UserRow
	userRow, userExists :=  database.FindUser(db,user.Login); 
	if !userExists{
		panic("O usuario não existe!")
	}
	userFromDB = userRow
	if match := compareHash(user.Password,userFromDB.Password); match == true{
		if context.GetHeader("api") != "" {
			if context.GetHeader("api") == apiKey {
				token := createJWT()
				context.IndentedJSON(http.StatusNonAuthoritativeInfo,gin.H{"message":token})
			}else{
				context.IndentedJSON(http.StatusNonAuthoritativeInfo,gin.H{"message":"Api key incorreta!"})
			}
		}else{
				context.IndentedJSON(http.StatusNonAuthoritativeInfo,gin.H{"message":"APIKey em branco!"})
		}
	}else{
		fmt.Printf("O resultado da comparação do hash é: %v\n",match)
		context.IndentedJSON(http.StatusNotAcceptable,gin.H{"message":"Senha incorreta!"})
	}
}

func compareHash(password, hash string)bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil{
		return false
	}
	return true
}

var Test string = "test"

func errHandler(err error){
	if err != nil {
		panic(err.Error())
	}
}

func createJWT()(string){
	token := jwt.New(jwt.SigningMethodHS256)
	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	//assigniment
	tokenStr, err := token.SignedString(secret)
	if err != nil{
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println(tokenStr)
	return tokenStr
}

func ValidateJWT(token string,context *gin.Context){
	if token != ""{
		token, err := jwt.Parse(token, func(t *jwt.Token)(interface{}, error){
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				context.IndentedJSON(http.StatusNonAuthoritativeInfo,gin.H{"message":"Token invalido"})
			}else{
				fmt.Println(ok)
			}
			return secret, nil
		})
		if err != nil{ 
			context.IndentedJSON(http.StatusNotAcceptable,gin.H{"message":"Token expirado"})
		}
		if token.Valid{
			context.IndentedJSON(http.StatusOK,gin.H{"message":"Autorizado"})
		}
	}else{
		context.IndentedJSON(http.StatusNotAcceptable,gin.H{"message":"Token em branco"})
	}
}
	
func Home(context *gin.Context){
	if context.GetHeader("api") != ""{
		if context.GetHeader("api") == apiKey{
			ValidateJWT(context.GetHeader("token"),context)
		}
	}else{
		context.IndentedJSON(http.StatusNotAcceptable,gin.H{"message":"Api key em branco"})
	}
}