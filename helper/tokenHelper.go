package helper

import (
	"context"
	"log"
	"os"
	"resturnat-management/database"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.RegisteredClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SecretKey = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email, firstName, lastName, uid string) (string, string, error) {
	claims := &SignedDetails{
		First_Name: firstName,
		Last_Name:  lastName,
		Email:      email,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(24 * time.Hour)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(168 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId}
	update := bson.M{
		"$set": bson.M{
			"token":         signedToken,
			"refresh_token": signedRefreshToken,
		},
	}

	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx, filter, update, &opts)
	if err != nil {
		log.Printf("Failed to update tokens for user %s: %v", userId, err)
		return

	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	// parsing the token
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (any, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil || token == nil {
		if err != nil {
			msg = err.Error()
		} else {
			msg = "invalid token"
		}
		return
	}

	// if it is not vali
	claims, ok := token.Claims.(*SignedDetails)
	if !ok || claims == nil {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt == nil {
		msg = "token has no expiry"
		return
	}

	// if it expired
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}

	return claims, msg
}
