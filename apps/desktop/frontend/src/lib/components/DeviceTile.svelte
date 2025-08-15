<script lang="ts">
  import Color from "$lib/color";
  import type { device } from "$lib/wailsjs/go/models";

  export let deviceInfo: device.DeviceInfo;

  let color = Color.fromGO(deviceInfo.Color);
  // name, color, addressname, address, type
</script>

<div class="device-tile-container">
  <div class="profile">
    <span class="image" style="background-color: {color.toHexString()};"></span>
  </div>
</div>

<style lang="sass">
div.device-tile-container
  display: flex
  div.profile
    span.image
      display: inline-block
      height: 100px
      width: 100px

</style>
