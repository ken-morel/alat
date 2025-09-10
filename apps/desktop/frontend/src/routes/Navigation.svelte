<script lang="ts" module>
</script>

<script lang="ts">
  import { page } from "$app/state";
  import { Navigation } from "@skeletonlabs/skeleton-svelte";
  import { connectedDevice } from "$lib/store";
  import guessIcon from "$lib/icons";
  import AlatLogo from "$lib/assets/logo.svg";

  import IconDashboard from "@lucide/svelte/icons/hotel";
  import IconDevices from "@lucide/svelte/icons/phone";
  import IconSettings from "@lucide/svelte/icons/settings";
</script>

<div class="">
  <Navigation.Rail>
    {#snippet header()}
      <Navigation.Tile
        selected={page.url.pathname == "/about"}
        title="About alat"
        classes="bg-primary-700"
      >
        <img src={AlatLogo} alt="Alat Logo" />
      </Navigation.Tile>
    {/snippet}
    {#snippet tiles()}
      <Navigation.Tile
        selected={page.url.pathname == "/dashboard"}
        label="Dashboard"
        href="/dashboard"
      >
        <IconDashboard />
      </Navigation.Tile>
      <Navigation.Tile
        selected={page.url.pathname == "/devices"}
        label="Devices"
        href="/devices"
      >
        <IconDevices />
      </Navigation.Tile>
      {#if $connectedDevice}
        {@const Icon = guessIcon($connectedDevice.Info.Type)}
        <Navigation.Tile
          selected={page.url.pathname == "/dashboard/device"}
          label={$connectedDevice.Info.Name}
          href="/dashboard/device"
        >
          <Icon color={$connectedDevice.Info.Color.Hex} />
        </Navigation.Tile>
      {/if}
    {/snippet}
    {#snippet footer()}
      <Navigation.Tile
        selected={page.url.pathname.startsWith("/setup")}
        labelExpanded="Settings"
        href="/setup"
        title="Setup"
      >
        <IconSettings />
      </Navigation.Tile>
    {/snippet}
  </Navigation.Rail>
</div>
