<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetUniversalClipboardSettings } from "$lib/wails/wailsjs/go/app/App";
  import { createCheckbox, createSwitch } from "@melt-ui/svelte";

  let { data } = $props();

  prevUrl.set("/setup/services/fileshare");
  nextUrl.set("/setup/services/notifications");

  let settings = $state(data.settings);

  const {
    elements: { root: enabledSwitch, thumb: enabledThumb },
    states: { checked: enabledChecked },
  } = createSwitch({
    defaultChecked: settings.Enabled,
  });
  const {
    elements: { root: syncTextCheckbox, input: syncTextInput },
    states: { checked: syncTextChecked },
  } = createCheckbox({
    defaultChecked: settings.SyncText,
  });
  const {
    elements: { root: syncImagesCheckbox, input: syncImagesInput },
    states: { checked: syncImagesChecked },
  } = createCheckbox({
    defaultChecked: settings.SyncImages,
  });
  const {
    elements: {
      root: ignorePasswordManagersCheckbox,
      input: ignorePasswordManagersInput,
    },
    states: { checked: ignorePasswordManagersChecked },
  } = createCheckbox({
    defaultChecked: settings.IgnorePasswordManagers,
  });

  $effect(() => {
    settings.Enabled = $enabledChecked;
    settings.SyncText = $syncTextChecked;
    settings.SyncImages = $syncImagesChecked;
    settings.IgnorePasswordManagers = $ignorePasswordManagersChecked;
    return () => {
      SettingsSetUniversalClipboardSettings(settings).catch((err: any) => {
        console.error("Failed to save clipboard settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Universal Clipboard</h1>
    <p class="text-sm text-surface-500">
      Sync your clipboard across all your devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Universal Clipboard</label
      >
      <button use:enabledSwitch class="switch">
        <span use:enabledThumb />
      </button>
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        <div class="flex items-center justify-between">
          <label for="sync-text" class="font-medium">Sync Text</label>
          <button class="checkbox" use:syncTextCheckbox>
            <input use:syncTextInput />
          </button>
        </div>
        <div class="flex items-center justify-between">
          <label for="sync-images" class="font-medium">Sync Images</label>
          <button class="checkbox" use:syncImagesCheckbox>
            <input use:syncImagesInput />
          </button>
        </div>
        <div class="flex items-center justify-between">
          <label for="ignore-password-managers" class="font-medium"
            >Ignore Password Managers</label
          >
          <button class="checkbox" use:ignorePasswordManagersCheckbox>
            <input use:ignorePasswordManagersInput />
          </button>
        </div>
      </div>
    {/if}
  </div>
</div>
