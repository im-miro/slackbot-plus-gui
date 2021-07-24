package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Command_number   int
	Command          string
	Command_response string
	Token_number     int
}

type Slacktoken struct {
	gorm.Model
	Token_number             int
	Token_identifier_message string
	Token                    string
}

func dbGetall_command() []Command {
	db, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var command []Command
	db.Select("*").Order("created_at desc").Find(&command)
	return command
}

func dbGetall_slacktoken() []Slacktoken {
	db, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var slacktoken []Slacktoken
	db.Select("*").Order("created_at desc").Find(&slacktoken)
	return slacktoken
}

func dbcreate_command(command_number int, command string, command_response string, token_number int) Command {
	db, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Command{Command_number: command_number, Command: command, Command_response: command_response, Token_number: token_number})
	return Command{}
}

func dbcreate_slacktoken(token_number int, token_identifier_message string, token string) Slacktoken {
	db, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Slacktoken{Token_number: token_number, Token_identifier_message: token_identifier_message, Token: token})
	return Slacktoken{}
}

func get_slacktoken_number() (time_int int) {
	time := time.Now()
	const layout_for_time_to_token_number = "20060102150405"
	time_int, _ = strconv.Atoi(time.Format(layout_for_time_to_token_number))
	return time_int

}

func main() {

	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*.html")
	engine.GET("/token_register", func(c *gin.Context) {
		token_data := dbGetall_slacktoken()
		fmt.Print(token_data)
		c.HTML(http.StatusOK, "token.html", gin.H{
			"data_tokens": token_data,
		})
	})

	engine.POST("/post_token", func(c *gin.Context) {
		in_token_name := c.PostForm("token_name")
		in_token := c.PostForm("token")
		time_int := get_slacktoken_number()
		fmt.Print(time_int, in_token_name, in_token)
		dbcreate_slacktoken(time_int, in_token_name, in_token)
		c.Redirect(302, "/token_register")
	})

	engine.Static("/templates", "./templates")

	engine.Run(":8080")

	////////////////
	// DB初期化/////
	////////////////

	// db_command, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db_command.AutoMigrate(&Command{})

	// db_command.Create(&Command{Command_number: 0, Command: "reset DB", Command_response: "reset DB", Token_number: 0})

	// db_slacktoken, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db_slacktoken.AutoMigrate(&Slacktoken{})

	// db_slacktoken.Create(&Slacktoken{Token_number: 0, Token_identifier_message: "this is test token", Token: "xoxb-779435090240-2277363653056-3emRAthdzTNSLn3mtNIgaCba"})

}
