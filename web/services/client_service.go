package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/henriquepw/imperium-tattoo/web/db"
	"github.com/henriquepw/imperium-tattoo/pkg/customid"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/hash"
	"github.com/henriquepw/imperium-tattoo/pkg/validate"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientService interface {
	CreateClient(ctx context.Context, dto types.ClientCreateDTO) (*types.Client, error)
	ListClients(ctx context.Context) ([]types.Client, error)
	GetClientById(ctx context.Context, id string) (*types.Client, error)
	UpdateClinetById(ctx context.Context, id string, dto types.ClientUpdateDTO) (*types.Client, error)
}

type clientService struct {
	store db.ClientStore
}

func NewClientService(store db.ClientStore) *clientService {
	return &clientService{store}
}

func (s *clientService) CreateClient(ctx context.Context, payload types.ClientCreateDTO) (*types.Client, error) {
	if err := validate.CheckPayload(payload); err != nil {
		return nil, err
	}

	id, err := customid.New()
	if err != nil {
		return nil, errors.Internal("Não foi possível inserir o cliente")
	}

	if s.store.ExistCPF(ctx, payload.CPF) {
		return nil, errors.InvalidRequestData(map[string]string{"cpf": "CPF já cadastrado"})
	}

	brithday, err := time.Parse(time.DateOnly, payload.Brithday)
	if err != nil {
		return nil, errors.InvalidRequestData(map[string]string{"brithday": "Data inválida"})
	}

	now := time.Now()
	client := types.Client{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Brithday:  brithday,
		Name:      payload.Name,
		CPF:       payload.CPF,
		Instagram: payload.Instagram,
		Phone:     payload.Phone,
		Email:     payload.Email,
		Address:   payload.Address,
	}

	cpf := strings.Split(payload.CPF, ".")
	passwordHash, err := hash.HashPassword(cpf[0] + cpf[1])
	if err != nil {
		return nil, err
	}

	err = s.store.Insert(ctx, client, passwordHash)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Internal("Não foi possível inserir o cliente")
	}

	return &client, nil
}

func (s *clientService) GetClientById(ctx context.Context, id string) (*types.Client, error) {
	c, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Cliente não encontrado")
	}

	return c, nil
}

func (s *clientService) ListClients(ctx context.Context) ([]types.Client, error) {
	items, err := s.store.List(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Internal("Não foi possível listar os clientes")
	}

	return items, nil
}

func (s *clientService) UpdateClinetById(ctx context.Context, id string, dto types.ClientUpdateDTO) (*types.Client, error) {
	if err := validate.CheckPayload(dto); err != nil {
		return nil, err
	}

	if err := s.store.Update(ctx, id, dto); err != nil {
		log.Println(err)
		return nil, errors.Internal("Não foi possível atualizar os dados do cliente")
	}

	c, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Cliente não encontrado")
	}

	return c, err
}
