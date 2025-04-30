package npcapi

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCSCoordinates(context *gin.Context) ([2]int, error) {
	var err error = nil

	// Validate Querystring parameters
	// TODO(wholesomeow): Implement success bool in GetQuery
	in_coords, success := context.GetQuery("cs_coords")
	if !success {
		status, response := Response500("request failed: failed to query querystring parameter")
		context.JSON(status, response)
	}

	query_coords := strings.Split(in_coords, ",")

	// Check length
	if len(query_coords) <= 0 || len(query_coords) == 1 {
		msg := fmt.Sprintf("passed querystring parameter bad length: %s", query_coords)
		err = errors.New(msg)
	}
	if err != nil {
		// Return error message if querystring parameters don't pass
		msg := fmt.Sprintf("request failed: %s", err)
		status, response := Response500(msg)
		context.JSON(status, response)
	}

	// Check type
	for _, i := range query_coords {
		if reflect.TypeOf(i).Kind() != reflect.String {
			err = errors.New("passed querystring does not contain string")
		}
	}
	if err != nil {
		// Return error message if querystring parameters don't pass
		msg := fmt.Sprintf("request failed: %s", err)
		status, response := Response500(msg)
		context.JSON(status, response)
	}

	// Convert passed strings to int
	cs_coords := []int{}
	for _, i := range query_coords {
		val, _ := strconv.Atoi(i)
		cs_coords = append(cs_coords, val)
	}

	return [2]int(cs_coords), nil
}
