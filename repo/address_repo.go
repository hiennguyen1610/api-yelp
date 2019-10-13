package repo

import (
	"github.com/hiennguyen1610/api-yelp/model"
)

type AddressRepo struct {}

func NewAddressRepo() AddressRepo {
	return AddressRepo{}
}

func (repo AddressRepo) FindAll() []model.Address {
	return []model.Address{
	{"McDonald's", "Restaurant", "Paya Lebar", "Fast food"}, 
	{"Burger King", "Restaurant", "Paya Lebar", "Fast food"}}
}