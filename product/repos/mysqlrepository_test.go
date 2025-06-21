package repos_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/Phuong-Hoang-Dai/DStore/product/repos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, mock, gormDB, err := SetupMockDB()
	if err != nil {
		t.Fatalf("There is something wrong. Erorr: %v", err)
	}
	defer db.Close()

	data := product.Product{
		Name:  "Haha",
		Price: 99,
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `Product`").
		WithArgs(data.Name, "", data.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	t.Run("Test Create", func(t *testing.T) {
		product_Repos := repos.NewMysqlProductRepo(gormDB)
		id, err := product_Repos.CreateProduct(&data)
		if err != nil {
			t.Errorf("Create is incorrect. id = %v \n err: %v", id, err)
		}
	})

}

func TestUpdateProduct(t *testing.T) {
	db, mock, gormDB, err := SetupMockDB()
	if err != nil {
		t.Fatalf("There is something wrong. Erorr: %v", err)
	}
	defer db.Close()

	testData := []struct {
		product.Product
		rowsaffected int64
		wantErr      bool
	}{
		{Product: product.Product{Id: 3, Name: "MuaHahahoo", Price: 69}, rowsaffected: 1, wantErr: false},
		{Product: product.Product{Id: 4, Name: "MuaHahahoo", Price: 659}, rowsaffected: 5, wantErr: false},
		{Product: product.Product{Id: 2, Name: "MuaHaha", Price: 69}, rowsaffected: 0, wantErr: true},
	}
	for _, data := range testData {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `Product` SET `name`=?,`price`=? WHERE `id` = ?")).
			WithArgs(data.Name, data.Price, data.Id).
			WillReturnResult(sqlmock.NewResult(1, data.rowsaffected))
		mock.ExpectCommit()

		t.Run("Test Update", func(t *testing.T) {
			product_Repos := repos.NewMysqlProductRepo(gormDB)
			err := product_Repos.UpdateProduct(&data.Product)
			if err != nil && !data.wantErr {
				t.Errorf("Update is incorrect. \n err: %v", err)
			}
		})
	}
}

func SetupMockDB() (*sql.DB, sqlmock.Sqlmock, *gorm.DB, error) {
	db, mock, err := sqlmock.New()
	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err2 != nil && err == nil {
		err = err2
	}
	return db, mock, gormDB, err
}
