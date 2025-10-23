<script lang="ts">
  import { connectedDevice } from "$lib/store";
  import guessIcon from "$lib/icons";
  import DeviceBattery from "$lib/components/DeviceBattery.svelte";
  import { sendingDevices, isSendingTo } from "$lib/filesend";
  import { goto } from "$app/navigation";
  let dev = $derived($connectedDevice);
</script>

<div class="h-full w-full grid place-items-center">
  {#if dev}
    {@const Icon = guessIcon(dev.info.type)}
    <div
      class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
    >
      <header class="">
        <div class="flex p-4 border-b border-surface-200-800">
          <div class="flex p-4">
            <Icon size={50} color={dev.info.color.hex} />
            <h3 class="ml-8 h3">{dev.info.name}</h3>
          </div>
          <div class="flex flex-2 pr-1">
            <div class="flex grow"></div>
            <DeviceBattery {dev} />
          </div>
        </div>
        <button
          class="btn preset-filled-primary-700-300"
          onclick={() => {
            if (!isSendingTo(dev))
              sendingDevices.update((devices) => devices.concat([dev]));
            goto("/filesend");
          }}>Send files</button
        >
      </header>
    </div>
  {:else}
    <div class="card preset-filled-error-400-600 p-8">
      <p class="h2 text-error-700-300">No device selected.</p>
    </div>
  {/if}
</div>
