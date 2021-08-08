import styled from 'styled-components';

export const Container = styled.div`
    display: grid;
    max-width: 1000px;
    margin: 0 auto;
`;

export const Flex = styled.div`
    display: flex;
    margin-top: 1rem;
`;

export const Image = styled.img`
    @media only screen and (max-width: 500px) {
        max-width: 250px;
    }
`;

export const Divider = styled.div`
  width: 100%;
  height: 1px;
  margin-top: 10px;
  margin-bottom: 10px;
`;

export const Grid = styled.div`
    display: flex;
    flex-direction: column;
    flex-basis: 100%;
    flex: 1;
`;

export const Margin = styled(Grid)`
    margin-left: 1rem;
`;

export const Header = styled.h2`
    align-items: center;
    margin: 1rem;
`;

export const AlignedHeader = styled.h3`
    text-align: left;
`;

export const AlignedText = styled.p`
    text-align: left;
`;

export const Label = styled.p`
	text-align: left;
	font-weight: bold;
`;

export const Bold = styled.p`
    font-weight: bold;
`;

export const ButtonStack = styled.div`
    width: 500px;
    margin: auto;
    flex-direction: row;
    align-items: center;
    justify-content: center;
`;

export const KitchenGrid = styled.div`
    display: grid;
    grid-gap: 10px;
    grid-template-columns: repeat(4, 1fr);
`;