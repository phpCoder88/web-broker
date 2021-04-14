package service

import (
	"log"

	"github.com/phpCoder88/web-broker/internal/pkg/model"
	web_broker "github.com/phpCoder88/web-broker/pkg/web-broker"
)

func (s *Service) Put(req *web_broker.PutValueReq) error {
	if err := s.repo.Put(&model.PutValue{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		log.Printf("service/Put: put repo err: %v", err)
		return err
	}

	return nil
}
