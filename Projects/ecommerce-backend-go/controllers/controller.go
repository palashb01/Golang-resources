package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/palashb01/ecommerce-backend-go/models"
	"github.com/palashb01/ecommerce-backend-go/database"
	"github.com/palashb01/ecommerce-backend-go/tokens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

// HashPassword
// VerifyPassword
// Login
// SignUP
// ProductViewerAdmin
// SearchProduct()
// SearchProductByQuery()

func HashPassword(password string) string {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panicln(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	if err != nil {
		return false, "Login or password is incorrect"
	}
	return true, ""
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validateErr := Validate.Struct(user)
		if validateErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}
		count,err := UserCollection.CountDocuments(ctx,bson.M{"email": user.Email},nil)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0{
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}
		count,err = UserCollection.CountDocuments(ctx,bson.M{"phone": user.Phone},nil)
		defer cancel()
		if err!=nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone number already exists"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := tokens.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_ , err = UserCollection.InsertOne(ctx, user)
		if err!=nil{
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user item was not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully signed-in")
	}
}

func Login() gin.HandlerFunc {
	return func (c *gin.Context){
		ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Second)
		defer cancel()
		var user models.User
		if err:=c.BindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var founduser models.User
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}
		PassIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !PassIsValid{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := tokens.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
		defer cancel()
		tokens.UpdateAllTokens(token,refreshToken,founduser.User_ID)
		c.JSON(http.StatusFound,founduser)
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productlist []models.Product
		ctx,cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err:= ProductCollection.Find(ctx, bson.D{})
		if err!=nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		err = cursor.All(ctx, &productlist)
		if err!=nil{
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError,err)
			return
		}
		defer cursor.Close(ctx)
		if err := cursor.Err(); err!=nil{
			log.Panicln(err)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, productlist)
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context){
		var SearchProducts []models.Product
		queryParam := c.Query("name")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		searchquerydb, err := ProductCollection.Find(ctx, bson.M{"name":bson.M{"$regex":queryParam}})
		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "Something went wrong while fetching the data")
			return
		}
		err = searchquerydb.All(ctx, &SearchProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "Invalid request")
			return
		}
		defer searchquerydb.Close(ctx)
		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(500, "Internal server error")
			return
		}
		c.IndentedJSON(200, SearchProducts)
	}
}
