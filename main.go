package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Company struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    CEO string  `json:"ceo"`
    Revenue  string `json:"revenue"`
}

var companies = []Company{
    {ID: "1", Name: "Dell", CEO: "Michael Dell", Revenue: "92.2 billion"},
    {ID: "2", Name: "Netflix", CEO: "Reed Hastings", Revenue: "20.2 billion"},
    {ID: "3", Name: "Microsoft", CEO: "Satya Nadella", Revenue: "320 million"},
}



func HomepageHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message":"Welcome to the Tech Company listing API with Golang"})
}

func NewCompanyHandler(c *gin.Context) {
	var newCompany Company

	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newCompany.ID = xid.New().String()
	companies = append(companies, newCompany)
	c.JSON(http.StatusCreated,  newCompany)
}


func GetCompaniesHandler(c *gin.Context) {
    c.JSON(http.StatusOK, companies)
}

func UpdateCompanyHandler(c *gin.Context) {
	id := c.Param("id")
	var company Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1

	for i := 0; i < len(companies); i++ {
		if companies[i].ID == id {
			index = 1
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}

	companies[index] = company
	c.JSON(http.StatusOK, company)
}


func DeleteCompanyHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1

	for i := 0; i < len(companies); i++ {
		if companies[i].ID == id {
			index = 1
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}

	companies = append(companies[:index], companies[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Company has been deleted",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", HomepageHandler)
	router.GET("/companies", GetCompaniesHandler)
	router.POST("/company", NewCompanyHandler)
	router.PUT("/company/:id", UpdateCompanyHandler)
	router.DELETE("/company/:id", DeleteCompanyHandler)
	router.Run()
}