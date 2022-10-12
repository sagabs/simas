package zoomhistory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateZoomHistory(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := ZoomHistoryRequest{
		Nama:       "cayo",
		Email:      "cayo@gmail.com",
		Kategori:   "Kredit",
		Keterangan: "gatau",
	}
	res, _, err := service.CreateZoomHistory(data)
	fmt.Println("test", data.Nama, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// data = RegisterRequest{
	// 	Username: "remasertu",
	// 	Password: "123456",
	// 	Name:     "rema",
	// }
	// res, _, err = service.Login(data)
	// assert.NotNil(t, err)

}
