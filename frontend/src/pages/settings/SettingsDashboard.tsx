import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Table, Pane } from "evergreen-ui";
import { Container, Divider, Flex, Grid, Header } from "../../components/utils/reusable-components";
import { useHistory } from "react-router";
import { useServices } from "../../components/services/service-context";
import {SettingOption, User} from "../../components/utils/data-types";
import SettingSidebar from "../../components/settings/setting-sidebar";
import PersonalSettings from "../../components/settings/personal-settings";
import BrandingSettings from "../../components/settings/branding-settings";
import RestaurantSettings from "../../components/settings/restaurant-settings";
import HealthSettings from "../../components/settings/health-settings";
import {useAuthenticationContext} from "../../components/services/authentication-context";

export default function SettingsDashboard() {
	const { myMenuService } = useServices();
	const [userType, setUser] = useState<User | undefined>();
	const [view, setView] = useState<SettingOption>('personal');
	const [auth, authActions] = useAuthenticationContext();

	const changeView = (option: SettingOption) => {
		setView(option);
	};

	return (
		<Container>
			<Flex>
				<Grid>
					<SettingSidebar changeView={changeView} userType={auth.userType}/>
				</Grid>
				<Grid>
					{view == 'personal' && <PersonalSettings/>}
					{view == 'health' && <HealthSettings/>}
					{view == 'restaurant' && <RestaurantSettings/>}
					{view == 'branding' && <BrandingSettings/>}
				</Grid>
				<Grid/>
			</Flex>
		</Container>
	);
}
