package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pagination(data, total, limit, offset interface{}) (int, interface{}) { //newbaseresponsepagination
	return http.StatusOK, map[string]interface{}{
		"data":   data,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	}
}

func Scroll(data interface{}, scrollId string) (int, interface{}) { //newbaseresponsescroll
	return http.StatusOK, map[string]interface{}{
		"items":     data,
		"scroll_id": scrollId,
	}
}

func Data(code int, data interface{}) (int, interface{}) { //newresponse
	return code, gin.H{
		"data": data,
	}
}

func OK(data interface{}) (int, interface{}) { //newOKresponse
	return http.StatusOK, data
}

func Created(data map[string]interface{}) (int, interface{}) {
	result := map[string]interface{}{
		"created": true,
	}
	for key, value := range data {
		result[key] = value
	}
	return http.StatusCreated, result
}
func Error(code int, msg interface{}) (int, interface{}) {
	return code, gin.H{
		"error": msg,
	}
}

func ServiceUnavailable() (int, interface{}) {
	return http.StatusServiceUnavailable, gin.H{
		"error": http.StatusText(http.StatusServiceUnavailable),
	}
}

func ServiceUnavailableMsg(msg interface{}) (int, interface{}) {
	return http.StatusServiceUnavailable, gin.H{
		"error": msg,
	}
}

func BadRequest() (int, interface{}) {
	return http.StatusBadRequest, gin.H{
		"error": http.StatusText(http.StatusBadRequest),
	}
}

func BadRequestMsg(msg interface{}) (int, interface{}) {
	return http.StatusBadRequest, gin.H{
		"error": msg,
	}
}

func NotFound() (int, interface{}) {
	return http.StatusNotFound, gin.H{
		"error": http.StatusText(http.StatusNotFound),
	}
}

func NotFoundMsg(msg interface{}) (int, interface{}) {
	return http.StatusNotFound, gin.H{
		"error": msg,
	}
}

func Forbidden() (int, interface{}) {
	return http.StatusForbidden, gin.H{
		"error": "Do not have permission for the request.",
	}
}

func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	}
}
