import React, { useState } from "react";
import { Button, Dialog, Pane, TextInput } from "evergreen-ui";
import { User } from "../utils/data-types";
import styled from "styled-components";
import { useServices } from "../services/service-context";
import { useHistory } from "react-router-dom";
import TextField from "../text-field";
import {useAuthenticationContext} from "../services/authentication-context";
import {useCarts} from "../services/cart-context";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

interface Props {
    restaurantId: string;
}

export default function ExistingRestaurantWarningDialog(props: Props) {
    const { restaurantId } = props;
    const [isShown, setIsShown] = useState(true);
    const [cart, cartActions] = useCarts();

    const clear = () => {
        cartActions.clearCart();
        cartActions.restaurantCheckIn(restaurantId);
        setIsShown(false);
    };

    return (
        <Pane>
            <DialogContainer>
                <Dialog
                    isShown={isShown}
                    title="Warning!"
                    onCancel={() => setIsShown(false)}
                    onConfirm={clear}
                    confirmLabel="Clear"
                    onCloseComplete={() => setIsShown(false)}
                >
                    You have existing items in your cart. Would you like to clear it to order from this restaurant?
                </Dialog>
            </DialogContainer>
        </Pane>
    );
}
