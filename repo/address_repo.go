package repo

import (
	"github.com/hiennguyen1610/api-yelp/model"
)

type AddressRepo struct {}

func NewAddressRepo() AddressRepo {
	return AddressRepo{}
}

func (repo AddressRepo) FindAll() []model.Address {
	return []model.Address{{"a1", "a2", "a3", "a4"}, 
	{"b1", "b2", "b3", "b4"}}
}