import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Table } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header } from "../../components/utils/reusable-components";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";
import {Restaurant} from "../../components/utils/data-interfaces";

const ButtonContainer = styled.div`
	margin: 1rem;
	padding: 1rem;
	align-items: center;
	justify-content: center;
`;

export default function RestaurantDashboard() {
	const { myMenuService } = useServices();
	const [isRestaurant, toggleRestaurant] = useState<boolean>(myMenuService.hasRestaurant);
	const history = useHistory();
	const [data, setData] = useState<Restaurant[]>([]);
	const [filtered, setFiltered] = useState<Restaurant[]>([]);
	
	const getRestaurants = async () => {
		const json = await myMenuService.getRestaurants("true");
		if (typeof json !== "undefined" && json.data != null) {
			setData(json.data);
			setFiltered(json.data);
		}
	};

	const filterSubstr = (substr: string) => {
		const filteredData: Restaurant[] = data.filter((item: Restaurant) => item.name.startsWith(substr));
		setFiltered(filteredData);
	};

	const addRestaurant = () => {
		history.push("/admin/restaurant/add");
	};

	const addMenu = () => {
		history.push("/admin/restaurant/menu");
	};

	const redirectToRestaurant = (resID: string) => {
		myMenuService.setCurrRestaurant(resID);
		history.push("/admin/restaurant");
	}

	useEffect(() => {
		myMenuService.setCurrRestaurant('');
		getRestaurants();
	}, []);

	return (
		<Container>
			<Flex>
				<Grid>
					<Header>Welcome back</Header>
				</Grid>
				<Grid>
					<ButtonContainer>
						<Button onClick={addRestaurant} width={200}>
							Add New Restaurant
						</Button>
					</ButtonContainer>
				</Grid>
			</Flex>
			<Table>
				<Table.Head>
					<Table.SearchHeaderCell placeholder="Search by restaurant name..." onChange={filterSubstr} />
					<Table.TextHeaderCell>Location</Table.TextHeaderCell>
					<Table.TextHeaderCell>Type</Table.TextHeaderCell>
				</Table.Head>
				<Table.VirtualBody height={300}>
					{filtered.map((r: Restaurant) => (
						<Table.Row isSelectable onSelect={() => redirectToRestaurant(r.id)}>
							<Table.TextCell>{r.name}</Table.TextCell>
							<Table.TextCell>{r.location}</Table.TextCell>
							<Table.TextCell>{r.type}</Table.TextCell>
						</Table.Row>
					))}
				</Table.VirtualBody>
			</Table>
		</Container>
	);
}
