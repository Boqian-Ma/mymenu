import React, {useState} from 'react';
import {Button} from "evergreen-ui";
import {CategoryOption} from "../utils/data-types";

interface Props {
    children: string;
    selected: string;
    option: string;
    changeOption(option: string): void;
}

export default function CategoryButton(props: Props) {
    const { selected, option, changeOption } = props;

    return(
        <Button appearance={selected == option ? 'primary' : 'default'} intent={selected == option ? 'danger' : 'none'} onClick={() => changeOption(option)}>{props.children}</Button>
    )
}