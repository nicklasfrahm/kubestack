<script>
  import { GoogleOAuthProvider } from 'google-oauth-gsi';

  export let clientID;

  const googleProvider = new GoogleOAuthProvider({
    clientId: clientID,
    onScriptLoadError: () => console.log('onScriptLoadError'),
    onScriptLoadSuccess: () => console.log('onScriptLoadSuccess'),
  });

  const login = googleProvider.useGoogleLogin({
    flow: 'implicit',
    clientId: clientID,
    onSuccess: (tokenResponse) => console.log(tokenResponse),
  });
</script>

<div
  class="p-4 rounded-md border border-slate-800 hover:bg-blue-600 hover:text-white cursor-pointer"
  on:click={login}
  on:keypress={login}
>
  Login with Google
</div>
