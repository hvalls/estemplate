package main

import (
	"io/ioutil"

	"estemplate/command"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/templates/render", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		cmd, err := command.NewRenderTemplateCommand(jsonData)
		if err != nil {
			panic(err)
		}
		filename, cleanUp, err := command.Execute(cmd)
		if err != nil {
			panic(err)
		}
		defer cleanUp()

		c.File(filename)
	})

	r.Run()
}
