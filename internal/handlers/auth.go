package handlers

import (
    "blog-api/internal/models"
    "blog-api/internal/storage"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AuthHandler struct {
    store *storage.MemoryStorage
}

func NewAuthHandler(s *storage.MemoryStorage) *AuthHandler {
    return &AuthHandler{store: s}
}

func (h *AuthHandler) Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Используем новый метод ValidateUser вместо прямого доступа к users
    if !h.store.ValidateUser(user.Username, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := "sample-jwt-token"
    c.JSON(http.StatusOK, gin.H{"token": token})
}