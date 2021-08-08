import {storiesOf} from "@storybook/react";
import RestaurantCard from "../components/customer/restaurant-card";
import OrderMenuItem from "../components/customer/order-menu-item";
import {ANDIAMO_TRATTORIA, EXAMPLE_ORDER, HOKKIEN_NOODLES} from "./test-data";
import {ServiceProvider} from "../components/services/service-context";
import CartPanel from "../components/customer/cart-panel";
import CartCheckout from "../pages/customer/CartCheckout";
import OrderHistoryCard from "../components/customer/order-history-card";

storiesOf('Customer', module)
    .add('Restaurant card', () => {
        const viewMenu = () => {};

        return (
            <RestaurantCard restaurant={ANDIAMO_TRATTORIA} viewMenu={viewMenu} />
        )
    })
    .add('Cart panel', () => {
        return (
            <ServiceProvider>
                <CartPanel/>
            </ServiceProvider>
        )
    })
    .add('Shopping cart checkout', () => {
        return (
            <CartCheckout/>
        )
    })
    .add('Order history card', () => {
        return (
            <OrderHistoryCard order={EXAMPLE_ORDER} />
        )
    });