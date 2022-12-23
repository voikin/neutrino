package repository

type Repository struct {
	TgUserRepository
}

type TgUserRepository interface {
	SaveTgUser() error
}
