package main


import (
	"log"
	"net/http"
	"os"


	"database/sql"

	_ "github.com/lib/pq"


	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
	"fmt"

	_ "io/ioutil"

)

type Product struct {
	// gorm.Model
	Code string
	Price uint
	Fullname string
	Email string
}

//func repeatHandler(c *gin.Context) {
//	var buffer bytes.Buffer
//	for i := 0; i < repeat; i++ {
//		buffer.WriteString("Hello from Go!\n")
//	}
//	c.String(http.StatusOK, buffer.String())
//}


func main() {

	// 		"host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")

	var args string
	args += "host=" + os.Getenv("myhost") + " "
	args += "user=" + os.Getenv("user") + " "
	args += "dbname=" + os.Getenv("dbname") + " "
	args += "sslmode=disable "
	args += "password=" + os.Getenv("password")
	log.Println("args is: ", args)


	var URI = os.Getenv("URI")
	db, errDB := sql.Open("postgres", URI)
	defer db.Close()

	if errDB != nil {
		log.Fatalf("Error connecting to the DB")
	} else {
		log.Println("Connection is successful!")
	}

	rows, errQuery := db.Query(`SELECT 123 * 321 as result `)
	if errQuery != nil {
		log.Println(errQuery)
	}
	defer rows.Close()

	for rows.Next() {
		var result int
		if err := rows.Scan(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result is %d\n", result)
	}

	//if !db.HasTable("products") {
	//	db.CreateTable(&Product{})
	//	db.AutoMigrate(&Product{})
	//
	//}

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})


	router.GET("/onlinetraces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})


	router.GET("/traces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/fileupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileupload.tmpl.html", nil)
	})

	router.GET("/online", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.POST("/fileupload", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("File is uploaded!"))
	})

	router.Run(":" + port)
}