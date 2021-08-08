import {storiesOf} from "@storybook/react";
import AddMenuDetails from "../pages/menus/AddMenuDetails";

storiesOf('Menus', module)
    .add('Add Menu', () => {
        return (
            <AddMenuDetails/>
        )
    });