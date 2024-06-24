package controller

import (
	"flover-market/model"
	"fmt"
	"log"
	"time"
)

type PurchaseAdmin struct {
	PurchaseId    int                  `json:"purchase_id"`
	PurchaseDate  time.Time            `json:"purchase_date"`
	CustomerName  string               `json:"customer_name"`
	CustomerPhone string               `json:"customer_phone"`
	PurchasePrice int                  `json:"purchase_price"`
	PurchaseItems []model.PurchaseItem `json:"purchase_items"`
}

func GetPurchasesAdmin() ([]PurchaseAdmin, error) {
	purchases, err := model.GetPurchases()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var purchasesAdmin []PurchaseAdmin
	for _, purchase := range purchases {
		purchaseAdmin := PurchaseAdmin{
			PurchaseId:   purchase.PurchaseId,
			PurchaseDate: time.Unix(int64(purchase.PurchaseDate), 0),
		}
		purchaseItems, err := model.GetItemsPurchase(purchase.PurchaseId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		purchaseAdmin.PurchaseItems = purchaseItems
		for _, item := range purchaseItems {
			purchaseAdmin.PurchasePrice += item.ProductCount * item.ProductPrice
		}
		customer, err := model.GetCustomer(int64(purchase.CustomerId))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		purchaseAdmin.CustomerName = customer.CustomerName
		purchaseAdmin.CustomerPhone = customer.CustomerPhone
		purchasesAdmin = append(purchasesAdmin, purchaseAdmin)
	}
	return purchasesAdmin, nil
}

func AddPurchase(token string, storeId int, productsId []int) (model.Purchase, error) {
	ses, err := model.GetSession(token)
	if err != nil {
		log.Println(err)
		return model.Purchase{}, err
	}
	purchaseItems := make(map[int]*model.PurchaseItem)
	for _, productId := range productsId {
		if _, ok := purchaseItems[productId]; ok {
			purchaseItems[productId].ProductCount++
			continue
		}
		purchaseItems[productId] = &model.PurchaseItem{
			ProductId:    productId,
			ProductCount: 1,
			ProductPrice: 0,
		}
		price, err := model.GetLastPriceChange(productId)
		if err != nil {
			return model.Purchase{}, err
		}
		purchaseItems[productId].ProductPrice = price.NewPrice
	}
	purchaseItemsList := make([]model.PurchaseItem, 0, len(purchaseItems))
	for _, item := range purchaseItems {
		purchaseItemsList = append(purchaseItemsList, *item)
	}
	return model.AddPurchase(model.Purchase{CustomerId: int(ses.CustomerId), StoreId: storeId, PurchaseDate: int(time.Now().Unix())}, purchaseItemsList)

}

func GetDetailsPurchase(token string, purchaseId int) (model.Purchase, []model.PurchaseItem, error) {
	ses, err := model.GetSession(token)
	if err != nil {
		log.Println(err)
		return model.Purchase{}, nil, err
	}
	purchase, err := model.GetPurchase(purchaseId)
	if err != nil {
		log.Println(err)
		return model.Purchase{}, nil, err
	}
	if purchase.CustomerId != int(ses.CustomerId) {
		log.Println("Not your purchase")
		return model.Purchase{}, nil, fmt.Errorf("not your purchase")
	}
	itemsPurchase, err := model.GetItemsPurchase(purchaseId)
	if err != nil {
		log.Println(err)
		return model.Purchase{}, nil, err
	}
	return purchase, itemsPurchase, nil
}
