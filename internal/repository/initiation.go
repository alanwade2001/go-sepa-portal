package repository

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-portal/internal/repository/entity"
)

type Initiation struct {
	persist *db.Persist
}

type IInitiation interface {
	FindAll() ([]*entity.Initiation, error)
	FindByID(id string) (*entity.Initiation, error)
	Perist(initn *entity.Initiation) (*entity.Initiation, error)
}

func NewInitiation(persist *db.Persist) IInitiation {
	initiation := &Initiation{
		persist: persist,
	}

	return initiation
}

func (s *Initiation) FindAll() ([]*entity.Initiation, error) {
	var initiations []*entity.Initiation
	if err := s.persist.DB.Find(&initiations).Error; err != nil {
		return nil, err
	}

	return initiations, nil
}

func (s *Initiation) FindByID(id string) (*entity.Initiation, error) {
	initiation := &entity.Initiation{}
	if err := s.persist.DB.First(initiation, id).Error; err != nil {
		return nil, err
	}

	return initiation, nil
}

func (s *Initiation) Perist(initn *entity.Initiation) (*entity.Initiation, error) {
	tx := s.persist.DB.Save(initn)
	err := tx.Error
	return initn, err
}
