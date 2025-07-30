package infrastructure_test

import (
    "testing"
    "task_manager_Testing/Infrastructure"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"
)

type JwtTokenServiceSuite struct {
    suite.Suite
    service *infrastructure.JwtTokenService
}

func (s *JwtTokenServiceSuite) SetupTest() {
    s.service = infrastructure.NewJwtTokenService("mySecretKey")
}

func (s *JwtTokenServiceSuite) TestIGenerateAccessToken_Success() {
    tokenResp, err := s.service.IGenerateAccessToken("user123", "admin")
    assert.NoError(s.T(), err)
    assert.NotEmpty(s.T(), tokenResp.AccessToken)
}

func (s *JwtTokenServiceSuite) TestIVerifyToken_Success() {
    tokenResp, _ := s.service.IGenerateAccessToken("user123", "admin")
    claims, err := s.service.IVerifyToken(tokenResp.AccessToken)

    assert.NoError(s.T(), err)
    assert.Equal(s.T(), "user123", claims["sub"])
    assert.Equal(s.T(), "admin", claims["role"])
}

func (s *JwtTokenServiceSuite) TestIVerifyToken_Failure_InvalidToken() {
    tamperedToken := "invalid.token.string"
    claims, err := s.service.IVerifyToken(tamperedToken)

    assert.Error(s.T(), err)
    assert.Nil(s.T(), claims)
}

func TestJwtTokenServiceSuite(t *testing.T) {
    suite.Run(t, new(JwtTokenServiceSuite))
}
