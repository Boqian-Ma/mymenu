import React, { useEffect, useState } from "react";
import { Container, Grid } from "../utils/reusable-components";
import { Checkbox, Table } from "evergreen-ui";
import { useServices } from "../services/service-context";
import { Button } from "evergreen-ui";
import EditMenuItem from "./edit-item-display";
import { MenuItem } from "../utils/data-interfaces"

export default function MenuTable() {
    const { myMenuService } = useServices();
    const [data, setData] = useState<MenuItem[]>([]);
    const [filtered, setFiltered] = useState<MenuItem[]>([]);

    const getMenu = async () => {
        const json = await myMenuService.getMenuItems();
        if (json && json.data) {
            setData(json.data);
            setFiltered(json.data);
        } else {
            setData([])
            setFiltered([])
        }
    }

    const filterSubstr = (substr: string) => {
		const filteredData: MenuItem[] = data.filter((item: MenuItem) => item.name.startsWith(substr));
		setFiltered(filteredData);
	};

	const deleteMenuItem = (itemId: string) => {
		// remove element via api
    
		myMenuService.deleteMenuItem(itemId).then(() => {
			getMenu();
		});
	};

	useEffect(() => {
		getMenu();
	}, []);

    return (
        <Grid>
            <Table>
                <Table.Head>
                    <Table.SearchHeaderCell flexBasis={200} flexShrink={0} flexGrow={0} placeholder="Search by item name..." onChange={filterSubstr}/>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}>Description</Table.TextHeaderCell>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}>Price ($)</Table.TextHeaderCell>
                    <Table.SelectMenuCell flexBasis={110} flexShrink={0} flexGrow={0}>Category</Table.SelectMenuCell>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}>Chef Special</Table.TextHeaderCell>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}>On Menu</Table.TextHeaderCell>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}/>
                    <Table.TextHeaderCell flexBasis={110} flexShrink={0} flexGrow={0}/>
                </Table.Head>
                <Table.VirtualBody height={600}>
                    {filtered.map((food) => (
                        <Table.Row key={food.id}>
                            <Table.TextCell flexBasis={200} flexShrink={0} flexGrow={0}>{food.name}</Table.TextCell>
                            <Table.TextCell flexBasis={110} flexShrink={0} flexGrow={0}>{food.description}</Table.TextCell>
                            <Table.TextCell flexBasis={110} flexShrink={0} flexGrow={0}>{food.price}</Table.TextCell>
                            <Table.TextCell flexBasis={110} flexShrink={0} flexGrow={0}>{food.category_name}</Table.TextCell>
                            <Table.Cell justifyContent="center" flexBasis={110} flexShrink={0} flexGrow={0}><Checkbox checked={food.is_special} /></Table.Cell>
                            <Table.Cell justifyContent="center" flexBasis={110} flexShrink={0} flexGrow={0}><Checkbox checked={food.is_menu} /></Table.Cell>
                            <Table.Cell><EditMenuItem
                                id={food.id}
                                name={food.name}
                                description={food.description}
                                price={food.price}
                                category_id={food.category_id}
                                is_special={food.is_special}
                                is_menu={food.is_menu}
                                allergy={food.allergy}
                                filename={food.file}
                                callback={getMenu}
                            /></Table.Cell>
                            <Table.Cell><Button onClick={() => deleteMenuItem(food.id)}>DELETE</Button></Table.Cell>
                        </Table.Row>
                    ))}
                </Table.VirtualBody>
            </Table>
        </Grid>
    )
}
