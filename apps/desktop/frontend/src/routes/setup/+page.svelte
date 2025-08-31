<script lang="ts">
  import { nextUrl, prevUrl } from "./wizard.svelte";
  prevUrl.set(null);
  nextUrl.set("/setup/device");
</script>

<div class="content-center">
  <h3 class="h3">Setup assistant</h3>
  <p class="text-surface-500">Let's get your device set up.</p>
</div>
