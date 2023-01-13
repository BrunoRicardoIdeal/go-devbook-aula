package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitemtpy"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	erro := usuario.formatar(etapa)
	if erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("Nome obrigatório.")
	}
	if usuario.Nick == "" {
		return errors.New("Nick obrigatório.")
	}
	if usuario.Email == "" {
		return errors.New("Email obrigatório.")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("Invalid e-mail format")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha obrigatória.")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil
}
