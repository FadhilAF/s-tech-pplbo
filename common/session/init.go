package session

import (
  "github.com/gin-contrib/sessions"
)
type Session interface {
  store sessions.Store
}
