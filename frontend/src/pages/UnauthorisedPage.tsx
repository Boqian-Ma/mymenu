import React from "react";
import styled from "styled-components";
import { ButtonStack, Flex, Grid, Image } from "../components/utils/reusable-components";
import RegistrationDialog from "../components/registration/registration-dialog";
import LoginDialog from "../components/login/login-dialog";
import { Button, Dialog, Pane, Alert } from "evergreen-ui";
import { useServices } from "../components/services/service-context";
import { ReactComponent as Logo } from "../logo/myMenuLogo.svg";
import { useHistory } from "react-router-dom";

const GraphicsContainer = styled.div`
	height: 500px;
`;

const TextContainer = styled.div`
	padding: 1rem;
`;

export default function UnauthorisedPage() {
	const { myMenuService } = useServices();
	const history = useHistory();

	const redirectToHome = () => {
		history.replace("/");
	};

	return (
		<Pane>
			<Flex style={{ justifyContent: "center", paddingTop: 125 }}>
				<Pane elevation={1} background="tint1" padding={50}>
					<Pane fontSize={30}>
						<Logo width={300} />
						<TextContainer>Unauthorised Access</TextContainer>
					</Pane>
					<Pane paddingTop={30}>
						<Alert intent="danger">
							Please login with an account that can access this page.
							<br />
							<br />
							<Button onClick={redirectToHome}>Home</Button>
						</Alert>
					</Pane>
				</Pane>
			</Flex>
		</Pane>
	);
}
