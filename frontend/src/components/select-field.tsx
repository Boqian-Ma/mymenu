import React, {useState} from 'react';
import {Select} from "evergreen-ui";

interface Props {
    categories: string[];
    value: string,
    updateValue(e: any): any
}

export default function SelectField(props: Props) {
    const { categories, value, updateValue } = props;

    //const [value, setValue] = useState('');

    return (
        <Select value={value} onChange={(e: any) => updateValue(e.target.value)} width={500}>
            {categories.map((category) => {
                return (
                    <option value={category}>
                        {category}
                    </option>
                )
            })}
        </Select>
    )
}