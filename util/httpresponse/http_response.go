/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    httpresponse
	@Date    2022/4/19 14:25
	@Desc
*/

package httpresponse

import (
	"errors"
	"net/http"

	"github.com/NextSmartShip/common/util/log"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Bad request", http.StatusBadRequest)
}

func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	log.GetLogEntry(r).WithError(err).WithField("error-slug", slug).Warn(logMSg)
	resp := Error{slug, status}

	if err := render.Render(w, r, resp); err != nil {
		logrus.Errorf("httpRespondWithError err %v", err)

		return
	}
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e Error) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.Code)

	return nil
}

func ErrRsp(w http.ResponseWriter, r *http.Request, status int, message string) {
	errRsp := Error{
		Code:    status,
		Message: message,
	}

	w.WriteHeader(http.StatusBadRequest)
	render.Respond(w, r, errRsp)

	return
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	var slugError CommonError

	if ok := errors.Is(err, slugError); !ok {
		InternalError("internal-server-error", err, w, r)

		return
	}

	switch slugError.ErrorType() {
	case ErrorTypeAuthorization:
		Unauthorised(slugError.Slug(), slugError, w, r)
	case ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}
