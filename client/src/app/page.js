import { useState } from 'react';
import Link from 'next/link'
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

export default function Home() {

    const [loginStatus, setLoginStatus] = useState(false);


    function checkLogin() {

    }


    return (
        <div className=" flex flex-Col items-center  w-full h-lvh  ">
            <Navbar bg="dark" data-bs-theme="dark">
                <Container>
                    <Navbar.Brand href="#home">ADA Map</Navbar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="#home">Home</Nav.Link>
                    </Nav>
                    <Nav className="me-auto">
                        <Nav.Link href="#home">Login</Nav.Link>
                    </Nav>
                </Container>
            </Navbar>
        </div>
    );
}
