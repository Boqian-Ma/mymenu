import React, { useEffect, useState } from "react";
import { Container, Flex, Grid, Label, Divider } from "../utils/reusable-components";
import TextField from "../text-field";
import SelectField from "../select-field";
import { Button, FilePicker, Checkbox, Select } from "evergreen-ui";
import { useServices } from "../services/service-context";
import { Category } from "../utils/data-interfaces"
import styled from "styled-components";

const Logo = styled.div`
    width: 100px;
    height: 100px;
    border: 2px solid gray;
`;

const Image = styled.img`
    width: 100%;
    height: 100%;
`;

export default function AddItemDisplay() {
	const { myMenuService } = useServices();
	const id = myMenuService.getCurrRestaurant();
	const [categories, setCategories] = useState<Category[]>([]);
	const [name, setName] = useState("");
	const [description, setDescription] = useState("");
	const [priceStr, setPriceStr] = useState("");
	const [chefSpecial, setChefSpecial] = useState(false);
	const [onMenu, setOnMenu] = useState(false);
	const allergies = ['', 'Nuts', 'Sea Food', 'Shellfish', 'Dairy', 'Eggs']
	const [allergy, setAllergy] = useState("");
	const [categoryID, setCategoryID] = useState("");
	const [failure, setFailure] = useState("");
	const [file, setFile] = useState<File | undefined>(undefined);
	const [isAdded, setIsAdded] = useState(false);
	const [lastAdded, setLastAdded] = useState("");
	const [src, setSrc] = useState("");

	const addItem = async () => {
		const price = parseFloat(priceStr);
		if (name === "" || description === "") {
			setFailure("Name or Description cannot be empty");
		} else if (isNaN(price) || price <= 0) {
			setFailure("Invalid Price");
		} else if (file === undefined || !file.name.endsWith(".png")) {
			setFailure("Invalid file - Select a .png image");
		} else {
			await myMenuService.addMenuItem(name, description, price, chefSpecial, onMenu, categoryID, allergy, file.name, src);
			setIsAdded(true);
			setLastAdded(name);
			setFailure("");
			clearFields();
		}
	};

	const toBase64 = (file: File) => new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(file);
		reader.onload = () => resolve(reader.result);
		reader.onerror = error => reject(error);
	});

	const clearFields = () => {
		setName("");
		setDescription("");
		setPriceStr("");
		setChefSpecial(false);
		setOnMenu(false);
	};

	const getCategories = async () => {
		const json = await myMenuService.getMenuItemCategories();
		if (typeof json !== "undefined" && json.data != null) {
			setCategories(json.data);
			setCategoryID(json.data[0].id);
		}
	};

	const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

	const srcChange = async (f: File) => {
		const s = await toBase64(f);
		setSrc(String(s));
	}

	useEffect(() => {
		getCategories();
	}, []);

	return (
		<Container>
			<Label>Image</Label>
			<Logo><Image src={src} /></Logo>
			<FilePicker
                width={200}
                onChange={files => {
					setFile(files[0])
					srcChange(files[0])
				}}
                placeholder="Upload Logo"
            />
			<TextField label="Name" input={name} updateInput={setName} type="text" />
			<TextField label="Description" input={description} updateInput={setDescription} type="text" />
			<Label>Allergies</Label>
			<SelectField categories={allergies} value={allergy} updateValue={setAllergy} />
			<TextField label="Price" input={priceStr} updateInput={setPriceStr} type="number" />
			<Grid>
				<Label>Type</Label>
				{/* <SelectField categories={categories.map((c) => c.name)} value={categoryID} updateValue={setCategoryID}/> */}
				{/* <Select value={category.name} onChange={e => setCategory({id:e.target.id, name:e.target.value})} width={500}> */}
				<Select value={categoryID} onChange={(e) => setCategoryID(e.target.value)} width={500}>
					{categories.map((category) => {
						return <option value={category.id}>{category.name}</option>;
					})}
				</Select>
			</Grid>
			<Grid>
				<Label>Chef Special</Label>
				<Checkbox checked={chefSpecial} onChange={(e) => setChefSpecial(e.target.checked)}></Checkbox>
			</Grid>
			<Grid>
				<Label>On Menu</Label>
				<Checkbox checked={onMenu} onChange={(e) => setOnMenu(e.target.checked)}></Checkbox>
			</Grid>
			{failure !== "" && invalidMsg(failure)}
			{isAdded && <Label style={{ color: "green" }}>Item {lastAdded} added</Label>}
			<Flex>
				<Button appearance="primary" intent="success" onClick={addItem}>
					Add
				</Button>
				<Button onClick={clearFields}>Clear</Button>
			</Flex>
		</Container>
	);
}
