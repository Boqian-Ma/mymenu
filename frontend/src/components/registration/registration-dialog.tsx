import React, { useState } from "react";
import { Button, Dialog, Pane, TextInput } from "evergreen-ui";
import { User } from "../utils/data-types";
import styled from "styled-components";
import { useServices } from "../services/service-context";
import { useHistory } from "react-router-dom";
import TextField from "../text-field";
import { useAuthenticationContext } from "../services/authentication-context";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

interface Props {
	userType: User;
	callbackFunc?: any;
}

export default function RegistrationDialog(props: Props) {
	const { userType, callbackFunc } = props;
	const [isShown, setIsShown] = useState(false);
	const [name, setName] = useState("");
	const [password, setPassword] = useState("");
	const [passwordConfirm, setPasswordConfirm] = useState("");
	const [email, setEmail] = useState("");
	const [failure, setFailure] = useState("");
	const { myMenuService } = useServices();
	const [auth, authActions] = useAuthenticationContext();
	const history = useHistory();

	const open = () => {
		setIsShown(true);
		setName("");
		setPassword("");
		setEmail("");
		setPasswordConfirm("");
	};

	const emailCheck = (email: string) => {
		let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

		return !re.test(email);
	};

	const submit = async () => {
		if (password.length < 8) {
			setFailure("Password needs to be at least 8 characters long.");
		} else if (password !== passwordConfirm) {
			setFailure("Passwords need to match.");
		} else if (emailCheck(email)) {
			setFailure("Enter a valid email.");
		} else {
			try {
				const jsonObj = await myMenuService.register(email, userType, name, password);
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
				setFailure("Server Error");
			}
		}
	};

	const invalidMsg = (msg: String) => {
		return <h4 style={{ color: "red" }}>{msg}</h4>;
	};

	return (
		<Pane>
			<DialogContainer>
				<Dialog
					isShown={isShown}
					title="Registration"
					onCancel={() => setIsShown(false)}
					onConfirm={submit}
					confirmLabel="Confirm"
					onCloseComplete={() => setIsShown(false)}
				>
					<TextField label="Name" input={name} updateInput={setName} type="text" />
					<TextField label="Password" type="password" input={password} updateInput={setPassword} />
					<TextField
						label="Confirm Password"
						input={passwordConfirm}
						updateInput={setPasswordConfirm}
						type="password"
					/>
					<TextField label="Email" input={email} updateInput={setEmail} type="email" />
					{failure !== "" && invalidMsg(failure)}
				</Dialog>

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
