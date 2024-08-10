package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgrijalva/jwt-go"

)

type User struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Name		    *string			`json:"name" validate:"required,min=2,max=100"`
	Username		*string			`json:"username" validate:"required,min=2,max=100"`
	Password		*string			`json:"Password" validate:"required,min=6"`
	Email			*string			`json:"email" validate:"email,required"`
	User_type		string			`json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Created_at		time.Time		`json:"created_at"`
	Updated_at		time.Time		`json:"updated_at"`
	User_id			string			`json:"user_id"`
}

type UserClaim struct{
	User_id			string			
	Username		string
	Email			string
	User_type		string
	jwt.StandardClaims
}