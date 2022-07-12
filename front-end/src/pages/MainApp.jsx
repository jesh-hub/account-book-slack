import '@/pages/MainApp.scss';
import { Container, Nav, Navbar } from 'react-bootstrap';
import { Link, Outlet, useHref } from 'react-router-dom';
import { RiGroup2Fill } from 'react-icons/ri';
import { useCallback, useState } from 'react';

export default function MainApp() {
  const to = useHref('/group');

  const [navBarHeight, setNavBarClient] = useState(0);
  const navBar = useCallback(node => {
    if (node !== null)
      setNavBarClient(node.getBoundingClientRect().height);
  }, []);
  return (
    <>
      <main
        className="main-app"
        style={{ 'paddingBottom': navBarHeight }}
      >
        <Outlet />
      </main>
      <Navbar
        variant="light"
        fixed="bottom"
        bg="light"
        ref={navBar}
      >
        <Container>
          <Nav className="main-nav">
            <Nav.Link
              as={Link}
              className="nav-item"
              to={to}
            ><RiGroup2Fill /></Nav.Link>
          </Nav>
        </Container>
      </Navbar>
    </>
  );
}
