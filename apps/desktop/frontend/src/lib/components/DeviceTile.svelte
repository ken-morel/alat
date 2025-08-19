<script lang="ts">
  import Color from "$lib/color";
  import { device, service } from "$lib/wailsjs/go/models";
  import { ICONS, NAMES } from "$lib";

  let {
    services = undefined,
    deviceInfo,
  }: { deviceInfo: device.DeviceInfo; services?: service.Service[] } = $props();
  let color = Color.fromGO(deviceInfo.color);
  // name, color, addressname, address, type
</script>

<div class="device-tile-container">
  <div class="profile">
    <span
      class="image"
      style="background-color: {color.toHexString()};"
      title={deviceInfo.code}
    >
      {ICONS[deviceInfo.type]}
    </span>
  </div>
  <div class="info">
    <h5>{deviceInfo.name}</h5>
    <code
      title={deviceInfo.address.ip + ":" + deviceInfo.address.port.toString()}
    >
      {deviceInfo.address.phrase}
    </code>
  </div>
  {#if services}
    <div class="services">
      {#each services as service}
        <span
          style:color={service.enabled ? "lightgreen" : "red"}
          style:display="block"
          title={service.name}>{service.name}</span
        >
      {/each}
    </div>
  {/if}
</div>

<style lang="sass">
@use '$lib/styles/theme'

div.device-tile-container
  $msep: 2px theme.$border-dark solid
  $mseph:  2px theme.$border-light solid
  display: flex
  background-color: theme.$primary-d3
  width: fit-content
  border: $msep
  &:hover
    border: $mseph
    background-color: theme.$primary-d2
    div.info
      border-left: $mseph
  div.profile
    span.image
      display: inline-block
      height: 100px
      width: 100px
      text-align: center
      padding-top: calc(50% - 30px)
      font-size: 60px
  div.info
    border-left: $msep
    padding-left: 10px
    padding-right: 10px
    text-align: right
    h5
      font-size: xx-large
    code
      font-size: small
      background-color: #8888
      padding: 2px 4px
  div.services
    border-left: $msep
    padding: 5px 8px
    color: theme.$text-primary
    background-color: theme.$tertiary-d4

</style>
