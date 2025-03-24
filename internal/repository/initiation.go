package repository

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-portal/internal/data"
)

type Initiation struct {
	persist *db.Persist
}

type IInitiation interface {
	FindAll() ([]*data.Initiation, error)
	FindByID(id string) (*data.Initiation, error)
	Persist(initn *data.Initiation) (*data.Initiation, error)
}

func NewInitiation(persist *db.Persist) IInitiation {
	initiation := &Initiation{
		persist: persist,
	}

	return initiation
}

func (s *Initiation) FindAll() ([]*data.Initiation, error) {
	var initiations []*data.Initiation
	if err := s.persist.DB.Find(&initiations).Error; err != nil {
		return nil, err
	}

	return initiations, nil
}

func (s *Initiation) FindByID(id string) (*data.Initiation, error) {
	initiation := &data.Initiation{}
	if err := s.persist.DB.First(initiation, id).Error; err != nil {
		return nil, err
	}

	return initiation, nil
}

func (s *Initiation) Persist(initn *data.Initiation) (*data.Initiation, error) {
	tx := s.persist.DB.Save(initn)
	err := tx.Error
	return initn, err
}
