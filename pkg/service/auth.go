package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "fdklskl324kfdo"
	signingKey = "dsmvi4k5o2gjrbv40jf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, userpassword string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, generatePasswordHash(userpassword))
	if err != nil {
		return "", err //TODO заменить потом на acsess b reftoken
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
