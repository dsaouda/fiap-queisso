package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func PostLogin(c *gin.Context) {

	//recebendo post
	login := models.Login{}
	c.Bind(&login)

	db := c.MustGet("db").(*mgo.Database)

	result := models.Login{}
	db.C("login").Find(bson.M{"email": login.Email, "password": login.Password}).One(&result)

	if (models.Login{}) == result {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login e/ou senha não é válido"})
		return
	}

	c.JSON(http.StatusOK, result)
}

