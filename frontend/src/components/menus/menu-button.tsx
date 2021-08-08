import React, {useState} from 'react';
import {Button} from "evergreen-ui";
import {MenuOption} from "../utils/data-types";

interface Props {
    children: string;
    selected: MenuOption;
    option: MenuOption;
    changeOption(option: MenuOption): void;
}

export default function MenuButton(props: Props) {
    const { selected, option, changeOption } = props;

    return(
        <Button appearance={selected == option ? 'primary' : 'default'} intent={selected == option ? 'danger' : 'none'} onClick={() => changeOption(option)}>{props.children}</Button>
    )
}