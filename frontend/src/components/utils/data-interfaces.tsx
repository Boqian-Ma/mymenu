import {TableStatus} from "./data-types";

export interface Restaurant {
    id: string;
    location: string;
    name: string;
    type: string;
    email: string;
    phone: string;
    website: string;
    businessHours: string;
    file: string;
    cuisine: string;
}

export interface MenuItem {
    id: string;
    name: string;
    price: number;
    description: string;
    category_id: string;
    category_name: string;
    is_special: boolean;
    is_menu: boolean;
    file: string;
    allergy: string;
}

export interface Category {
    id: string;
    name: string;
}

export interface CreateOrderRequest {
    items: Map<string, number>;
    restaurant_id: string;
    status: string;
    total_cost: number;
}

export interface OrderItem {
    item_id: string;
    item_name: string;
    item_price: number;
    quantity: number;
}

export interface Order {
    id: string;
    items: OrderItem[];
    restaurant_id: string;
    status: string;
    table_num: number;
    total_cost: number;
    user_id: string;
    created_at: string;
}

export interface Table {
    table_num: number;
    num_seats: number;
    status: TableStatus;
    restaurant_id: string
}