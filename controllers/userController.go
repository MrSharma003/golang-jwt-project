package controllers

import (
	"github.com/MrSharma003/GO-JWT/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate *validator.Validate = validator.New()

func HashPassword()

func VerifPassword()

func Signup()

func Login()

func GetUsers()

func GetUserById()
