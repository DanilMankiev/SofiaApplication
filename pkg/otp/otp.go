package otp

import (
	"github.com/xlzd/gotp"
)

type CodeGenerator interface {
	GenerateCode(length int) string 
}

type OtpGenerator struct{}

func NewOTPGenerator() *OtpGenerator{
	return &OtpGenerator{}
}

func (otp *OtpGenerator) GenerateCode(length int) string{
	return gotp.RandomSecret(length)
}