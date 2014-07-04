package app

import (
  //"log"
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/owais/gamerhub/globals"
  "github.com/owais/gamerhub/routes"
)

func init() {
  e := gin.New()
  globals.Application = e
  routes.Include(e)
  http.Handle("/", e)
}
