import React, {useState} from 'react';
import {TagInput} from "evergreen-ui";

interface Props {
    placeholderText: string;
    existingValues: string[];
    updateValues(e: string[]): any;
}

export default function TagInputField(props: Props) {
    const { placeholderText, existingValues, updateValues } = props;
    const [values, setValues] = useState<string[]>(existingValues);

    const changeValues = (values: string[]) => {
        setValues(values);
        updateValues(values);
    };

    return (
        <TagInput
            inputProps={{ placeholder: placeholderText }}
            values={values}
            tagSubmitKey="enter"
            onChange={(values: string[]) => changeValues(values)}
        />
    )
}