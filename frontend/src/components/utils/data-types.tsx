export type User = "Manager" | "Customer";

export type MenuOption = "category" | "add-item" | "add-menu" | "table";

export type NavigationOption =
	| "home"
	| "tables"
	| "kitchen"
	| "menus"
	| "settings";

export type CategoryOption = "special" | "entree" | "main" | "desserts" | "drinks";

export type SettingOption = "personal" | "health" | "security" | "payment-methods" | "restaurant" | "reports" | "branding";

export type TableStatus = "-" | "Taken" | "Free"