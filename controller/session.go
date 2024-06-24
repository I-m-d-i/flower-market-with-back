package controller

import "flover-market/model"

// GetSession возвращает сессию по ее токену
func GetSession(token string) (model.Session, error) {
	return model.GetSession(token)
}
