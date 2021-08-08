import React, { useState } from "react";
import {
	BookIcon,
	BookmarkIcon,
	Button,
	HistoryIcon,
	HomeIcon,
	LogOutIcon,
	PanelTableIcon,
	PeopleIcon,
	ShoppingCartIcon,
	WrenchIcon,
} from "evergreen-ui";
import { Container, Grid } from "./utils/reusable-components";
import styled from "styled-components";
import { User, NavigationOption } from "./utils/data-types";
import { useHistory } from "react-router";
import { useServices } from "./services/service-context";
import { useAuthenticationContext } from "./services/authentication-context";
import { useCarts } from "./services/cart-context";

const NavigationContainer = styled(Container)`
	display: flex;
	width: 1000px;
	margin-left: auto;
	padding: 1rem;
	box-shadow: 0 1px 4px rgba(255, 255, 255, 1);
`;

const NavbarButton = styled.button`
	font-size: 1em;
	padding: 0.25em 1em;
	background: none;
	width: 150px;
	color: white;
	cursor: pointer;
	border: none;
`;

const NavContainer = styled.div`
	width: 100%;
	box-shadow: 0 1px 4px rgba(255, 255, 255, 1);
`;

const LogoGrid = styled(Grid)`
	padding: 0 30px 0 10px;
`;

interface Props {
	type: User | undefined;
}

export default function NavigationBar(props: Props) {
	const [userType, setType] = useState<User | undefined>(props.type);
	const [auth, authActions] = useAuthenticationContext();
	const { myMenuService } = useServices();
	const [cart, cartActions] = useCarts();
	const history = useHistory();
	let [selected, setSelection] = useState<NavigationOption>("home");

	const logOut = () => {
		authActions.logOut();
		myMenuService.logOut();
		history.push("/");
	};

	const goHome = () => {
		myMenuService.setCurrRestaurant("");
		switch (userType) {
			case "Manager":
				history.push("/admin/dashboard");
				break;
			case "Customer":
				history.push("/customer/dashboard");
				break;
		}
	};

	const goToTablesAdmin = () => {
		if (myMenuService.getCurrRestaurant() !== "") {
			history.push("/admin/restaurant/tables");
		}
	};

	const goToKitchenAdmin = () => {
		if (myMenuService.getCurrRestaurant() !== "") {
			history.push("/admin/restaurant/kitchen");
		}
	};

	const goToRecentMenu = () => {
		if (cart.restaurantId) {
			history.push(`/customer/${cart.restaurantId}/menu`);
		}
	};

	const goToOrderHistory = () => {
		history.push("/customer/orders");
	};

	const goToCart = () => {
		history.push("/customer/checkout");
	};

	const goToMenus = () => {
		if (myMenuService.getCurrRestaurant() !== "") {
			history.push("/admin/restaurant/menu");
		}
		setSelection("menus");
	};

	const goToSettings = () => {
		// Will handle user type at dashboard
		history.push("/settings");
		setSelection("settings");
	};

	const renderNavBar = () => {
		switch (userType) {
			case "Manager":
				return managerNavigation();
			case "Customer":
				return customerNavigation();
			case undefined:
				return;
		}
	};

	const managerNavigation = () => {
		return (
			<NavigationContainer>
				<Grid>
					<Button iconBefore={HomeIcon} width={150} onClick={goHome} appearance="primary" intent="danger">
						Home
					</Button>
				</Grid>
				<Grid>
					<Button
						iconBefore={PanelTableIcon}
						width={150}
						onClick={goToTablesAdmin}
						appearance="primary"
						intent="danger"
					>
						Manage Tables
					</Button>
				</Grid>
				<Grid>
					<Button
						iconBefore={PeopleIcon}
						width={150}
						onClick={goToKitchenAdmin}
						appearance="primary"
						intent="danger"
					>
						Kitchen
					</Button>
				</Grid>
				<Grid>
					<Button iconBefore={BookIcon} width={150} onClick={goToMenus} appearance="primary" intent="danger">
						Menus
					</Button>
				</Grid>
				<Grid>
					<Button
						iconBefore={WrenchIcon}
						width={150}
						onClick={goToSettings}
						appearance="primary"
						intent="danger"
					>
						Settings
					</Button>
				</Grid>
				<Grid>
					<Button iconBefore={LogOutIcon} width={150} onClick={logOut} appearance="primary" intent="danger">
						Log Out
					</Button>
				</Grid>
			</NavigationContainer>
		);
	};

	const customerNavigation = () => {
		return (
			<NavigationContainer>
				<NavigationContainer>
					<Grid>
						<Button
							iconBefore={HomeIcon}
							width={150}
							onClick={goHome}
							appearance="primary"
							intent="success"
						>
							Home
						</Button>
					</Grid>
					<Grid>
						<Button
							iconBefore={HistoryIcon}
							width={150}
							onClick={goToOrderHistory}
							appearance="primary"
							intent="success"
						>
							Order History
						</Button>
					</Grid>
					<Grid>
						<Button
							iconBefore={WrenchIcon}
							width={150}
							onClick={goToSettings}
							appearance="primary"
							intent="success"
						>
							Settings
						</Button>
					</Grid>
					<Grid>
						<Button
							iconBefore={ShoppingCartIcon}
							width={150}
							onClick={goToCart}
							appearance="primary"
							intent="success"
						>
							Cart
						</Button>
					</Grid>
					<Grid>
						<Button
							iconBefore={LogOutIcon}
							width={150}
							onClick={logOut}
							appearance="primary"
							intent="success"
						>
							Log Out
						</Button>
					</Grid>
				</NavigationContainer>
			</NavigationContainer>
		);
	};

	return <Container>{userType != undefined && renderNavBar()}</Container>;
}
