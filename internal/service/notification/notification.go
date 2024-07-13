package notification

import pa "github.com/DanilMankiev/SofiaApplication/internal/service/pap"

type NotificationService struct {
}

func NewNoti() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) Get() error {
	pa.NewNoti()
	return nil
}