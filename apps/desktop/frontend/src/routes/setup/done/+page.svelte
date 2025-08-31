<script lang="ts">
  import { nextUrl, prevUrl } from "../wizard.svelte";
  import { Check } from "lucide-svelte";

  prevUrl.set("/setup/services/foldersync");
  nextUrl.set("/");
</script>

<div class="flex flex-col items-center space-y-4 text-center">
  <div
    class="flex h-16 w-16 items-center justify-center rounded-full bg-success-500 text-white"
  >
    <Check size={48} />
  </div>
  <h1 class="text-2xl font-bold">Setup Complete</h1>
  <p class="text-sm text-surface-500">
    You're all set up and ready to go. You can change these settings at any time
    in the app's settings.
  </p>
</div>
