import React, {useEffect, useState} from 'react';
import {AlignedHeader, Container, Divider, Grid, Header, Label} from "../utils/reusable-components";
import TextField from "../text-field";
import {Button, toaster} from "evergreen-ui";
import { useServices } from '../services/service-context';

interface UserDetails {
    name: string;
    email: string;
}

export default function PersonalSettings() {
    const { myMenuService } = useServices();
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword] = useState("");
    const [failure, setFailure] = useState("");

    const emailCheck = (email: string) => {
		let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

		return !re.test(email);
	};

    const updateDetails = async () => {
        if (oldPassword.length > 0) {
            // Updating user info with password
            if (oldPassword !== newPassword) {
                setFailure("Passwords need to match");
            } else if (oldPassword.length < 8) {
                setFailure("Password needs to be at least 8 characters");
            } else {
                const response = await myMenuService.updateCurrentUser(email, name, newPassword);
                if (response.hasOwnProperty("message")) {
                    setFailure(response["message"]);
                }
                toaster.success("User Details Updated")
                setFailure("");
            }
        } else if (emailCheck(email)) {
            setFailure("Invalid Email");
        } else if (name.length == 0) {
            setFailure("Invalid name");
        } else {
            // update user info without password
            const response = await myMenuService.updateCurrentUser(email, name, newPassword);
            if (response.hasOwnProperty("message")) {
                setFailure(response["message"]);
            } else {
                toaster.success("User Details Updated");
                setFailure("");
            }
        }
    };

    const getCurrentUser = async () => {
        const userData = await myMenuService.getCurrentUser();
        if (userData && userData.item) {
            const userDetails = userData.item.details;
            setName(userDetails.name);
            setEmail(userDetails.email);
        }
    };

    const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

    useEffect(() => {
        getCurrentUser();
    }, [])

    return (
        <Container>
            <AlignedHeader>
                Basic Information
            </AlignedHeader>
            <Grid>
                <TextField label="Name" input={name} updateInput={setName} type="text"/>
                <TextField label="Email" input={email} updateInput={setEmail} type="email"/>
            </Grid>
            <Divider/>
            <AlignedHeader>
                Password
            </AlignedHeader>
            <Grid>
                <TextField label="Old Password" input={oldPassword} updateInput={setOldPassword} type="password"/>
                <TextField label="New Password" input={newPassword} updateInput={setNewPassword} type="password"/>
            </Grid>
            <Divider/>
            {failure !== "" && invalidMsg(failure)}
            {/* {updated && <Label style={{color:"green"}}>User details updated</Label>} */}
            <Button onClick={updateDetails} intent="success">
                Update Details
            </Button>
        </Container>
    )
}