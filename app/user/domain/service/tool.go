package service

import "judgeMore_server/pkg/crypt"

func (uc *UserService) PasswordHash(password string) (string, error) {
	return crypt.PasswordHash(password)
}
func (uc *UserService) PasswordVerify(password, hash string) bool {
	return crypt.VerifyPassword(password, hash)
}
