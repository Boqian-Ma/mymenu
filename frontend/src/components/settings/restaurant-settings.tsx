import React, { useEffect, useState } from "react";
import { AlignedHeader, Container, Divider, Flex, Grid, Header, Label } from "../utils/reusable-components";
import TextField from "../text-field";
import { Button, FilePicker, Textarea, toaster } from "evergreen-ui";
import styled from "styled-components";
import { Restaurant } from "../../components/utils/data-interfaces";
import { useServices } from "../services/service-context";
import SelectField from "../select-field";

const Logo = styled.div`
	width: 100px;
	height: 100px;
	border: 2px solid gray;
`;

const Image = styled.img`
    width: 100%;
    height: 100%;
`;

export default function RestaurantSettings() {
	const { myMenuService } = useServices();
	const [restaurantInfo, setRestaurantInfo] = useState<Restaurant>({
		id: "",
		name: "",
		location: "",
		type: "",
		email: "",
		phone: "",
		website: "",
		businessHours: "",
		file: "",
		cuisine: "",
	});
	const [restaurants, setRestaurants] = useState<Restaurant[]>([]);
	const [newName, setNewName] = useState("");
	const [newLocation, setNewLocation] = useState("");
	const [newType, setNewType] = useState("");
	const [currRestaurant, setCurrRestaurant] = useState<Restaurant>({
		id: "",
		name: "",
		location: "",
		type: "",
		email: "",
		phone: "",
		website: "",
		businessHours: "",
		file: "",
		cuisine: "",
	});
	const [selectedRestaurant, setSelectedRestaurant] = useState("");
	const [newEmail, setNewEmail] = useState("");
	const [newPhone, setNewPhone] = useState("");
	const [newWebsite, setNewWebsite] = useState("");
	const [newFile, setNewFile] = useState<File | undefined>(undefined);
	const [newFilename, setNewFilename] = useState("");
	const [newBusinessHours, setNewBusinessHours] = useState("");
	const [failure, setFailure] = useState("");
	const [src, setSrc] = useState("");

	const cuisines = ['Asian', 'Indian', 'European', 'Mediterranean', 'North American', 'South American']
	const [newCuisine, setNewCuisine] = useState("");

	const formatName = (name: string, location: string) => {
		return `${name} - ${location}`;
	};

	const toBase64 = (file: File) => new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(file);
		reader.onload = () => resolve(reader.result);
		reader.onerror = error => reject(error);
	});

	const updateDetails = async () => {
		// need to error check
		if (newName === "" || newLocation === "" || newType === "") {
			setFailure("Name, Location and Type fields must be filled in");
		} else if (newFilename === "" || (newFilename !== "" && !newFilename.endsWith(".png"))) {
			setFailure("Invalid file - Select a .png image");
		} else {
			let file64 = "";
			if (newFile) {
				file64 = String(await toBase64(newFile));
			}
			const response = await myMenuService.editRestaurant(
				currRestaurant.id,
				newName,
				newType,
				newLocation,
				newEmail,
				newPhone,
				newWebsite,
				newBusinessHours,
				newFilename,
				file64,
				newCuisine
			);
			if (response.hasOwnProperty("message")) {
				setFailure(response["message"]);
			}
			toaster.success("Restaurant Details Updated");
			setFailure("");
		}
	};

	const getUserRestaurants = async () => {
		const json = await myMenuService.getUserRestaurants();
		if (typeof json !== "undefined" && json.data != null) {
			setRestaurants(json.data);
			setSelectedRestaurant(formatName(json.data[0].name, json.data[0].location));
		}
	};

	useEffect(() => {
		getUserRestaurants();
	}, []);

	const invalidMsg = (msg: String) => {
		return <Label style={{ color: "red" }}>{msg}</Label>;
	};

	const srcChange = async (f: File) => {
		const s = await toBase64(f);
		setSrc(String(s));
	}

	useEffect(() => {
		const [name, location] = selectedRestaurant.split(" - ");
		for (const i in restaurants) {
			const r = restaurants[i];
			if (r.name == name && r.location == location) {
				setCurrRestaurant(r);
				setNewName(r.name.trim());
				setNewLocation(r.location.trim());
				setNewType(r.type.trim());
				setNewEmail(r.email.trim());
				setNewPhone(r.phone.trim());
				setNewWebsite(r.website.trim());
				setNewBusinessHours(r.businessHours);
				setNewFilename(r.file);
				setNewCuisine(r.cuisine);
				setSrc(process.env.PUBLIC_URL + `/assets/images/${r.file}`)
			}
		}
	}, [selectedRestaurant]);

	const basicSettings = (
		<Grid>
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
			<Divider />
			<Grid>
				<TextField label="Business Name" input={newName} updateInput={setNewName} type="text" />
			</Grid>
			<Divider />
			<Grid>
				<TextField label="Business Location" input={newLocation} updateInput={setNewLocation} type="text" />
			</Grid>
			<Divider />
			<Grid>
				<TextField label="Business Type" input={newType} updateInput={setNewType} type="text" />
			</Grid>
            <Divider/>
            <Grid>
                <Label>Cuisine</Label>
				<SelectField categories={cuisines} value={newCuisine} updateValue={setNewCuisine}/>
            </Grid>
		</Grid>
	);

	const contactSettings = (
		<div style={{ position: "relative", left: 150 }}>
			<Label>Contact Details</Label>
			<TextField label="Email" input={newEmail} updateInput={setNewEmail} type="text" />
			<TextField label="Phone" input={newPhone} updateInput={setNewPhone} type="text" />
			<TextField label="Website" input={newWebsite} updateInput={setNewWebsite} type="text" />
			<Label>Business Hours</Label>
			<Textarea
				onChange={(e: any) => setNewBusinessHours(e.target.value)}
				value={newBusinessHours}
				style={{ height: 150 }}
			/>
		</div>
	);

	return (
		<Container>
			<AlignedHeader>Restaurant</AlignedHeader>
			<SelectField
				categories={restaurants.map((r: Restaurant) => formatName(r.name, r.location))}
				value={selectedRestaurant}
				updateValue={setSelectedRestaurant}
			/>
			<Divider />
			<Flex>
				{basicSettings}
				{contactSettings}
			</Flex>
			<Divider />
			{failure !== "" && invalidMsg(failure)}
			<Divider />
			<Button width={200} onClick={updateDetails}>
				Update Details
			</Button>
		</Container>
	);
}
