<script lang="ts">
  import "../app.css";
  import favicon from "$lib/assets/favicon.svg";

  let { children } = $props();

  import { Switch } from "@skeletonlabs/skeleton-svelte";

  let checked = $state(false);

  $effect(() => {
    const mode = localStorage.getItem("mode") || "light";
    checked = mode === "dark";
  });

  const onCheckedChange = (event: { checked: boolean }) => {
    const mode = event.checked ? "nosh" : "alat";
    document.documentElement.setAttribute("data-theme", mode);
    localStorage.setItem("mode", mode);
    checked = event.checked;
  };
</script>

<Switch {checked} {onCheckedChange}></Switch>

<svelte:head>
  <link rel="icon" href={favicon} />
  <script>
    const mode = localStorage.getItem("mode") || "light";
    document.documentElement.setAttribute("data-theme", mode);
  </script>
</svelte:head>

{@render children?.()}
