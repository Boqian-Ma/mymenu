import React, {useEffect, useState} from 'react';
import {Container, Divider, Grid, Header} from "../../components/utils/reusable-components";
import styled from "styled-components";
import {useServices} from "../../components/services/service-context";
import {Order} from "../../components/utils/data-interfaces";
import OrderHistoryCard from "../../components/customer/order-history-card";

const GridContainer = styled.div`
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    position: relative;
    width: 100%;
    height: 100%;
`;

export default function OrderHistory() {
    const { myMenuService } = useServices();
    const [orders, setOrders] = useState<Order[] | undefined>();
    const [userId, setUserId] = useState<string>("");

    useEffect(() => {
        getUserId();
    }, []);

    useEffect(() => {
        getOrders();
    }, [userId]);

    const getUserId = async () => {
        const userData = await myMenuService.getCurrentUser();
        console.log(userData.item.id);
        if (userData && userData.item) {
            setUserId(userData.item.id);
        }
    };

    // TODO: not working yet
    const getOrders = async () => {
        const response = await myMenuService.getOrderHistory(userId);
        if (response && response.data) {
            setOrders(response.data);
        }
    };

    const loadOrders = () => {
        if (orders == undefined) {
            return;
        } else {
            return orders.map((o: Order) => <OrderHistoryCard order={o}/>);
        }
    };

    return (
        <Container>
            <Grid>
                <Header>Order History</Header>
            </Grid>
            <GridContainer>
                {loadOrders()}
            </GridContainer>
        </Container>
    )
}