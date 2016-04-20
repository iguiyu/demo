package main

import (
	"log"
	"net/http"

	"github.com/kobeld/goutils"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kobeld/mgowrap"
)

// ========================
// === Model definition ===
//
type User struct {
	Id     bson.ObjectId `bson:"_id"`
	Name   string
	Email  string
	Avatar string
}

// To use the mgowrap, should implement following two methods defined in the "mgowrap.PersistentObject" interface.
func (this *User) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *User) CollectionName() string {
	return "users"
}

// Set the user avatar with Gravatar service
func (this *User) setAvatar() {
	if this.Avatar == "" {
		this.Avatar = goutils.UrlSize(this.Email, 60)
	}
}

// =========================
// === Handler functions ===
//
func userIndex(c *gin.Context) {

	var (
		err   error
		users = []*User{}
	)

	// Find all users
	err = mgowrap.FindAll(bson.M{}, &users)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Set avatar for all users
	for _, user := range users {
		user.setAvatar()
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"users": users})
}

func userNew(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}

func userShow(c *gin.Context) {

	var (
		err   error
		user  = &User{}
		idHex = c.Param("id")
	)

	// Find a user by plain Id
	err = mgowrap.FindByIdHex(idHex, &user)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	user.setAvatar()

	c.HTML(http.StatusOK, "show.html", user)

}

func userCreate(c *gin.Context) {

	var (
		err  error
		user = &User{
			Id:    bson.NewObjectId(),
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
		}
	)

	err = mgowrap.Save(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/user/"+user.Id.Hex())
}

// ==================================
// The main logic for this simple app
//
func main() {

	// Setup DB
	mgowrap.SetupDatbase("localhost", "userdb")

	// Init Gin and setup the router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	{
		router.GET("/", userIndex)
		router.GET("/new", userNew)
		router.GET("/user/:id", userShow)
		router.POST("/user", userCreate)
	}

	log.Printf("Listening and serving http on %s\n", ":8081")
	router.Run(":8081")
}
