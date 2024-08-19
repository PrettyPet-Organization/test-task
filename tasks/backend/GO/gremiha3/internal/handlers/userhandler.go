package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/KozlovNikolai/test-task/internal/model"
	"github.com/KozlovNikolai/test-task/internal/pkg/utils"
	"github.com/KozlovNikolai/test-task/internal/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler is ...
type UserHandler struct {
	logger *zap.Logger
	repoWR store.IUserRepository
	repoRO store.IUserRepository
}

// NewUserHandler is ...
func NewUserHandler(logger *zap.Logger, repoWR store.IUserRepository, repoRO store.IUserRepository) *UserHandler {
	return &UserHandler{
		logger: logger,
		repoWR: repoWR,
		repoRO: repoRO,
	}
}

// CreateUser is ...
// CreateUserTags		godoc
// @Summary				Register user
// @Description			Save register data of user in Repo.
// @Param				user body model.AddUser true "Create user. Логин указывается в формате электронной почты. Пароль не меньше 6 символов. Роль: super или regular"
// @Produce				application/json
// @Tags				user
// @Success				200 {object} model.User
// @failure				400 {string} err.Error()
// @failure				500 {string} err.Error()
// @Router				/user/register [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	var addUser model.AddUser
	// Заполняем структуру addUser данными из запроса
	if err := c.ShouldBindJSON(&addUser); err != nil {
		u.logger.Error("Error binding JSON-addUser", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	// Заполняем структуру User данными из addUser
	user.Login = addUser.Login
	user.Password = addUser.Password
	user.Role = addUser.Role

	// Валидация данных пользователя
	if err := user.Validate(); err != nil {
		u.logger.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хеширование пароля
	if err := user.HashPassword(); err != nil {
		u.logger.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Сохраняем пользователя в базе с проверкой на уникальность email
	id, err := u.repoWR.Create(context.TODO(), user)
	if err != nil {
		u.logger.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// загружаем из базы информацию о сохраненном пользователе:
	user, err = u.repoRO.GetByID(context.TODO(), id)
	if err != nil {
		u.logger.Error("Error load user from db after saving", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUser is ...
// GetUserTags 		godoc
// @Summary			Get user by user id.
// @Description		Return user with "id" number.
// @Param			user_id path int true "User ID"
// @Tags			Auth
// @Security     	BearerAuth
// @Success			200 {object} model.User
// @failure			404 {string} err.Error()
// @Router			/user/{user_id} [get]
func (u *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	authLogin, authRole := utils.GetLevel(c)
	u.logger.Debug("принятые логин и роль из токена", zap.String("login", authLogin), zap.String("role", authRole))

	// если запрос делает суперпользователь, то ему можно всё
	if authRole == "super" {
		user, err := u.repoRO.GetByID(context.TODO(), id)
		if err != nil {
			u.logger.Error("Error getting user", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	} else if authRole == "regular" { // если запрос делает обычный пользователь, то ему можно смотреть только собственные данные
		user, err := u.repoRO.GetByLogin(context.TODO(), authLogin)
		if err != nil {
			u.logger.Error("Error getting user", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if user.ID != id {
			u.logger.Error("не хватает уровня доступа")
			c.JSON(http.StatusNotFound, gin.H{"error": "не хватает уровня доступа"})
			return

		}
		c.JSON(http.StatusOK, user)
	} else {
		u.logger.Error("Error level role", zap.String("level error", "role not super and not regular"))
		c.JSON(http.StatusNotFound, gin.H{"level error": "role not super and not regular"})
	}

}

// LoginUser is ...
// LoginUserTags 		godoc
// @Summary				Login user by login=email/password.
// @Description			Returns the authorization status
// @Param				user body model.LoginUser true "Login user. Логин указывается в формате электронной почты. Пароль не меньше 6 символов. Роль: super или regular"
// @Accept       		json
// @Produce				json
// @Tags				Auth
// @Success      		200 {string} string "Logged in"
// @failure				400 {string} err.Error()
// @failure				500 {string} err.Error()
// @Router				/user/login [post]
func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginUser model.LoginUser
	if err := c.BindJSON(&loginUser); err != nil {
		u.logger.Error("Error binding JSON-loginUser", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find login in DB
	user, err := u.repoRO.GetByLogin(context.TODO(), loginUser.Login)
	if err != nil {
		u.logger.Error("Error. User not find", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check username and password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		u.logger.Error("Error. Password is wrong", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error. Password is wrong"})
		return
	}
	// Create JWT token
	// expirationTime := time.Now().Add(config.Cfg.TokenTimeDuration)
	tokenString, err := utils.GenerateJWT(user.Login, user.Role)
	if err != nil {
		u.logger.Error("Error. Could not create token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
	}
	// // Set cookie
	// http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:     "token",
	// 	Value:    tokenString,
	// 	Expires:  expirationTime,
	// 	HttpOnly: true,
	// })

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "token": tokenString})
}

// GetUsers is ...
// GetUsersTags 		godoc
// @Summary			Get all users.
// @Description		Return users list.
// @Tags			user
// @Security     	BearerAuth
// @Produce      json
// @Success			200 {object} []model.User
// @failure			404 {string} err.Error()
// @Router			/users [get]
func (u *UserHandler) GetUsers(c *gin.Context) {
	authLogin, authRole := utils.GetLevel(c)
	u.logger.Debug("принятые логин и роль из токена", zap.String("login", authLogin), zap.String("role", authRole))

	// если запрос делает суперпользователь, то ему можно всё
	if authRole == "super" {
		users, err := u.repoRO.GetAll(context.TODO())
		if err != nil {
			u.logger.Error("Error getting users", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
			return
		}
		c.JSON(http.StatusOK, users)
		return
	}
	u.logger.Error("не хватает уровня доступа")
	c.JSON(http.StatusNotFound, gin.H{"error": "не хватает уровня доступа"})

}
