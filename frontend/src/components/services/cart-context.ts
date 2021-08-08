import {createStore, createHook, StoreActionApi} from 'react-sweet-state';
import {MenuItem} from "../utils/data-interfaces";

interface State {
    cart: MenuItem[];
    totalCost: number;
    restaurantId: string | undefined;
}

type StoreApi = StoreActionApi<State>;
type Actions = typeof actions;

const initialState: State = {
    cart: [],
    totalCost: 0,
    restaurantId: undefined
};

const calculateTotalCost = (cart: MenuItem[]): number => {
    let cost = 0;
    cart.forEach((i: MenuItem) => {
        cost += i.price;
    });

    return Math.round(cost*100)/100;
};

const actions = {
    restaurantCheckIn: (restaurantId: string) => ({ setState }: StoreApi) => {
        setState({
           restaurantId: restaurantId
        });
    },

    addItemToCart: (item: MenuItem) => ({ setState, getState }: StoreApi) => {
        const { cart } = getState();

        cart.push(item);
        setState({
            totalCost: calculateTotalCost(cart)
        });
    },

    removeItemFromCart: (item: MenuItem, index: number) => ({ setState, getState }: StoreApi) => {
        const { cart, totalCost } = getState();
        cart.splice(index, 1);
        setState({
            totalCost: calculateTotalCost(cart)
        });
    },

    formatCart: () => ({ getState }: StoreApi) => {
        const { cart } = getState();
        const map = new Map<string, number>();

        cart.forEach((i: MenuItem) => {
            // @ts-ignore
            map.set(i.id, map.get(i.id) == undefined ? 1 : map.get(i.id) + 1);
        });
        return map;
    },

    clearCart: () => ({ setState }: StoreApi) => {
        setState({
            cart: [],
            totalCost: 0
        });
    },

    getCart: () => ({ getState }: StoreApi) => {
        return getState().cart;
    },

    getTotalCost: () => ({ getState }: StoreApi) => {
        return getState().totalCost;
    },

    getRestaurantId: () => ({ getState }: StoreApi) => {
        return getState().restaurantId;
    },

    isNewRestaurant: () => ({ getState }: StoreApi) => {
        return getState().restaurantId == undefined;
    },

    isEmpty: () => ({ getState }: StoreApi) => {
        return getState().cart.length == 0;
    },

};

const Store = createStore<State, Actions>({
    initialState,
    actions
});

export const useCarts = createHook(Store);