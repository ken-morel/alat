<script lang="ts" module>
</script>

<script lang="ts">
  import { page } from "$app/state";
  import { Navigation } from "@skeletonlabs/skeleton-svelte";
  import { connectedDevice } from "$lib/store";
  import { slide } from "svelte/transition";
  import guessIcon from "$lib/icons";
  import AlatLogo from "$lib/assets/logo.svg";

  import IconDashboard from "@lucide/svelte/icons/hotel";
  import IconDevices from "@lucide/svelte/icons/phone";
  import IconSettings from "@lucide/svelte/icons/settings";
  import IconFileSend from "@lucide/svelte/icons/file";
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
      <Navigation.Tile
        selected={page.url.pathname == "/filesend"}
        label="Send files"
        href="/filesend"
      >
        <IconFileSend />
      </Navigation.Tile>
      {#if $connectedDevice}
        {@const Icon = guessIcon($connectedDevice.info.type)}
        <div transition:slide class="w-full">
          <Navigation.Tile
            selected={page.url.pathname == "/dashboard/device"}
            label={$connectedDevice.info.name}
            href="/dashboard/device"
          >
            <Icon color={$connectedDevice.info.color.hex} />
          </Navigation.Tile>
        </div>
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
