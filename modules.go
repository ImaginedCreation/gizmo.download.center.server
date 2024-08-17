package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	RESCODEOK    RESCODE = 200
	RESCODEERROR RESCODE = 500
)

type RES[T any] struct {
	Code    RESCODE `json:"code"`
	Message string  `json:"message"`
	Data    T       `json:"data"`
}

type Register_P struct {
	UserName        string `json:"username" binding:"required"`
	NickName        string `json:"nick_name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type Login_P struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CONFIG struct {
	PORT          string         `json:"port"`
	SQL           SQL            `json:"sql"`
	TIMELOCALTION *time.Location `json:"time_local"`
}

type SQL struct {
	IP        string `json:"ip"`
	PORT      string `json:"port"`
	USER      string `json:"user"`
	PASS      string `json:"pass"`
	DATABASE  string `json:"database"`
	IDLE      int    `json:"idle"`
	OPENCOUNT int    `json:"open_count"`
	LIFETIME  int    `json:"life_time"`
}

type Token struct {
	USERNAME string `json:"username"`
	VERSION  string `json:"version"`
}

type Claims struct {
	Token
	jwt.RegisteredClaims
}

func GenerateToken(p *Token) (string, error) {
	claims := Claims{
		*p,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t_s, err := token.SignedString([]byte(G_SECRET))

	if err != nil {
		return "", err
	}
	return t_s, nil
}

func OnLoad() {
	__InitSQL()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("time location load failed %v \n", err.Error())
	} else {
		time.Local = loc
		G_CONFIG.TIMELOCALTION = loc
	}
}

var __DB__ *gorm.DB

func __InitSQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		G_CONFIG.SQL.USER,
		G_CONFIG.SQL.PASS,
		G_CONFIG.SQL.IP,
		G_CONFIG.SQL.PORT,
		G_CONFIG.SQL.DATABASE,
	)

	logfile, err := os.Create("sql.log")
	if err != nil {
		fmt.Printf("log file create failed \n")
		panic(err.Error())
	}

	__logger__ := logger.New(
		log.New(logfile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: __logger__,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Printf("sql connect failed \n")
		panic(err.Error())
	}

	sql_db, _ := db.DB()

	sql_db.SetMaxIdleConns(G_CONFIG.SQL.IDLE)

	sql_db.SetMaxOpenConns(G_CONFIG.SQL.OPENCOUNT)

	sql_db.SetConnMaxLifetime(time.Duration(time.Duration(G_CONFIG.SQL.LIFETIME) * time.Hour))

	__DB__ = db

	__DB__.AutoMigrate(&USER{})

	fmt.Printf("sql connect successfully \n")

}

func NewDBClient(c *context.Context) *gorm.DB {
	return __DB__.WithContext(*c)
}

func ParseValidatorError(err error) string {
	if err.Error() == "EOF" {
		return "please fill request parameters ."
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		return "paramter validator failed ."
	}
	return err.Error()
}

type RESCODE int

func RES_OK_FN[T any](data T) RES[T] {
	return RES[T]{
		Code:    RESCODEOK,
		Message: "ok",
		Data:    data,
	}
}

func RES_ERROR_FN(error_str string) RES[string] {
	return RES[string]{
		Code:    RESCODEERROR,
		Message: error_str,
		Data:    "",
	}
}
