import React, { useEffect, useState } from "react";
import {
	AlignedHeader,
	AlignedText,
	Bold,
	Container,
	Divider,
	Flex,
	Grid,
	Header,
} from "../../components/utils/reusable-components";
import TextField from "../../components/text-field";
import CartItem from "../../components/customer/cart-item";
import styled from "styled-components";
import { Button, DeleteIcon, ShoppingCartIcon } from "evergreen-ui";
import { CreateOrderRequest, MenuItem } from "../../components/utils/data-interfaces";
import { useCarts } from "../../components/services/cart-context";
import { useServices } from "../../components/services/service-context";

const CheckoutButton = styled(Button)`
	width: 400px;
	margin-left: 1rem;
`;

const EmptyCartText = styled(AlignedText)`
	color: red;
`;

export default function CartCheckout() {
	const { myMenuService } = useServices();
	const [carts, cartActions] = useCarts();
	const [tableNumber, setTableNumber] = useState(myMenuService.getTableNum() || 1);
	const [name, setName] = useState("");
	const [email, setEmail] = useState("");

	const createOrderRequest = (
		items: Map<string, number>,
		restaurantId: string,
		status: string,
		totalCost: number
	) => {
		let orderRequest: CreateOrderRequest = {
			items: items,
			restaurant_id: restaurantId,
			status: status,
			total_cost: totalCost,
		};
		return orderRequest;
	};

	const renderItemsInCart = () => {
		return cartActions.getCart().map((item: MenuItem, idx: number) => {
			return <CartItem item={item} index={idx} />;
		});
	};

	const checkoutCart = async () => {
		// TODO: create and send invoice
		// TODO: backend API calls
		const items = cartActions.formatCart();
		const totalCost = cartActions.getTotalCost();
		const orderRequest: CreateOrderRequest = createOrderRequest(items, "123", "ordered", totalCost);
		try {
			const obj = Object.fromEntries(items);
			await myMenuService.createOrderCust(cartActions.getRestaurantId(), tableNumber, obj);
			await myMenuService.updateTableStatusCust(cartActions.getRestaurantId(), tableNumber, "occupy");
			cartActions.clearCart();
		} catch (e) {
			throw e;
		} finally {
		}
	};

	const clearCart = () => {
		cartActions.clearCart();
	};

	useEffect(() => {
		renderItemsInCart();
	}, [carts.cart]);

	return (
		<Container>
			<Header>Checkout</Header>
			{tableNumber && <Header>Table #{tableNumber}</Header>}
			{!cartActions.isEmpty() && (
				<Grid>
					<AlignedHeader>Enter Details</AlignedHeader>
					<TextField
						label="Table Number"
						input={String(tableNumber)}
						updateInput={setTableNumber}
						type="number"
					/>
					<TextField label="Name" input={name} updateInput={setName} type="text" />
					<TextField label="Email" input={email} updateInput={setEmail} type="email" />
				</Grid>
			)}
			<Divider />
			{renderItemsInCart()}
			{cartActions.isEmpty() && <EmptyCartText>No items in cart</EmptyCartText>}
			<Divider />
			<Bold>Total: ${cartActions.getTotalCost()}</Bold>
			<Flex>
				<CheckoutButton
					iconBefore={ShoppingCartIcon}
					appearance="primary"
					disabled={cartActions.isEmpty()}
					onClick={checkoutCart}
				>
					Checkout
				</CheckoutButton>
				<CheckoutButton
					iconBefore={DeleteIcon}
					appearance="primary"
					intent="danger"
					disabled={cartActions.isEmpty()}
					onClick={clearCart}
				>
					Clear Cart
				</CheckoutButton>
			</Flex>
		</Container>
	);
}
