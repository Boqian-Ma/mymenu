import React, { useState } from "react";
import { Button, Dialog, Pane, TextInput } from "evergreen-ui";
import { User } from "../utils/data-types";
import styled from "styled-components";
import { useServices } from "../services/service-context";
import { useHistory } from "react-router-dom";
import TextField from "../text-field";
import PasswordReset from "./login-password-reset";
import { Flex } from "../utils/reusable-components";
import { useAuthenticationContext } from "../services/authentication-context";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

interface Props {
	historyRoute: string;
}

export default function GuestLoginButton(props: Props) {
	const userType = "Customer";
	const { historyRoute } = props;
	const [failure, setFailure] = useState("");
	const { myMenuService } = useServices();
	const [auth, authActions] = useAuthenticationContext();
	const history = useHistory();

	const login = async () => {
		try {
			const jsonObj = await myMenuService.loginGuest();
			if (jsonObj.hasOwnProperty("message")) {
				setFailure(jsonObj["message"]);
			} else {
				myMenuService.updateAuthToken(jsonObj["token"]);
				localStorage.setItem("token", jsonObj["token"]);
				localStorage.setItem("userType", userType);
				authActions.updateUser(userType);
				setFailure("");
				history.push(historyRoute);
			}
		} catch (e) {
			console.log(e);
		}
	};

	const invalidMsg = (msg: String) => {
		return <h4 style={{ color: "red" }}>{msg}</h4>;
	};

	return (
		<Pane>
			<DialogContainer>
				<Button appearance="primary" onClick={login}>
					Guest
				</Button>
				{failure != "" && invalidMsg(failure)}
			</DialogContainer>
		</Pane>
	);
}
