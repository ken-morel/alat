<script lang="ts">
  import { SettingsSetSetupComplete } from "$lib/wails/wailsjs/go/app/App";
  import { nextUrl, prevUrl } from "../wizard.svelte";
  import { Check } from "lucide-svelte";

  prevUrl.set("/setup/services/foldersync");
  nextUrl.set("/");

  SettingsSetSetupComplete(true);
</script>

<div class="flex flex-col items-center space-y-4 text-center">
  <div
    class="flex h-32 w-32 items-center justify-center rounded-full bg-success-500 text-white"
  >
    <Check size={100} />
  </div>
  <h1 class="text-2xl font-bold">Setup Complete</h1>
  <p class="text-sm text-surface-300">
    You're all set up and ready to go. You can change these settings at any time
    in the app's settings.
  </p>
</div>
