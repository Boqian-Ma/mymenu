import { Pane } from "evergreen-ui";
import { useEffect, useState } from "react";
import {Order} from "../../components/utils/data-interfaces";
import { useServices } from "../services/service-context";
import { Bold, Flex, Grid, Label } from "../utils/reusable-components";
import OrderHistoryItems from "./order-history-items";

export default function OrderHistoryRestaurant() {
    const { myMenuService } = useServices();
    const [orders, setOrders] = useState<Order[]>([]);

    const getOrders = async () => {
        const response = await myMenuService.getOrderHistoryRestaurant();
        if (response && response.data) {
            setOrders(response.data);
        }
    };

    const formatDate = (date_str: string) => {
        const dateObj = new Date(date_str);
        return dateObj.toLocaleString('en-AU')
    }

    useEffect(() => {
        getOrders();
    }, []);

    return (
        <Pane width={500} border>
            <Label>Order History</Label>
            {orders.map((order) => (<Grid>
                                        <Flex>
                                        <Bold>Table #{order.table_num} - {order.status} - ${order.total_cost}</Bold>
                                        <Bold style={{marginLeft: "auto"}}>{formatDate(order.created_at)}</Bold>
                                        </Flex>
                                        <OrderHistoryItems items={order.items} />
                                    </Grid>))}
        </Pane>
    )
}