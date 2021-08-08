import React, { useEffect, useState } from "react";
import { RouteComponentProps, useLocation } from "react-router-dom";
import queryString from "query-string";
import { useHistory } from "react-router";
import { Container, Grid, Flex } from "../../components/utils/reusable-components";
import { useServices } from "../../components/services/service-context";
import RegistrationDialog from "../../components/registration/registration-dialog";
import LoginDialog from "../../components/login/login-dialog";
import { Button } from "evergreen-ui";
import GuestLoginButton from "../../components/login/login-guest-button";
import { useAuthenticationContext } from "../../components/services/authentication-context";

export default function CustomerLanding() {
	const userType = "Customer";
	const { myMenuService } = useServices();
	const history = useHistory();
	const { search } = useLocation();
	const values = queryString.parse(search);
	const [restaurantName, setRestaurantName] = useState("");
	const [auth, authActions] = useAuthenticationContext();

	const getRestaurantName = async () => {
		const json = await myMenuService.loginGuest();
		if (json.hasOwnProperty("message")) {
			console.log("Authentication as guest error");
		} else {
			const userType = "Customer";
			myMenuService.updateAuthToken(json["token"]);
			localStorage.setItem("token", json["token"]);
			localStorage.setItem("userType", userType);
			authActions.updateUser(userType);
		}
		if (values.restaurantid != null) {
			const json = await myMenuService.getRestaurantDetails(values.restaurantid as string);
			console.log(json);
			if (json && json.item !== null) {
				setRestaurantName(json.item.name);
			}
		}
	};

	const callbackFunc = () => {
		if (Number(values.tablenum) === NaN) {
			console.log("Error: Invalid table number format");
			return;
		} else {
			myMenuService.setTableNum(Number(values.tablenum));
			history.push(`/customer/${values.restaurantid}/menu`);
		}
	};

	useEffect(() => {
		getRestaurantName();
	}, []);

	return (
		<Container>
			<Flex>
				<Grid>
					<h1>Welcome to {restaurantName}</h1>
					<div style={{ height: 50 }}></div>
					<Container>
						Would you like to sign up?{" "}
						<RegistrationDialog userType="Customer" callbackFunc={callbackFunc} />
					</Container>
					<Container>
						Or login <LoginDialog userType="Customer" callbackFunc={callbackFunc} />
					</Container>
					<Container>
						Or continue as a guest{" "}
						<Flex style={{ justifyContent: "center" }}>
							<Button appearance="primary" onClick={callbackFunc}>
								Guest
							</Button>
						</Flex>
					</Container>
				</Grid>
			</Flex>
		</Container>
	);
}
