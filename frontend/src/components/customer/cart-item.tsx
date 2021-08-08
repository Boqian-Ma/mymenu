import React, {useEffect, useState} from 'react';
import {Bold, Flex, Grid} from "../utils/reusable-components";
import {IconButton, MinusIcon, PlusIcon} from "evergreen-ui";
import styled from "styled-components";
import {MenuItem} from "../utils/data-interfaces";
import {useCarts} from "../services/cart-context";

const ImageContainer = styled.div`
    border: 2px solid gray;
    width: 100px;
    height: 100px;
`;

const Title = styled.div`
    display: flex;
    margin-left: 1rem;
`;

interface Props {
    index: number;
    item: MenuItem;
}

export default function CartItem(props: Props) {
    const [carts, cartActions] = useCarts();
    const { index, item } = props;
    const [image, setImage] = useState("");

    useEffect(() => {
        setFilePath();
    }, []);

    const removeItem = () => {
        cartActions.removeItemFromCart(item, index);
    };

    const setFilePath = () => {
        setImage(process.env.PUBLIC_URL + `/assets/images/${item.file}`);
    };

    return (
        <Flex>
            <ImageContainer><img src={image} width={100} height={100} /></ImageContainer>
            <Grid>
                <Title>
                    <Bold>{item.name}</Bold>
                </Title>
            </Grid>
            <Grid>
                <Title>
                    <p>${item.price}</p>
                </Title>
            </Grid>
            <Grid>
                <IconButton icon={MinusIcon} intent="success" onClick={removeItem} />
            </Grid>
        </Flex>
    )
}