package services

import (
	"context"
	"database/sql"
	"errors"
	"gofiber-restapi/domain"
	"gofiber-restapi/dto"
	"time"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customerData, nil
}

func (c customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.customerRepository.Save(ctx, &customer)
}

func (c customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	if _, err := uuid.Parse(req.ID); err != nil {
		return errors.New("invalid UUID format")
	}

	persisted, err := c.customerRepository.FindById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data customer not found")
		}
		return err
	}

	if persisted.ID == "" {
		return errors.New("data customer not found")
	}

	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return c.customerRepository.Update(ctx, &persisted)
}

func (c customerService) Delete(ctx context.Context, id string) error {
	exist, err := c.customerRepository.FindById(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data customer not found")
		}
		return err
	}
	if exist.ID == "" {
		return errors.New("data customer not found")
	}
	return c.customerRepository.Delete(ctx, id)
}

func (c customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	if _, err := uuid.Parse(id); err != nil {
		return dto.CustomerData{}, errors.New("invalid UUID format")
	}
	persisted, err := c.customerRepository.FindById(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.CustomerData{}, errors.New("data customer not found")
		}
		return dto.CustomerData{}, err
	}

	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("data customer not found")
	}

	return dto.CustomerData{
		ID:   persisted.ID,
		Name: persisted.Name,
		Code: persisted.Code,
	}, nil
}
