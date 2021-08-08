import React from "react";
import NavigationBar from "../components/nav-bar";
import {storiesOf} from "@storybook/react";

storiesOf('Globals', module)
    .add('Manager Navigation bar', () => {
        return (
            <NavigationBar type="Manager"/>
        )
    })
    .add('Customer Navigation bar', () => {
        return (
            <NavigationBar type="Customer"/>
        )
    });