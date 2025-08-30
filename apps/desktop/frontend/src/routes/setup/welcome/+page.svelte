<script lang="ts">
  import { wizardState } from "../wizard-state";

  wizardState.prevUrl = null; // No "Back" button on the first step
  wizardState.nextUrl = "/setup/appearance";
</script>

<div class="space-y-2">
  <h3 class="h3">Welcome to Alat</h3>
  <p class="text-surface-500">Let's get your device set up.</p>
</div>

<!-- Form for Device Name -->
<label class="label">
  <span>Device Name</span>
  <input class="input" type="text" placeholder="My Awesome Laptop" />
</label>

