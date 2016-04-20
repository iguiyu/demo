package helpers

import (
	"github.com/kobeld/goutils"

	"github.com/kobeld/mgowrap"
)

// Get mongo server address from consul service
func SetupDBFromConsul(database string) {
	dbAddr, err := goutils.GetConsulServiceAddress("mongo")
	if err != nil {
		panic(err)
	}

	mgowrap.SetupDatbase(dbAddr, database)
}

func DropDB() {
	err := mgowrap.DropDatabase()
	if err != nil {
		panic(err)
	}
}

func DropCollection(name string) {
	err := mgowrap.DropCollection(name)
	if err != nil {
		panic(err)
	}
}
