package service

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/kardianos/service"
)

const (
	sysconfigDir = "/etc/sysconfig"
)

type Service struct {
	Name        string
	Binary      string
	Arguments   []string
	Environment map[string]string

	svc service.Service
}

type dummy struct{}

func (p *dummy) Start(_ service.Service) error {
	return nil
}

func (p *dummy) Stop(_ service.Service) error {
	return nil
}

func (s *Service) Install() error {
	if err := os.MkdirAll(sysconfigDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(sysconfigDir, s.Name), generateEnvFile(s.Environment), 0644); err != nil {
		return err
	}

	cfg := &service.Config{
		Name:        s.Name,
		DisplayName: s.Name,
		Option:      map[string]interface{}{},
		Executable:  s.Binary,
		Arguments:   s.Arguments,
	}

	if u, err := user.Current(); err == nil && u.Uid != "0" {
		cfg.UserName = u.Username
	}

	svc, err := service.New(&dummy{}, cfg)
	if err != nil {
		return err
	}

	s.svc = svc

	// TODO: rewrite and reload if changed!
	if err := svc.Install(); err != nil {
		log.Printf("error: %s", err.Error())
	}

	return nil
}

func generateEnvFile(env map[string]string) []byte {
	result := ""

	for k, v := range env {
		result += fmt.Sprintf("%s=%s\n", k, v)
	}

	return []byte(result)
}

func (s *Service) Start() error {
	return s.svc.Start()
}

func (s *Service) Stop() error {
	return s.svc.Stop()
}
