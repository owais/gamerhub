package routes

import (
  "github.com/go-martini/martini"

  "github.com/owais/gamerhub/apps/games/views"
)


func Patterns(r martini.Router) {
  r.Get("/", views.Home)
}
