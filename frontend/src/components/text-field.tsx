import React from "react";
import styled from "styled-components";
import {Grid, Label} from "./utils/reusable-components";
import { TextInput } from "evergreen-ui";

interface Props {
	label: string;
	input: string;
	updateInput(e: any): any;
	type: string;
}

export default function TextField(props: Props) {
	const { label, input, updateInput, type } = props;

	return (
		<Grid>
			<Label>{label}</Label>
			<TextInput
				onChange={(e: any) => updateInput(e.target.value)}
				value={input}
				type={type}
			/>
		</Grid>
	);
}
