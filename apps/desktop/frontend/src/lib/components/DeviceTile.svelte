<script lang="ts">
  import Color from "$lib/color";
  import { device } from "$lib/wailsjs/go/models";

  export let deviceInfo: device.DeviceInfo;
  const NAMES = ["DESKTOP", "MOBILE", "TV", "WEB"];

  let color = Color.fromGO(deviceInfo.Color);
  // name, color, addressname, address, type
</script>

<div class="device-tile-container">
  <div class="profile">
    <span
      class="image"
      style="background-color: {color.toHexString()};"
      title={deviceInfo.Code}
    >
      {NAMES[deviceInfo.Type]}
    </span>
  </div>
  <div class="info">
    <h5>{deviceInfo.Name}</h5>
    <code
      title={deviceInfo.Address.IP + ":" + deviceInfo.Address.Port.toString()}
    >
      {deviceInfo.Address.Phrase}
    </code>
  </div>
</div>

<style lang="sass">
@use '$lib/styles/theme'

div.device-tile-container
  $msep: 2px theme.$secondary-d2 solid
  $mseph:  2px theme.$secondary solid
  display: flex
  background-color: theme.$secondary-d3
  width: fit-content
  border: $msep
  &:hover
    border: $mseph
    background-color: theme.$secondary-d2
    div.info
      border-left: $mseph
  div.profile
    span.image
      display: inline-block
      height: 100px
      width: 100px
      text-align: center
      padding-top: calc(50% - 12px)
  div.info
    border-left: $msep
    padding-left: 20px
    padding-right: 20px
    h5
      font-size: xx-large
    code
      font-size: small

</style>
