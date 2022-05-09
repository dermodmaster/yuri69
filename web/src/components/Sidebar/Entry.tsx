import { useNavigate } from 'react-router';
import styled from 'styled-components';
import { Button } from '../Button';
import { Styled } from '../props';
import { LinearGradient } from '../styleParts';

type Props = Styled & {
  icon: JSX.Element;
  label: JSX.Element | string;
  to?: string;
  action?: () => void;
  color?: string;
  disabled?: boolean;
};

const EntryContainer = styled(Button)<{ color?: string }>`
  display: flex;
  align-items: center;
  justify-content: left;
  transition: all 0.2s ease;
  padding: 0;
  width: 100%;

  ${(p) => LinearGradient(p.color ?? p.theme.background3)}

  &:enabled:hover {
    filter: brightness(1.2);
  }
`;

const Icon = styled.div`
  min-width: 4em;
  width: 4em;
  height: 4em;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 1em;

  > svg {
    width: 50%;
    height: 50%;
  }
`;

const Label = styled.span`
  font-weight: 600;
  text-transform: uppercase;
  font-size: 1.3rem;
  padding: 0 1em 0 0;
  white-space: nowrap;
`;

export const Entry: React.FC<Props> = ({ icon, label, to, action = () => {}, color, ...props }) => {
  const nav = useNavigate();

  const _onClick = () => {
    if (!!to) nav(to);
    else action();
  };

  const _label = typeof label === 'string' ? <Label>{label}</Label> : label;

  return (
    <EntryContainer onClick={_onClick} color={color} {...props}>
      <Icon>{icon}</Icon>
      {_label}
    </EntryContainer>
  );
};
