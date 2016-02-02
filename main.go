package main

import (
	"encoding/json"
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/jebrial/learnlink/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type Conf struct {
	Url string
}

var (
	SECRET = "SECRETHERE"
)

func dbWare() gin.HandlerFunc {
	//load the config
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic(err)
	}
	decoder := json.NewDecoder(file)
	conf := Conf{}

	err = decoder.Decode(&conf)
	if err != nil {
		log.Panic(err)
	}
	// connect to the database
	db, err := models.NewDB(conf.Url)
	if err != nil {
		log.Panic(err)
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {

	// Set up server
	ginServer := gin.Default()
	ginServer.Use(dbWare())

	//login public routes
	ginServer.POST("/login", UserLogin)
	ginServer.POST("/signup", UserAdd)
	//set up private routes
	private := ginServer.Group("/api")
	private.Use(jwt.Auth(SECRET))
	// user private /api routes
	private.DELETE("/user/delete/:email", UserRemove)

	//course rivate /api routes
	private.GET("/course/all/:email", CourseIndex)
	private.POST("/course/new", CourseAdd)
	private.PUT("/course/update/:id", CourseUpdate)
	private.DELETE("/course/delete/:id", CourseRemove)

	ginServer.Run(":3001")
}

func UserLogin(ctx *gin.Context) {
	user, err := models.FindUser(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "error loging in"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ctx.PostForm("password")+"my secret pepper")) != nil {
		ctx.JSON(401, gin.H{"error": "User not Authorized"})
		return
	}
	token := jwtLib.New(jwtLib.GetSigningMethod("HS256"))
	token.Claims["ID"] = user.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // expires in a week
	tokenString, err2 := token.SignedString([]byte(SECRET))
	if err2 != nil {
		ctx.JSON(500, gin.H{"error": "Problem generating token"})
	}
	user.Password = ""
	ctx.JSON(200, gin.H{"user": user, "token": tokenString})
}

func UserAdd(ctx *gin.Context) {
	user, err := models.AddUser(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	token := jwtLib.New(jwtLib.GetSigningMethod("HS256"))
	token.Claims["ID"] = user.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // expires in a week
	tokenString, err2 := token.SignedString([]byte(SECRET))
	if err2 != nil {
		ctx.JSON(500, gin.H{"error": "Problem generating token"})
	}
	user.Password = ""
	ctx.JSON(200, gin.H{"user": user, "token": tokenString, "success": "New user added"})
}

func UserRemove(ctx *gin.Context) {
	_, err := models.RemoveUser(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "User removed"})
}

func CourseIndex(ctx *gin.Context) {
	courses, err := models.AllCourses(ctx)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(404, gin.H{"error": "No courses found"})
		return
	}
	ctx.JSON(200, courses)
}

func CourseAdd(ctx *gin.Context) {
	_, err := models.AddCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "New course added"})
}

func CourseUpdate(ctx *gin.Context) {
	_, err := models.UpdateCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error processing request"})
		return
	}
	ctx.JSON(200, gin.H{"success": "Course updated"})
}

func CourseRemove(ctx *gin.Context) {
	_, err := models.RemoveCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "Course removed"})
}
