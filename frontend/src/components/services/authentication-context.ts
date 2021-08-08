import {createStore, createHook, StoreActionApi} from 'react-sweet-state';
import {User} from "../utils/data-types";

interface State {
    userType: User | undefined;
}

type StoreApi = StoreActionApi<State>;
type Actions = typeof actions;

const initialState: State = {
    userType: undefined
};

const actions = {
    updateUser: (type: User) => ({ setState }: StoreApi) => {
        setState({
            userType: type
        });
    },

    logOut: () => ({ setState }: StoreApi) => {
        setState({
            userType: undefined
        });
    },
};

const Store = createStore<State, Actions>({
    initialState,
    actions
});

export const useAuthenticationContext = createHook(Store);