package usercontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	var user models.User

	// Bind other form data fields
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	if user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email dan password tidak boleh kosong"})
		return
	}

	// hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to hash password",
		})
		return
	}
	
	user = models.User{Email: user.Email, Password: string(hash)}
	
	result := models.DB.Create(&user)


	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"user": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}


func Login(c *gin.Context){
	var req_user models.User

	// Bind other form data fields
	req_user.Email = c.PostForm("email")
	req_user.Password = c.PostForm("password")

	if req_user.Email == "" || req_user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email dan password tidak boleh kosong"})
		return
	}

	
	var user models.User
	models.DB.First(&user, "email = ?", req_user.Email)

	if user.Id == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Email tidak ditemukan",
		})

		return
	}

	err :=	bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req_user.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to create token",
		})

		return
	}

	// send it back

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "","", false, true)
		
	c.JSON(http.StatusOK, gin.H{
		"token":tokenString,
	})
}
func Logout(c *gin.Context) {
	// Clear the authentication cookie
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func Validate( c *gin.Context){
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message" : user,
	})
}