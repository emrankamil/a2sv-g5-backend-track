package data

import (
	"context"
	"errors"
	"log"
	"os"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var userCollection *mongo.Collection
var validate = validator.New()

func InitUser(){
	userCollection = OpenCollection(Client, "users")
}

func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = "password is incorrect"
		check=false
	}
	return check, msg
}

func ValidateToken(signedToken string) (claims *models.UserClaim, err error){
	token, msg := jwt.ParseWithClaims(
		signedToken, 
		&models.UserClaim{}, 
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if msg != nil || !token.Valid{
		err = msg
		return
	}

	claims, ok:= token.Claims.(*models.UserClaim)
	if !ok{
		err = errors.New("the token is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		err = errors.New("token is expired")
		return
	}
	return claims, err
}

func GenerateJWTToken(user_id string, username string, email string, user_type string) (signedToken, signedRefreshToken string, err error){
	claims := &models.UserClaim{
		User_id: user_id,
		Username: username,
		Email: email,
		User_type: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &models.UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
    if err != nil {
        return "", "", err
    }

    signedRefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
    if err != nil {
        return "", "", err
    }

    return
}

func CheckUserType(c *gin.Context, role string) (err error) {
    userType := c.GetString("user_type")
    err = nil
    if userType != role {
        err = errors.New("unauthorized to access this resource")
        return err
    }
    return
}

func HandleSignup(user *models.User) error{
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if validationErr := validate.Struct(user) ; validationErr != nil{
		return errors.New(validationErr.Error())
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"username":user.Username})
	
	if err != nil {
		log.Panic(err)
		return errors.New(err.Error())
	}
	if count > 0{
		return errors.New("this username already exists")
	}

	count, err = userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
	// defer cancel()
	if err!= nil {
		log.Panic(err)
		return errors.New(err.Error())
	}

	if count > 0{
		return errors.New("this email already exists. ")
	}

	password := HashPassword(*user.Password)
	user.Password = &password
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex() 

	countUsers, err := userCollection.CountDocuments(ctx, bson.M{})
	if err!= nil {
		log.Panic(err)
		return errors.New(err.Error())
	}
	if countUsers == 0{
		user.User_type = "ADMIN"
	}else{
	user.User_type = "USER"
	}

	_, insertionErr := userCollection.InsertOne(ctx, user)
	if insertionErr != nil{
		return errors.New("user registration failed")
	}
	var found error
	found = userCollection.FindOne(ctx, bson.M{"username":user.Username}).Decode(&found)
	// defer cancel()
	return nil
}

func HandleLogin(user *models.User) (signedToken, signedRefreshToken string, err error){
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var foundUser models.User

	msg := userCollection.FindOne(ctx, bson.M{"username":user.Username}).Decode(&foundUser)
	if msg != nil{
		err = errors.New("user not found")
		return
	}
	if foundUser.Username == nil{
		err = errors.New("user not found")
		return
	}

	check, verifMsg := VerifyPassword(*user.Password, *foundUser.Password)
	// defer cancel()
	if !check{
		err = errors.New(verifMsg)
		return
	}

	if foundUser.Username == nil || foundUser.Email == nil || foundUser.User_type == "" {
		err = errors.New("invalid user data")
		return
	}	

	token, refreshToken, err := GenerateJWTToken(foundUser.User_id, *foundUser.Username, *foundUser.Email,  foundUser.User_type)
	return token, refreshToken, err
}

func HandleGetUsers() (users []models.User, err error){
	cur, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	for cur.Next(context.TODO()) {
		var user models.User

		val := cur.Decode(&user)
		if val != nil {
			users = []models.User{}
			err = errors.New(val.Error())
			return
		}

		users = append(users, user)
	}

	if cur_err := cur.Err(); cur_err != nil {
		users = []models.User{}
		err = errors.New(cur_err.Error())
		return
	}

	cur.Close(context.TODO())
	return users, nil
}

func UpdateUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("INVALID ID")
    }

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "user_type", Value: "ADMIN"},
		}},
	}
	updateResult, result := userCollection.UpdateOne(context.TODO(), filter, update)
	if result != nil {
		return errors.New(result.Error())
	}
	if updateResult.MatchedCount == 0{
		return errors.New("USER NOT FOUND")
	}
	return nil
}
