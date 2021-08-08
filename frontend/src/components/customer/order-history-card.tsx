import React, {useEffect, useState} from 'react';
import {AlignedHeader, Container, Divider, Grid, Label} from "../utils/reusable-components";
import {Button, Dialog} from "evergreen-ui";
import styled from "styled-components";
import {Order, Restaurant} from "../utils/data-interfaces";
import {useServices} from "../services/service-context";
import KitchenOrderItems from "../kitchen/kitchen-order-items";

const CardContainer = styled.div`
    display: grid;
    background-color: rgb(244, 245, 247);
    color: rgb(52, 69, 99);
    border-radius: 5px;
    width: 250px;
    padding: 0px 8px;
`;

const CardImage = styled.div`
    border: 2px solid gray;
    width: 100%;
    height: 80%;
`;

const Image = styled.img`
    width: 100%;
    height: 100%;
`;

interface Props {
    order: Order;
}

export default function OrderHistoryCard(props: Props) {
    const [restaurant, setRestaurant] = useState<Restaurant>();
    const [image, setImage] = useState("");
    const [isShown, setIsShown] = useState(false);
    const { myMenuService } = useServices();
    const { order } = props;

    useEffect(() => {
        getRestaurant();
    }, []);

    useEffect(() => {
        setFilePath();
    }, [restaurant]);

    const setFilePath = () => {
        if (restaurant) {
            setImage(process.env.PUBLIC_URL + `/assets/images/${restaurant.file}`);
        }
    };

    const getRestaurant = async () => {
        const response = await myMenuService.getRestaurantCust(order.restaurant_id);
        if (response && response.item) {
            setRestaurant(response.item);
        }
    };

    const formatDate = (date_str: string) => {
        const dateObj = new Date(date_str);
        return dateObj.toLocaleString('en-AU')
    };

    const loadOrder = () => {
        if (restaurant == undefined) {
            return;
        } else {
            return (
                <Container>
                    <Dialog
                        isShown={isShown}
                        title="Order"
                        onCloseComplete={() => setIsShown(false)}
                        hasFooter={false}
                    >
                        <KitchenOrderItems items={order.items} />
                        <Divider />
                    </Dialog>

                    <CardContainer>
                        <AlignedHeader>
                            {restaurant.name}
                        </AlignedHeader>
                        <CardImage>
                            <Image src={image} />
                        </CardImage>
                        <Grid>
                            <Label>{formatDate(order.created_at)}</Label>
                            <Label>Table #{order.table_num}</Label>
                            <Label>Status: {order.status}</Label>
                            <Label>Total Cost: ${order.total_cost}</Label>
                            <Divider />
                            <Button onClick={() => setIsShown(true)} intent="success">Show Items</Button>
                            <Divider/>
                        </Grid>
                    </CardContainer>
                    <Divider/>
                </Container>
            )
        }
    };

    return (
        <Container>
            {loadOrder()}
        </Container>
    )
}