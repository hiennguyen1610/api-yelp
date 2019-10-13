package service

import (
	"github.com/hiennguyen1610/api-yelp/model"
	"github.com/hiennguyen1610/api-yelp/repo"
)

type AddressService struct {
	addressRepo repo.AddressRepo
}

func NewAddressService() AddressService {
	addressService := AddressService{}
	addressService.addressRepo = repo.NewAddressRepo()
	return addressService
}

func (srvc AddressService) Find(name string, location string) []model.Address {
	return srvc.addressRepo.FindAll()
}