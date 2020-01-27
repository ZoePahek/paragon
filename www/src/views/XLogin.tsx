import Cookies from "js-cookie";
import * as React from "react";
import { Redirect } from "react-router-dom";
import {
  Button,
  Form,
  Grid,
  Header,
  Icon,
  Image,
  Segment
} from "semantic-ui-react";

const XLogin = () => {
  let userId = Cookies.get("pg-userid");
  if (userId) {
    return <Redirect to="/" />;
  }

  return (
    <Grid textAlign="center" style={{ height: "100vh" }} verticalAlign="middle">
      <Grid.Column style={{ maxWidth: 450 }}>
        <Header as="h2" color="blue" textAlign="center">
          <Image src="/app/logo512.png" /> Log-in or Sign up
        </Header>
        <Form size="large">
          <Segment stacked>
            <Button
              href="/oauth/login"
              icon
              basic
              labelPosition="right"
              size="large"
            >
              <Icon name="google" />
              Sign Up
            </Button>
          </Segment>
        </Form>
      </Grid.Column>
    </Grid>
  );
};
export default XLogin;
