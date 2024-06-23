package usecase

import (
	"context"
	"errors"

	"next-learn-go-sqlc/db/sqlc"
	"next-learn-go-sqlc/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user sqlc.User) error
	Login(user sqlc.User) (LoginResponse, error)
	GetUserById(userId pgtype.UUID) (UserResponse, error)
	GetUserByEmail(email string) (UserResponse, error)
	RefreshToken(refreshTokenString string) (TokenResponse, error)
}

type userUsecase struct {
	uq sqlc.Querier
	uv validator.IUserValidator
}

type LoginResponse struct {
	ID           pgtype.UUID `json:"id"`
	Email        string      `json:"email"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserResponse struct {
	ID       pgtype.UUID `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserUsecase(uq sqlc.Querier, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{uq, uv}
}

func (uu *userUsecase) SignUp(user sqlc.User) error {
	if err := uu.uv.UserValidate(user); err != nil {
		return err
	}

	_, err := uu.uq.GetUserByEmail(context.Background(), user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	arg := sqlc.CreateUserParams{Name: user.Name, Email: user.Email, Password: string(hash)}
	_, err = uu.uq.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) Login(user sqlc.User) (LoginResponse, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		return LoginResponse{}, err
	}
	storedUser := sqlc.User{}
	user, err := uu.uq.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return LoginResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return LoginResponse{}, err
	}

	resLogin := LoginResponse{
		ID:           storedUser.ID,
		Email:        storedUser.Email,
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}
	return resLogin, nil
}

func (uu *userUsecase) GetUserById(userId pgtype.UUID) (UserResponse, error) {
	ctx := context.Background()
	user, err := uu.uq.GetUserById(ctx, userId)
	if err != nil {
		return UserResponse{}, err
	}
	resUser := UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}

func (uu *userUsecase) GetUserByEmail(email string) (UserResponse, error) {
	ctx := context.Background()
	user, err := uu.uq.GetUserByEmail(ctx, email)
	if err != nil {
		return UserResponse{}, err
	}
	resUser := UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return resUser, nil
}

func (uu *userUsecase) RefreshToken(refreshTokenString string) (TokenResponse, error) {

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return TokenResponse{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return TokenResponse{}, errors.New("invalid refresh token")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	newTokenString, err := newToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		AccessToken:  newTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
