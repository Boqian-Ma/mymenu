import React, {useEffect, useState} from 'react';
import styled from "styled-components";
import {Badge} from "evergreen-ui";
import {Restaurant} from "../utils/data-interfaces";
import {AlignedHeader, Container} from "../utils/reusable-components";

const ContainerWithPadding = styled(Container)`
    padding: 0.5rem;
`;

const CardContainer = styled.div`
    background-color: rgb(244, 245, 247);
    color: rgb(52, 69, 99);
    border-radius: 5px;
    width: 300px;
    height: 200px;
    padding: 0px 8px;
`;

const InfoContainer = styled.div`
    display: flex;
    position: relative;
    justify-content: space-between;
`;

const CardImage = styled.div`
    border: 2px solid gray;
    width: 100%;
    height: 60%;
`;

const Image = styled.img`
    width: 100%;
    height: 100%;
`;

interface Props {
    restaurant: Restaurant;
    viewMenu(id: any): any;
}

export default function RestaurantCard(props: Props) {
    const { restaurant, viewMenu } = props;
    const [image, setImage] = useState("");

    useEffect(() => {
        setFilePath();
    }, []);

    const setFilePath = () => {
        setImage(process.env.PUBLIC_URL + `/assets/images/${restaurant.file}`);
    };

    return (
        <ContainerWithPadding onClick={() => viewMenu(restaurant.id)}>
            <CardContainer>
                <CardImage>
                    <Image src={image} />
                </CardImage>
                <AlignedHeader>
                    {restaurant.name}
                </AlignedHeader>
                <InfoContainer>
                    <Badge color="green">{restaurant.type}</Badge>
                    <Badge color="red">{restaurant.location}</Badge>
                </InfoContainer>
            </CardContainer>
        </ContainerWithPadding>
    );
}