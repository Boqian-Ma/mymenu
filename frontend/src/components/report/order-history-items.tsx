import React, {useState} from 'react';
import {Divider, Flex, Grid, Label} from "../utils/reusable-components";
import { OrderItem } from "../utils/data-interfaces"

interface Props {
    items: OrderItem[];
}

export default function OrderHistoryItems(props: Props) {
    const { items } = props;

    return (
        <div>
            {items.map((item) => (
                <>
                    <Label style={{marginLeft:20}}>{item.quantity} x {item.item_name}</Label>
                </>
            ))}
        </div>
    );
}