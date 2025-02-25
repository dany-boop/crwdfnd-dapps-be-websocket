package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"

)

type DonationRepository interface {
	CreateDonation(donation *models.Donation) error
	GetDonationsByUserID(userID uint) ([]models.Donation, error)
	GetAllDonations() ([]models.Donation, error)
}

type donationRepository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) DonationRepository {
	return &donationRepository{db}
}

func (r *donationRepository) CreateDonation(donation *models.Donation) error {
	return r.db.Create(donation).Error
}

func (r *donationRepository) GetDonationsByUserID(userID uint) ([]models.Donation, error) {
	var donations []models.Donation
	err := r.db.Where("user_id = ?", userID).Find(&donations).Error
	return donations, err
}

func (r *donationRepository) GetAllDonations() ([]models.Donation, error) {
	var donations []models.Donation
	err := r.db.Find(&donations).Error
	return donations, err
}
