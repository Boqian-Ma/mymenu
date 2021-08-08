import React, { useState } from "react";
import { Button } from "evergreen-ui";
import { NavigationOption } from "./utils/data-types";

interface Props {
	children: string;
	selected: NavigationOption;
	option: NavigationOption;
	disabled: boolean;
	redirect(): void;
}

export default function NavigationButton(props: Props) {
	const { selected, option, disabled, redirect } = props;

	return (
		<Button
			color={selected == option ? "white" : "slategrey"}
			fontSize={"20"}
			width={150}
			appearance={selected == option ? "primary" : "minimal"}
			disabled={disabled}
			onClick={() => redirect()}
		>
			{props.children}
		</Button>
	);
}
