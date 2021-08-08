import React, { useEffect, useState } from "react";
import { Container, Flex, Header } from "../../components/utils/reusable-components";
import styled from "styled-components";
import SelectField from "../../components/select-field";
import { MenuOption } from "../../components/utils/data-types";
import MenuButton from "../../components/menus/menu-button";
import AddItemDisplay from "../../components/menus/add-item-display";
import MenuTable from "../../components/menus/menu-table-view";
import { useLocation } from "react-router-dom";
import { useServices } from "../../components/services/service-context";
import CreateCategoryDisplay from "../../components/menus/create-category-display";
import { Restaurant } from "../../components/utils/data-interfaces"

export default function AddMenuDetails() {
	const { myMenuService } = useServices();
	const [selected, setSelected] = useState<MenuOption>("table");
	const id = myMenuService.getCurrRestaurant();
	const [data, setData] = useState<Restaurant>({id:"",location:"",name:"",type:"",email:"",phone:"",website:"",businessHours:"",file:"",cuisine:""});
	const categories = ["Breakfast", "Lunch", "Dinner", "Desserts", "Drinks"];

	const getRestaurantDetails = async () => {
		if (id !== null) {
			const json = await myMenuService.getRestaurant();
			if (typeof json !== "undefined" && json.item != null) {
				setData(json.item);
			}
		}
	};

	useEffect(() => {
		getRestaurantDetails();
	}, []);

	const changeOption = (option: MenuOption) => {
		setSelected(option);
	};

	return (
		<Container>
			<Header>Menu {data.name}</Header>
			<Flex>
				<MenuButton selected={selected} option="table" changeOption={() => changeOption("table")}>
					Menu View
				</MenuButton>
				<MenuButton selected={selected} option="category" changeOption={() => changeOption("category")}>
					Create Category
				</MenuButton>
				<MenuButton selected={selected} option="add-item" changeOption={() => changeOption("add-item")}>
					Add Item
				</MenuButton>
			</Flex>
			{selected == "table" && <MenuTable />}
			{selected == "category" && <CreateCategoryDisplay />}
			{selected == "add-item" && <AddItemDisplay />}
		</Container>
	);
}
