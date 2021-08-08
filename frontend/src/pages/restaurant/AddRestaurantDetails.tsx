import React, { useEffect, useState } from "react";
import {
	Container,
	Divider,
	Flex,
	Grid,
	Header,
	Label,
} from "../../components/utils/reusable-components";
import styled from "styled-components";
import TextField from "../../components/text-field";
import { Button, FilePicker, Textarea } from "evergreen-ui";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";
import SelectField from "../../components/select-field";

const Footer = styled.div`
	display: grid;
	align-items: center;
	margin-top: 4rem;
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

export default function AddRestaurantDetails() {
	const [name, setName] = useState("");
	const [type, setType] = useState("");
	const [location, setLocation] = useState("");
	const [file, setFile] = useState<File | undefined>(undefined);
	const [email, setEmail] = useState("");
	const [phone, setPhone] = useState("");
	const [website, setWebsite] = useState("");
	const [src, setSrc] = useState("");

	const cuisines = ['Asian', 'Indian', 'European', 'Mediterranean', 'North American', 'South American']
	const [cuisine, setCuisine] = useState(cuisines[0]);

	const [businessHours, setBusinessHours] = useState("");
	const { myMenuService } = useServices();
	const [failure, setFailure] = useState('');
	const history = useHistory();

	const createRestaurant = async () => {
		// TODO: call API and create a restaurant
		if (name.trim() === "" || type.trim() === "" || location.trim() === "") {
			setFailure("All fields must be filled")
		} else if (file === undefined || !file.name.endsWith(".png")) {
			setFailure("Invalid file - Select a .png image");
		} else {
			//const file64 = await toBase64(file);
			//console.log(file64);
			await myMenuService.addRestaurant(name, type, location, email, phone, website, businessHours, file.name, src, cuisine);
			history.push("/admin/dashboard");
		}
	};
	
	const toBase64 = (file: File) => new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(file);
		reader.onload = () => resolve(reader.result);
		reader.onerror = error => reject(error);
	});

	const invalidMsg = (msg:String) => {
        return (
            <Label style={{color: "red"}}>
                {msg}
            </Label>
        )
    }

	const srcChange = async (f: File) => {
		const s = await toBase64(f);
		setSrc(String(s));
	}

	const basicSettings = (
        <Grid>
            <Logo><Image src={src} /></Logo>
            <Divider/>
            <FilePicker
                width={200}
                onChange={files => {
					setFile(files[0])
					srcChange(files[0])
				}}
                placeholder="Upload Logo"
            />
            <Divider/>
            <Grid>
                <TextField label="Business Name" input={name} updateInput={setName} type="text"/>
            </Grid>
            <Divider/>
            <Grid>
                <TextField label="Business Location" input={location} updateInput={setLocation} type="text"/>
            </Grid>
            <Divider/>
            <Grid>
                <TextField label="Business Type" input={type} updateInput={setType} type="text"/>
            </Grid>
            <Divider/>
            <Grid>
                <Label>Cuisine</Label>
				<SelectField categories={cuisines} value={cuisine} updateValue={setCuisine}/>
            </Grid>
        </Grid>
    )

    const contactSettings = (
        <Grid>
			<Label>Contact Details</Label>
            <TextField label="Email" input={email} updateInput={setEmail} type="text"/>
            <TextField label="Phone" input={phone} updateInput={setPhone} type="text"/>
            <TextField label="Website" input={website} updateInput={setWebsite} type="text"/>
            <Label>Business Hours</Label>
            <Textarea onChange={(e: any) => setBusinessHours(e.target.value)} value={businessHours} style={{height:150}}/>
        </Grid>
    )

	return (
		<Container>
			<Header>
				<b>Register My Restaurant</b>
			</Header>
			<Flex>
				{basicSettings}
				{contactSettings}
			</Flex>
			{failure !== "" && invalidMsg(failure)}
			<Footer>
				<Button width={200} onClick={createRestaurant}>
					Sign Up
				</Button>
			</Footer>
		</Container>
	);
}
