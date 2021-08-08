import React from 'react';
import {Container} from "../utils/reusable-components";
import {Button, Pane} from "evergreen-ui";
import {SettingOption, User} from "../utils/data-types";

interface Props {
    userType: User | undefined;
    changeView(option: SettingOption): any;
}

export default function SettingSidebar(props: Props) {
    const { userType, changeView } = props;

    const userView = () => {
        return (
            <>
                <Pane padding={5} display="flex" justifyContent="left">
                    User Settings
                </Pane>
                <Pane display="flex" flexDirection="column" textAlign="right">
                    <Button onClick={() => changeView('personal')} appearance="minimal" display="flex" justifyContent="right">
                        Personal Information
                    </Button>
                    <Button onClick={() => changeView('health')} appearance="minimal" display="flex" justifyContent="right">
                        Health Information
                    </Button>
                </Pane>
            </>
        )
    };

    const adminView = () => {
        return (
            <>
                <Pane padding={5} display="flex" justifyContent="left">
                    Admin Settings
                </Pane>
                <Pane display="flex" flexDirection="column" textAlign="right">
                    <Button onClick={() => changeView('restaurant')} appearance="minimal" display="flex" justifyContent="right">
                        Restaurant Information
                    </Button>
                    <Button onClick={() => changeView('reports')} appearance="minimal" display="flex" justifyContent="right">
                        Reports
                    </Button>
                    <Button onClick={() => changeView('branding')} appearance="minimal" display="flex" justifyContent="right">
                        Branding
                    </Button>
                </Pane>
            </>
        )
    };

    return (
        <Pane paddingTop={50}>
            <Pane flexDirection="column" elevation={2} float="left" display="flex" padding={10}>
                {userView()}
                <Pane padding={10}></Pane>
                {userType == "Manager" && adminView()}
            </Pane>
        </Pane>
    )
}