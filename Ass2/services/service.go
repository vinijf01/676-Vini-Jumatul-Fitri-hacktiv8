package services

import (
	"ass2/databases"
	"ass2/models"
	"time"
)

func SaveOrder(req models.Request) (models.Response, error) {
	var res models.Response
	db := databases.GetDb()
	var items []models.Items
	for _, vitem := range req.Items {
		var item models.Items
		item.ItemCode = vitem.ItemCode
		item.Description = vitem.Description
		item.Quantity = vitem.Quantity
		items = append(items, item)
	}

	order := models.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    time.Now(),
		Items:        items,
	}

	errdb := db.Create(&order).Error
	if errdb != nil {
		return res, errdb
	}

	return models.Response{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}, nil
}

func GetOrder() ([]models.Order, error) {
	var order []models.Order
	db := databases.GetDb()
	err := db.Preload("Items").Order("id_order").Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetOrderByID(IdOrder uint) (models.Order, error) {
	var order models.Order
	db := databases.GetDb()
	err := db.Where("id_order = ?", IdOrder).Preload("Items").Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func UpdateOrder(req models.Request, IdOrder uint) (models.Order, error) {
	var order models.Order
	db := databases.GetDb()
	err := db.Where("id_order = ?", IdOrder).Preload("Items").Find(&order).Error
	if err != nil {
		return order, err
	}

	order.CustomerName = req.CustomerName
	order.OrderedAt = time.Now()
	// order.Items = req.Items

	err = db.Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil

}

func DeleteOrder(IdOrder uint) (string, error) {
	var order models.Order
	var item models.Items
	db := databases.GetDb()
	err := db.Where("id_order = ?", IdOrder).Preload("Items").Find(&order).Error
	if err != nil {
		return "Data tidak ditemukan", err
	}
	err = db.Where("order_id", IdOrder).Delete(item).Error
	if err != nil {
		return "Data Item Gagal dihapus", err
	}
	err = db.Where("id_order = ?", IdOrder).Delete(order).Error
	if err != nil {
		return "Data Order Gagal dihapus", err
	}
	return "message : Data Order berhasil di hapus", nil

}
