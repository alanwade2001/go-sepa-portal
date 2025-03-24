package service

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log/slog"

	"net/http"

	"strings"
	"time"

	"github.com/alanwade2001/go-sepa-portal/data"
	utils "github.com/alanwade2001/go-sepa-utils"
)

type Store struct {
	Address string
	client  http.Client
}

type IStore interface {
	StoreDocument(content string) (doc *data.Document, err error)
}

func NewStore() IStore {

	address := utils.Getenv("DOCS_ADDRESS", "https://0.0.0.0:8443")
	slog.Info("docs address", "address", address)

	s := &Store{
		Address: address,
		client: http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}

	return s
}

func (s *Store) StoreDocument(content string) (doc *data.Document, err error) {
	doc = &data.Document{}

	reader := strings.NewReader(content)

	slog.Debug("storing document")
	response, err := s.client.Post(s.Address+"/documents", "text/plain", reader)

	if err != nil {
		slog.Error("failed to post", "error", err)
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		err = errors.New("failed to create document: " + response.Status)
		slog.Error("failed to store", "error", err)
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(doc)
	if err != nil {
		slog.Error("failed to decode", "error", err)
		return nil, err
	}

	slog.Info("Store Doc", "Doc ID", doc.ID)

	return doc, nil
}
