package services

import (
	"golang_starter_kit_2025/app/models"
	"golang_starter_kit_2025/facades"
)

type FarmerService struct{}

func (*FarmerService) GetAllFarmers() ([]models.Farmer, error) {
	var farmers []models.Farmer
	if err := facades.DB.Find(&farmers).Error; err != nil {
		return nil, err
	}
	return farmers, nil
}

func (*FarmerService) Find(id uint64) (models.Farmer, error) {
	var farmer models.Farmer
	if err := facades.DB.First(&farmer, "id = ?", id).Error; err != nil {
		return farmer, err
	}
	return farmer, nil
}

func (*FarmerService) Put(farmer models.Farmer) (models.Farmer, error) {
	var existing models.Farmer
	err := facades.DB.First(&existing, "id = ?", farmer.ID).Error
	if err == nil {
		if err := facades.DB.Model(&existing).Updates(farmer).Error; err != nil {
			return farmer, err
		}
		return existing, nil
	}
	if err := facades.DB.Create(&farmer).Error; err != nil {
		return farmer, err
	}
	return farmer, nil
}

func (*FarmerService) Delete(id uint64) error {
	var farmer models.Farmer
	if err := facades.DB.First(&farmer, "id = ?", id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&farmer).Error
}
