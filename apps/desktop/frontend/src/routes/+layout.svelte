<script lang="ts">
  import favicon from "$lib/assets/logo.svg";

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>
<header class="w3-bar w3-top">
  <a class="w3-button" href="/dashboard">Dashboard</a>
</header>
{@render children?.()}
