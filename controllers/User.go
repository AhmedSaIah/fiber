package controllers

import (
	"fmt"
	"github.com/AhmedSaIah/fiber/models"
	"github.com/AhmedSaIah/fiber/repository"
	"github.com/AhmedSaIah/fiber/util"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/asaskevich/govalidator.v9"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	//GetUser(ctx *fiber.Ctx) (models.User, error)
	//GetUsers(ctx *fiber.Ctx) ([]models.User, error)
	//PutUser(ctx *fiber.Ctx) error
}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{userRepo: userRepo}
}

func (u *userController) SignUp(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error parsing input": err})
	}
	user.Email = NormalizeEmail(user.Email)
	if !govalidator.IsEmail(user.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "please enter an actual email address"})
	}
	userExists, err := u.userRepo.GetByEmail(user.Email)
	if err == mongo.ErrNoDocuments {
		if strings.TrimSpace(user.Password) == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password can not be empty"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password does not match"})
		}
		user.Password = string(hashedPassword)
		if err = u.userRepo.Save(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user creation failed"})
		}
		return ctx.Status(fiber.StatusCreated).JSON(user)
	}
	if userExists != nil {
		err = util.ErrEmailAlreadyExists
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Signing up failed"})
}

func (u *userController) SignIn(ctx *fiber.Ctx) error {
	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

//func (u *userController) GetUser(ctx *fiber.Ctx) (models.User, error) {
//	return models.User{}, nil
//}
//
//func (u *userController) GetUsers(ctx *fiber.Ctx) ([]models.User, error) {
//	return nil
//}
//
//func (u *userController) PutUser(ctx *fiber.Ctx) error {
//	return nil
//}

//func UserRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
//	id := ctx.Params("id")
//	objectID, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized user"})
//	}
//	token := ctx.Locals("user").(*jwt.Token)
//	payload, ok := token.Claims(*jwt.)
//
//}

//func NewUserController(userRepo *repository.UserRepository) *UserController {
//	return &UserController{userRepo}
//}

//func (u *UserController) SignUp(c *fiber.Ctx) error {
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
//
//	userExists, err := u.userRepository.FindByEmail(user.Email)
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
//
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
//func (u *UserController) Login(c *fiber.Ctx) error {
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

func validateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" "+err.Tag())
		}
		return fmt.Errorf(strings.Join(errors, ", "))
	}
	return nil
}

func NewToken(user models.User) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	tokenSign := jwt.New(jwt.SigningMethodHS256)
	expires := time.Now().Add(time.Hour * 24).Unix()

	claims := tokenSign.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := tokenSign.SignedString(secret)
	if err != nil {
		return "", 0, nil
	}

	return token, expires, err
}
