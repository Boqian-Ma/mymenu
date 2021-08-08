package reports

import (
	"time"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/menu_items"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/orders"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/categories"
	"github.com/gin-gonic/gin"
)

type Service interface {
	// Get a type of report
	Get(c *gin.Context, restaurantID string, reportType string) (*entity.HomePageReport, error)
}

type service struct {
	repo         Repository
	userRepo     users.Repository
	menuItemRepo menu_items.Repository
	orderRepo    orders.Repository
	categoriesRepo categories.Repository
}

type PopularItem struct {
	ItemID   string `db:"item_id"`
	Quantity int    `db:"sum"`
}

func NewService(repo Repository, userRepo users.Repository, menuItemRepo menu_items.Repository, orderRepo orders.Repository, categoriesRepo categories.Repository) service {
	return service{repo, userRepo, menuItemRepo, orderRepo, categoriesRepo}
}
func (s service) Get(c *gin.Context, restaurantID string, reportType string) (*entity.HomePageReport, error) {
	if err := auth.IsManagerOf(c, restaurantID); err != nil {
		return nil, err
	}

	if reportType != "home" {
		return nil, errors.BadRequest("Query parameter not found")
	}

	var daily *entity.HomePageReportItem
	var weekly *entity.HomePageReportItem
	var monthly *entity.HomePageReportItem
	var quarterly *entity.HomePageReportItem
	var yearly *entity.HomePageReportItem
	var all_time *entity.HomePageReportItem

	now := time.Now()

	daily, err := popularItemHelper(c, s, restaurantID, now, 0, 0, -1)
	if err != nil {
		return nil, err
	}

	weekly, err = popularItemHelper(c, s, restaurantID, now, 0, 0, -7)
	if err != nil {
		return nil, err
	}

	monthly, err = popularItemHelper(c, s, restaurantID, now, 0, -1, 0)
	if err != nil {
		return nil, err
	}

	quarterly, err = popularItemHelper(c, s, restaurantID, now, 0, -3, 0)
	if err != nil {
		return nil, err
	}

	yearly, err = popularItemHelper(c, s, restaurantID, now, -1, 0, 0)
	if err != nil {
		return nil, err
	}

	all_time, err = popularItemHelper(c, s, restaurantID, now, 0, 0, 0)
	if err != nil {
		return nil, err
	}

	report := &entity.HomePageReport{
		Daily:       daily,
		Weekly:      weekly,
		Monthly:     monthly,
		Quarterly:   quarterly,
		Yearly:      yearly,
		AllTime:     all_time,
		GeneratedAt: now,
	}
	return report, nil
}

// Helper function to retrive popular items, total revenue and popular category
func popularItemHelper(c *gin.Context, s service, restaurantID string, end_time time.Time, year int, month int, day int) (*entity.HomePageReportItem, error) {

	var return_item entity.HomePageReportItem

	if year == 0 && month == 0 && day == 0 {
		popular_item, err := s.repo.GetMostPopularItem(c, end_time, time.Time{}, restaurantID)
		if err != nil {
			return nil, err
		}

		if popular_item.ItemID != "N/A" {
			menu_item, err := s.menuItemRepo.Get(c, restaurantID, popular_item.ItemID)

			if err != nil {
				return nil, err
			}
			return_item.MostOrderedItem = menu_item

			category, err := s.categoriesRepo.Get(c, restaurantID, menu_item.CategoryID)
			if err != nil {
				return nil, err
			}
			return_item.MostOrderedCategory = category

		} else {
			// If no popular item exists, send an empty item and category
			menu_item := &entity.MenuItem{
				ID: "",
				Name: "",
				Description: "",
				Price: 0.0,
				IsSpecial: false,
				IsMenu: false,
				CategoryID: "",
			}
			return_item.MostOrderedItem = menu_item

			category := &entity.MenuItemCategory{
				ID: "",
				Name: "",
			}

			return_item.MostOrderedCategory = category
		}

		revenue, err := s.repo.GetTotalRevenue(c, restaurantID, end_time, time.Time{})
		if err != nil {
			return nil, err
		}

		return_item.MostOrderedItemQuantity = popular_item.Quantity
		return_item.TotalRevenue = revenue

	} else {

		popular_item, err := s.repo.GetMostPopularItem(c, end_time, end_time.AddDate(year, month, day), restaurantID)

		if err != nil {
			return nil, err
		}

		if popular_item.ItemID != "N/A" {
			menu_item, err := s.menuItemRepo.Get(c, restaurantID, popular_item.ItemID)

			if err != nil {
				return nil, err
			}
			return_item.MostOrderedItem = menu_item

			category, err := s.categoriesRepo.Get(c, restaurantID, menu_item.CategoryID)
			if err != nil {
				return nil, err
			}
			return_item.MostOrderedCategory = category
		} else {
			// If no popular item exists, send an empty item and category
			menu_item := &entity.MenuItem{
				ID: "",
				Name: "",
				Description: "",
				Price: 0.0,
				IsSpecial: false,
				IsMenu: false,
				CategoryID: "",
			}
			return_item.MostOrderedItem = menu_item

			category := &entity.MenuItemCategory{
				ID: "",
				Name: "",
			}

			return_item.MostOrderedCategory = category
		}


		revenue, err := s.repo.GetTotalRevenue(c, restaurantID, end_time, end_time.AddDate(year, month, day))
		if err != nil {
			return nil, err
		}

		return_item.MostOrderedItemQuantity = popular_item.Quantity
		return_item.TotalRevenue = revenue
	}

	return &return_item, nil
}
