import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Pane, Table } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header, Label } from "../../components/utils/reusable-components";
import { TableStatus } from "../utils/data-types";
import { useServices } from "../services/service-context";
import CreateTable from "./create-table-view";
import UpdateTable from "./update-table-view";
import { Order, Table as TableInterface } from "../utils/data-interfaces";
import QRCodeTable from "./qrcode-table-view";

export default function TablesList() {
	const { myMenuService } = useServices();
	const [tables, setTables] = useState<TableInterface[]>([]);
	const [orders, setOrders] = useState<Order[]>([]);
	const [filteredTables, setFilteredTables] = useState<TableInterface[]>([]);
	const [statusFilter, setStatusFilter] = useState("-");
	const [formView, setFormView] = useState("create");
	const [nextTableNum, setNextTableNum] = useState(1);
	const [updateTableNum, setUpdateTableNum] = useState(1);
	const statusTypes: TableStatus[] = ["-", "Taken", "Free"];

	const getTablesDetails = async () => {
		const json = await myMenuService.getTables();
		if (json && json.data) {
			setTables(json.data);
			setFilteredTables(
				json.data.filter((table: TableInterface) => statusFilter == "-" || table.status == statusFilter)
			);
			setNextTableNum(json.data.length + 1);
		}

		const ordersJson = await myMenuService.getOrders("", true);
		if (ordersJson && ordersJson.data) {
			setOrders(ordersJson.data);
		} else {
			setOrders([]);
		}

		//console.log((orders.filter((order:Order) => order.table_num === 1)))
	};

	const createTable = () => {
		setFormView("create");
	};

	const updateTable = (tableNum: number) => {
		setFormView("update");
		setUpdateTableNum(tableNum);
	};

	const qrcodeTable = (tablenum: number) => {
		setFormView("qrcode");
		setUpdateTableNum(tablenum);
	};

	const getOrderPrice = (tableNum: number) => {
		const order = orders.filter((order: Order) => order.table_num === tableNum)[0];
		if (order) {
			return order.total_cost;
		}
		return 0;
	};

	const getOrderStatus = (tableNum: number) => {
		const order = orders.filter((order: Order) => order.table_num === tableNum)[0];
		if (order) {
			return order.status;
		}
		return "-";
	};

	useEffect(() => {
		getTablesDetails();
	}, [statusFilter]);

	return (
		<Container>
			<Flex>
				<Table width={500}>
					<Table.Head>
						<Table.TextHeaderCell>Table No.</Table.TextHeaderCell>
						<Table.TextHeaderCell>No. Seats</Table.TextHeaderCell>
						<Table.TextHeaderCell>Total Cost</Table.TextHeaderCell>
						<Table.SelectMenuCell
							selectMenuProps={{
								options: statusTypes.map((label) => ({ label, value: label })),
								selected: statusFilter,
								onSelect: (item) => setStatusFilter(item.value.toString()),
								hasTitle: false,
								hasFilter: false,
								closeOnSelect: true,
							}}
						>
							{statusFilter}
						</Table.SelectMenuCell>
						<Table.TextHeaderCell>Order Status</Table.TextHeaderCell>
						<Table.TextHeaderCell />
						<Table.TextHeaderCell />
					</Table.Head>
					<Table.VirtualBody height={400}>
						{filteredTables.map((table) => (
							<Table.Row key={table.table_num}>
								<Table.TextCell>{table.table_num}</Table.TextCell>
								<Table.TextCell>{table.num_seats}</Table.TextCell>
								<Table.TextCell>{getOrderPrice(table.table_num)}</Table.TextCell>
								<Table.TextCell
									textProps={table.status === "Taken" ? { color: "red" } : { color: "green" }}
								>
									{table.status}
								</Table.TextCell>
								<Table.TextCell>{getOrderStatus(table.table_num)}</Table.TextCell>
								<Table.Cell flexBasis={40}>
									<Button onClick={() => updateTable(table.table_num)} padding={20}>
										UPDATE
									</Button>
								</Table.Cell>
								<Table.Cell flexBasis={40}>
									<Button onClick={() => qrcodeTable(table.table_num)} padding={20}>
										QR CODE
									</Button>
								</Table.Cell>
							</Table.Row>
						))}
					</Table.VirtualBody>
				</Table>
				<div style={{ width: 75 }}></div>
				<Pane border paddingLeft={10} paddingRight={10}>
					{formView === "create" && <CreateTable tableNumber={nextTableNum} callback={getTablesDetails} />}
					{formView === "update" && <UpdateTable tableNumber={updateTableNum} callback={getTablesDetails} />}
					{formView === "qrcode" && <QRCodeTable tableNumber={updateTableNum} />}
				</Pane>
			</Flex>
			<Button onClick={createTable} style={{ width: 100 }} fontSize={15} padding={20}>
				Create Table
			</Button>
		</Container>
	);
}
