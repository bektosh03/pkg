package auth

import (
	"github.com/Blank-Xu/sql-adapter"
	"github.com/bektosh03/pkg/models"
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type JwtAuth struct {
	signingKey []byte
	enforcer   *casbin.Enforcer
}

type JwtAuthArgs struct {
	SigningKey   []byte
	UseFile      bool
	FileAdapter  fileadapter.Adapter
	SqlAdapter   sqladapter.Adapter
	PathToPolicy string
}

func New(args JwtAuthArgs) JwtAuth {
	if len(args.SigningKey) < 1 {
		panic("signing key cannot have length of zero")
	}
	var (
		enforcer *casbin.Enforcer
		err      error
	)
	if args.UseFile {
		enforcer, err = casbin.NewEnforcer(args.PathToPolicy, args.FileAdapter)
	} else {
		enforcer, err = casbin.NewEnforcer(args.PathToPolicy, args.SqlAdapter)
	}
	if err != nil {
		log.Fatalln("cannot create new enforcer", err)
	}
	initEnforcer(enforcer)

	return JwtAuth{
		signingKey: args.SigningKey,
		enforcer:   enforcer,
	}
}

func (a JwtAuth) Authentication(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	claims, err := a.ExtractClaims(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
			Error:   err.Error(),
			Message: "Invalid access token",
		})
	}
	role := claims["role"]
	path := c.Request.URL.Path
	method := c.Request.Method
	allowed, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
			Error: err.Error(),
			Message: "Could not enforce policy",
		})
	}
	if !allowed {
		c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
			Error:   "not allowed",
			Message: "Not enough rights to perform this action",
		})
	}
}

func initEnforcer(enforcer *casbin.Enforcer) {
	err := enforcer.LoadPolicy()
	if err != nil {
		log.Fatalln("error while loading policy", err)
	}
	err = enforcer.LoadModel()
	if err != nil {
		log.Fatalln("error while loading model", err)
	}
	err = enforcer.SavePolicy()
	if err != nil {
		log.Fatalln("error while saving policy", err)
	}
}
