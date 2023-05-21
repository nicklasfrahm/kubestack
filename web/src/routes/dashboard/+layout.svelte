<nav class="navigation">
  <a href="/">🧀 <span class="underline">bri3</span></a>
</nav>

<main class="content">
  <slot />
</main>

<style lang="postcss">
  :global(html) {
    @apply p-4;
  }

  .navigation {
    @apply rounded-md p-4 mx-auto max-w-screen-xl bg-white border border-slate-800;
  }

  .navigation a {
    @apply text-xl font-bold text-slate-800;
  }

  .content {
    @apply mx-auto max-w-screen-xl pt-4;
  }
</style>
