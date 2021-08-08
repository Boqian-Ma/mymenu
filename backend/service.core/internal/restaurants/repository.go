package restaurants

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access restaurant data in the DB
type Repository interface {
	// Creates a new restaurant
	Create(c *gin.Context, restaurant *entity.Restaurant) error
	// Creates a new relationship between a user and a restaurant
	CreateMember(c *gin.Context, resMember *entity.RestaurantMember) error
	// Returns a restaurants details
	Get(c *gin.Context, restaurantID string) (*entity.Restaurant, error)
	// Returns a list of all restaurants
	List(c *gin.Context) ([]*entity.Restaurant, error)
	// Updates a restaurants details
	Update(c *gin.Context, restaurant *entity.Restaurant) error
	// Returns a list of all the restaurants the user has access to
	ListRestaurantsForUser(c *gin.Context) ([]*entity.Restaurant, error)
	// Returns a list of
	RecommendedList(c *gin.Context) ([]*entity.Restaurant, error)

	// Retuans 3 top cuisines
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access restaurant data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) Create(c *gin.Context, restaurant *entity.Restaurant) error {
	return r.db.Insert(c, "restaurants", restaurant)
}

func (r repository) CreateMember(c *gin.Context, resMember *entity.RestaurantMember) error {
	return r.db.Insert(c, "restaurant_members", resMember)
}

func (r repository) Update(c *gin.Context, restaurant *entity.Restaurant) error {
	return r.db.Update(c, restaurant, "update restaurants set ... where id=$1", restaurant.ID)
}

func (r repository) Get(c *gin.Context, restaurantID string) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant

	if err := r.db.Select(c, &restaurant, "select * from restaurants where id=$1", restaurantID); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// todo pagination &/ filtering &/ sorting as req for FE
func (r repository) List(c *gin.Context) ([]*entity.Restaurant, error) {
	var restaurants []*entity.Restaurant

	if err := r.db.Select(c, &restaurants, "select * from restaurants"); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r repository) ListRestaurantsForUser(c *gin.Context) ([]*entity.Restaurant, error) {
	var resMembers []*entity.RestaurantMember

	if err := r.db.Select(c, &resMembers, "select * from restaurant_members where user_id=$1", c.GetString("userID")); err != nil {
		return nil, err
	}

	var restaurantIDs []string
	for _, member := range resMembers {
		restaurantIDs = append(restaurantIDs, member.RestaurantID)
	}

	return r.constructRestaurantData(c, restaurantIDs)
}

// todo this should be done in a single query just slightly tricker parsing
func (r repository) constructRestaurantData(c *gin.Context, restaurantIDs []string) ([]*entity.Restaurant, error) {
	var restaurants []*entity.Restaurant
	for _, id := range restaurantIDs {
		restaurant, err := r.Get(c, id)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func (r repository) RecommendedList(c *gin.Context) ([]*entity.Restaurant, error) {

	restaurants := []*entity.Restaurant{}

	cuisines, err := getTopCuisinesByUser(r, c)

	if err != nil {
		fmt.Println("yeet")
		return nil, err
	}

	for _, cuisine := range cuisines {
		var res []*entity.Restaurant
		res, err := getTwoRestaurantsWithCuisine(r, c, cuisine.Cuisine)
		if err != nil {
			fmt.Println("yeet1")

			return nil, err
		}
		restaurants = append(restaurants, res...)
	}

	return restaurants, nil
}

func getTwoRestaurantsWithCuisine(r repository, c *gin.Context, cuisine string) ([]*entity.Restaurant, error) {

	var restaurants []*entity.Restaurant

	if err := r.db.Select(c, &restaurants, "select * from restaurants where restaurants.cuisine=$1 limit 2", cuisine); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func getTopCuisinesByUser(r repository, c *gin.Context) ([]*entity.Cuisine, error) {
	// Get top three cuisines
	var cuisines []*entity.Cuisine

	if err := r.db.Select(c, &cuisines, "select restaurants.cuisine from restaurants join orders on (orders.restaurant_id=restaurants.id) join users on (orders.user_id=users.id) where users.id=$1 and orders.status='completed' group by restaurants.cuisine order by count(restaurants.cuisine) DESC Limit 3", c.GetString("userID")); err != nil {
		return nil, err
	}

	return cuisines, nil
}
