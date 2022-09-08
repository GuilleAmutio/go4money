package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	// Execute Go tests and then finish
	os.Exit(m.Run())
}
