package usecase

import (
	"database/sql"
	"github.com/brunorfeliciano/gointensivo/internal/infra/database"
	"github.com/brunorfeliciano/gointensivo/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CalculatePriceUseCaseSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository //entity.OrderRepositoryInterface
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, primary key (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseSuite))
}

func (suite *CalculatePriceUseCaseSuite) TestGivenAnOrder_WhenCalculateFinalPrice_ThenShouldCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	order.CalculateFinalPrice()

	calculateFInalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFInalPriceInput)

	suite.NoError(err)
	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
