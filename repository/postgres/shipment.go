package postgres

import (
	"context"
	"errors"
	"order/entity"

	"gorm.io/gorm"
)

func (p *PostgresDB) CreateShipment(ctx context.Context, shipmentEntity entity.Shipment) (entity.Shipment, error) {
	shipment := mapShipmentEntityToShipment(shipmentEntity)
	err := p.db.WithContext(ctx).Create(&shipment).Error
	if err != nil {
		return entity.Shipment{}, err
	}
	return mapShipmentToShipmentEntity(shipment), nil
}

func (p *PostgresDB) UpdateShipment(ctx context.Context, shipmentEntity entity.Shipment) error {
	shipment := mapShipmentEntityToShipment(shipmentEntity)
	return p.db.WithContext(ctx).Save(&shipment).Error
}

func (p *PostgresDB) FindShipmentByOrderID(ctx context.Context, orderID uint) (entity.Shipment, bool, error) {
	var shipment Shipment
	err := p.db.WithContext(ctx).Where("order_id = ?", orderID).First(&shipment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Shipment{}, false, nil
		}
		return entity.Shipment{}, false, err
	}
	return mapShipmentToShipmentEntity(shipment), true, nil
}
