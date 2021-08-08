import React, {useEffect, useState} from 'react';
import {Container, Flex, Grid, Label, Divider} from "../utils/reusable-components";
import TextField from "../text-field";
import SelectField from "../select-field";
import {Button, FilePicker, Checkbox, Table, TableRow} from "evergreen-ui";
import { useServices } from '../services/service-context';
import { Category } from "../utils/data-interfaces"

export default function CreateCategoryDisplay() {
    const { myMenuService } = useServices();
    const [category, setCategory] = useState("");
    const [failure, setFailure] = useState("");
    const [categories, setCategories] = useState<Category[]>([])

    const clearFields = () => {
        setCategory("");
    }

    const addCategory = async () => {
        setCategory(category.trim());
        if (category === "") {
            setFailure("Category cannot be empty!!");
        } else {
            const result = await myMenuService.createMenuItemCategory(category);
            if (result.hasOwnProperty('message')) {
                setFailure(result['message']);
            } else {
                setFailure("");
                clearFields();
                getCategories();
            }
        }
    };

    const getCategories = async () => {
        const json = await myMenuService.getMenuItemCategories();
        if (typeof json !== "undefined" && json.data != null) {
            setCategories(json.data);
        } 
    }
    const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

    useEffect(() => {
		getCategories();
	}, []);

    return (
        <Container>
            <TextField
                label="Category Name"
                input={category}
                updateInput={setCategory}
                type="text" />

            {failure !== "" && invalidMsg(failure)}
            <Divider />
            <Button appearance="primary" intent="success" onClick={addCategory}>Add</Button>
            <Divider/>
            <Table>
                <Table.Head>
                    <Table.TextHeaderCell>Category name</Table.TextHeaderCell>
                </Table.Head>
                <Table.VirtualBody height={500}>{categories.map((category) => (
                    <Table.Row key={category.id}>
                        <Table.TextCell>{category.name}</Table.TextCell>
                    </Table.Row>
                ))}
                </Table.VirtualBody>
            </Table>
        </Container>
    )
}