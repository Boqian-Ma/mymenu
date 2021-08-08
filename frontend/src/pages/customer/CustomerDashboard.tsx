import React, {useEffect, useState} from "react";
import styled from "styled-components";
import {Container, Grid, Header} from "../../components/utils/reusable-components";
import {useHistory} from "react-router";
import {useServices} from "../../components/services/service-context";
import RestaurantCard from "../../components/customer/restaurant-card";
import { Restaurant } from "../../components/utils/data-interfaces";

const GridContainer = styled.div`
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    position: relative;
    width: 100%;
    height: 100%;
`;

export default function CustomerDashboard() {
    const {myMenuService} = useServices();
    const history = useHistory();
    const [restaurants, setRestaurants] = useState<Restaurant[]>([]);
    const [recommended, setRecommended] = useState<Restaurant[]>([]);

    const viewMenu = (id: string) => {
        history.push(`/customer/${id}/menu`);
    };

    const getRestaurants = async () => {
        const restaurants = await myMenuService.getRestaurants("");
        setRestaurants(restaurants.data)
    };

    const getRecommended = async () => {
        const restaurants = await myMenuService.getRecommendedRestaurants();
        setRecommended(restaurants.data)
    };

    const loadRestaurantCards = (restaurantList: Restaurant[]) => {
        return restaurantList.map((r: Restaurant) => <RestaurantCard restaurant={r} viewMenu={viewMenu} />)
    };

    useEffect(() => {
        getRestaurants();
        getRecommended();
    }, []);

    return (
        <Container>
            <Grid>
                <Header>Welcome Back</Header>
            </Grid>
            <Grid>
                <Header>Recommended Restaurants</Header>
            </Grid>
            <GridContainer>
                {loadRestaurantCards(recommended)}
            </GridContainer>
            <Grid>
                <Header>All Restaurants</Header>
            </Grid>
            <GridContainer>
                {loadRestaurantCards(restaurants)}
            </GridContainer>
        </Container>
    );
}
