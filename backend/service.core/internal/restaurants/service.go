package restaurants

import (
	"fmt"
	"math/rand"
	"time"
	"encoding/base64"
	"image/png"
	"os"
	"strings"
	"bytes"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
)

// Service encapsulates the usage logic for the restaurants service
type Service interface {
	// Creates a new restaurant attached to the current user's account
	Create(c *gin.Context, req CreateRestaurantRequest) (*entity.Restaurant, error)
	// Lists all the restaurants a user has access to
	List(c *gin.Context, mine bool) ([]*entity.Restaurant, error)
	// Returns a restaurant's details
	Get(c *gin.Context, restaurantID string) (*entity.Restaurant, error)
	// Updates a restaurants details
	Update(c *gin.Context, restaurantID string, req UpdateRestaurantRequest) (*entity.Restaurant, error)

	// Gets a list of restaurants based on a user's order history
	RecommendedList(c *gin.Context) ([]*entity.Restaurant, error)
}

type CreateRestaurantRequest struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Cuisine       string `json:"cuisine"`
	Location      string `json:"location"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Website       string `json:"website"`
	BusinessHours string `json:"businessHours"`
	File          string `json:"file"`
	File64		  string `json:"file64"`
}

type UpdateRestaurantRequest = CreateRestaurantRequest

func (m CreateRestaurantRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required),
		validation.Field(&m.Type, validation.Required),
		validation.Field(&m.Location, validation.Required),
		validation.Field(&m.File, validation.Required),
	)
}

type service struct {
	repo     Repository
	userRepo users.Repository
}

// NewService creates a new restaurants service
func NewService(repo Repository, userRepo users.Repository) Service {
	return service{repo, userRepo}
}

func saveImage(b64encode string, filename string) {
	//b := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAASMAAAEsCAYAAACIdtX4AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAA+ESURBVHhe7d0xiBzZmQfwmYvkTIECDVyghQ1GsIEMZ1iBL5jgDhQ4kGADLygRnOEMNsighT3wBRfswgYrcLAHPlBysA4Mo1DgZIILtODAyjTZbiaFylbZnGv0hnld3a/0urqq6+vu3w+En7ySurpn6uP7f6+qZv8ff/HobA9gYv+Q/hdgUooREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISwf+2T35+lNcBkdEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEhKEZACIoREIJiBISgGAEh7H94f3tvlH3w+lla9fPk+p202h2rfmZtPsPl7eJn1tAZASEoRkAIihEQwtbNjIaeeSxrU/K+z6mOz2l9dEZACIoREMJWxLRSK331xtu0eufwgw/Sqs53J6/SalzrasVrIsfnf/xzWg3jy199klZlkd7/ED4+OkirOqfff59W77z54Upazdr2yKYzAkJQjIAQdiqmlSwb37qMGe36tOmlz+bzh1+l1d/dHO79d3p5GUe+fPxZWs0a8j0OYdnI1aUdx0rENIAJKUZACGJah9r4dnTrYVr1U4osQ5uJZiVDR7YsmpWEev8dTl48TqtutXGsREwDmJBiBISwMTHtl1d/klbzfnJ6nFazVo1pbf9+9w9pNY0+cWbVaLJum/ge//vpb9NqGKWY9uPhvbSa96c3P6bV5tIZASEoRkAIihEQwv7vf/PF2dR5s2sedOGDm3fTat7rp/fTataqM6OpZ0Rd3r44Tau9vccnT9Jq3jbNjB4ePUirvb0rtw7TKp5VZ0ilmdH1u/+bVvO+f/k0rcqin+c6IyAExQgIYf/JNy/PxmrxauJXoyuC1SjFtLZlY1vkmFaKM+uKZd89O0mrvb2T0/KNqkeH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"
	b64data := b64encode[strings.IndexByte(b64encode, ',')+1:]
	//b64data := b[strings.IndexByte(b, ',')+1:]
	unbased, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		fmt.Println("Cannot decode b64")
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		fmt.Println("Bad png")
		panic("Bad png")
	}

	//f, err := os.OpenFile("../frontend/public/assets/images/" + filename + ".png", os.O_WRONLY|os.O_CREATE, 0777)
	f, err := os.OpenFile("../frontend/public/assets/images/" + filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Cannot open file")
		fmt.Println("../frontend/public/assets/images/" + filename + ".png")
		panic("Cannot open file")
	}
	//fmt.Println("Saving to ../frontend/public/assets/images/" + filename + ".png")
	fmt.Println("Saving to ../frontend/public/assets/images/" + filename)
	png.Encode(f, im)
}

func (s service) Create(c *gin.Context, req CreateRestaurantRequest) (*entity.Restaurant, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}
	resID := entity.GenerateRestaurantID()
	now := time.Now()
	restaurant := &entity.Restaurant{
		ID:            resID,
		CreatedAt:     now,
		UpdatedAt:     now,
		Name:          req.Name,
		Type:          req.Type,
		Cuisine:       req.Cuisine,
		Location:      req.Location,
		Email:         req.Email,
		Phone:         req.Phone,
		Website:       req.Website,
		BusinessHours: req.BusinessHours,
		File:          req.File,
	}
	if err := s.repo.Create(c, restaurant); err != nil {
		return nil, err
	}
	fmt.Println("HI")
	resMember := &entity.RestaurantMember{
		UserID:       c.GetString("userID"),
		RestaurantID: restaurant.ID,
	}
	if err := s.repo.CreateMember(c, resMember); err != nil {
		return nil, err
	}
	
	//saveImage(req.File64, resID)
	saveImage(req.File64, req.File)

	return restaurant, nil
}

func (s service) List(c *gin.Context, mine bool) ([]*entity.Restaurant, error) {
	if !mine {
		return s.repo.List(c)
	}

	user, err := s.userRepo.GetUser(c, c.GetString("userID"))
	if err != nil {
		return nil, err
	}

	if user.AccountType != entity.Manager {
		return nil, errors.BadRequest("Only managers can list their restaurants")
	}

	return s.repo.ListRestaurantsForUser(c)
}

func (s service) Get(c *gin.Context, restaurantID string) (*entity.Restaurant, error) {
	return s.repo.Get(c, restaurantID)
}

func (s service) Update(c *gin.Context, restaurantID string, req UpdateRestaurantRequest) (*entity.Restaurant, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	if err := auth.IsManagerOf(c, restaurantID); err != nil {
		return nil, err
	}

	restaurant, err := s.repo.Get(c, restaurantID)
	if err != nil {
		return nil, err
	}
	restaurant.Name = req.Name
	restaurant.Location = req.Location
	restaurant.Type = req.Type
	restaurant.UpdatedAt = time.Now()
	restaurant.Email = req.Email
	restaurant.Phone = req.Phone
	restaurant.Website = req.Website
	restaurant.BusinessHours = req.BusinessHours
	restaurant.File = req.File
	restaurant.Cuisine = req.Cuisine
	
	if err := s.repo.Update(c, restaurant); err != nil {
		return nil, err
	}
	
	//saveImage(req.File64, restaurantID)
	if req.File64 != "" {
		saveImage(req.File64, req.File)
	}

	return restaurant, nil
}

func (s service) RecommendedList(c *gin.Context) ([]*entity.Restaurant, error) {
	// Makes sure user exists
	_, err := s.userRepo.GetUser(c, c.GetString("userID"))

	if err != nil {
		return nil, err
	}

	rand_num := rand.Intn(9)

	chance := 1

	if rand_num <= chance {

		fmt.Println("randomised")
		// get all restaurants

		restaurants, err := s.repo.List(c)

		if err != nil {
			return nil, err
		}

		rand.Shuffle(len(restaurants), func(i, j int) {
			restaurants[i], restaurants[j] = restaurants[j], restaurants[i]
		})

		return restaurants[:6], nil

	} else {
		return s.repo.RecommendedList(c)
	}
}
