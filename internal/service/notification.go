package service

type NotificationService struct {
}

func NewNoti() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) Get() error {
	return nil
}