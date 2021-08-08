import React, {createContext, ReactNode, useContext} from 'react';
import {myMenuService} from "./mymenu-service";

const service = {
  myMenuService: new myMenuService(),
};

export const ServiceProvider = (props: { children: ReactNode }) => {
  return <ServicesContext.Provider value={service}>{props.children}</ServicesContext.Provider>
};

export const ServicesContext = createContext(service);

export const useServices = () => {
  return useContext(ServicesContext);
};
