package controllers

import "github.com/gin-gonic/gin"

// Defines the common interface for all REST controllers
type RestController interface {
    // Registers the REST endpoints of the controller
    RegisterHandlers(router *gin.Engine)
}