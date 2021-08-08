import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Pane, Table } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header, Label } from "../../components/utils/reusable-components";
import { TableStatus } from "../../components/utils/data-types";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";

import { useLocation } from "react-router";
import SelectField from "../../components/select-field";
import TablesList from "../../components/tables/tables-list";

const ButtonContainer = styled.div`
	margin: 1rem;
	padding: 1rem;
	align-items: center;
	justify-content: center;
`;

interface Restaurant {
	id: string;
	location: string;
	name: string;
	type: string;
}

export default function ManageTables() {
	const { myMenuService } = useServices();
	const history = useHistory();
	const [data, setData] = useState<Restaurant>({ id: "", location: "", name: "", type: "" });
	const [qrcode, setQrCode] = useState<string>("");
	const QRCode = require("qrcode.react");

	const getRestaurantDetails = async () => {
		const json = await myMenuService.getRestaurant();
		if (typeof json !== "undefined" && json.item != null) {
			setData(json.item);
		}
	};

	useEffect(() => {
		getRestaurantDetails();
	}, []);

	return (
		<Container>
			<Header>Manage Tables - {data.name}</Header>
			<Divider></Divider>
			<Flex>
				<TablesList />
			</Flex>
		</Container>
	);
}
