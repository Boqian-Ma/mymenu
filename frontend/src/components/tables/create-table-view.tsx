import { Button, toaster } from "evergreen-ui";
import { useState } from "react";
import { useServices } from "../services/service-context";
import TextField from "../text-field";
import { Divider, Grid, Label } from "../utils/reusable-components";

interface Props {
    tableNumber: number;
    callback(): any;
}

export default function CreateTable(props: Props) {
    const { tableNumber, callback } = props;
    const { myMenuService } = useServices();
    const [numSeats, setNumSeats] = useState("");
    const [failure, setFailure] = useState("");

    const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};


    const createTable = async () => {
        if (isNaN(parseInt(numSeats))) {
            setFailure("Invalid number of seats");
        } else {
            await myMenuService.createTable(tableNumber, parseInt(numSeats));
            toaster.success(`Table ${tableNumber} Created`);
            callback();
        }
    }

    return (
        <Grid>
            <Label>Creating table {tableNumber}</Label>
            <TextField label="No. Seats" type="number" input={numSeats} updateInput={setNumSeats} />
            <Divider />
            {failure !== "" && invalidMsg(failure)}
            <Button appearance="primary" intent="success" onClick={createTable}>Create</Button>
        </Grid>
    );
}