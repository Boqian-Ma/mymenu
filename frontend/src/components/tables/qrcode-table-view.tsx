import {
	Button,
	PlusIcon,
	Select,
	Switch,
	Table,
	TableCell,
	TableRow,
	toaster,
	TrashIcon,
	Dialog,
	Pane,
} from "evergreen-ui";
import { useEffect, useState } from "react";
import styled from "styled-components";
import SelectField from "../select-field";
import { useServices } from "../services/service-context";
import TextField from "../text-field";
import { TableStatus } from "../utils/data-types";
import { Divider, Flex, Grid, Label } from "../utils/reusable-components";
import { MenuItem, Order } from "../utils/data-interfaces";

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

interface Props {
	tableNumber: number;
}

export default function QRCodeTable(props: Props) {
	const { tableNumber } = props;
	const { myMenuService } = useServices();
	const [qrcode, setQrCode] = useState("");
	const QRCode = require("qrcode.react");

	const generateQRCode = (tableid: number) => {
		var myString =
			myMenuService.getRootUrl() +
			"/customer/landing?restaurantid=" +
			myMenuService.getCurrRestaurant() +
			"&tablenum=" +
			tableid;
		// TODO: getCurrRestaurant currently not working
		setQrCode(myString);

		return;
	};

	useEffect(() => {
		generateQRCode(tableNumber);
	}, [tableNumber]);

	return (
		<Pane width={280}>
			<Grid>
				<Table>
					<Table.Head>
						<Table.TextHeaderCell>QR Code</Table.TextHeaderCell>
					</Table.Head>
					<Table.VirtualBody height={220}>
						<TableCell height={220} justifyContent={"center"}>
							<QRCode value={qrcode} size={200} />
						</TableCell>
					</Table.VirtualBody>
				</Table>
			</Grid>
		</Pane>
	);
}
