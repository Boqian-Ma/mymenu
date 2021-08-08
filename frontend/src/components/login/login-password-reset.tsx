import React, { useState } from "react";
import { Button, Dialog, Pane, TextInput } from "evergreen-ui";
import { User } from "../utils/data-types";
import styled from "styled-components";
import { useServices } from "../services/service-context";
import { useHistory } from "react-router-dom";
import TextField from "../text-field";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

export default function PasswordReset() {
	const [isShown, setIsShown] = useState(true);
	const [failure, setFailure] = useState("");
	const [email, setEmail] = useState("");
	const [newPassword, setNewPassword] = useState("");
	const [newPasswordConfirm, setNewPasswordConfirm] = useState("");
	const history = useHistory();
	const { myMenuService } = useServices();

	const submit = async () => {
		if (newPassword.length < 8) {
			setFailure("Password must be at least 8 characters long");
		} else if (newPassword !== newPasswordConfirm) {
			setFailure("New passwords do not match");
		} else {
			const jsonObj = await myMenuService.resetPassword(email, newPassword);
			if (jsonObj.hasOwnProperty("message")) {
				setFailure(jsonObj["message"]);
			} else {
				setFailure("");
				setIsShown(false);
			}
		}
	};

	const invalidMsg = (msg: String) => {
		return <h4 style={{ color: "red" }}>{msg}</h4>;
	};
	return (
		<Dialog
			isShown={isShown}
			title="Reset Password"
			onCancel={() => setIsShown(false)}
			onConfirm={submit}
			confirmLabel="Confirm"
			onCloseComplete={() => setIsShown(false)}
		>
			<TextField label="Email" input={email} updateInput={setEmail} type="email" />
			<TextField label="New Password" input={newPassword} updateInput={setNewPassword} type="password" />
			<TextField
				label="Re-enter New Password"
				input={newPasswordConfirm}
				updateInput={setNewPasswordConfirm}
				type="password"
			/>
			{failure !== "" && invalidMsg(failure)}
		</Dialog>
	);
}
