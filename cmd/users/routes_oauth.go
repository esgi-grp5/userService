package main

import (
	"go-micro/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// OAuth godoc
// @Summary Show the status of server.
// @Description get access to User application.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Unauthorized 401 {object} map[string]interface{}
// @Router /oauth [post]
func (s *server) OAuth(c *gin.Context) {
	// Init var
	var (
		oauthReq  config.OAuthRequest
		oauthResp config.OAuthResponse
	)

	// Get JSON body
	if err := c.ShouldBindJSON(&oauthReq); err != nil {
		log.Err(err).Msg("Error in OAuth")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}

	// Verify information with server
	if s.oauth.OAuthRequest.ID != oauthReq.ID || s.oauth.OAuthRequest.Secret != oauthReq.Secret {
		log.Warn().Interface("OAuth request", oauthReq).Msg("Invalid ID or Secret")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID or Secret"})
		return
	}

	// Generate OAuthResponse
	oauthResp.AccessToken = RandomString(20)
	oauthResp.TokenType = "bearer_token"

	// Put on server
	s.oauth.OAuthResponse.AccessToken = oauthResp.AccessToken
	s.oauth.OAuthResponse.TokenType = oauthResp.TokenType

	// Return JSON reposne
	res := map[string]interface{}{
		"access_token": oauthResp.AccessToken,
		"token_type":   oauthResp.TokenType,
	}
	c.JSON(http.StatusOK, res)
}
