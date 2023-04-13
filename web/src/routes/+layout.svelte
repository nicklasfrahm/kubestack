<script>
  import '../app.css';
</script>

<nav class="navigation">
  <a href="/">Kubestack</a>
</nav>

<main class="content">
  <slot />
</main>

<style lang="postcss">
  .navigation {
    @apply rounded-md p-4 mx-auto max-w-screen-xl bg-white border-2 border-slate-800;
  }

  .content {
    @apply mx-auto max-w-screen-xl pt-4;
  }
</style>
