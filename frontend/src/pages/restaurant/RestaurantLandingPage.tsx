import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Pane, Table } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header, Label } from "../../components/utils/reusable-components";
import { User } from "../../components/utils/data-types";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";

import { useLocation } from "react-router";
import SelectField from "../../components/select-field";
import { Restaurant, MenuItem, Category } from "../../components/utils/data-interfaces";
import RestaurantReport from "../../components/report/restaurant-report";
import OrderHistoryRestaurant from "../../components/report/order-history";

const ButtonContainer = styled.div`
	margin: 1rem;
	padding: 1rem;
	align-items: center;
	justify-content: center;
`;

interface ReportItem {
    most_ordered_item: MenuItem;
    most_ordered_item_quantity: number;
    most_ordered_category: Category;
    total_revenue: number;
}

export default function RestaurantLandingPage() {
    const { myMenuService } = useServices();
    const history = useHistory();
    const id = myMenuService.getCurrRestaurant();
    const [data, setData] = useState<Restaurant>({id:'', location:'',name:'',type:'',email:'',phone:'',website:'',businessHours:'',file:'',cuisine:''});
    
    const getRestaurantDetails = async () => {
        const json = await myMenuService.getRestaurant();
        if (typeof json !== "undefined" && json.item != null) {
            setData(json.item);
        } 
    }

    const redirectToMenu = () => {
		history.push("/admin/restaurant/menu");
	}

    const redirectToTable = () => {
		history.push("/admin/restaurant/tables");
	}

    const redirectToKitchen = () => {
        history.push("/admin/restaurant/kitchen");
    }

	useEffect(() => {
		getRestaurantDetails();
	}, []);

    return (
        <Container>
            <Flex>
                <Grid>
                    <Header>{data.name}</Header>
                </Grid>
                <Grid>
                    <ButtonContainer>
						<Button width={200} onClick={redirectToTable}>
							Manage Tables
						</Button>
						<Divider></Divider>
						<Button width={200} onClick={redirectToKitchen}>
							View Kitchen
						</Button>
                        <Divider></Divider>
						<Button width={200} onClick={redirectToMenu}>
							View Menu
						</Button>
					</ButtonContainer>
                </Grid>
            </Flex>
            <Flex>
                <RestaurantReport />
                <div style={{width:20}} />
                <OrderHistoryRestaurant />
            </Flex>
        </Container>
    );
}