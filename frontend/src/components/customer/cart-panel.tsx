import React, {useEffect} from 'react';
import {Bold, Container, Flex, Grid, Margin} from "../utils/reusable-components";
import {IconButton, ShoppingCartIcon} from "evergreen-ui";
import styled from "styled-components";
import {useHistory} from "react-router";
import {useCarts} from "../services/cart-context";

const CartContainer = styled.div`
	display: flex;
	background-color: white;
	position: fixed;
	bottom: 0;
	width: 1000px;
	margin-left: auto;
	height: 50px;
	padding: 1rem;
	border-radius: 5px;
	border: 1px solid grey;
	z-index: 1;
`;

const ShoppingCartButton = styled(IconButton)`
    margin: auto;
`;

export default function CartPanel() {
    const [carts, cartActions] = useCarts();
    const history = useHistory();

    const checkoutCart = () => {
        history.push("/customer/checkout");
    };

    useEffect(() => {}, [carts.totalCost]);

    return (
        <CartContainer>
            <Grid>
                <Bold>Total: ${cartActions.getTotalCost()}</Bold>
            </Grid>
            <Grid>
                <ShoppingCartButton icon={ShoppingCartIcon} intent="success" onClick={checkoutCart} />
            </Grid>
        </CartContainer>
    );
}