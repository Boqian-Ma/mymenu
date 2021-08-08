import React, { useEffect, useState } from "react";
import { Button, Dialog, Pane, TextInput, Checkbox, Select, FilePicker } from "evergreen-ui";
import { User } from "../utils/data-types";
import styled from "styled-components";
import { useServices } from "../services/service-context";
import { useHistory } from "react-router-dom";
import TextField from "../text-field";
import { Divider, Grid, Label } from "../utils/reusable-components";
import SelectField from "../select-field";
import { Category } from "../utils/data-interfaces"

const DialogContainer = styled.div`
	padding: 0.5rem;
`;

const Logo = styled.div`
	width: 100px;
	height: 100px;
	border: 2px solid gray;
`;

const Image = styled.img`
    width: 100%;
    height: 100%;
`;

interface Props {
	id: string;
	name: string;
	description: string;
	price: number;
	is_special: boolean;
	is_menu: boolean;
	category_id: string;
	allergy: string;
	filename: string;
	callback(): any;
}

export default function EditMenuItem(prop: Props) {
	const { id, name, description, price, is_special, is_menu, category_id, allergy, filename, callback } = prop;
	const [isShown, setIsShown] = useState(false);
	const [failure, setFailure] = useState("");
	const { myMenuService } = useServices();

	const [categories, setCategories] = useState<Category[]>([]);
	const [newName, setNewName] = useState(name);
	const [newDescription, setNewDescription] = useState(description);
	const [newPriceStr, setNewPriceStr] = useState(price.toString());
	const [newChefSpecial, setNewChefSpecial] = useState(is_special);
	const [newOnMenu, setNewOnMenu] = useState(is_menu);
	const [newCategoryID, setNewCategoryID] = useState(category_id);
	const [newFilename, setNewFilename] = useState(filename);
	const [newFile, setNewFile] = useState<File | undefined>(undefined);
	const allergies = ['', 'Nuts', 'Sea Food', 'Shellfish', 'Dairy', 'Eggs']
	const [newAllergy, setNewAllergy] = useState(allergy);
	const [src, setSrc] = useState(process.env.PUBLIC_URL + `/assets/images/${filename}`);

	const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

	const open = () => {
		setIsShown(true);
	};

	const toBase64 = (file: File) => new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(file);
		reader.onload = () => resolve(reader.result);
		reader.onerror = error => reject(error);
	});

	const submit = async () => {
		const newPrice = parseFloat(newPriceStr);
		if (name === "" || description === "") {
			setFailure("Name or Description cannot be empty!");
		} else if (isNaN(price) || price <= 0) {
			setFailure("Invalid Prices");
		} else if (newFilename === "" || (newFilename !== "" && !newFilename.endsWith(".png"))) {
			setFailure("Invalid file - Select a .png image");
		} else {
			let file64 = "";
			if (newFile) {
				file64 = String(await toBase64(newFile));
			}
			await myMenuService.editMenuItem(id, newName, newDescription, newPrice, newChefSpecial, newOnMenu, newCategoryID, newAllergy, newFilename, file64);
			setIsShown(false);
			callback();
		}
	};

	const getCategories = async () => {
		const json = await myMenuService.getMenuItemCategories();
		if (typeof json !== "undefined" && json.data != null) {
			setCategories(json.data);
		}
	};

	const srcChange = async (f: File) => {
		const s = await toBase64(f);
		setSrc(String(s));
	}

	useEffect(() => {
		getCategories();
	}, []);

	return (
		<Pane>
			<DialogContainer>
				<Dialog
					isShown={isShown}
					title="Edit menu item"
					onCancel={() => setIsShown(false)}
					onConfirm={submit}
					confirmLabel="Confirm"
					onCloseComplete={() => setIsShown(false)}
				>
					<Logo><Image src={src} /></Logo>
					<Divider />
					<FilePicker
						width={200}
						onChange={(files) => {
							setNewFile(files[0]);
							setNewFilename(files[0].name);
							srcChange(files[0])
						}}
						placeholder={newFilename || "Placeholder text"}
					/>
					<TextField label="Name" input={newName} updateInput={setNewName} type="text" />
					<TextField label="Description" input={newDescription} updateInput={setNewDescription} type="text" />
					<Label>Allergies</Label>
					<SelectField categories={allergies} value={newAllergy} updateValue={setNewAllergy} />
					<TextField label="Price" input={newPriceStr} updateInput={setNewPriceStr} type="number" />
					<Grid>
						<Label>Type</Label>
						{/* <SelectField categories={categories} value={newCategory} updateValue={setNewCategory}/> */}
						<Select value={newCategoryID} onChange={(e) => setNewCategoryID(e.target.value)} width={500}>
							{categories.map((category) => {
								return <option value={category.id}>{category.name}</option>;
							})}
						</Select>
					</Grid>
					<Grid>
						<Label>Chef Special</Label>
						<Checkbox
							checked={newChefSpecial}
							onChange={(e) => setNewChefSpecial(e.target.checked)}
						></Checkbox>
					</Grid>
					<Grid>
						<Label>On Menu</Label>
						<Checkbox
							checked={newOnMenu}
							onChange={(e) => setNewOnMenu(e.target.checked)}
						></Checkbox>
					</Grid>

					{failure !== "" && invalidMsg(failure)}
				</Dialog>

				<Button appearance="minimal" onClick={open}>
					EDIT
				</Button>
			</DialogContainer>
		</Pane>
	);
}
