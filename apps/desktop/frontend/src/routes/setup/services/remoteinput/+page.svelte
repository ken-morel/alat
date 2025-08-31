<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetRemoteInputSettings } from "$lib/wails/wailsjs/go/app/App";
  import { createCheckbox, createSlider, createSwitch } from "@melt-ui/svelte";

  let { data } = $props();

  prevUrl.set("/setup/services/media");
  nextUrl.set("/setup/services/foldersync");

  let settings = $state(data.settings);

  const {
    elements: { root: enabledSwitch, thumb: enabledThumb },
    states: { checked: enabledChecked },
  } = createSwitch({
    defaultChecked: settings.Enabled,
  });

  const {
    elements: { root: naturalScrollingCheckbox, input: naturalScrollingInput },
    states: { checked: naturalScrollingChecked },
  } = createCheckbox({
    defaultChecked: settings.NaturalScrolling,
  });

  const {
    elements: { root: slider, range, thumb },
    states: { value },
  } = createSlider({
    defaultValue: [settings.MouseSensitivity],
    max: 5,
    min: 0.1,
    step: 0.1,
  });

  $effect(() => {
    settings.Enabled = $enabledChecked;
    settings.NaturalScrolling = $naturalScrollingChecked;
    settings.MouseSensitivity = $value[0];
    return () => {
      SettingsSetRemoteInputSettings(settings).catch((err: any) => {
        console.error("Failed to save remote input settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Remote Input</h1>
    <p class="text-sm text-surface-500">
      Use your phone as a remote control for your computer.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Remote Input</label>
      <button use:enabledSwitch class="switch">
        <span use:enabledThumb />
      </button>
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        <div class="flex items-center justify-between">
          <label for="natural-scrolling" class="font-medium"
            >Natural Scrolling</label
          >
          <button class="checkbox" use:naturalScrollingCheckbox>
            <input use:naturalScrollingInput />
          </button>
        </div>
        <div class="space-y-2">
          <label for="mouse-sensitivity" class="font-medium"
            >Mouse Sensitivity ({settings.MouseSensitivity})</label
          >
          <span use:slider class="slider">
            <span use:range />
            <span use:thumb />
          </span>
        </div>
      </div>
    {/if}
  </div>
</div>
