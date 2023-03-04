import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import Card from "@mui/material/Card";
import { CardActions, CardHeader, CardContent, Avatar } from "@mui/material";
import { Google, CloudRounded } from "@mui/icons-material";
import useAuth from "./AuthProvider";

const GOOGLE_OAUTH_URL = "https://accounts.google.com/o/oauth2/v2/auth";

const getLoginLink = () => {
  const params = new URLSearchParams({
    client_id: process.env.REACT_APP_GOOGLE_CLIENT_ID,
    redirect_uri: window.location.origin,
    response_type: "token",
    scope: "openid email profile",
    // prompt: "select_account",
  });
  return `${GOOGLE_OAUTH_URL}?${params.toString()}`;
};

function App() {
  const { auth } = useAuth();

  return (
    <>
      <div
        style={{
          backgroundColor: "#efefef",
          minHeight: "100vh",
          display: "flex",
        }}
      >
        <Grid
          style={{ height: "100vh" }}
          spacing={2}
          container
          justifyContent="center"
          alignItems="center"
        >
          <Grid item xs={6}>
            <Card variant="outlined">
              <CardHeader
                avatar={
                  <Avatar style={{ backgroundColor: "#ff9800" }}>
                    <CloudRounded />
                  </Avatar>
                }
                title="Kubestack portal"
                subheader="Manage your infrastructure"
              ></CardHeader>
              <CardContent>
                <pre>{JSON.stringify(auth, null, 2)}</pre>
              </CardContent>
              <CardActions>
                <Button
                  startIcon={<Google />}
                  href={getLoginLink()}
                  variant="outlined"
                  style={{ textTransform: "none" }}
                >
                  Sign in with Google
                </Button>
              </CardActions>
            </Card>
          </Grid>
        </Grid>
      </div>
    </>
  );
}

export default App;
