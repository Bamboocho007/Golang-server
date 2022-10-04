package services

import (
	"errors"
	"vasya_project/db"
	"vasya_project/dtos"
	"vasya_project/helpers"
	"vasya_project/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) dtos.Response[dtos.TokenResponseDto] {
	env_config, exists := helpers.GetRouterData[utils.EnvConfig](c, "envConfig")

	if !exists {
		c.Abort()
		panic(errors.New("envConfig not exists"))
	}

	form := dtos.LoginFormDto{}
	bind_err := c.BindJSON(&form)

	if bind_err != nil {
		return dtos.Response[dtos.TokenResponseDto]{
			Meta_data: dtos.MetaData{Message: "invalid form"},
		}
	}

	user, error := db.GetUserByEmail(env_config.Db_connection, form.Email)

	if error != nil {
		return dtos.Response[dtos.TokenResponseDto]{
			Meta_data: dtos.MetaData{Message: "invalid email or password"},
		}
	}

	token := NewJwt(user.Id, 1, env_config.Jwt_secret)

	return dtos.Response[dtos.TokenResponseDto]{
		Data: &dtos.TokenResponseDto{Token: token},
	}
}
