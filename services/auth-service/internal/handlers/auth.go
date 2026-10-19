package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"errors"
	"time"

	"auth-service/internal/validations"

	"net/http"

	"auth-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	Service *service.AuthService
}

func NewAuthHandlers(service *service.AuthService) *AuthHandlers {
	return &AuthHandlers{
		Service: service,
	}
}

func (h *AuthHandlers) SignUp(c *gin.Context) {

	var request validations.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body request",
		})
		return
	}

	if err:= validations.ValidateUser(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return 
	}

	account := &models.Account{
		ID: id.String(),
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
		AccountType: request.AccountType,
	}

	user, err := h.Service.SignUpUser(account)

	if err != nil {
		if errors.Is(err, repository.ErrAccountExists){
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user": user,
	})

}

func (h *AuthHandlers) LoginUser(c *gin.Context){
	jwtcfg := config.JWTLoadConfig()
	var request validations.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := validations.ValidateLogin(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	account, err := h.Service.LoginUser(request.Email, request.Password)
	
	
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Invalid Email or password",
        })
        return
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": account.ID,
		"account_type": account.AccountType,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtcfg.JWT_Secret))

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create authentication token",
        })
        return
    }

	secureCookie := true

	c.SetCookie(
		"token",
		tokenString,
		86400,
		"/",
		"",
		secureCookie,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Login successful",
        "user": gin.H{
            "id":          account.ID,
            "username":    account.Username,
            "email":       account.Email,
            "account_type": account.AccountType,
        },
    })
}

func (h *AuthHandlers) Logout (c *gin.Context) {
	 c.SetCookie(
        "token",
        "",
        -1,  // Expire the cookie
        "/",
        "",
        true, // Secure: true in production with HTTPS
        true, // HttpOnly
    )

    c.SetSameSite(http.SameSiteLaxMode)

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Logout successful",
    })
}
