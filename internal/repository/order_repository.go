package repository

import (
	"context"
	"database/sql"
	"order-service/internal/entity"
	"order-service/internal/repository/source"

	"github.com/sirupsen/logrus"
)

type IOrderRepository interface {
	Add(order *entity.Order) error
	UpdateStatus(orderId string, status string) error
}
type orderRepository struct {
	db  *sql.DB
	log *logrus.Logger
	ctx context.Context
}

func NewOrderRepository(db *sql.DB, log *logrus.Logger, ctx context.Context) IOrderRepository {
	return &orderRepository{
		db:  db,
		log: log,
		ctx: ctx,
	}
}
func (r *orderRepository) Add(order *entity.Order) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.ORDER_ADD)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &order.Id, &order.UserId, &order.ProductId, &order.Price, &order.Quantity, &order.Status)
	if err != nil {
		r.log.WithError(err).Error("failed exec context")
	}
	return nil
}
func (r *orderRepository) UpdateStatus(orderId string, status string) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.ORDER_ADD)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &orderId, &status)
	if err != nil {
		r.log.WithError(err).Error("failed exec context")
	}
	return nil
}
