<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import favicon from "$lib/assets/logo.svg";
  import { IsSetupComplete } from "$lib/wailsjs/go/app/App";
  import { onMount } from "svelte";

  let { children } = $props();

  onMount(async () => {
    const isSetup = await IsSetupComplete();
    if (!isSetup && $page.url.pathname !== "/setup") {
      await goto("/setup");
    }
  });
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>
<header class="w3-bar w3-top">
  <a class="w3-button" href="/dashboard">Dashboard</a>
</header>
{@render children?.()}
