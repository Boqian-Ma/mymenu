import React, {useState} from 'react';
import {AlignedHeader, Container, Divider, Flex, Grid, Header, Label} from "../utils/reusable-components";
import TextField from "../text-field";
import {Button, FilePicker} from "evergreen-ui";
import styled from "styled-components";

const Logo = styled.div`
    width: 100px;
    height: 100px;
    border: 2px solid gray;
`;

export default function BrandingSettings() {
    const [selectedFile, setFile] = useState<FileList | undefined>(undefined);

    const updateDetails = () => {
        // TODO: API call
    };

    return (
        <Container>
            <Logo>{/* TODO: Show logo here */}</Logo>
            <Divider/>
            <FilePicker
                width={200}
                onChange={files => {
                    setFile(files);
                    console.log(files);
                }}
                placeholder="Upload Logo"
            />
        </Container>
    )
}