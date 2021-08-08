package entity

import "time"

type HomePageReport struct {
	Daily       *HomePageReportItem `json:"daily"`
	Weekly      *HomePageReportItem `json:"weekly"`
	Monthly     *HomePageReportItem `json:"monthly"`
	Quarterly   *HomePageReportItem `json:"quarterly"`
	Yearly      *HomePageReportItem `json:"yearly"`
	AllTime     *HomePageReportItem `json:"all_time"`
	GeneratedAt time.Time           `json:"generated_at"`
}

type HomePageReportItem struct {
	MostOrderedItem         *MenuItem         `json:"most_ordered_item"`
	MostOrderedItemQuantity int               `json:"most_ordered_item_quantity"`
	MostOrderedCategory     *MenuItemCategory `json:"most_ordered_category"`
	TotalRevenue            float64           `json:"total_revenue"`
}

type TotalRevenue struct {
	Revenue float64 `db:"total_cost"`
}
