package loader

import "github.com/datactf/dataloader/pkg/data"

func LoadMongoDBAccounts(n int, dbName, collName string) error {
	accounts := make([]interface{}, n)
	for i := 0; i < n; i++ {
		account, err := data.GenerateAccount()
		if err != nil {
			return err
		}
		accounts[i] = account
	}

	return LoadMongoDB(dbName, collName, accounts)
}

func init() {

}
