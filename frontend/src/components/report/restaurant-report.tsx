import { Pane } from "evergreen-ui";
import { useEffect, useState } from "react";
import SelectField from "../select-field";
import { useServices } from "../services/service-context";
import { Category, MenuItem } from "../utils/data-interfaces";
import { Container, Divider, Label } from "../utils/reusable-components";

interface ReportItem {
    most_ordered_item: MenuItem;
    most_ordered_item_quantity: number;
    most_ordered_category: Category;
    total_revenue: number;
}

export default function RestaurantReport() {
    const { myMenuService } = useServices();
    const [currReport, setCurrReport] = useState<ReportItem>({most_ordered_item:{id:"",name:"",description:"",price:0,is_special:false,is_menu:false,category_id:"",file:"",category_name:"",allergy:""}, most_ordered_category:{id:"", name:""}, most_ordered_item_quantity:0, total_revenue:0})
    const categories: {[key: string]: string;} = {"Today":"daily", "This Week": "weekly", "This Month":"monthly", "Last 3 Months":"quarterly", "This Year":"yearly", "All Time":"all_time"}
    const [category, setCategory] = useState("Today");

    const getReport = async () => {
        const reportResponse = await myMenuService.getReport();
        if (reportResponse && reportResponse.item && (reportResponse.item)[categories[category]]) {
            setCurrReport((reportResponse.item)[categories[category]])
        }
    }
    
    useEffect(() => {
        getReport()
    }, [category])

    return (
        <Container>
            <SelectField categories={Object.keys(categories)} value={category} updateValue={setCategory}/>
            <Divider></Divider>
            <Pane border paddingLeft={20} width={500}>
                <Label>Gross Sales</Label>
                <Label>${currReport.total_revenue}</Label>
            </Pane>
            <Divider></Divider>
            <Pane border paddingLeft={20} width={500}>
                <Label>Most Popular Product</Label>
                <Label>{currReport.most_ordered_item.name}</Label>
                </Pane>
            <Divider></Divider>
            <Pane border paddingLeft={20} width={500}>
                <Label>Top Category By Sales</Label>
                <Label>{currReport.most_ordered_category.name}</Label>
            </Pane>
        </Container>
    )
}