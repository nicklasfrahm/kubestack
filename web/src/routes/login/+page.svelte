<script>
  import { PUBLIC_GOOGLE_CLIENT_ID } from '$env/static/public';
  import { onMount } from 'svelte';

  let redirectURI = '';
  let params = {};

  onMount(() => {
    redirectURI = `${window.location.origin}/login`;
    params = parseHash(window.location.hash);
    if (Object.keys(params).length > 0) {
      window.location.hash = '';
    }
  });

  const parseHash = (hash) => {
    const data = hash.substring(1);

    if (!data) {
      return null;
    }

    return data.split('&').reduce((parsed, pair) => {
      const params = { ...parsed };
      const [key, value] = pair.split('=');
      params[decodeURIComponent(key)] = decodeURIComponent(value);
      return params;
    }, {});
  };

  const GOOGLE_OAUTH_URL = 'https://accounts.google.com/o/oauth2/v2/auth';

  const getLoginLink = (redirectURI) => {
    console.log(redirectURI);
    const params = new URLSearchParams({
      client_id: PUBLIC_GOOGLE_CLIENT_ID,
      redirect_uri: redirectURI,
      response_type: 'token',
      scope: 'openid email profile',
      // prompt: "select_account",
    });
    return `${GOOGLE_OAUTH_URL}?${params.toString()}`;
  };
</script>

<div class="h-screen w-screen flex justify-center items-center">
  <div class="p-4 rounded-md border border-slate-800 overflow-x-clip">
    <a href={getLoginLink(redirectURI)}>Login with Google</a>
    <pre>
      {JSON.stringify(params, null, 2)}
    </pre>
  </div>
</div>
