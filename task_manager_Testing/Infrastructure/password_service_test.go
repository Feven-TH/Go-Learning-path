package infrastructure_test

import (
    "testing"
    "task_manager_Testing/Infrastructure"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"
)

type BcryptPasswordServiceSuite struct {
    suite.Suite
    service *infrastructure.BcryptPasswordService
}

func (s *BcryptPasswordServiceSuite) SetupTest() {
    s.service = infrastructure.NewBcryptPasswordService()
}

func (s *BcryptPasswordServiceSuite) TestHashPassword_Success() {
    hashed, err := s.service.IHashPassword("secret123")
    assert.NoError(s.T(), err)
    assert.NotEmpty(s.T(), hashed)
}

func (s *BcryptPasswordServiceSuite) TestComparePassword_Success() {
    plain := "secret123"
    hashed, _ := s.service.IHashPassword(plain)
    err := s.service.IComparePassword(hashed, plain)
    assert.NoError(s.T(), err)
}

func (s *BcryptPasswordServiceSuite) TestComparePassword_Failure() {
    hashed, _ := s.service.IHashPassword("original")
    err := s.service.IComparePassword(hashed, "wrong-password")
    assert.Error(s.T(), err)
}

func TestBcryptPasswordServiceSuite(t *testing.T) {
    suite.Run(t, new(BcryptPasswordServiceSuite))
}
