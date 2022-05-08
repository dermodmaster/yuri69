import styled, { createGlobalStyle, ThemeProvider } from 'styled-components';
import { DarkTheme } from './theme/theme';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { MainRoute } from './routes/Main';
import { LoginRoute } from './routes/Login';
import { SoundsRoute } from './routes/Sounds';

const GlobalStyle = createGlobalStyle`
  @import url("https://fonts.googleapis.com/css2?family=Rubik:wght@300;500;600&display=swap");

  body {
    font-family: 'Rubik', sans-serif;
    background-color: ${(p) => p.theme.background};
    color: ${(p) => p.theme.text};
    padding: 0;
    margin: 0;
  }

  * {
    box-sizing: border-box;
  }
`;

const Outlet = styled.div`
  width: 100vw;
  height: 100vh;
`;

const App: React.FC = () => {
  return (
    <ThemeProvider theme={DarkTheme}>
      <Outlet>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<MainRoute />}>
              <Route index element={<SoundsRoute />} />
              <Route path="settings" element={<p>pog</p>} />
            </Route>
            <Route path="/login" element={<LoginRoute />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </BrowserRouter>
      </Outlet>
      <GlobalStyle />
    </ThemeProvider>
  );
};

export default App;
