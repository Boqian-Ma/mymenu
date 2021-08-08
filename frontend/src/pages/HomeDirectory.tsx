import React from "react";
import styled from "styled-components";
import { ButtonStack, Flex, Grid, Image } from "../components/utils/reusable-components";
import RegistrationDialog from "../components/registration/registration-dialog";
import LoginDialog from "../components/login/login-dialog";
import { Button, Dialog, Pane } from "evergreen-ui";

const GraphicsContainer = styled.div`
	position: relative;
	height: 100%;
	height: calc(100vh - 40px);
	justify-content: center;
	text-align: center;
`;

const TextStyle = {
	color: "white",
	fontSize: "24px",
	paddingTop: "10px",
	filter: "drop-shadow(10px 0px 10px #000000)",
};

const HomeSplash = styled(Image)`
	width: 100%;
	min-height: 500px;
	height: 105%;
	object-fit: cover;
	filter: saturate(300%) brightness(30%) hue-rotate(180deg);
`;

const BackgroundFloat = styled.div`
	position: absolute;
	display: block;
	top: 0px;
	height: 90%;
	width: 100%;
`;

const FloatingLogoImage = styled(Image)`
	margin-left: auto;
	margin-right: auto;
	width: 30%;
	max-width: 400px;
	filter: drop-shadow(5px 0px 4px #000000);
`;

const UserFlowDiv = styled.div`
	left: 50px;
	top: 50%;
	padding-top: 30px;
`;

const CentreColumns = styled(Grid)`
	justify-content: center;
`;

export default function HomeDirectory() {
	return (
		<GraphicsContainer>
			<HomeSplash src="assets/homesplash.jpg" />
			<BackgroundFloat>
				<Flex style={{ height: "100%", justifyContent: "center" }}>
					<CentreColumns>
						<FloatingLogoImage src="assets/myMenuWhite.svg" />
						<h2 style={TextStyle}> the simplest way to order your food</h2>
					</CentreColumns>
					<CentreColumns>
						<UserFlowDiv>
							<h3 style={TextStyle}>Sign up for:</h3>
							<RegistrationDialog userType="Manager" />
							<RegistrationDialog userType="Customer" />
						</UserFlowDiv>
						<UserFlowDiv>
							<h3 style={TextStyle}>Login:</h3>
							<LoginDialog userType="Manager" />
							<LoginDialog userType="Customer" />
						</UserFlowDiv>
					</CentreColumns>
				</Flex>
			</BackgroundFloat>
		</GraphicsContainer>
	);
}
