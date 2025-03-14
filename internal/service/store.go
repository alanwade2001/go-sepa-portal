package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"strings"
	"time"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	utils "github.com/alanwade2001/go-sepa-utils"
)

type Store struct {
	Address string
	client  http.Client
}

func NewStore() *Store {

	address := utils.Getenv("DOCUMENTS_ADDRESS", "http://0.0.0.0:8081")

	s := &Store{
		Address: address,
		client:  http.Client{Timeout: 10 * time.Second},
	}

	return s
}

func (s *Store) StoreDocument(content string) (doc *model.Document, err error) {
	doc = &model.Document{}

	reader := strings.NewReader(content)

	log.Println("storing document")
	response, err := s.client.Post(s.Address+"/documents", "text/plain", reader)

	if err != nil {
		log.Printf("failed to post, %v", err)
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		err = errors.New("failed to create document: " + response.Status)
		log.Printf("failed to store, %v", err)
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(doc)
	if err != nil {
		log.Printf("failed to decode, %v", err)
		return nil, err
	}

	log.Printf("Document ID: %d", doc.ID)

	return doc, nil
}
