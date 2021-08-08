import React, {useEffect, useState} from 'react';
import {Bold, Container, Flex, Grid} from "../utils/reusable-components";
import styled from "styled-components";
import {MenuItem} from "../utils/data-interfaces";
import {IconButton, PlusIcon, StarIcon, Tooltip, WarningSignIcon} from "evergreen-ui";
import {useCarts} from "../services/cart-context";

const ImageContainer = styled.div`
    border: 2px solid gray;
    width: 100px;
    height: 100px;
`;

const DescriptionContainer = styled.div`
    border: 2px solid gray;
    width: 300px;
    height: 50px;
    margin-left: 0.5rem;
`;

const DescriptionText = styled.p`
    font-size: 12px;
`;

const Title = styled.div`
    display: flex;
    min-width: 300px;
`;

const Name = styled(Bold)`
    margin-left: 1rem;
`;

const Price = styled.p`
    margin-left: 2rem;
`;

const Special = styled(StarIcon)`
    margin-top: 1rem;
    margin-left: auto;
    margin-right: 2rem;
    color: gold;
`;

const AddButton = styled(IconButton)`
    margin-top: 0.5rem;
`;

interface Props {
    item: MenuItem;
    canOrder: boolean;
}

// TODO: fix styling and structure GUH
export default function OrderMenuItem(props: Props) {
    const { item, canOrder } = props;
    const [carts, cartActions] = useCarts();
    const [image, setImage] = useState("");

    useEffect(() => {
        setFilePath();
    }, []);

    const addToCart = () => {
        cartActions.addItemToCart(item);
    };

    const setFilePath = () => {
        setImage(process.env.PUBLIC_URL + `/assets/images/${item.file}`);
    };

    const warning = () => {
        const msg = `Contains: ${item.allergy}`;
        return (
            <Tooltip content={msg}>
                <WarningSignIcon marginTop={15} marginRight={30} color="warning" />
            </Tooltip>
        )
    }

    return (
        <Container>
            <Flex>
                <ImageContainer><img src={image} width={100} height={100} ></img></ImageContainer>
                <Grid>
                    <Title>
                        <Name>{item.name}</Name>
                        <Price>${item.price}</Price>
                        {item.is_special && <Special/>}
                    </Title>
                    <DescriptionContainer>
                        <DescriptionText>{item.description}</DescriptionText>
                    </DescriptionContainer>
                </Grid>
                <Grid>
                    {item.allergy !== "" && warning()}
                </Grid>
                <Grid>
                    {canOrder && <AddButton icon={PlusIcon} intent="success" onClick={addToCart} />}
                </Grid>
            </Flex>
        </Container>
    )
}