import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Group, Pane } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header, KitchenGrid, Label } from "../../components/utils/reusable-components";
import { TableStatus } from "../../components/utils/data-types";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";

import { useLocation } from "react-router";
import SelectField from "../../components/select-field";
import TablesList from "../../components/tables/tables-list";
import KitchenOrderItems from "../../components/kitchen/kitchen-order-items";
import {Order} from "../../components/utils/data-interfaces";
import { Restaurant, Table } from "../../components/utils/data-interfaces"

const ButtonContainer = styled.div`
	margin: 1rem;
	padding: 1rem;
	align-items: center;
	justify-content: center;
`;

export default function ManageKitchen() {
    const { myMenuService } = useServices();
    const history = useHistory();
    const id = myMenuService.getCurrRestaurant();
    const [restaurant, setRestaurant] = useState<Restaurant>({id:'', location:'',name:'',type:'',email:'',phone:'',website:'',businessHours:'',file:'',cuisine:''});
    const [tables, setTables] = useState<Table[]>([]);
    const [orders, setOrders] = useState<Order[]>([]);
    const [filteredOrders, setFilteredOrders] = useState<Order[]>([]);
    const orderStatuses = ["ordered", "served"]
    const [orderStatus, setOrderStatus] = useState("");
    
    const getRestaurantDetails = async () => {
        const restaurant_response = await myMenuService.getRestaurant();
        if (restaurant_response && restaurant_response.item) {
            setRestaurant(restaurant_response.item);
        }

        const tablesResponse = await myMenuService.getTables();
        if (tablesResponse && tablesResponse.data) {
            setTables(tablesResponse.data);
        }

        const ordersResponse = await myMenuService.getOrders("", true);
        if (ordersResponse && ordersResponse.data) {
            setOrders(ordersResponse.data);
            if (orderStatus === "") {
                setFilteredOrders(ordersResponse.data);
            } else {
                setFilteredOrders(ordersResponse.data.filter((order: Order) => order.status === orderStatus))
            }
        }
    }

    const toggleOrderStatus = (option: string) => {
        if (orderStatus === option) {
            setOrderStatus("");
            setFilteredOrders(orders);
        } else {
            setOrderStatus(option);
            setFilteredOrders(orders.filter((order) => order.status === option))
        }
    }

    const serveOrder = async (order_id: string) => {
        await myMenuService.serveOrder(order_id);
        getRestaurantDetails();
    }


	useEffect(() => {
		getRestaurantDetails();
	}, []);



    return (
        <Container>
            <Header>Kitchen - {restaurant.name}</Header>
            <Group>
                {orderStatuses.map((option) => (
                    <Button size="large" onClick={() => toggleOrderStatus(option)} isActive={orderStatus === option}>{option.charAt(0).toUpperCase() + option.slice(1)}</Button>
                ))}
            </Group>
            <Divider></Divider>
            {filteredOrders.length === 0 && <Label>NO ORDERS!!</Label>}
            <KitchenGrid>
                {filteredOrders.map((order) => (
                    <Pane border marginBottom={30} paddingBottom={10} paddingLeft={20} paddingRight={20}>
                        <Label>Table #{order.table_num}</Label>
                        <Label>Status: {order.status}</Label>
                        <Divider />
                        <Label>Order Items</Label>
                        <KitchenOrderItems items={order.items} />
                        <Button appearance="primary" intent="success" disabled={order.status === "served"} onClick={() => serveOrder(order.id)}>Serve Order</Button>
                    </Pane>
                ))}
            </KitchenGrid>
        </Container>
    );
}