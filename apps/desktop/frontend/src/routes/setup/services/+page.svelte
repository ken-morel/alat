<script lang="ts">
  import { nextUrl, prevUrl } from "../wizard.svelte";
  prevUrl.set("/setup/device");
  nextUrl.set("/setup/services/sysinfo");
</script>

<div class="">
  <h3 class="h3">Services</h3>
  <p class="text-surface-300 p">
    Services are what you allow other (paired) devices to access.
  </p>
  <p class="text-surface-300 p">
    The services are already enabled and preconfigured you can review and modify
    them now or later in the settings.
  </p>
</div>
