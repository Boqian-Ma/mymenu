import { Button, PlusIcon, Select, Switch, Table, TableCell, TableRow, toaster, TrashIcon, Dialog, Pane } from "evergreen-ui";
import { useEffect, useState } from "react";
import styled from "styled-components";
import SelectField from "../select-field";
import { useServices } from "../services/service-context";
import TextField from "../text-field";
import { TableStatus } from "../utils/data-types";
import { Divider, Flex, Grid, Label } from "../utils/reusable-components";
import {MenuItem, Order, Table as TableInterface} from "../utils/data-interfaces";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

interface Props {
    tableNumber: number;
    callback(): any;
}

export default function UpdateTable(props: Props) {
    const { tableNumber, callback } = props;
    const { myMenuService } = useServices();
    const [tables, setTables] = useState<TableInterface[]>([]);
    const [table, setTable] = useState<TableInterface>({table_num:0,num_seats:0,status:"-",restaurant_id:""});
    const [newTableNum, setNewTableNum] = useState(tableNumber);
    const [newNumSeats, setNewNumSeats] = useState("0");
    const [newStatus, setNewStatus] = useState<TableStatus>("Free");
    const emptyOrder: Order = {id:'', items:[], restaurant_id:'', status:'', table_num:0, total_cost:0, user_id:'', created_at:''}
    const [tableOrder, setTableOrder] = useState<Order>(emptyOrder)
    const [createFailure, setCreateFailure] = useState("");
    const [moveFailure, setMoveFailure] = useState("");
    const [failure, setFailure] = useState("");
    const [showAddItem, setShowAddItem] = useState(false);

    const [menuItems, setMenuItems] = useState<MenuItem[]>([]);
    const [quantity, setQuantity] = useState("0");
    const [newItem, setNewItem] = useState("");


    const updateTable = async () => {
        if (isNaN(parseInt(newNumSeats))) {
            setCreateFailure("Invalid number of seats");
        } else {
            await myMenuService.updateTable(table.table_num, parseInt(newNumSeats), newStatus);
            toaster.success(`Table ${table.table_num} Updated`);
            callback();
        }
    }

    const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

    const getTable = async (tableNum: number) => {
        const tableJson = await myMenuService.getTable(tableNum);
        if (tableJson && tableJson.item) {
            setTable(tableJson.item);
            setNewNumSeats((tableJson.item.num_seats).toString());
            setNewStatus(tableJson.item.status);
            const tableOrderResponse = await myMenuService.getTableOrder(tableNum);
            if (tableOrderResponse && tableOrderResponse.item && tableOrderResponse.item.id !== "") {
                setTableOrder(tableOrderResponse.item)
            } else {
                setTableOrder(emptyOrder)
            }
        }
        const itemsJson = await myMenuService.getMenuItems();
        if (itemsJson && itemsJson.data) {
            setMenuItems(itemsJson.data);
            setNewItem(itemsJson.data[0].id);
        }
    }

    const getTables = async () => {
        const json = await myMenuService.getTables();
        if (json && json.data) {
            setTables(json.data);
        }
    }

    useEffect(() => {
		setNewTableNum(tableNumber);
        getTable(tableNumber);
	}, [tableNumber]);

    useEffect(() => {
        getTables();
    }, [])

    const submitItem = async () => {
        await myMenuService.createOrder(table.table_num, {[newItem]: parseInt(quantity)});
        //await myMenuService.updateTableStatus(table.table_num, "occupy");
        setShowAddItem(false);
        getTable(table.table_num);
        setQuantity("0");
        callback();
    }


    const addItemToOrder = async () => {
        await myMenuService.addItemToOrder(tableOrder.id, newItem, parseInt(quantity));
        //await myMenuService.updateTableStatus(table.table_num, "occupy");
        setShowAddItem(false);
        getTable(table.table_num);
        setQuantity("0");
        callback();
    }

    const removeItemFromOrder = async (item_id: string) => {
        await myMenuService.removeItemFromOrder(tableOrder.id, item_id)
        getTable(table.table_num);
        callback();
    }

    const moveTable = async () => {
        const newTable = await myMenuService.getTable(newTableNum);
        if (newTable && newTable.item) {
            if (newTable.item.status == "Taken") {
                setMoveFailure("Cannot move to table that is already taken")
            } else {
                await myMenuService.moveOrder(tableOrder.id, newTableNum);
                getTable(newTableNum);
                callback();
                setMoveFailure("")
            }
        } else {
            setMoveFailure("Invalid new table number")
        }
    }

    const cancelOrder = async (order_id: string) => {
        await myMenuService.cancelOrder(order_id);
        getTable(table.table_num);
        callback();
    }

    const completeOrder = async (order_id: string) => {
        await myMenuService.completeOrder(order_id);
        getTable(table.table_num);
        callback();
    }


    const addItem = () => {
        return (
            <Dialog
			isShown={showAddItem}
			title={"Add Item"}
			onCancel={() => setShowAddItem(false)}
			onConfirm={tableOrder.id === "" ? submitItem : addItemToOrder}
			confirmLabel="Confirm"
			onCloseComplete={() => setShowAddItem(false)}
		    >
                <Label>Item</Label>
                <Select value={newItem} onChange={(e: any) => {setNewItem(e.target.value)}} width={500}>
                    {menuItems.map((item) => {
                        return (
                            <option value={item.id}>
                                {item.name} - ${item.price}
                            </option>
                        )
                    })}
                </Select>
                <TextField
                    label="Quantity"
                    input={quantity}
                    updateInput={setQuantity}
                    type="number"
                />
            </Dialog>
        )
    }

    return (
        <Pane>
            <DialogContainer>
                {showAddItem && addItem()}
            </DialogContainer>
            <Grid>
                <Label>Updating table {table.table_num}</Label>
                <Label>Taken</Label>
                <Switch height={20} checked={newStatus === "Taken"} onChange={(e: any) => setNewStatus(e.target.checked ? "Taken" : "Free")} />
                <TextField label="No. Seats" type="number" input={newNumSeats} updateInput={setNewNumSeats} />
                <Divider />
                {createFailure !== "" && invalidMsg(createFailure)}
                <Button appearance="primary" intent="success" onClick={updateTable}>Update</Button>
                <Divider />
                <Label>Move to table</Label>
                {moveFailure !== "" && invalidMsg(moveFailure)}
                <Select value={newTableNum} onChange={(e: any) => setNewTableNum(parseInt(e.target.value))} width="auto">
                    {tables.map((table) => {
                        return (
                            <option value={table.table_num}>
                                {table.table_num}
                            </option>
                        )
                    })}
                </Select>
                <Divider />
                <Button appearance="primary" intent="success" onClick={moveTable}>Move table</Button>
                <Label>Update Order</Label>
                <Table>
                    <Table.Head>
                        <Table.TextHeaderCell>Item</Table.TextHeaderCell>
                        <Table.TextHeaderCell>Quantity</Table.TextHeaderCell>
                        <Table.TextHeaderCell onClick={() => setShowAddItem(true)} isSelectable cursor={"pointer"} flexBasis={40} justifyContent={"center"}><PlusIcon /></Table.TextHeaderCell>
                    </Table.Head>
                    {<Table.VirtualBody height={150}>
                        {(tableOrder.items).map((item) => (
                            <TableRow>
                                <Table.TextCell>{item.item_name}</Table.TextCell>
                                <Table.TextCell>{item.quantity}</Table.TextCell>
                                <TableCell onClick={() => removeItemFromOrder(item.item_id)} isSelectable cursor={"pointer"} flexBasis={40} justifyContent={"center"}><TrashIcon /></TableCell>
                            </TableRow>
                        ))}
                    </Table.VirtualBody>}
                    <Table.VirtualBody height={50}>
                        <TableCell><Label>Total Price ${tableOrder.total_cost}</Label></TableCell>
                    </Table.VirtualBody>
                </Table>
                {tableOrder.id !== "" && <Flex>
                    <Button appearance="primary" intent="danger" onClick={() => cancelOrder(tableOrder.id)}>Cancel Order</Button>
                    <div style={{width:10}}></div>
                    <Button appearance="primary" intent="success" onClick={() => completeOrder(tableOrder.id)}>Complete Order</Button>
                    {failure !== "" && invalidMsg(failure)}
                </Flex>}
            </Grid>
        </Pane>
    );
}