package postgres

import (
	"context"
	"order/entity"

	"github.com/google/uuid"
)

func (p *PostgresDB) CreateOrder(ctx context.Context, orderEntity entity.Order) (entity.Order, entity.OrderOutbox, error) {
	tx := p.db.WithContext(ctx).Begin()

	order := mapOrderEntityToOrder(orderEntity)

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return entity.Order{}, entity.OrderOutbox{}, err
	}

	orderOutbox := entity.OrderOutbox{
		OrderID:        order.ID,
		IdempotencyKey: uuid.NewString(),
	}

	if err := tx.Create(&orderOutbox).Error; err != nil {
		tx.Rollback()
		return entity.Order{}, entity.OrderOutbox{}, nil
	}

	tx.Commit()

	return mapOrderToOrderEntity(order), orderOutbox, nil
}

func (p *PostgresDB) FindOrderByID(ctx context.Context, id uint) (entity.Order, error) {
	var order entity.Order
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&order).Error; err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
