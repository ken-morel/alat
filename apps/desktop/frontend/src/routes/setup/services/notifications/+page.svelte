<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetNotificationSyncSettings } from "$lib/wails/wailsjs/go/app/App";
  import { createCheckbox, createSwitch } from "@melt-ui/svelte";

  let { data } = $props();

  prevUrl.set("/setup/services/clipboard");
  nextUrl.set("/setup/services/media");

  let settings = $state(data.settings);

  const {
    elements: { root: enabledSwitch, thumb: enabledThumb },
    states: { checked: enabledChecked },
  } = createSwitch({
    defaultChecked: settings.Enabled,
  });

  const {
    elements: { root: quickRepliesCheckbox, input: quickRepliesInput },
    states: { checked: quickRepliesChecked },
  } = createCheckbox({
    defaultChecked: settings.QuickReplies,
  });

  $effect(() => {
    settings.Enabled = $enabledChecked;
    settings.QuickReplies = $quickRepliesChecked;
    return () => {
      SettingsSetNotificationSyncSettings(settings).catch((err: any) => {
        console.error("Failed to save notification settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Notification Sync</h1>
    <p class="text-sm text-surface-500">
      Sync notifications between your devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Notification Sync</label>
      <button use:enabledSwitch class="switch">
        <span use:enabledThumb />
      </button>
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        <div class="flex items-center justify-between">
          <label for="quick-replies" class="font-medium"
            >Enable Quick Replies</label
          >
          <button class="checkbox" use:quickRepliesCheckbox>
            <input use:quickRepliesInput />
          </button>
        </div>
      </div>
    {/if}
  </div>
</div>
