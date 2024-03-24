package controller

import "github.com/gin-gonic/gin"
import "go-code-generate/quick-start-demo/example/service/vo"
import "go-code-generate/quick-start-demo/example/controller/r"

func CreateAuthorsHandler(c *gin.Context) {
	if err := NewAuthorsController(c).DoCreateAuthors(); err != nil {
		r.BadResponse(c, err.Error())
	}
}

func QueryAuthorsHandler(c *gin.Context) {
	if err := NewAuthorsController(c).DoQueryAuthors(); err != nil {
		r.BadResponse(c, err.Error())
	}
}

func UpdateAuthorsHandler(c *gin.Context) {
	if err := NewAuthorsController(c).DoUpdateAuthors(); err != nil {
		r.BadResponse(c, err.Error())
	}
}

func DeleteAuthorsHandler(c *gin.Context) {
	if err := NewAuthorsController(c).DoDeleteAuthors(); err != nil {
		r.BadResponse(c, err.Error())
	}
}

func bindAuthorsVO(c *gin.Context) error {
	v := &vo.AuthorsVO{}
	if err := c.ShouldBind(v); err != nil {
		return err
	}
	return nil
}

type AuthorsController struct {
	*gin.Context
}

func NewAuthorsController(context *gin.Context) *AuthorsController {
	return &AuthorsController{Context: context}
}

func (s *AuthorsController) DoCreateAuthors() error {

	return nil
}

func (s *AuthorsController) DoQueryAuthors() error {

	return nil
}

func (s *AuthorsController) DoUpdateAuthors() error {

	return nil
}

func (s *AuthorsController) DoDeleteAuthors() error {

	return nil
}
