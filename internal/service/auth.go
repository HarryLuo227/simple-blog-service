package service

import "errors"

type AuthRequest struct {
	User     string `header:"user" binding:"required"`
	Password string `header:"password" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.User, param.Password)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
