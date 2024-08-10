package main
import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"

	"github.com/Jarntae/Sa_StockOrder/entity"
  )


func main() {
	db, err := gorm.Open(sqlite.Open("Stock.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(

	&entity.Category{},

	&entity.Employee{},

	&entity.Product{},

	&entity.Stock{},

	&entity.Supplier{},
)
	
}