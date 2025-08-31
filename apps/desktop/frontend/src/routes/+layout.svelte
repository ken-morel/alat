<script lang="ts">
  import "../app.css";
  import { Navigation } from "@skeletonlabs/skeleton-svelte";
  import AlatLogo from "$lib/assets/logo.svg";

  import IconDashboard from "@lucide/svelte/icons/hotel";
  import IconDevices from "@lucide/svelte/icons/phone";
  import IconSettings from "@lucide/svelte/icons/settings";

  let { children, data } = $props();
</script>

<div class="flex h-full">
  {#if data.ready}
    <div class=" h-full">
      <Navigation.Rail>
        {#snippet header()}
          <Navigation.Tile title="About alat" classes="bg-primary-700">
            <img src={AlatLogo} alt="Alat Logo" />
          </Navigation.Tile>
        {/snippet}
        {#snippet tiles()}
          <Navigation.Tile label="Dashboard" href="/dashboard">
            <IconDashboard />
          </Navigation.Tile>
          <Navigation.Tile label="Devices" href="devices">
            <IconDevices />
          </Navigation.Tile>
        {/snippet}
        {#snippet footer()}
          <Navigation.Tile labelExpanded="Settings" href="/setup" title="Setup">
            <IconSettings />
          </Navigation.Tile>
        {/snippet}
      </Navigation.Rail>
    </div>
  {/if}

  {@render children?.()}
</div>
