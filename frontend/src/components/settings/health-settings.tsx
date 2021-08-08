import React, {useState} from 'react';
import {AlignedHeader, Container, Divider, Flex, Grid, Header, Label} from "../utils/reusable-components";
import TextField from "../text-field";
import {Button, FilePicker} from "evergreen-ui";
import styled from "styled-components";
import TagInputField from "../tag-input-field";

export default function HealthSettings() {
    const [location, setLocation] = useState("");
    const [allergies, setAllergies] = useState<string[]>([]);
    const [cuisines, setCuisines] = useState<string[]>([]);

    const updateDetails = () => {
        // TODO: API call
    };

    return (
        <Container>
            <Label>Allergies</Label>
            <TagInputField placeholderText="Add allergies" existingValues={allergies} updateValues={setAllergies}/>
            <Divider/>
            <Label>Food Preferences</Label>
            <TagInputField placeholderText="Add cuisines" existingValues={cuisines} updateValues={setCuisines}/>
            <Divider/>
            <Button intent="success">Update Details</Button>
        </Container>
    )
}