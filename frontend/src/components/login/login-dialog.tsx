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
	userType: User;
	callbackFunc?: any;
}

export default function LoginDialog(props: Props) {
	const { userType, callbackFunc } = props;
	const [isShown, setIsShown] = useState(false);
	const [failure, setFailure] = useState("");
	const [showReset, setShowReset] = useState(false);
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const { myMenuService } = useServices();
	const [auth, authActions] = useAuthenticationContext();
	const history = useHistory();

	const open = () => {
		setIsShown(true);
		setShowReset(false);
		setEmail("");
		setPassword("");
	};

	const reset = () => {
		setIsShown(false);
		setShowReset(true);
	};

	const resetView = <PasswordReset />;

	const emailCheck = (email: string) => {
		let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

		return !re.test(email);
	};

	const submit = async () => {
		if (emailCheck(email)) {
			setFailure("Enter a valid email.");
		} else {
			try {
				const jsonObj = await myMenuService.login(email, userType, password);
				if (jsonObj.hasOwnProperty("message")) {
					setFailure(jsonObj["message"]);
				} else {
					myMenuService.updateAuthToken(jsonObj["token"]);
					localStorage.setItem("token", jsonObj["token"]);
					localStorage.setItem("userType", userType);
					authActions.updateUser(userType);
					setFailure("");
					setIsShown(false);
					if (callbackFunc !== undefined) {
						callbackFunc();
						return;
					}
					switch (userType) {
						case "Manager":
							history.push("/admin/dashboard");
							break;
						case "Customer":
							history.push("/customer/dashboard");
							break;
					}
				}
			} catch (e) {
				console.log(e);
			}
		}
	};

	const invalidMsg = (msg: String) => {
		return <h4 style={{ color: "red" }}>{msg}</h4>;
	};

	const loginView = (
		<Dialog
			isShown={isShown}
			title={userType + " Login"}
			onCancel={() => setIsShown(false)}
			onConfirm={submit}
			confirmLabel="Confirm"
			onCloseComplete={() => setIsShown(false)}
		>
			<TextField label="Email" input={email} updateInput={setEmail} type="email" />
			<TextField label="Password" input={password} updateInput={setPassword} type="password" />
			<Flex>
				<Button appearance="primary" onClick={reset}>
					Reset Password
				</Button>
			</Flex>
			{failure !== "" && invalidMsg(failure)}
		</Dialog>
	);

	return (
		<Pane>
			<DialogContainer>
				{showReset && resetView}
				{isShown && loginView}
				<Button
					appearance="primary"
					intent={userType === "Customer" ? "success" : "danger"}
					onClick={open}
					fontSize={20}
					padding={20}
				>
					{userType}
				</Button>
			</DialogContainer>
		</Pane>
	);
}
