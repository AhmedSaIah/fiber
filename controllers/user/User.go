package user

//
//import (
//	"fmt"
//	"strings"
//	"time"
//
//	"github.com/dgrijalva/jwt-go"
//	"github.com/go-playground/validator/v10"
//	"github.com/gofiber/fiber/v2"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"golang.org/x/crypto/bcrypt"
//
//	"github.com/AhmedSaIah/fiber/database"
//	"github.com/AhmedSaIah/fiber/models"
//)
//
//type UController struct {
//	userRepository *database.UserRepository
//}
//
//func NewUserController(userRepository *database.UserRepository) *UController {
//	return &UController{
//		userRepository: userRepository,
//	}
//}
//
//func (u *UController) SignUp(c *fiber.Ctx) error {
//	var user models.User
//	err := c.BodyParser(&user)
//	if err != nil {
//		return err
//	}
//
//	err = validateStruct(&user)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	}
//	userExists, err := u.userRepository.FindByEmail(user.Email)
//
//	if err != nil {
//		return err
//	}
//
//	if userExists != nil {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//	user.Password = string(hashedPassword)
//	user.ID = primitive.NewObjectID()
//	err = u.userRepository.Save(&models.User{
//		ID:       user.ID,
//		Name:     user.Name,
//		Email:    user.Email,
//		Password: user.Password,
//	})
//	if err != nil {
//		return err
//	}
//	return c.JSON(fiber.Map{"message": "User saved!"})
//}
//
//func (u *UController) Login(c *fiber.Ctx) error {
//	var user models.User
//	err := c.BodyParser(&user)
//	if err != nil {
//		return err
//	}
//	userExists, err := u.userRepository.FindByEmail(user.Email)
//	if err != nil {
//		return err
//	}
//	if userExists == nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user or password"})
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))
//	if err != nil {
//		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
//	}
//
//	token, err := createToken(userExists.ID)
//	if err != nil {
//		return err
//	}
//	return c.JSON(fiber.Map{"token": token})
//}
//
//func validateStruct(s interface{}) error {
//	validate := validator.New()
//	err := validate.Struct(s)
//	if err != nil {
//		var errors []string
//		for _, err := range err.(validator.ValidationErrors) {
//			errors = append(errors, err.Field()+" "+err.Tag())
//		}
//		return fmt.Errorf(strings.Join(errors, ", "))
//	}
//	return nil
//}
//
//func createToken(userID primitive.ObjectID) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//	claims["userID"] = userID
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	secret := []byte("secret-token")
//	tokenString, err := token.SignedString(secret)
//	if err != nil {
//		return "", nil
//	}
//
//	return tokenString, err
//}
