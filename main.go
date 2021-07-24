package main

import (
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
	Token_number             int
	Token_identifier_message string
	Token                    string
}

// func dbGetall_command() []Command {
// 	db, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	var command []Command
// 	db.Select("*").Order("created_at desc").Find(&command)
// 	return command
// }

// func dbGetall_slacktoken() []Slacktoken {
// 	db, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	var slacktoken []Slacktoken
// 	db.Select("*").Order("created_at desc").Find(&slacktoken)
// 	return slacktoken
// }

// func dbcreate_command(command_number int, command string, command_response string, token_number int) Command {
// 	db, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.Create(&Command{Command_number: command_number, Command: command, Command_response: command_response, Token_number: token_number})
// 	return Command{}
// }

// func dbcreate_slacktoken(token_number int, token_identifier_message string, token string) Slacktoken {
// 	db, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.Create(&Slacktoken{Token_number: token_number, Token_identifier_message: token_identifier_message, Token: token})
// 	return Slacktoken{}
// }

func main() {

	// engine := gin.Default()

	// user_agent := ""

	// engine.Use(func(c *gin.Context) {
	// 	user_agent = c.GetHeader("User-Agent")
	// 	rex := regexp.MustCompile(".*")
	// 	user_agent = rex.FindString(user_agent)
	// 	c.Next()
	// })

	// engine.LoadHTMLGlob("templates/*.html")
	// engine.GET("/show_msg", func(c *gin.Context) {
	// 	info_100 := dbGet100all()
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"data_name": info_100,
	// 	})
	// })

	// engine.POST("/post", func(c *gin.Context) {
	// 	in_name := c.PostForm("Name")
	// 	in_msg := c.PostForm("Message")
	// 	dbcreate(in_name, in_msg)
	// 	c.Redirect(302, "/show_msg")
	// })

	// engine.Static("/img", "./img")
	// engine.Static("/templates", "./templates")

	// engine.Run(":8080")

	////////////////
	// DB初期化/////
	////////////////

	db_command, err := gorm.Open(sqlite.Open("command.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db_command.AutoMigrate(&Command{})

	db_command.Create(&Command{Command_number: 0, Command: "reset DB", Command_response: "reset DB", Token_number: 0})

	db_slacktoken, err := gorm.Open(sqlite.Open("slacktoken.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db_slacktoken.AutoMigrate(&Slacktoken{})

	db_slacktoken.Create(&Slacktoken{Token_number: 0, Token_identifier_message: "this is test token", Token: "xoxb-779435090240-2277363653056-3emRAthdzTNSLn3mtNIgaCba"})

}
