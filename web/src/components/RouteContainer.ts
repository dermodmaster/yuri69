import styled, { css } from 'styled-components';

type Props = {
  maxWidth?: string;
};

export const RouteContainer = styled.div<Props>`
  padding-top: 1.5em;
  padding-right: 1.5em;
  width: 100%;
  height: 100%;

  ${(p) =>
    p.maxWidth &&
    css`
      max-width: ${p.maxWidth};
    `}
`;
