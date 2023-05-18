package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"strconv"
	"context"
	"net/http"
	"github.com/go-playground/validator/v10"
	helper "github.com/alvo254/locker/helpers"
	"github.com/alvo254/locker/models"
	"github.com/alvo254/locker/helpers"
	"golang.org/x/crypto/bcrypt"

)

var userCollection *mongo.Collection = database.OpenCollection(database.client, "user")
var validate = validator.New()

func HashPassword(){

}

func VerifyPassword(){

}

func SignUp(){

}

func Login(){

}

func GetUsers(){

}

//Only admin can acess
func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		//user_id comes from routes (specified :user_id)
		userId := c.Param("user_id")
		helper.MatchUserTypeToUid(c, userId); error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		}
	}

}




