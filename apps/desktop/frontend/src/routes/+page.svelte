<script lang="ts">
  import AnimatedLogo from "$lib/assets/logo-animated.svg";
  import { goto } from "$app/navigation";
  // import { setLocale } from '$lib/paraglide/runtime';
  // import { page } from '$app/state';
  // import { m } from '$lib/paraglide/messages.js';
  goto("/dashboard");
</script>

<section class="hero w3-container">
  <div class="left">
    <img src={AnimatedLogo} alt="alat's logo" width="400" />
  </div>
  <div class="right">
    <h2>alat</h2>
  </div>
</section>

<style lang="sass">
@use '$lib/styles/theme.sass'

section.hero
  width: 100%
  color: theme.$primary
</style>
