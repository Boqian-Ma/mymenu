export const ANDIAMO_TRATTORIA = {
    id: "123",
    location: "Summer Hill",
    name: "Andiamo Trattoria",
    type: "Italian",
    email: "italian@gmail.com",
    phone: "1234567890",
    website: "test.com",
    businessHours: "closed",
    file: "italy.jpg",
    cuisine: "European"
};

export const MCDONALDS = {
    id: "234",
    location: "Kensington",
    name: "McDonald's",
    type: "Fast Food",
    email: "fastfood@gmail.com",
    phone: "1234567890",
    website: "test.com",
    businessHours: "closed",
    file: "fastfood.jpg",
    cuisine: "North America"
};

export const IPPUDO = {
    id: "345",
    location: "Sydney",
    name: "Ippudo",
    type: "Japanese",
    email: "japan@gmail.com",
    phone: "1234567890",
    website: "test.com",
    businessHours: "closed",
    file: "japan.jpg",
    cuisine: "Asian"
};

export const HOKKIEN_NOODLES = {
    id: "123",
    name: "Hokkien Noodles",
    price: 10.99,
    description: "Hokkien-styled noodles"
};

export const RAMEN = {
    id: "234",
    name: "Tonkotsu ramen",
    price: 12.49,
    description: "Fresh pork broth noodles"
};

export const SASHIMI = {
    id: "345",
    name: "Sashimi",
    price: 9.99,
    description: "Fresh 5 pc sashimi"
};

export const PIZZA = {
    id: "456",
    name: "Pizza",
    price: 14.00,
    description: "Large pizza"
};

export const LASAGNA = {
    id: "567",
    name: "Lasagna",
    price: 12.00,
    description: "Fresh beef lasagna"
};

export const CHEESEBURGER = {
    id: "678",
    name: "Cheeseburger",
    price: 13.99,
    description: "Australian beef patty with melted cheese, lettuce, tomato, pickles, onions and secret sauce"
};

export const COKE = {
    id: "789",
    name: "Coca Cola",
    price: 2.99,
    description: "375 mL can of Coke"
};

export const MELON_SODA = {
    id: "890",
    name: "Melon Soda",
    price: 2.99,
    description: "Japanese soft drink"
};

export const EXAMPLE_ORDER = {
    id: "123",
    items: [
        {
            item_id: "123",
            item_name: "Cheeseburger",
            item_price: 8,
            quantity: 2
        },
        {
            item_id: "234",
            item_name: "Chicken Nuggets",
            item_price: 12,
            quantity: 24
        },
        {
            item_id: "345",
            item_name: "Frozen Coke",
            item_price: 2,
            quantity: 2
        },
    ],
    restaurant_id: "234",
    status: "completed",
    table_num: 3,
    total_cost: 22,
    user_id: "12345",
    created_at:''
};

export const EXAMPLE_RESTAURANT_LIST = [ANDIAMO_TRATTORIA, MCDONALDS, IPPUDO];

export const ANDIAMO_MENU = [PIZZA, LASAGNA, COKE];

export const MCDONALDS_MENU = [CHEESEBURGER, COKE];

export const IPPUDO_MENU = [RAMEN, SASHIMI, COKE, MELON_SODA];