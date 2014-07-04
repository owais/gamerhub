package routes

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/owais/gamerhub/globals"
  "github.com/owais/gamerhub/views"
  //"github.com/owais/gamerhub/apps/games/routes"
)


func Include(e *gin.Engine) {
  globals.Application.LoadHTMLTemplates("templates/*")
  e.GET("/",  views.Home);
  e.ServeFiles("/static/*filepath", http.Dir("static"))
  //e.Group("/games", routes.Patterns);
}
