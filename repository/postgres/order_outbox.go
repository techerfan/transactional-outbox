package postgres

import (
	"context"
	"order/entity"
)

func (p *PostgresDB) UpdateOrderOutbox(ctx context.Context, orderOutboxEntity entity.OrderOutbox) error {
	orderOutbox := mapOrderOutboxEntityToOrderOutbox(orderOutboxEntity)

	return p.db.WithContext(ctx).Save(&orderOutbox).Error
}

func (p *PostgresDB) DeleteOrderOutbox(ctx context.Context, id uint) error {
	return p.db.WithContext(ctx).Unscoped().Delete(&OrderOutbox{}, id).Error
}

func (p *PostgresDB) GetFirstOrderFromOutbox(ctx context.Context) (entity.OrderOutbox, error) {
	var orderOutbox entity.OrderOutbox

	if err := p.db.WithContext(ctx).Where("pushed = ?", false).Order("id ASC").First(&orderOutbox).Error; err != nil {
		return entity.OrderOutbox{}, err
	}

	return orderOutbox, nil
}

func (p *PostgresDB) FindOrderOutboxByID(ctx context.Context, id uint) (entity.OrderOutbox, error) {
	var orderOutbox entity.OrderOutbox
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&orderOutbox).Error; err != nil {
		return entity.OrderOutbox{}, err
	}
	return orderOutbox, nil
}
