package db

import (
	"database/sql"
	"log"
	"fmt"
	"time"
    "errors"
    "strconv"

    "github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
    "Auth-Service-Rest-Api/internal/models"
    "Auth-Service-Rest-Api/internal/auth"
    "golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var jwtKey = []byte("1234")

func ConnectPostgresDB() error {
    connStr := "postgres://postgres:1234@host.docker.internal:5432/auth_service?sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        return err
    }
    err = db.Ping()
    if err != nil {
        return err
    }
    log.Println("Connected to PostgreSQL database")
    return nil
}

func ClosePostgresDB() {
    if db != nil {
        db.Close()
        log.Println("Disconnected from PostgreSQL database")
    }
}

func GetPostgresDB() (*sql.DB, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func WaitWhileDBNotReady() {
    fmt.Println("Waiting for database to be ready...")
    for {
        if err := db.Ping(); err == nil {
            fmt.Println("Database is ready!")
            break
        }
        fmt.Println("Database is not ready, waiting...")
        time.Sleep(time.Second)
    }
}

func RegisterUser(user models.User) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE Email = $1", user.Email).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("ошибка при проверке существования пользователя: %v", err)
	}
	if count > 0 {
		return 0, errors.New("Email уже занят")
	}

	if !auth.IsEmailValid(user.Email) || !auth.IsPasswordSafe(user.Password) {
		return 0, errors.New("Некорректный email или пароль")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("ошибка при генерации пароля: %v", err)
	}

	query := "INSERT INTO users(Email, Password) VALUES($1, $2) RETURNING Id"
	err = db.QueryRow(query, user.Email, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при регистрации пользователя: %v", err)
	}

	return user.ID, nil
}

func AuthorizeUser(user models.User) (string, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT Password FROM Users WHERE Email = $1", user.Email).Scan(&hashedPassword)
	if err != nil {
		return "", fmt.Errorf("ошибка при поиске пользователя: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", errors.New("неправильный пароль")
	}

	var userID int
	err = db.QueryRow("SELECT Id FROM users WHERE Email = $1", user.Email).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении ID пользователя: %v", err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(userID),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %v", err)
	}

	return tokenString, nil
}
