import { TableStatus, User } from "../utils/data-types";
import { CreateOrderRequest } from "../utils/data-interfaces";

export class myMenuService {
	rootUrl: string;
	frontendUrl: string;
	token: string; // Authentication token
	userType: User | undefined;
	hasRestaurant: boolean;
	currRestaurant: string;
	tableNum: number | undefined;

	constructor() {
		this.rootUrl = "http://localhost:5000";
		this.frontendUrl = "http://localhost:3000";
		this.token = "";
		this.userType = undefined;
		this.hasRestaurant = false;
		this.currRestaurant = "";
		this.tableNum = undefined;

		if (localStorage.getItem("token")) {
			this.token = localStorage.getItem("token")!!;
		}
	}

	public logOut = async () => {
		return fetch(`${this.rootUrl}/api/v1/logout`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				this.token = "";
				localStorage.removeItem("token");
				localStorage.removeItem("userType");
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurant:", err);
			});
	};

	public getCurrentUser = async () => {
		return fetch(`${this.rootUrl}/api/v1/users/current`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurant:", err);
			});
	};

	public updateCurrentUser = async (email: string, name: string, password: string) => {
		let payload;
		if (password !== "") {
			payload = { details: { email: email, name: name }, newPassword: password };
		} else {
			payload = { details: { email: email, name: name } };
		}
		console.log(payload);
		return fetch(`${this.rootUrl}/api/v1/users/current`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify(payload),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot reset password:", err);
			});
	};

	public login = async (email: string, userType: string, password: string) => {
		return fetch(`${this.rootUrl}/api/v1/login`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer`,
			},
			body: JSON.stringify({
				email: email,
				userType: userType.toLowerCase(),
				password: password,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400 || response.status === 409) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot login:", err);
			});
	};

	public loginGuest = async () => {
		return fetch(`${this.rootUrl}/api/v1/login/guest`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer`,
			},
			body: JSON.stringify({}),
		})
			.then((response) => {
				if (response.ok || response.status === 400 || response.status === 409) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot login:", err);
			});
	};

	public register = async (email: string, userType: string, name: string, password: string) => {
		return fetch(`${this.rootUrl}/api/v1/register`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer`,
			},
			body: JSON.stringify({
				accountType: userType.toLowerCase(),
				details: { email: email, name: name },
				password: password,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400 || response.status === 409) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot register account:", err);
			});
	};

	public resetPassword = async (email: string, newPassword: string) => {
		return fetch(`${this.rootUrl}/api/v1/reset`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer`,
			},
			body: JSON.stringify({
				email: email,
				newPassword: newPassword,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot reset password:", err);
			});
	};

	public addRestaurant = async (
		name: string,
		type: string,
		location: string,
		email: string,
		phone: string,
		website: string,
		businessHours: string,
		file: string,
		file64: string,
		cuisine: string
	) => {
		this.hasRestaurant = true;
		return fetch(`${this.rootUrl}/api/v1/restaurants`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				name: name,
				type: type,
				location: location,
				email: email,
				phone: phone,
				website: website,
				businessHours: businessHours,
				file: file,
				file64: file64,
				cuisine: cuisine,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot register:", err);
			});
	};

	public getRestaurants = async (mine: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants?mine=${mine}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurants:", err);
			});
	};

	public getRestaurant = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurant:", err);
			});
	};

	public getRestaurantCust = async (res_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${res_id}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurant:", err);
			});
	};

	public getUserRestaurants = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants?mine=true`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get restaurants:", err);
			});
	};

	public editRestaurant = async (
		resID: string,
		name: string,
		type: string,
		location: string,
		email: string,
		phone: string,
		website: string,
		businessHours: string,
		file: string,
		file64: string,
		cuisine: string
	) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${resID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				name: name,
				type: type,
				location: location,
				email: email,
				phone: phone,
				website: website,
				businessHours: businessHours,
				file: file,
				file64: file64,
				cuisine: cuisine,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot edit restaurant:", err);
			});
	};

	public getTables = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/tables`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get tables:", err);
			});
	};

	public getTable = async (tableNum: number) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/tables/${tableNum}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get tables:", err);
			});
	};

	public createTable = async (tableNum: number, numSeats: number) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/tables`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				table_num: tableNum,
				num_seats: numSeats,
				restaurant_id: this.currRestaurant,
				status: "Free",
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot create table:", err);
			});
	};

	public updateTable = async (tableNum: number, numSeats: number, status: TableStatus) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/tables/${tableNum}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				num_seats: numSeats,
				status: status,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot update table:", err);
			});
	};

	public updateTableStatus = async (tableNum: number, action: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/tables/${tableNum}/${action}`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot toggle status:", err);
			});
	};

	public updateTableStatusCust = async (res_id: string | undefined, tableNum: number, action: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${res_id}/tables/${tableNum}/${action}`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot toggle status:", err);
			});
	};

	public getMenuItems = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/menu_items`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get menu items:", err);
			});
	};

	public getMenuCust = async (res_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${res_id}/menu`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get menu items:", err);
			});
	};

	public getMenuItem = async (item_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/menu_items/${item_id}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get menu items:", err);
			});
	};

	public addMenuItem = async (
		name: string,
		description: string,
		price: number,
		is_special: boolean,
		is_menu: boolean,
		categoryID: string,
		allergy: string,
		file: string,
		file64: string
	) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/menu_items`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				name: name,
				description: description,
				price: price,
				is_special: is_special,
				is_menu: is_menu,
				category_id: categoryID,
				allergy: allergy,
				file: file,
				file64: file64
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot edit menu item:", err);
			});
	};

	public editMenuItem = async (
		itemID: string,
		name: string,
		description: string,
		price: number,
		is_special: boolean,
		is_menu: boolean,
		categoryID: string,
		allergy: string,
		file: string,
		file64: string
	) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/menu_items/${itemID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				name: name,
				description: description,
				price: price,
				is_special: is_special,
				is_menu: is_menu,
				category_id: categoryID,
				allergy: allergy,
				file: file,
				file64: file64
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot add menu item:", err);
			});
	};

	public deleteMenuItem = async (itemID: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/menu_items/${itemID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot delete menu item:", err);
			});
	};

	public getMenuItemCategories = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/categories`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get menu item categories:", err);
			});
	};

	public getMenuItemCategoriesCust = async (restaurantId: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${restaurantId}/categories`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get menu item categories:", err);
			});
	};

	public getMenuItemCategory = async (cat_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/categories/${cat_id}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot get menu item category:", err);
			});
	};

	public createMenuItemCategory = async (name: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/categories`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				name: name,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot edit menu item:", err);
			});
	};

	public createOrder = async (table_num: number, item: any) => {
		return fetch(`${this.rootUrl}/api/v1/orders`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				restaurant_id: this.currRestaurant,
				table_num: table_num,
				items: item,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public createOrderCust = async (res_id: string | undefined, table_num: number, item: any) => {
		return fetch(`${this.rootUrl}/api/v1/orders`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				restaurant_id: res_id,
				table_num: parseInt(String(table_num)),
				items: item,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public moveOrder = async (order_id: string, new_table_num: number) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}/move`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				restaurant_id: this.currRestaurant,
				new_table_number: new_table_num,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public serveOrder = async (order_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}/serve`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public cancelOrder = async (order_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}/cancel`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public completeOrder = async (order_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}/complete`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot create order:", err);
			});
	};

	public removeItemFromOrder = async (order_id: string, item_id: string) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}/items/${item_id}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public addItemToOrder = async (order_id: string, item_id: string, quantity: number) => {
		return fetch(`${this.rootUrl}/api/v1/orders/${order_id}`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
			body: JSON.stringify({
				order_id: order_id,
				item_id: item_id,
				quantity: quantity,
			}),
		})
			.then((response) => {
				if (response.ok || response.status === 400) {
					return response.json();
				} else {
					throw new Error(`Status: ${response.status}`);
				}
			})
			.catch((err) => {
				console.error("Cannot add item to order:", err);
			});
	};

	public getOrders = async (status: string, active: boolean) => {
		return fetch(`${this.rootUrl}/api/v1/orders?res_id=${this.currRestaurant}&status=${status}&active=${active}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public getTableOrder = async (table_num: number) => {
		return fetch(`${this.rootUrl}/api/v1/order?res_id=${this.currRestaurant}&table_num=${table_num}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public getOrderHistory = async (userId: string) => {
		console.log(userId);
		return fetch(`${this.rootUrl}/api/v1/orders?usr_id=${userId}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public getOrderHistoryRestaurant = async () => {
		return fetch(`${this.rootUrl}/api/v1/orders?res_id=${this.currRestaurant}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	}

	public getReport = async () => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${this.currRestaurant}/report?type=home`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public updateAuthToken = (token: string) => {
		this.token = token;
	};

	public getUserType = () => {
		return this.userType;
	};

	public getCurrRestaurant = () => {
		return this.currRestaurant;
	};

	public setCurrRestaurant = (resID: string) => {
		this.currRestaurant = resID;
	};

	public getRestaurantDetails = (resID: string) => {
		return fetch(`${this.rootUrl}/api/v1/restaurants/${resID}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get order:", err);
			});
	};

	public getRecommendedRestaurants = async () => {
		return fetch(`${this.rootUrl}/api/v1/recommended-restaurants`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${this.token}`,
			},
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Status: ${response.status}`);
				}
				return response.json();
			})
			.catch((err) => {
				console.error("Cannot get recommended restaurants:", err);
			});
	};

	public getRootUrl = () => {
		return this.frontendUrl;
	};

	public setTableNum = (tableNum: number) => {
		this.tableNum = tableNum;
	};

	public getTableNum = () => {
		return this.tableNum;
	};
}
