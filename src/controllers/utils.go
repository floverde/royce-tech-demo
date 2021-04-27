package controllers

import (
	"github.com/gin-gonic/gin"
	"math/bits"
	"net/http"
	"strconv"
)

// Fetchs an unsigned integer (uint) parameter from the URL of
// an HTTP request Returns true if successful, false otherwise.
func fetchUintURLParam(cxt *gin.Context, name string, value *uint) bool {
    // Retrieves the requested user ID from the request URL
    retVal, err := strconv.ParseUint(cxt.Param(name), 10, bits.UintSize)
	// Check that the previous operation was successful
	if err != nil {
		// Returns an error 400 BAD REQUEST attaching the error message
        cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
		// Indicates that the operation was unsuccessful
        return false
	} else {
		// Sets the value of the output parameter
		*value = uint(retVal)
		// Indicates that the operation was successful
		return true
	}
}