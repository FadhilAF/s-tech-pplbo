package session

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func main(db *sql.DB, r *gin.Engine) {

	store, err := postgres.NewStore(db, []byte("secret"))
	if err != nil {
		// handle err
	}

	//tinggal pake line di bawah ini bae
	r.Use(sessions.Sessions("mysession", store))
	//tinggal pake line di atas ini bae

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
