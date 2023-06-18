package service

import (
	"github.com/golang-jwt/jwt/v5"
	"main/dtos"
	"main/utils"
	"time"
)

type ApiKeyService struct {
}

func (service *ApiKeyService) GenerateApiKey(user, role string, apiKeyDto *dtos.ApiKeyDto) (string, error) {
	var secretKey = []byte("SECRETAPIKEYFORBOOKINGAPP")
	var claims utils.Claims
	if apiKeyDto.Limited {
		claims = limitedApiKey(user, role, apiKeyDto.TimeLimit)
	} else {
		claims = lastingApiKey(user, role)
	}

	apiKeyString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	apiKey, err := apiKeyString.SignedString(secretKey)

	return apiKey, err
}

func limitedApiKey(user, role string, hours int32) utils.Claims {
	duration := time.Duration(hours) * time.Hour
	return utils.Claims{
		CustomClaims: map[string]string{
			"username": user,
			"role":     role,
			"goal":     "buyingTickets",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}

func lastingApiKey(user, role string) utils.Claims {
	return utils.Claims{
		CustomClaims: map[string]string{
			"username": user,
			"role":     role,
			"goal":     "buyingTickets",
		},
	}
}
