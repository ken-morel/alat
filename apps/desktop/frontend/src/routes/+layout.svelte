<script lang="ts">
  import "../app.css";
  import favicon from "$lib/assets/logo.svg";
  import Toaster from "$lib/Toaster.svelte";

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>

<Toaster />

{@render children?.()}
