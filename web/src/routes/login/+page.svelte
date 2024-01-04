<script>
  import { PUBLIC_GOOGLE_CLIENT_ID } from '$env/static/public';
  import GoogleLogin from './GoogleLogin.svelte';

  let params = {};

  const handleSignInSuccess = (idToken) => {
    // Handle the ID token as needed.
    console.log('ID Token:', idToken);
  };
</script>

<div class="h-screen w-screen flex justify-center items-center">
  <div class="p-4 rounded-md border border-slate-800 overflow-x-clip">
    <GoogleLogin clientID={PUBLIC_GOOGLE_CLIENT_ID} />

    <pre>
      {JSON.stringify(params, null, 2)}
    </pre>
  </div>
</div>
