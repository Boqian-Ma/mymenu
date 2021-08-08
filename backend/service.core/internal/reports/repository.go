package reports

import (
	"time"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	// Get the most popular item between two times
	GetMostPopularItem(c *gin.Context, end_time time.Time, start_time time.Time, restaurantID string) (*PopularItem, error)
	// Get total revenue between two times
	GetTotalRevenue(c *gin.Context, restaurantID string, end_time time.Time, start_time time.Time) (float64, error)
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access report data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) GetMostPopularItem(c *gin.Context, end_time time.Time, start_time time.Time, restaurantID string) (*PopularItem, error) {
	//var menu_item *entity.MenuItem
	var item PopularItem
	// Get item id

	if end_time.IsZero() {
		err := r.db.Select(c,
			&item,
			"select item_id , sum(quantity) as sum from orders_items join orders on (orders_items.order_id = orders.id)  where orders.restaurant_id=$1 group by item_id order by sum(quantity) DESC limit 1",
			restaurantID,
		)

		if err != nil {
			if err.Error() == "Not Found" {
				item.ItemID = "N/A"
				item.Quantity = 0
			} else {
				return nil, err
			}
		}

	} else {

		err := r.db.Select(c,
			&item,
			"select orders_items.item_id , sum(quantity) as sum from orders_items join orders on (orders_items.order_id = orders.id) where orders.restaurant_id=$1 and DATE(orders.updated_at) between Date($2) and Date($3) group by orders_items.item_id order by sum(quantity) DESC limit 1",
			restaurantID, start_time, end_time,
		)

		if err != nil {

			if err.Error() == "Not Found" {
				item.ItemID = "N/A"
				item.Quantity = 0
			} else {
				return nil, err
			}
		}

	}

	return &item, nil
}

func (r repository) GetTotalRevenue(c *gin.Context, restaurantID string, end_time time.Time, start_time time.Time) (float64, error) {

	var re entity.TotalRevenue
	re.Revenue = 0.0

	if start_time.IsZero() {
		err := r.db.Select(c,
			&re,
			"select sum(total_cost) as total_cost from orders where orders.restaurant_id=$1",
			restaurantID,
		)

		if err != nil {
			return 0, err
		}

	} else {

		err := r.db.Select(c,
			&re,
			"select sum(total_cost) as total_cost from orders where orders.restaurant_id=$1 and DATE(orders.updated_at) between DATE($2) and DATE($3)",
			restaurantID, start_time, end_time,
		)
		if err != nil {
			return 0, err
		}
	}
	return re.Revenue, nil
}
