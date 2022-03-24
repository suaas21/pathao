package api

import (
	"encoding/json"
	"github.com/suaas21/pathao/api/response"
	"github.com/suaas21/pathao/logger"
	"github.com/suaas21/pathao/model"
	"github.com/suaas21/pathao/service"
	"io/ioutil"
	"net/http"
)

// UserController ...
type UserController struct {
	schema service.User
	lgr    logger.StructLogger
}

// NewUserController ...
func NewUserController(schema service.User, lgr logger.StructLogger) *UserController {
	return &UserController{
		schema: schema,
		lgr:    lgr,
	}
}

func (u *UserController) UserController(res http.ResponseWriter, req *http.Request) {
	incomingReq := model.GraphQLIncomingRequest{}
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyBytes, &incomingReq)
	if err != nil {
		response.ServeJSON(res, err, http.StatusInternalServerError)
		return
	}
	result, err := u.schema.UserSchema(incomingReq)
	if err != nil {
		response.ServeJSON(res, result, http.StatusInternalServerError)
		return
	}
	if result.HasErrors() {
		var errors []string
		for _, e := range result.Errors {
			errors = append(errors, e.Error())
		}
		response.ServeJSON(res, result, http.StatusBadRequest)
		return
	}

	response.ServeJSON(res, result.Data, http.StatusOK)
	return
}
