import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Container, Flex, Header } from "../../components/utils/reusable-components";
import { Category, MenuItem, Restaurant } from "../../components/utils/data-interfaces";
import { CategoryOption } from "../../components/utils/data-types";
import CategoryButton from "../../components/customer/category-button";
import CartPanel from "../../components/customer/cart-panel";
import {
	ANDIAMO_MENU,
	ANDIAMO_TRATTORIA,
	IPPUDO,
	IPPUDO_MENU,
	MCDONALDS,
	MCDONALDS_MENU,
} from "../../stories/test-data";
import OrderMenuItem from "../../components/customer/order-menu-item";
import { useCarts } from "../../components/services/cart-context";
import ExistingRestaurantWarningDialog from "../../components/customer/existing-restaurant-warning-dialog";
import styled from "styled-components";
import { myMenuService } from "../../components/services/mymenu-service";
import { useServices } from "../../components/services/service-context";

const CartPanelDivider = styled.div`
	width: 100%;
	height: 50px;
	margin-top: 10px;
	margin-bottom: 10px;
`;

export default function CustomerRestaurantMenu() {
	const { myMenuService } = useServices();
	let { restaurantId } = useParams() as any;
	const [restaurant, setRestaurant] = useState<Restaurant | undefined>();
	const [categories, setCategories] = useState<Category[]>([]);
	const [selected, setSelectedCategory] = useState<string>("all");
	const [menu, setMenu] = useState<MenuItem[]>([]);
	const [cart, cartActions] = useCarts();
	const [isExistingRestaurant, changeRestaurantView] = useState(
		!cartActions.isNewRestaurant() && cart.restaurantId != restaurantId && !cartActions.isEmpty()
	);
	const tablenum = myMenuService.getTableNum();

	useEffect(() => {
		if (cartActions.isNewRestaurant() || cartActions.isEmpty()) {
			cartActions.restaurantCheckIn(restaurantId);
		}
		getRestaurant();
		getMenu();
		getCategories();
	}, []);

	const getRestaurant = async () => {
		const json = await myMenuService.getRestaurantCust(restaurantId);
		if (json && json.item) {
			setRestaurant(json.item);
		}
	};

	const getMenu = async () => {
		const json = await myMenuService.getMenuCust(restaurantId);
		if (json && json.data) {
			setMenu(json.data);
		}
	};

	const getCategories = async () => {
		const json = await myMenuService.getMenuItemCategoriesCust(restaurantId);
		if (json && json.data) {
			setCategories(json.data);
			console.log(json.data);
		}
	};

	const loadCategories = () => {
		return categories.map((c: Category) => (
			<CategoryButton selected={selected} option={c.id} changeOption={() => changeOption(c.id)}>
				{c.name}
			</CategoryButton>
		));
	};

	const changeOption = (category: string) => {
		setSelectedCategory(category);
	};

	const loadMenu = () => {
		return menu.map((i: MenuItem) => {
			if (!i.is_special && (i.category_id == selected || selected == "all")) {
				return <OrderMenuItem item={i} canOrder={restaurantId == cart.restaurantId} />;
			}
		});
	};

	const loadMenuSpecial = () => {
		return menu.map((i: MenuItem) => {
			if (i.is_special && (i.category_id == selected || selected == "all")) {
				return <OrderMenuItem item={i} canOrder={restaurantId == cart.restaurantId} />;
			}
		});
	};

	const loadChefsSpecial = () => {
		return menu.map((i: MenuItem) => {
			if (selected == "special" && i.is_special) {
				return <OrderMenuItem item={i} canOrder={restaurantId == cart.restaurantId} />;
			}
		});
	};

	const loadRestaurantDetails = () => {
		if (restaurant == undefined) {
			return;
		}
		return (
			<Container>
				<Header>{restaurant.name}</Header>
				{tablenum && <p>Table #{tablenum}</p>}
				<Flex>
					<CategoryButton selected={selected} option="all" changeOption={() => changeOption("all")}>
						All
					</CategoryButton>
					<CategoryButton selected={selected} option="special" changeOption={() => changeOption("special")}>
						Chef's Special
					</CategoryButton>
					{loadCategories()}
				</Flex>
				{selected == "special" && loadChefsSpecial()}
				{selected != "special" && loadMenuSpecial()}
				{selected != "special" && loadMenu()}
			</Container>
		);
	};

	return (
		<Container>
			{isExistingRestaurant && <ExistingRestaurantWarningDialog restaurantId={restaurantId} />}
			{loadRestaurantDetails()}
			<CartPanelDivider />
			<CartPanel />
		</Container>
	);
}
