import React, { useEffect } from "react";
import "./App.css";
import AddRestaurantDetails from "./pages/restaurant/AddRestaurantDetails";
import HomeDirectory from "./pages/HomeDirectory";
import RestaurantDashboard from "./pages/restaurant/RestaurantDashboard";
import RestaurantLandingPage from "./pages/restaurant/RestaurantLandingPage";
import ManageTables from "./pages/restaurant/ManageTables";
import ManageKitchen from "./pages/kitchen/ManageKitchen";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import NavigationBar from "./components/nav-bar";
import AddMenuDetails from "./pages/menus/AddMenuDetails";
import SettingsDashboard from "./pages/settings/SettingsDashboard";
import { Divider } from "./components/utils/reusable-components";
import CustomerDashboard from "./pages/customer/CustomerDashboard";
import CustomerRestaurantMenu from "./pages/customer/CustomerRestaurantMenu";
import CustomerLanding from "./pages/customer/CustomerLanding";
import CartCheckout from "./pages/customer/CartCheckout";
import { useAuthenticationContext } from "./components/services/authentication-context";
import OrderHistory from "./pages/customer/OrderHistory";
import UnauthorisedPage from "./pages/UnauthorisedPage";

export default function App() {
	const [auth, authActions] = useAuthenticationContext();

	const renderNavBar = () => {
		return <NavigationBar type={auth.userType} />;
	};

	useEffect(() => {}, [auth.userType]);

	useEffect(() => {
		if (localStorage.getItem("token")) {
			if (localStorage.getItem("userType")) {
				authActions.updateUser(localStorage.getItem("userType"));
			}
		}
	}, []);

	return (
		<div className="App">
			<Router>
				<Switch>
					<Route exact path="/">
						<HomeDirectory />
					</Route>
					<Route exact path="/admin/dashboard">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<RestaurantDashboard />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/admin/restaurant">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<RestaurantLandingPage />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/admin/restaurant/add">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<AddRestaurantDetails />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/admin/restaurant/menu">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<AddMenuDetails />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/admin/restaurant/tables">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<ManageTables />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/admin/restaurant/kitchen">
						{auth.userType === "Manager" ? (
							<div>
								{renderNavBar()}
								<ManageKitchen />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/customer/dashboard">
						{auth.userType === "Customer" ? (
							<div>
								{renderNavBar()}
								<CustomerDashboard />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route path="/customer/landing" children={<CustomerLanding />}></Route>
					<Route
						path="/customer/:restaurantId/menu"
						children={
							<div>
								{renderNavBar()}
								<CustomerRestaurantMenu />
							</div>
						}
					></Route>
					<Route exact path="/customer/orders">
						{auth.userType === "Customer" ? (
							<div>
								{renderNavBar()}
								<OrderHistory />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/customer/checkout">
						{auth.userType === "Customer" ? (
							<div>
								{renderNavBar()}
								<CartCheckout />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
					<Route exact path="/settings">
						{auth.userType !== undefined ? (
							<div>
								{renderNavBar()}
								<SettingsDashboard />
							</div>
						) : (
							<UnauthorisedPage />
						)}
					</Route>
				</Switch>
				<Divider />
			</Router>
		</div>
	);
}
